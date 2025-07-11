package api

import (
	"chiwita/global"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func ListaUsuariosCanal(w http.ResponseWriter, r *http.Request) {
	mapa := make(map[string]interface{})
	var plantillaListaServidores, err = template.ParseFiles("plantillas/listaUsuariosCanal.template")
	if err != nil {
		fmt.Print("Error al leer la plantilla")
	}
	global.MutexCanales.Lock()
	canal, existe := global.Canales[r.FormValue("canal")]
	if existe {
		for k, _ := range canal.Usuarios {
			mapa[k] = k
		}
	} else {
		mapa["error"] = "Error no existe el canal"
	}
	global.MutexCanales.Unlock()

	b, _ := json.Marshal(mapa)
	mapa["resultado"] = string(b)

	plantillaListaServidores.Execute(w, mapa)
}
