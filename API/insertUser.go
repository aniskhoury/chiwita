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

func InsertUser(w http.ResponseWriter, r *http.Request) {

	mapResult := make(map[string]interface{})
	var templateInsertUser, err = template.ParseFiles("templates/insertUser.template")
	if err != nil {
		fmt.Print("Error reading insertUser.template")
	}
	consulta := "INSERT INTO `usuarios` (`nick`, `password`, `email` ) VALUES (?,?,?)"
	_, err = global.Db.Query(consulta, r.FormValue("nick"), r.FormValue("password"), r.FormValue("email"))
	if err != nil {
		mapResult["result"] = "ERROR_NICK_EMAIL_IN_USE"
	} else {
		mapResult["result"] = "USUARIO_INSERTADO_OK"
	}

	templateInsertUser.Execute(w, mapResult)
}
