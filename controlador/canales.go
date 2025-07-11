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
