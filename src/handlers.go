package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ServerListAPI(rw http.ResponseWriter, req *http.Request) {
	servers := GetServers()

	out, err := json.Marshal(servers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Fprintln(rw, string(out))
}

func AddServerAPI(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var t struct {
		Status int
		Info   ServerInfo
	}
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(rw, "Invalid post data", 400)
		return
	}

	token := AddServer(t.Status, t.Info)

	fmt.Fprintln(rw, token)
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

	if IsServer(t.Token) {
		SetServer(t.Token, t.Status, t.Info)
		fmt.Fprintln(rw, "OK")
	} else {
		http.Error(rw, "Invalid token", 404)
		return
	}

}

func DeleteServerAPI(rw http.ResponseWriter, req *http.Request) {
	stuff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Invalid post data", 400)
		return
	}

	if DeleteServer(string(stuff)) {
		fmt.Fprintln(rw, "Goodbye!")
	} else {
		http.Error(rw, "Invalid token", 404)
	}
}
