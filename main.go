package main

import (
	"chiwita/controlador"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
)

var mutexUsers sync.Mutex

type user struct {
	name      string
	connected bool
	contador  int
}

// var users [string]websocket.Conn
var users = make(map[net.Conn]user)
var addr = flag.String("addr", "localhost:8080", "http service address")
var db *sql.DB
var upgrader = websocket.Upgrader{} // use default options

func gestorConexion(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	//que al final de la ejecución del websocket se ejecute el cierre del websocket
	defer c.Close()
	var resulAutentificacion, u = controlador.Autentificacion(c, db)

	if resulAutentificacion == false {
		fmt.Println("Error al autentificar")
		return
	}

	fmt.Println("Autentificacion correcta")

	//inicializar el usuario a la hash map
	//de usuarios global. Para garantizar que solo se accede una vez por hilo
	//hay que usar mutex
	users[c.NetConn()] = user{u.Nick, true, 0}
	fmt.Println(users)
	for {
		mt, message, err := c.ReadMessage()

		if entry, ok := users[c.NetConn()]; ok {

			mutexUsers.Lock()
			// Then we modify the copy
			entry.contador = entry.contador + 1
			users[c.NetConn()] = entry
			mutexUsers.Unlock()
		}

		if err != nil {
			log.Println("read:", err)
			println("%s", err.Error())
			break
		}
		fmt.Printf("%+v\n", users)

		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	fmt.Println("Desconexio del socket")
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/gestorConexion")
}

func main() {

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = "chiwita"
	cfg.Passwd = "chiwita"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "chiwita"

	// Obtener el manejador de conexión de la Base de Datos
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("¡Iniciando Servidor!")

	http.HandleFunc("/gestorConexion", gestorConexion)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script src = "https://cdnjs.cloudflare.com/ajax/libs/js-sha512/0.9.0/sha512.min.js"></script>
<script>  

window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var userClient = document.getElementById("userClient");
    var passwordClient = document.getElementById("passClient");
    var ws;

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
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
			/*//////////////////////////////////
			//   AUTENTIFICACION              //   
			//   USUARIO + SHA512(CONTRASENA) //
			//////////////////////////////////*/
			print("COMUNICACION ABIERTA")
            var resultsend =""+userClient.value+" "+sha512(passClient.value);
            ws.send(resultsend);
        }
        ws.onclose = function(evt) {
            print("COMUNICACION CERRADA");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
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
        print("SEND: " + input.value);
		comando = "MSG introduccion "+input.value;
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
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
User<input id="userClient" type="text" value="user1">
Pass<input id="passClient" type="text" value="pass">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
