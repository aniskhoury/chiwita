package global

import (
	"chiwita/estructuras"
	"database/sql"
	"flag"
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

/****************************************************
 *   HashMap que mapea el Usuario hacia un socket   *
 ***************************************************/

var Usuarios = make(map[string]net.Conn)
var MutexUsuarios sync.Mutex

/****************************************************
 *   HashMap que mapea el Usuario hacia un socket   *
 ***************************************************/

var SocketUsuarios = make(map[net.Conn]estructuras.Usuario)
var MutexSocketUsuarios sync.Mutex

/****************************************************
 *   HashMap que mapea el string de un canal hacia  *
 *              la estructura del canal             *
 ***************************************************/

var Canales = make(map[string]estructuras.Canal)
var MutexCanales sync.Mutex

var Addr = flag.String("addr", "localhost:8080", "http service address")
var Db *sql.DB
var Upgrader = websocket.Upgrader{} // use default options
