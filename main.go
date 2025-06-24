package main

import (
	"chiwita/estructuras"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
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
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	var u estructuras.Usuario
	//fmt.Printf("%+v\n", users)
	//fmt.Println(reflect.TypeOf(c))
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	//Try login with user & password
	var userClient string
	var passClient string
	//Try login with user & password
	mt, getUserPassword, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		println("%s", err.Error())
		return
	} else {
		var result = strings.Split(string(getUserPassword), " ")
		userClient = result[0]
		passClient = result[1]
		fmt.Println(userClient, passClient)
		//if there is no username or password, then stop
		if userClient == "" || passClient == "" {
			return
		} else {
			stmt, err := db.Prepare("select * from usuarios where nick = ? and contrasena = ?")
			filas, err := stmt.Query(userClient, passClient)
			if err != nil {
				return
			}
			var autentificado bool

			autentificado = false
			for filas.Next() {
				if err := filas.Scan(&u.Idusuario, &u.Nick, &u.Contrasena, &u.Email, &u.Fecha_registro, &u.Activado); err != nil {
					fmt.Println("Error autentificacion")
				}
				autentificado = true
			}
			if autentificado == false {
				return
			}
		}
	}
	users[c.NetConn()] = user{userClient, true, 0}

	for {
		mt, message, err := c.ReadMessage()

		if entry, ok := users[c.NetConn()]; ok {

			mutexUsers.Lock()
			// Then we modify the copy
			entry.contador = entry.contador + 1
			users[c.NetConn()] = entry
			//users[c.NetConn()].contador = users[c.NetConn()].contador + 1
			// Then we reassign map entry
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
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = "chiwita"
	cfg.Passwd = "chiwita"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "chiwita"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
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
            print("OPEN");
            var resultsend =""+userClient.value+" "+passClient.value;
            ws.send(resultsend);
        }
        ws.onclose = function(evt) {
            print("CLOSE");
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
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
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
