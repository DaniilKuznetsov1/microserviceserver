package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"testserver2/tcpserver/netengine"
	"testserver2/tcpserver/restserver"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("TCP Rest Server Started")

	paramSt := netengine.TParamSt{}
	file, err := os.Open("env.json")
	if err != nil {
		paramSt.ConnHost = "localhost"
		paramSt.ConnPort = "8082"
		paramSt.ConnPortRest = "8083"
		paramSt.RestWebServAddr1 = "http://localhost:8082/setMon"
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&paramSt)
	if err != nil {
		fmt.Println("Decoder error: ", err.Error())
	}

	resstr := "Host: " + paramSt.ConnHost + ", Port: " + paramSt.ConnPort + ", PortREST: " + paramSt.ConnPortRest
	fmt.Println(resstr)

	//Каналы передачи сообщений
	chanWebToRest := make(chan netengine.TChannelMonMes, 0)
	chanRestToTCP := make(chan netengine.TChannelMes, 0)

	restserver.InitRestServer(paramSt.ConnHost, paramSt.ConnPortRest, paramSt.RestWebServAddr1, "true")
	fmt.Println("Start REST TCP Server")
	go restserver.StartRestServer(chanWebToRest, chanRestToTCP)
	//go tcpserver1.StartTCPServer(chanWebToRest, chanRestToTCP)

	for {
		time.Sleep(time.Second)
	}
}
