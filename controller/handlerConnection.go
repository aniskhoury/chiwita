/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package controller

import (
	"chiwita/global"
	"chiwita/structure"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func Autentification(s *websocket.Conn, db *sql.DB) (bool, structure.User) {

	var nickUser string
	var passwordUser string
	var u structure.User

	//Obtener el nick y contrasena para autentificar
	mt, obtenerNickContrasena, err := s.ReadMessage()
	//mt == 1 es si el tipo de frame es TextMessage
	//mt == 2 es si el tipo de frame es BinaryMessage
	if err != nil {
		log.Print("upgrade:", err)
		return false, u
	}
	if mt == 1 || mt == 2 {
		var result = strings.Split(string(obtenerNickContrasena), " ")
		nickUser = result[0]
		passwordUser = result[1]
		stmt, err := db.Prepare("select * from users where nick = ? and password = ?")
		filas, err := stmt.Query(nickUser, passwordUser)
		if err != nil {
			return false, u
		}

		for filas.Next() {
			/* Obtención de valores de los usuarios desde la base de datos
			 */
			if err := filas.Scan(&u.Iduser, &u.Nick, &u.Password, &u.Email, &u.Date_register, &u.Activated, &u.Bot); err != nil {
				return false, u
			}
			/*Inicialización de los datos del usuario
			 */
			u.Connection = s
			u.Channels = make(map[string]string)
			return true, u
		}

	}
	return false, structure.User{}
}

func HandlerConnection(w http.ResponseWriter, r *http.Request) {
	c, err := global.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	//que al final de la ejecución del websocket se ejecute el cierre del websocket
	defer c.Close()
	var resulAuth, u = Autentification(c, global.Db)

	if resulAuth == false {
		return
	}

	global.MutexUsers.Lock()
	u.Connection = c
	global.Users[u.Nick] = u
	global.MutexUsers.Unlock()

	global.MutexCounterUsers.Lock()
	global.CounterUsers = global.CounterUsers + 1
	global.MutexCounterUsers.Unlock()
	c.WriteMessage(1, []byte("SUCCESSFUL_AUTHENTIFICATION"))
	for {
		mt, mensaje, err := c.ReadMessage()
		/*
			if entry, ok := global.Usuarios[c.NetConn()]; ok {

				mutexUsers.Lock()
				// Then we modify the copy
				entry.contador = entry.contador + 1
				users[c.NetConn()] = entry
				mutexUsers.Unlock()
			}
		*/
		if err != nil {
			break
		}

		/*###############################################
		  ##### Gestión del protocolo de chat-juego #####
		  ###############################################

		*/
		var command = strings.Split(string(mensaje), " ")
		var mensajeEnviar string

		if len(command) > 1 {
			switch command[0] {
			case "LEFT_CHANNEL":
				if len(command) == 2 {
					global.MutexChannels.Lock()
					channel, exist := global.Channels[command[1]]
					if exist {
						_, userExist := global.Channels[channel.Name].Users[u.Nick]
						if userExist {
							if entrada, ok := global.Channels[channel.Name]; ok {
								entrada.CounterUsers = entrada.CounterUsers - 1
								global.Channels[channel.Name] = entrada
							}
							delete(global.Channels[channel.Name].Users, u.Nick)
							delete(channel.Users, u.Nick)
							for claveNick, usuario := range global.Channels[command[1]].Users {
								mensajeEnviar = fmt.Sprintf("USER_LEFT_CHANNEL %s %s", command[1], usuario.Nick)
								channel.Users[claveNick].Connection.WriteMessage(mt, []byte(mensajeEnviar))
							}
						} else {
							mensajeEnviar = fmt.Sprintf(("ERROR_DOESNT_EXIST_CHANNEL"))
						}

					} else {
						mensajeEnviar = fmt.Sprintf("ERROR_LEFT_CHANNEL %s DOESNT EXIST", command[1])
						u.Connection.WriteMessage(mt, []byte(mensajeEnviar))
					}
					global.MutexChannels.Unlock()
				} else {
					mensajeEnviar = fmt.Sprintf("ERROR_LEFT_CHANNEL FOR LEFT CHANNEL USE LEFT_CHANNEL NAMECHANNEL")
					u.Connection.WriteMessage(mt, []byte(mensajeEnviar))
				}
			case "MESSAGE_CHANNEL":
				if len(command) > 2 {
					global.MutexChannels.Lock()
					channel, exist := global.Channels[command[1]]

					if exist {
						_, exist := channel.Users[u.Nick]
						if exist {
							for nick, usuario := range channel.Users {
								if u.Nick != nick {
									mensajeEnviar = fmt.Sprintf("MESSAGE_CHANNEL %s %s", command[1], strings.SplitAfterN(string(mensaje), " ", 3)[2])
									usuario.Connection.WriteMessage(mt, []byte(mensajeEnviar))
								}
							}
						} else {
							mensajeEnviar = fmt.Sprintf("ERROR_NEED_STAY_IN_CHANNEL %s for send messages", command[1])
							u.Connection.WriteMessage(mt, []byte(mensajeEnviar))
						}

					} else {
						mensajeEnviar = fmt.Sprintf("ERROR_MESSAGE_CHANNEL Error channel %s doesnt exist", channel.Name)
						u.Connection.WriteMessage(mt, []byte(mensajeEnviar))
					}
					global.MutexChannels.Unlock()
				} else {
					mensajeEnviar = fmt.Sprintf("ERROR_MESSAGE_CHANNEL Error while sending %s", mensaje)
					u.Connection.WriteMessage(mt, []byte(mensajeEnviar))
				}
			case "PRIVATE_MESSAGE":
				if len(command) > 2 {
					global.MutexUsers.Lock()
					user, exist := global.Users[command[1]]
					if exist {
						mensajeEnviar = strings.SplitAfterN(string(mensaje), " ", 3)[2]
						mensajeEnviar = fmt.Sprintf("PRIVATEMESSAGE %s %s", u.Nick, mensajeEnviar)
						user.Connection.WriteMessage(mt, []byte(mensajeEnviar))
					} else {
						mensajeEnviar = fmt.Sprintf("ERROR_USER_DISCONNECTED %s", command[1])
						u.Connection.WriteMessage(mt, []byte(mensajeEnviar))
					}
					global.MutexUsers.Unlock()
				}
			case "JOIN_CHANNEL":
				/*
					global.MutexCanales.Lock()
					var usuariosEnElCanal = global.Canales[comando[1]]
					usuariosEnElCanal.Usuarios = append(usuariosEnElCanal.Usuarios, u)
					//enviar mensaje de que el usuario ha entrado a los usuarios del canal
					for k, v := range usuariosEnElCanal.Usuarios {
						fmt.Println("Enviando mensaje de que ha entrado usuario %s al canal k", v, k)

					}
					global.MutexCanales.Unlock()*/
				global.MutexChannels.Lock()
				channel, ok := global.Channels[command[1]]
				if ok {
					if len(command) == 2 {

						var data = channel.Users

						for nick, usuario := range data {
							if u.Nick != nick {
								mensajeEnviar = fmt.Sprintf("JOINED %s %s", command[1], u.Nick)
								usuario.Connection.WriteMessage(mt, []byte(mensajeEnviar))
							}

						}
						if entrada, ok := global.Channels[channel.Name]; ok {
							entrada.CounterUsers = entrada.CounterUsers + 1
							global.Channels[channel.Name] = entrada
						}
						channel.Users[u.Nick] = u
						u.Channels[command[1]] = command[1]
					} else {
						c.WriteMessage(mt, []byte("ERROR_DOESNT_EXIST_CHANNEL"))
					}
				} else {
					c.WriteMessage(mt, []byte("ERROR_DOESNT_EXIST_CHANNEL"))
				}
				global.MutexChannels.Unlock()

			default:
				c.WriteMessage(mt, []byte("ERROR_COMMAND_NOT_FOUND"))
			}
		} else {
			c.WriteMessage(mt, []byte("ERROR_COMMAND_NOT_FOUND"))
		}

		/*err = c.WriteMessage(mt, mensaje)
		if err != nil {
			log.Println("write:", err)
			break
		}*/
	}
	for nameChannel, _ := range u.Channels {
		var mensajeEnviar string
		var channel = global.Channels[nameChannel]
		channel.MutexChannel.Lock()
		for nick, _ := range global.Channels[nameChannel].Users {
			if u.Nick != nick {
				mensajeEnviar = fmt.Sprintf("USER_LEFT_CHANNEL %s %s", nameChannel, u.Nick)
				global.Users[nick].Connection.WriteMessage(1, []byte(mensajeEnviar))
			}
		}
		channel.MutexChannel.Unlock()

	}
	global.MutexChannels.Lock()
	for k, _ := range u.Channels {

		if entry, ok := global.Channels[k]; ok {
			entry.CounterUsers = entry.CounterUsers - 1
			global.Channels[k] = entry
		}

		delete(global.Channels[k].Users, u.Nick)
	}
	global.MutexChannels.Unlock()
	global.MutexUsers.Lock()
	delete(global.Users, u.Nick)
	global.MutexUsers.Unlock()
	global.MutexCounterUsers.Lock()
	global.CounterUsers = global.CounterUsers - 1
	global.MutexCounterUsers.Unlock()
}
