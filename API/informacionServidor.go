/*
	    Autor: Anis Khoury Ribas
		Date creation :15/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package api

import (
	"chiwita/global"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func InformacionServidor(w http.ResponseWriter, r *http.Request) {

	var plantillaListaServidores, err = template.ParseFiles("plantillas/informacionServidor.template")
	if err != nil {
		fmt.Print("Error al leer la plantilla")
	}

	mapa := make(map[string]interface{})

	mapaServidor := make(map[string]interface{})
	global.MutexContadorUsuarios.Lock()
	fmt.Println(global.ContadorUsuarios)
	mapaServidor["contadorUsuarios"] = global.ContadorUsuarios
	global.MutexContadorUsuarios.Unlock()

	b, _ := json.Marshal(mapaServidor)
	mapa["resultado"] = string(b)

	plantillaListaServidores.Execute(w, mapa)

}
