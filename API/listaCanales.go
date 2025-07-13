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

func ListaCanales(w http.ResponseWriter, r *http.Request) {

	var plantillaListaServidores, err = template.ParseFiles("plantillas/listaCanales.template")
	if err != nil {
		fmt.Print("Error al leer la plantilla")
	}

	mapa := make(map[string]interface{})

	mapaCanales := make([]map[string]interface{}, 0, 0)
	global.MutexCanales.Lock()
	for clave, _ := range global.Canales {
		var mapaCanal = make(map[string]interface{})
		/*
			var contador = 0
			for _, _ = range global.Canales[clave].Usuarios {
				contador = contador + 1
			}*/

		mapaCanal["canal"] = clave
		//mapaCanal["contadorUsuarios"] = contador
		mapaCanal["contadorUsuarios"] = global.Canales[clave].ContadorUsuarios

		mapaCanales = append(mapaCanales, mapaCanal)
	}
	global.MutexCanales.Unlock()
	b, _ := json.Marshal(mapaCanales)
	mapa["resultado"] = string(b)

	plantillaListaServidores.Execute(w, mapa)

}
