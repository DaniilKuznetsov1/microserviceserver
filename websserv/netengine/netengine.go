package netengine

import (
	"encoding/json"
	"strconv"
)

type TParamSt struct {
	ConnHost         string `json:"connHost,omitempty"`
	ConnPort         string `json:"connPort,omitempty"`
	ConnPortWeb      string `json:"connPortWeb,omitempty"`
	ConnPortRest     string `json:"connPortRest,omitempty"`
	RestTcpServAddr1 string `json:"restTcpServAddr1,omitempty"`
	Debug1           string `json:"debug1,omitempty"`
}

type TWebSocketMes struct {
	Id       json.Number `json:"id,omitempty"`
	User     string      `json:"user"`
	ServerNo json.Number `json:"serverNo"`
	Command  string      `json:"command"`
	Result   string      `json:"result,omitempty"`
}

type TChannelMes struct {
	ServerNo int
	User     string
	Command  string
}

type TChannelMonMes struct {
	Id     int
	User   string
	Result string
}

type TRestMes struct {
	Id           int
	ServerTypeId int
	User         string
	Command      string
}

func WebSocketMesToChannelMes(message TWebSocketMes) TChannelMes {
	var chanMes TChannelMes
	chanMes.User = message.User
	n, _ := message.ServerNo.Int64()
	chanMes.ServerNo = int(n)
	chanMes.Command = message.Command
	return chanMes
}

func ChannelMonMesToWebSocketMes(message TChannelMonMes) TWebSocketMes {
	var chanMes TWebSocketMes
	chanMes.Id = json.Number(strconv.Itoa(message.Id))
	chanMes.User = message.User
	chanMes.Result = message.Result
	return chanMes
}
