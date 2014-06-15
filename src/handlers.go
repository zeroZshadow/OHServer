package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ServerListAPI(rw http.ResponseWriter, req *http.Request) {

	serverlist, count := GetServers(false)

	servers := struct {
		Message     string
		Servers     []ServerInfo
		Connections int
		Activegames int
	}{
		"Last version: 15062014", serverlist, count, len(serverlist),
	}

	out, err := json.Marshal(servers)
	if err != nil {
		if *debug {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Fprintf(rw, string(out))
}

func FullSListAPI(rw http.ResponseWriter, req *http.Request) {
	serverlist, _ := GetServers(true)

	out, err := json.Marshal(serverlist)
	if err != nil {
		if *debug {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Fprintf(rw, string(out))
}

func AddServerAPI(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var t ServerInfo
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(rw, "Invalid post data", 400)
		if *debug {
			fmt.Println(err.Error())
		}
		return
	}

	token := AddServer(t)
	fmt.Fprintf(rw, token)
}

func UpdateServerAPI(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var t struct {
		Token  string
		Status int
		Info   ServerInfo
	}
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(rw, "Invalid post data", 400)
		return
	}

	err = SetServer(t.Token, t.Status, t.Info)
	if err != nil {
		http.Error(rw, "Invalid token", 404)
		if *debug {
			fmt.Println("[" + t.Token + "] " + err.Error())
		}
		return
	} else {
		fmt.Fprintf(rw, "OK")
	}
}

func DeleteServerAPI(rw http.ResponseWriter, req *http.Request) {
	stuff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Invalid post data", 400)
		return
	}

	if DeleteServer(string(stuff)) {
		fmt.Fprintf(rw, "Goodbye!")
	} else {
		http.Error(rw, "Invalid token", 404)
		if *debug {
			fmt.Println("[" + string(stuff) + "] Tried to delete inexistant server")
		}
	}
}
