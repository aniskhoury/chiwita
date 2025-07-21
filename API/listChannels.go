/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
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

func ListChannels(w http.ResponseWriter, r *http.Request) {

	var templateListChannels, err = template.ParseFiles("templates/listChannels.template")
	if err != nil {
		fmt.Print("Error al leer la plantilla")
	}

	mapResult := make(map[string]interface{})

	mapChannels := make([]map[string]interface{}, 0, 0)
	global.MutexChannels.Lock()
	for key, _ := range global.Channels {
		var channel = make(map[string]interface{})

		channel["channel"] = key
		//mapaCanal["contadorUsuarios"] = contador
		channel["counterUsers"] = global.Channels[key].CounterUsers

		mapChannels = append(mapChannels, channel)
	}
	global.MutexChannels.Unlock()
	b, _ := json.Marshal(mapChannels)
	mapResult["result"] = string(b)

	templateListChannels.Execute(w, mapResult)

}
