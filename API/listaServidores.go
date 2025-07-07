package api

import (
	"chiwita/estructuras"
	"chiwita/global"
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
	var resultado string

	resultado = ""
	for filas.Next() {
		if err := filas.Scan(&servidor.IdServidor, &servidor.NombreServidor, &servidor.IpServidor); err != nil {
			fmt.Println("Error al leer de la base de datos la lista de servidores")
		} else {
			resultado = `{"idservidor":"` + fmt.Sprintf("%d", servidor.IdServidor) + `,`
			resultado = resultado + `"nombreservidor":"` + string(servidor.NombreServidor+`",`)
			resultado = resultado + `"ipservidor":` + string(servidor.IpServidor) + `"}`
		}
	}

	mapa["resultado"] = resultado
	plantillaListaServidores.Execute(w, mapa)

}
