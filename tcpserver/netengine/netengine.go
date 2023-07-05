package netengine

type TParamSt struct {
	ConnHost         string `json:"connHost,omitempty"`
	ConnPort         string `json:"connPort,omitempty"`
	ConnPortWeb      string `json:"connPortWeb,omitempty"`
	ConnPortRest     string `json:"connPortRest,omitempty"`
	RestWebServAddr1 string `json:"restWebServAddr1,omitempty"`
	Debug1           string `json:"debug1,omitempty"`
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
