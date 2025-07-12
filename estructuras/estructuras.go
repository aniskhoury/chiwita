/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
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
