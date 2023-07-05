package websockets

import (

	//"fmt"
	//"net/http"

	"fmt"
	"net/http"
	"testserver2/websserv/netengine"
	"testserver2/websserv/restserver"

	"github.com/gorilla/websocket"
)

var connHost = ""
var connPortWeb = ""

var debug = ""

type TClientConn struct {
	Connected bool
	UserId    int
}

func InitWS(cHost, cPortWeb, debug1 string) int {
	connHost = cHost
	connPortWeb = cPortWeb
	debug = debug1
	return 0
}

var upgrader = websocket.Upgrader{}

var clients = make(map[*websocket.Conn]TClientConn)

func disConnClose(conn *websocket.Conn, err error) {
	fmt.Println("Conn Close: " + conn.UnderlyingConn().RemoteAddr().String())
	conn.Close()
}

func handleMessageToWeb(chanRestToWeb chan netengine.TChannelMonMes) {
	for {
		msg := <-chanRestToWeb
		decMsg := netengine.ChannelMonMesToWebSocketMes(msg)
		for client := range clients {
			err := client.WriteJSON(decMsg)
			if err != nil {
				fmt.Println("error web cli ", err)
				client.Close()
				delete(clients, client)
			}
			fmt.Println("REST To WS: " + decMsg.Result + ", user: " + decMsg.User)
		}
	}
}

func WebSocketStart() {
	if debug == "true" {
		fmt.Println("Web Socket Start")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if debug == "true" {
			fmt.Println("handle func: " + r.URL.String())
		}
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error connection upgrade: ", err)
		}

		defer disConnClose(conn, err)
		clients[conn] = TClientConn{Connected: true, UserId: 0}

		defer delete(clients, conn)
		fmt.Println("Web client connected, event loop: " + conn.UnderlyingConn().RemoteAddr().String())

		for {
			var message = netengine.TWebSocketMes{}
			cliConn := clients[conn]
			userId := 0

			if debug == "true" {
				fmt.Println("read json")
			}

			err := conn.ReadJSON(&message) //read web command
			if debug == "true" {
				fmt.Println("Input WebMes cliCommand: " + message.Command)
			}
			if err != nil {
				fmt.Println("Error read webmes ", err)
				break
			}
			//Анализ команды
			chanMessage := netengine.WebSocketMesToChannelMes(message)
			if debug == "true" {
				fmt.Println("ChanMes cliCommand: " + chanMessage.Command)
			}
			//Отправляем команду
			//chanWebToRest <- chanMessage
			dataResp := restserver.GetReturnPostData(chanMessage)

			n, _ := message.Id.Int64()
			userId = int(n)
			if userId > 0 {
				cliConn.UserId = int(n)
				clients[conn] = cliConn
			}

			//resMes, err := json.Marshal(dataResp)
			//resMes := make(map[string]string)
			//resMes["return"] = "OK"
			//resMes["echomes"] = message.Command

			err = conn.WriteJSON(dataResp)
			if err != nil {
				fmt.Println("error write webmes ", err)
			}
		}
	})
	if debug == "true" {
		fmt.Println("start handle mess, _listen")
	}
	//go handleMessageToWeb(chanRestToWeb)
	err := http.ListenAndServe(connHost+":"+connPortWeb, nil)
	if err != nil {
		fmt.Println("Error start: ", err)
	}
}
