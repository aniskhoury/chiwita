package controlador

import (
	"chiwita/estructuras"
	"database/sql"
	"fmt"
	"log"
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
