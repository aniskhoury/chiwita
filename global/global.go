package global

import (
	"chiwita/estructuras"
	"database/sql"
	"flag"
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

var Usuarios = make(map[net.Conn]estructuras.Usuario)
var MutexUsuarios sync.Mutex

var Addr = flag.String("addr", "localhost:8080", "http service address")
var Db *sql.DB
var Upgrader = websocket.Upgrader{} // use default options
