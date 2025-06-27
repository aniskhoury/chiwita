package estructuras

import (
	"github.com/gorilla/websocket"
	"github.com/quartercastle/vector"
)

type Usuario struct {
	Idusuario      uint64
	Nick           string
	Contrasena     string
	Email          string
	Fecha_registro string
	Activado       bool
}
type AtributosUsuario struct {
	Conexion           *websocket.Conn
	Nivel              uint64
	AntiInunfacion     vector.Vector
	Canales            vector.Vector
	PrivilegiosCanales vector.Vector
}
