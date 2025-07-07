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
	//plantillaListaServidores.Execute(w, mapa)

	//plantillaListaServidores.Execute(w, jsonArrayServidores)
	plantillaListaServidores.Execute(w, mapa)

}
