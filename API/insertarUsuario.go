/*
	    Autor: Anis Khoury Ribas
		Date creation :14/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package api

import (
	"chiwita/global"
	"fmt"
	"html/template"
	"net/http"
)

func InsertarUsuario(w http.ResponseWriter, r *http.Request) {

	mapa := make(map[string]interface{})
	var plantillaListaServidores, err = template.ParseFiles("plantillas/insertarUsuario.template")
	if err != nil {
		fmt.Print("Error al leer la plantilla")
	}
	consulta := "INSERT INTO `usuarios` (`nick`, `contrasena`, `email` ) VALUES (?,?,?)"
	_, err = global.Db.Query(consulta, r.FormValue("nick"), r.FormValue("contrasena"), r.FormValue("email"))
	if err != nil {
		mapa["resultado"] = "Error usuario o email ya registrado"
	} else {
		mapa["resultado"] = "USUARIO_INSERTADO_OK"
	}

	plantillaListaServidores.Execute(w, mapa)
}
