/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package controlador

import (
	"chiwita/estructuras"
	"chiwita/global"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func Autentificacion(s *websocket.Conn, db *sql.DB) (bool, estructuras.Usuario) {

	var nickUsuario string
	var contrasenaUsuario string
	var u estructuras.Usuario

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
		nickUsuario = result[0]
		contrasenaUsuario = result[1]
		stmt, err := db.Prepare("select * from usuarios where nick = ? and contrasena = ?")
		filas, err := stmt.Query(nickUsuario, contrasenaUsuario)
		if err != nil {
			return false, u
		}

		for filas.Next() {
			/* Obtención de valores de los usuarios desde la base de datos
			 */
			if err := filas.Scan(&u.Idusuario, &u.Nick, &u.Contrasena, &u.Email, &u.Fecha_registro, &u.Activado, &u.Bot); err != nil {
				return false, u
			}
			/*Inicialización de los datos del usuario
			 */
			u.Conexion = s
			u.Canales = make(map[string]string)
			return true, u
		}

	}
	return false, estructuras.Usuario{}
}

func GestorConexion(w http.ResponseWriter, r *http.Request) {
	c, err := global.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	//que al final de la ejecución del websocket se ejecute el cierre del websocket
	defer c.Close()
	var resulAutentificacion, u = Autentificacion(c, global.Db)

	if resulAutentificacion == false {
		return
	}

	global.MutexUsuarios.Lock()
	u.Conexion = c
	global.Usuarios[u.Nick] = u

	global.MutexUsuarios.Unlock()
	/*	global.MutexSocketUsuarios.Lock()
		global.SocketUsuarios[c.NetConn()] = u
		global.MutexSocketUsuarios.Unlock()
	*/
	var msgAutentificacion = []byte("AUTENTIFICACION_CORRECTA")
	c.WriteMessage(1, msgAutentificacion)
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
			log.Println("read:", err)
			println("%s", err.Error())
			break
		}

		/*###############################################
		  ##### Gestión del protocolo de chat-juego #####
		  ###############################################

		*/
		var comando = strings.Split(string(mensaje), " ")
		var mensajeEnviar string

		if len(comando) > 1 {
			switch comando[0] {
			case "SALIRCANAL":
				if len(comando) == 2 {
					global.MutexCanales.Lock()
					canal, existe := global.Canales[comando[1]]
					if existe {
						_, existeUsuarioEnCanal := global.Canales[canal.Nombre].Usuarios[u.Nick]
						if existeUsuarioEnCanal {
							delete(global.Canales[canal.Nombre].Usuarios, u.Nick)
							delete(canal.Usuarios, u.Nick)
							for claveNick, usuario := range global.Canales[comando[1]].Usuarios {
								mensajeEnviar = fmt.Sprintf("USUARIO_SALE_CANAL %s %s", comando[1], usuario.Nick)
								canal.Usuarios[claveNick].Conexion.WriteMessage(mt, []byte(mensajeEnviar))
							}
						} else {
							mensajeEnviar = fmt.Sprintf(("ERROR_USUARIO_NO_EXISTE_CANAL"))
						}

					} else {
						mensajeEnviar = fmt.Sprintf("ERROR_SALIRCANAL Canal %s no existe", comando[1])
						u.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
					}
					global.MutexCanales.Unlock()
				} else {
					mensajeEnviar = fmt.Sprintf("ERROR_SALIRCANAL Para salir de un canal envia SALIRCANAL nombredelcanal")
					u.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
				}
			case "MENSAJECANAL":
				if len(comando) > 2 {
					global.MutexCanales.Lock()
					canal, existe := global.Canales[comando[1]]

					if existe {
						_, existe := canal.Usuarios[u.Nick]
						if existe {
							for claveNick, usuario := range canal.Usuarios {
								if u.Nick != claveNick {
									mensajeEnviar = fmt.Sprintf("MENSAJECANAL %s %s", comando[1], strings.SplitAfterN(string(mensaje), " ", 3)[2])
									usuario.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
								}
							}
						} else {
							mensajeEnviar = fmt.Sprintf("ERROR_DEBES_ESTAR_CANAL %s para enviar mensajes", comando[1])
							u.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
						}

					} else {
						mensajeEnviar = fmt.Sprintf("ERROR_MENSAJECANAL Error canal %s no existe", canal.Nombre)
						u.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
					}
					global.MutexCanales.Unlock()
				} else {
					mensajeEnviar = fmt.Sprintf("ERROR_MENSAJECANAL Error al enviar %s", mensaje)
					u.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
				}
			case "MENSAJEPRIVADO":
				if len(comando) > 2 {
					global.MutexUsuarios.Lock()
					usuario, existe := global.Usuarios[comando[1]]
					if existe {
						mensajeEnviar = strings.SplitAfterN(string(mensaje), " ", 3)[2]
						mensajeEnviar = fmt.Sprintf("MENSAJEPRIVADO %s %s", u.Nick, mensajeEnviar)
						usuario.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
					} else {
						mensajeEnviar = fmt.Sprintf("ERROR_Usuario_NOCONECTADO %s", comando[1])
						u.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
					}
					global.MutexUsuarios.Unlock()
				}
			case "ENTRARCANAL":
				/*
					global.MutexCanales.Lock()
					var usuariosEnElCanal = global.Canales[comando[1]]
					usuariosEnElCanal.Usuarios = append(usuariosEnElCanal.Usuarios, u)
					//enviar mensaje de que el usuario ha entrado a los usuarios del canal
					for k, v := range usuariosEnElCanal.Usuarios {
						fmt.Println("Enviando mensaje de que ha entrado usuario %s al canal k", v, k)

					}
					global.MutexCanales.Unlock()*/
				global.MutexCanales.Lock()
				canal, ok := global.Canales[comando[1]]
				if ok {
					if len(comando) == 2 {
						var datos = canal.Usuarios

						for nick, usuario := range datos {
							if u.Nick != nick {
								mensajeEnviar = fmt.Sprintf("HAENTRADO %s %s", comando[1], u.Nick)
								usuario.Conexion.WriteMessage(mt, []byte(mensajeEnviar))
							}

						}
						canal.Usuarios[u.Nick] = u
						u.Canales[comando[1]] = comando[1]
					} else {
						c.WriteMessage(mt, []byte("Error canal no existe"))
					}
				}
				global.MutexCanales.Unlock()

			default:
				c.WriteMessage(mt, []byte("ERROR_COMANDO comando no encontrado"))
				log.Printf("Error comando no encontrado")
			}
		} else {
			var resultado = "Error tamaño de comando más pequeño que 2"
			c.WriteMessage(mt, []byte(resultado))
		}

		/*err = c.WriteMessage(mt, mensaje)
		if err != nil {
			log.Println("write:", err)
			break
		}*/
	}
	for nombreCanal, _ := range u.Canales {
		var mensajeEnviar string
		var canal = global.Canales[nombreCanal]
		canal.MutexCanal.Lock()
		for claveNick, _ := range global.Canales[nombreCanal].Usuarios {
			if u.Nick != claveNick {
				mensajeEnviar = fmt.Sprintf("USUARIO_SALE_CANAL %s %s", nombreCanal, u.Nick)
				global.Usuarios[claveNick].Conexion.WriteMessage(1, []byte(mensajeEnviar))
			}
		}
		canal.MutexCanal.Unlock()

	}
	for k, _ := range u.Canales {
		delete(global.Canales[k].Usuarios, u.Nick)
	}
	delete(global.Usuarios, u.Nick)

	fmt.Println("Desconexio del socket")
}
