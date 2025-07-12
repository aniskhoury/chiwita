/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package controlador

import (
	"chiwita/estructuras"
	"chiwita/global"
	"errors"
)

func CargarCanales() (bool, error) {

	stmt, err := global.Db.Prepare("select nombrecanal from canales ")
	filas, err := stmt.Query()
	if err != nil {
		return false, errors.New("No se pudo cargar los canales desde la Base de Datos")
	}

	for filas.Next() {
		var canal estructuras.Canal
		canal.ContadorUsuarios = 0
		if err := filas.Scan(&canal.Nombre); err != nil {
			return false, errors.New("Error al leer las filas de la Base de Datos al cargar los canales")
		}
		/*Cargar a la HashMap de Canales el canal*/
		global.MutexCanales.Lock()
		canal.Usuarios = make(map[string]estructuras.Usuario)
		global.Canales[canal.Nombre] = canal

		global.MutexCanales.Unlock()
	}
	return true, nil
}
