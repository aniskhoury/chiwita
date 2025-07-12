/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package api

import (
	"chiwita/estructuras"
	"chiwita/global"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func ListaServidores(w http.ResponseWriter, r *http.Request) {

	var plantillaListaServidores, err = template.ParseFiles("plantillas/listaServidores.template")
	if err != nil {
		fmt.Print("Error al leer la plantilla")
	}

	stmt, err := global.Db.Prepare("select * from servidores")
	filas, err := stmt.Query()
	var servidor estructuras.Servidor

	mapa := make(map[string]interface{})
	if err != nil {
		fmt.Println("Error leyendo lista servidores")
	}

	parseData := make([]map[string]interface{}, 0, 0)
	for filas.Next() {
		if err := filas.Scan(&servidor.IdServidor, &servidor.NombreServidor, &servidor.IpServidor); err != nil {
			fmt.Println("Error al leer de la base de datos la lista de servidores")
		} else {

			var mapaServidor = make(map[string]interface{})
			mapaServidor["idservidor"] = fmt.Sprintf("%d", servidor.IdServidor)
			mapaServidor["nombreservidor"] = servidor.NombreServidor
			mapaServidor["ipservidor"] = servidor.IpServidor
			parseData = append(parseData, mapaServidor)
		}
	}
	b, _ := json.Marshal(parseData)
	mapa["resultado"] = string(b)

	plantillaListaServidores.Execute(w, mapa)

}
