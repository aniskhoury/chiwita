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
			if err := filas.Scan(&u.Idusuario, &u.Nick, &u.Contrasena, &u.Email, &u.Fecha_registro, &u.Activado); err != nil {
				fmt.Println("Error autentificacion")
				return false, u
			}
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
		fmt.Println("Error al autentificar")
		return
	}

	fmt.Println("Autentificacion correcta")

	//inicializar el usuario a la hash map
	//de usuarios global. Para garantizar que solo se accede una vez por hilo
	//hay que usar mutex
	global.MutexUsuarios.Lock()
	global.Usuarios[u.Nick] = c.NetConn()
	global.MutexUsuarios.Unlock()
	global.MutexSocketUsuarios.Lock()
	global.SocketUsuarios[c.NetConn()] = u
	global.MutexSocketUsuarios.Unlock()
	var msgAutentificacion = []byte("AUTENTIFICACION_CORRECTA")
	c.WriteMessage(1, msgAutentificacion)
	for {
		mt, message, err := c.ReadMessage()
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
		fmt.Printf("%+v\n", global.Usuarios)

		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	fmt.Println("Desconexio del socket")
}
