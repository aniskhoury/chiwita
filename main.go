/*
	    Autor: Anis Khoury Ribas
		Date creation :12/07/2025
	    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
	    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	api "chiwita/API"
	"chiwita/controlador"
	"chiwita/global"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.NewConfig()
	cfg.User = "chiwita"
	cfg.Passwd = "chiwita"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "chiwita"

	// Obtener el manejador de conexión de la Base de Datos
	var err error
	/*
		Inicialización de los vectores de canales
	*/
	//global.Canales["principal"].Usuarios = new(estructuras.Usuario)
	global.Db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := global.Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("¡Iniciando Servidor!")
	//Cargar canales al servidor
	_, err = controlador.CargarCanales()
	if err != nil {
		log.Printf("Error al cargar los canales")

	}
	http.HandleFunc("/gestorConexion", controlador.GestorConexion)
	http.HandleFunc("/", home)
	http.HandleFunc("/listaCanales", api.ListaCanales)
	http.HandleFunc("/listaServidores", api.ListaServidores)
	http.HandleFunc("/listaUsuariosCanal", api.ListaUsuariosCanal)
	http.HandleFunc("/insertarUsuario", api.InsertarUsuario)
	http.HandleFunc("/informacionServidor", api.InformacionServidor)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(*global.Addr, nil))
}
func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/gestorConexion")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script src = "https://cdnjs.cloudflare.com/ajax/libs/js-sha512/0.9.0/sha512.min.js"></script>

<link rel="stylesheet" href="estatico/css/css.css">
 <script type="text/javascript">
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var userClient = document.getElementById("userClient");
    var passwordClient = document.getElementById("passClient");
	var seleccionServidor = document.getElementById("seleccionServidor");
    var ws;
	var estado = 0;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
	




    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
		//hace la conexion con el servidor elegido
		switch(seleccionServidor.value){
			case "Neptuno":
				ws = new WebSocket("ws://localhost:8080/gestorConexion");
				break;
			default:
				ws = new WebSocket("ws://localhost:8080/gestorConexion");
				break;
		}

        ws.onopen = function(evt) {
			/*//////////////////////////////////
			//   AUTENTIFICACION              //   
			//   USUARIO + SHA512(CONTRASENA) //
			//////////////////////////////////*/
			print("COMUNICACION ABIERTA")
            var resultsend =""+usuario.value+" "+sha512(contrasena.value);
            ws.send(resultsend);
        }
        ws.onclose = function(evt) {
            print("COMUNICACION CERRADA");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
			if (evt.data == "AUTENTIFICACION_CORRECTA"){

				mostrarCanales.innerHTML = "CARGAR LISTA CANALES";
			}
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        var comando = "";
		if (!ws) {
            return false;
        }
        print("SEND: " + texto.value);
		comando = texto.value;
        ws.send(comando);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>
Para autentificarte y conectarte a la comunidad, dale a Acceder
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="texto" type="text" value="Texto">
User<input id="usuario" type="text" value="usuario">
Pass<input id="contrasena" type="text" value="contrasena">

<select id="seleccionServidor">
  <option value="Neptuno">Neptuno</option>
</select>
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
<div id="mostrarCanales">
</div>

</body>

</html>
`))
