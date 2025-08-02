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
	"chiwita/controller"
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
	fmt.Println("¡Starting server!")
	//Load channels
	_, err = controller.LoadChannels()
	if err != nil {
		log.Printf("Error loading channels")
	}
	http.HandleFunc("/handlerConnection", controller.HandlerConnection)
	http.HandleFunc("/", home)
	http.HandleFunc("/listChannels", api.ListChannels)
	http.HandleFunc("/listUsersChannel", api.ListUsersChannel)
	http.HandleFunc("/insertUser", api.InsertUser)
	http.HandleFunc("/informationServer", api.InformationServer)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(*global.Addr, nil))
}
func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/handlerConnection")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<style>
body{
  height: 100%;
  width: 100%;
  margin: 0;
}
#channels {
	width:100%;
	height: 100%;
	background-color:brown;
}
#showChannels{
	background-color:yellow;
	height: 100%;
	width: 100%;
	bottom: 10;
	right: 10;
	position:fixed;
}

/* Style the tab */
.tab {
  overflow: hidden;
  border: 1px solid #ccc;
  background-color: #f1f1f1;
}

/* Style the buttons inside the tab */
.tab button {
  background-color: inherit;
  float: left;
  border: none;
  outline: none;
  cursor: pointer;
  padding: 14px 16px;
  transition: 0.3s;
  font-size: 17px;
}

/* Change background color of buttons on hover */
.tab button:hover {
  background-color: #ddd;
}

/* Create an active/current tablink class */
.tab button.active {
  background-color: #ccc;
}

/* Style the tab content */
.tabcontent {
  display: none;
  padding: 6px 12px;
  border: 1px solid #ccc;
  border-top: none;
}

/* Style the close button */
.topright {
  float: right;
  cursor: pointer;
  font-size: 28px;
}

.topright:hover {
	color: red;
}
</style>
<meta charset="utf-8">
<script src = "https://cdnjs.cloudflare.com/ajax/libs/js-sha512/0.9.0/sha512.min.js"></script>

<link rel="stylesheet" href="static/css/css.css">
 <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.7.1/jquery.min.js"></script>
<script src="/static/js/channels.js"></script>
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
				ws = new WebSocket("ws://localhost:8080/handlerConnection");
				break;
			default:
				ws = new WebSocket("ws://localhost:8080/handlerConnection");
				break;
		}

        ws.onopen = function(evt) {
			/*//////////////////////////////////
			//   AUTENTIFICACION              //   
			//   USUARIO + SHA512(CONTRASENA) //
			//////////////////////////////////*/
			print("OPEN_COMMUNICATION")
            var resultsend =""+usuario.value+" "+sha512(contrasena.value);
            ws.send(resultsend);
        }
        ws.onclose = function(evt) {
            print("CLOSED COMMUNICATION");         
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
			if (evt.data == "SUCCESSFUL_AUTHENTIFICATION"){


				var showChannels = document.createElement("div");
				showChannels.setAttribute("id","showChannels");
				showChannels.setAttribute("class","showChannels");
				


				document.getElementById("content").appendChild(showChannels);
				var tabs = document.createElement("div");
				tabs.setAttribute("id","tab");
				tabs.setAttribute("class","tab");
				tabs.innerHTML = tabs.innerHTML + '<button class="tablinks" onclick="showListChannels(event)" id="defaultOpen">List of Channels</button>'

				tabs.innerHTML = tabs.innerHTML + '<button class="tablinks" onclick="openChannel(event,\'Filosofia\')">London</button>'
				document.getElementById("showChannels").appendChild(tabs);

				var contentTab = document.createElement("div");
				contentTab.setAttribute("id","contentTab");
				contentTab.setAttribute("class","tabcontent");
				document.getElementById("showChannels").appendChild(contentTab);

				contentTab.innerHTML = "<h3>List Channels</h3>";

				//document.getElementById("showChannels").appendChild(contentListChannels);


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
	</td></tr>
</table>
<div id="content">
</div>

</body>

</html>
`))
