package restserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testserver2/websserv/netengine"
	"time"
)

var ConnHost = ""
var ConnPortRest = ""
var RestTCPServAddr1 = ""
var debug = ""

func InitRestServer(cHost, cPort, RestTCPSAddr1, debug1 string) {
	ConnHost = cHost
	ConnPortRest = cPort
	RestTCPServAddr1 = RestTCPSAddr1
	debug = debug1
}

func GetReturnPostData(WebToRest netengine.TChannelMes) netengine.TChannelMonMes {
	var postData netengine.TChannelMonMes
	message1 := WebToRest
	jsonData, err := json.Marshal(message1)
	if err != nil {
		fmt.Println("Error build request: ", err)
	}
	if debug == "true" {
		fmt.Println("chan message comm: " + message1.Command)
	}
	RestTCPSA := ""
	switch message1.ServerNo {
	case 1:
		RestTCPSA = RestTCPServAddr1
	default:
		RestTCPSA = RestTCPServAddr1
	}

	//POST
	if debug == "true" {
		fmt.Println("REST ServerAddr: " + RestTCPSA)
	}

	request1, err := http.NewRequest("POST", RestTCPSA, bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println("Error prep request: ", err)
	}

	//Header Type data
	request1.Header.Set("Content-Type", "application/json")
	Client := http.Client{Timeout: 3 * time.Second}
	resp, err := Client.Do(request1)
	if err != nil {
		fmt.Println("Send request REST Error: ", err)
		MonMes := netengine.TChannelMonMes{Id: 1, User: message1.User, Result: "Send request REST Error:  " + err.Error()}
		return MonMes
	}
	fmt.Println("Status Code: ", resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&postData)
	if debug == "true" {
		fmt.Println("postData: " + postData.User + ", " + postData.Result)
	}

	//respBody1, err := io.ReadAll(resp.Body)
	//fmt.Println("res body: ", string(respBody1))
	//resp.Body.Close()

	return postData
}

/* func PostTCPServ(chanWebToRest chan netengine.TChannelMes, chanRestToWeb chan netengine.TChannelMonMes) { //пересылка из канала на TCP Server
	var postData netengine.TChannelMonMes
	for {
		message1 := <-chanWebToRest
		jsonData, err := json.Marshal(message1)
		if err != nil {
			fmt.Println("Error build request: ", err)
		}
		if debug == "true" {
			fmt.Println("chan message comm: " + message1.Command)
		}

		RestTCPSA := ""
		switch message1.ServerNo {
		case 1:
			RestTCPSA = RestTCPServAddr1
		default:
			RestTCPSA = RestTCPServAddr1
		}

		//POST
		if debug == "true" {
			fmt.Println("REST ServerAddr: " + RestTCPSA)
		}

		request1, err := http.NewRequest("POST", RestTCPSA, bytes.NewReader(jsonData))
		if err != nil {
			fmt.Println("Error prep request: ", err)
			continue
		}

		//Header Type data
		request1.Header.Set("Content-Type", "application/json")
		Client := http.Client{Timeout: 3 * time.Second}
		resp, err := Client.Do(request1)
		if err != nil {
			fmt.Println("Send request REST Error: ", err)
			MonMes := netengine.TChannelMonMes{Id: 1, User: message1.User, Result: "Send request REST Error:  " + err.Error()}
			chanRestToWeb <- MonMes
			continue
		}
		fmt.Println("Status Code: ", resp.StatusCode)
		err = json.NewDecoder(resp.Body).Decode(postData)
		chanRestToWeb <- postData

		respBody1, err := io.ReadAll(resp.Body)
		fmt.Println("res body: ", string(respBody1))
		resp.Body.Close()

	}
}*/

/* func StartRestServer(chanWebToRest chan netengine.TChannelMes, chanRestToWeb chan netengine.TChannelMonMes) {
	var getIndex = func(w http.ResponseWriter, r *http.Request) {
		message1 := netengine.TChannelMonMes{Id: 1, User: "u1", Result: "OK"}
		json.NewEncoder(w).Encode(message1)
	}
	var setMonintor = func(w http.ResponseWriter, r *http.Request) {
		var Data netengine.TChannelMonMes

		if r.Method == "POST" {
			body, _ := io.ReadAll(r.Body)
			err := json.Unmarshal(body, &Data)

			if err != nil {
				fmt.Println("Error monitor resp: ", err)
			}
			defer r.Body.Close()

		}
	}

	if debug == "true" {
		fmt.Println("REST Server Started")
	}

	http.HandleFunc("/index", getIndex)
	http.HandleFunc("/setMon", setMonintor)

	go PostTCPServ(chanWebToRest, chanRestToWeb)

	http.ListenAndServe(ConnHost+":"+ConnPortRest, nil)
} */
