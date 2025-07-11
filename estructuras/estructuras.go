package estructuras

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Usuario struct {
	Idusuario      uint64
	Nick           string
	Contrasena     string
	Email          string
	Fecha_registro string
	Activado       bool
	Bot            bool
	Conexion       *websocket.Conn
	Canales        map[string]string
}
type AtributosUsuario struct {
	Nivel              uint64
	AntiInunfacion     map[string]string
	Canales            map[string]Canal
	PrivilegiosCanales map[string]uint64
}

type Canal struct {
	Nombre           string
	Usuarios         map[string]Usuario
	ContadorUsuarios uint64
	MutexCanal       sync.Mutex
}

type Servidor struct {
	IdServidor     uint64
	NombreServidor string
	IpServidor     string
}

type JSONEstructura struct {
	Nombre string
	Valor  string
}
