/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package global

import (
	"chiwita/structure"
	"database/sql"
	"flag"
	"sync"

	"github.com/gorilla/websocket"
)

/****************************************************
 *   HashMap que mapea el Usuario hacia un socket   *
 ***************************************************/

// var Usuarios = make(map[string]net.Conn)
var Users = make(map[string]structure.User)
var MutexUsers sync.Mutex

/****************************************************
 *   HashMap que mapea el Usuario hacia un socket   *
 ***************************************************/

/*var SocketUsuarios = make(map[net.Conn]estructuras.Usuario)
var MutexSocketUsuarios sync.Mutex
*/
/****************************************************
 *   HashMap que mapea el string de un canal hacia  *
 *              la estructura del canal             *
 ***************************************************/

var Channels = make(map[string]structure.Channel)
var MutexChannels sync.Mutex

var CounterUsers = 0
var MutexCounterUsers sync.Mutex

var Addr = flag.String("addr", "localhost:8080", "http service address")
var Db *sql.DB
var Upgrader = websocket.Upgrader{} // use default options
