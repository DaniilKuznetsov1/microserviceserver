package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"testserver2/websserv/netengine"
	"testserver2/websserv/restserver"
	"testserver2/websserv/websockets"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("Server Started")

	paramSt := netengine.TParamSt{}
	file, err := os.Open("env.json")
	if err != nil {
		paramSt.ConnHost = "localhost"
		paramSt.ConnPort = "8080"
		paramSt.ConnPortWeb = "8081"
		paramSt.ConnPortRest = "8082"
		paramSt.RestTcpServAddr1 = "http://localhost:8083/setMon"
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&paramSt)
	if err != nil {
		fmt.Println("Decoder error: ", err.Error())
	}

	resstr := "Host: " + paramSt.ConnHost + ", Port: " + paramSt.ConnPort + ", PortWebSock: " + paramSt.ConnPortWeb + ", PortREST: " + paramSt.ConnPortRest
	fmt.Println(resstr)

	//Каналы передачи сообщений
	//chanWebToRest := make(chan netengine.TChannelMes, 0)
	//chanRestToWeb := make(chan netengine.TChannelMonMes, 0)

	//инициализация модулей
	websockets.InitWS(paramSt.ConnHost, paramSt.ConnPortWeb, "true")
	restserver.InitRestServer(paramSt.ConnHost, paramSt.ConnPortRest, paramSt.RestTcpServAddr1, "true")

	go websockets.WebSocketStart()
	//go restserver.StartRestServer(chanWebToRest, chanRestToWeb)
	for {
		time.Sleep(time.Second)
	}

}
