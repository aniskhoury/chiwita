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
	AntiInunfacion     map[string]string
	Canales            map[string]Canal
	PrivilegiosCanales map[string]uint64
}

type Canal struct {
	Nombre           string
	Usuarios         vector.Vector
	ContadorUsuarios uint64
}

type Canales struct {
	ListaCanales vector.Vector
}
