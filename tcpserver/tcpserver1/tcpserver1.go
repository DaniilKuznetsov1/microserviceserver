package tcpserver1

import (
	"fmt"
	"math/rand"
	"strings"
	"testserver2/tcpserver/netengine"
	"time"
)

//var CharArray string = "abcdefghijklmnopqrstuvwxyz"
var CharArrayR []rune = []rune("abcdefghijklmnopqrstuvwxyz")

//var CharArray2 string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var CharArrayR2 []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func InitTCPServ() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("TCP Server Init")
}

func getRandomStr(lenstr int) string {
	var b strings.Builder
	for i := 0; i < lenstr; i++ {
		b.WriteRune(CharArrayR[rand.Intn(26)])
	}
	return b.String()
}

func getRandomStr2(lenstr int) string {
	var b strings.Builder
	for i := 0; i < lenstr; i++ {
		b.WriteRune(CharArrayR2[rand.Intn(26)])
	}
	return b.String()
}

func StartTCPServer(chanTCPToRest chan netengine.TChannelMonMes, chanRestToTCP chan netengine.TChannelMes) {
	//Обработка данных, пришедших по TChannelMes каналу
	for {
		msg := <-chanRestToTCP
		Rescomm := ""
		fmt.Println("Command: " + msg.Command)
		var monMes netengine.TChannelMonMes
		switch msg.Command {
		case "comm1":
			Rescomm = getRandomStr(10)
		case "comm2":
			Rescomm = getRandomStr2(15)
		case "comm3":
			Rescomm = "___"
		default:
			Rescomm = getRandomStr(10)
		}
		fmt.Println("Result: " + Rescomm)
		monMes.Id = msg.ServerNo
		monMes.User = msg.User
		monMes.Result = Rescomm
		chanTCPToRest <- monMes
	}
}

func SetDataTCP(RestToTcp netengine.TChannelMes) netengine.TChannelMonMes {
	msg := RestToTcp
	Rescomm := ""
	fmt.Println("Command: " + msg.Command)
	var monMes netengine.TChannelMonMes
	switch msg.Command {
	case "comm1":
		Rescomm = getRandomStr(10)
	case "comm2":
		Rescomm = getRandomStr2(15)
	case "comm3":
		Rescomm = "___"
	default:
		Rescomm = getRandomStr(10)
	}
	fmt.Println("Result: " + Rescomm)
	monMes.Id = msg.ServerNo
	monMes.User = msg.User
	monMes.Result = Rescomm
	return monMes
}
