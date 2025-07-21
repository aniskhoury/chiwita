/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package structure

import (
	"sync"

	"github.com/gorilla/websocket"
)

type User struct {
	Iduser        uint64
	Nick          string
	Password      string
	Email         string
	Date_register string
	Activated     bool
	Bot           bool
	Connection    *websocket.Conn
	Channels      map[string]string
}

type Channel struct {
	Name         string
	Users        map[string]User
	CounterUsers uint64
	MutexChannel sync.Mutex
}
