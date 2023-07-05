package restserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testserver2/tcpserver/netengine"
	"testserver2/tcpserver/tcpserver1"
	"time"
)

var ConnHost = ""
var ConnPortRest = ""
var RestWebServAddr1 = ""
var debug = ""

func InitRestServer(cHost, cPort, RestWebSAddr1, debug1 string) {
	ConnHost = cHost
	ConnPortRest = cPort
	RestWebServAddr1 = RestWebSAddr1
	debug = debug1
}

func PostWebServ(chanTcpToRest chan netengine.TChannelMonMes) { //пересылка из канала на Web Server
	for {
		message1 := <-chanTcpToRest
		jsonData, err := json.Marshal(message1)
		if err != nil {
			fmt.Println("Error build request: ", err)
		}
		RestWebSA := ""
		switch message1.Id {
		case 1:
			RestWebSA = RestWebServAddr1
		default:
			RestWebSA = RestWebServAddr1
		}

		//POST data
		request1, err := http.NewRequest("POST", RestWebSA, bytes.NewReader(jsonData))
		if err != nil {
			fmt.Println("Error send request: ", err)
			continue
		}

		//Header Type data
		request1.Header.Set("Content-Type", "application/json")
		Client := http.Client{Timeout: 3 * time.Second}
		resp, err := Client.Do(request1)
		if err != nil {
			fmt.Println("Send request REST Error: ", err)
			continue
		}
		fmt.Println("Status Code: ", resp.StatusCode)
		respBody1, err := io.ReadAll(resp.Body)
		fmt.Println("res body: ", string(respBody1))
		resp.Body.Close()
	}
}

func StartRestServer(chanWebToRest chan netengine.TChannelMonMes, chanRestToTCP chan netengine.TChannelMes) {
	var getIndex = func(w http.ResponseWriter, r *http.Request) {
		message1 := netengine.TChannelMonMes{Id: 1, User: "u1", Result: "OK"}
		json.NewEncoder(w).Encode(message1)
	}
	var setMonintor = func(w http.ResponseWriter, r *http.Request) { //Приём от WebSocket сервера сообщений и передача их в канал TChannelMes
		var Data netengine.TChannelMes

		if r.Method == "POST" {
			fmt.Println("SetMonitor ChannelMes")
			body, _ := io.ReadAll(r.Body)
			err := json.Unmarshal(body, &Data)

			if err != nil {
				fmt.Println("Error monitor resp: ", err)
			}
			DataRes := tcpserver1.SetDataTCP(Data)
			json.NewEncoder(w).Encode(DataRes)
			//chanRestToTCP <- Data

			defer r.Body.Close()
		}
	}

	http.HandleFunc("/index", getIndex)
	http.HandleFunc("/setMon", setMonintor)

	//go PostWebServ(chanWebToRest)

	http.ListenAndServe(ConnHost+":"+ConnPortRest, nil)
}
