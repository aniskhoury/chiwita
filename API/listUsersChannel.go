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

func ListUsersChannel(w http.ResponseWriter, r *http.Request) {
	mapResult := make(map[string]interface{})
	var templateListUsersChannel, err = template.ParseFiles("templates/listUsersChannel.template")
	if err != nil {
		fmt.Print("Error al leer la plantilla")
	}
	global.MutexChannels.Lock()
	channel, exist := global.Channels[r.FormValue("channel")]
	if exist {
		for k, _ := range channel.Users {
			mapResult[k] = ""
		}
	} else {
		mapResult["error"] = "ERROR_CHANNEL_DOESNT_EXIST"
	}
	global.MutexChannels.Unlock()

	b, _ := json.Marshal(mapResult)
	mapResult["result"] = string(b)

	templateListUsersChannel.Execute(w, mapResult)
}
