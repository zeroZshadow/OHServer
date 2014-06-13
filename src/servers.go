package main

import (
	"errors"
	"time"
)

var Servers map[string]*ServerItem

const ExpirationTime = time.Second * 30 // Server Timeout
const TokenLength = 10                  // Token length in characters

type ServerItem struct {
	Status     int
	Expiration *time.Timer
	Info       ServerInfo
}

type ServerInfo struct {
	GUID           string
	Map            string
	CurrentPlayers int
	MaxPlayers     int
	Version        string
}

// We might change how we manage servers, so have some wrappers

// Creates the Server map
func InitServers() {
	Servers = make(map[string]*ServerItem)
}

// Adds a server to the map
func AddServer(s ServerInfo) string {
	token := RandStr(10)

	server := new(ServerItem)
	server.Status = 1
	server.Info = s
	server.Expiration = time.NewTimer(ExpirationTime)

	go func() {
		<-server.Expiration.C
		DeleteServer(token)
	}()

	Servers[token] = server

	return token
}

// Updates a server in the map
func SetServer(token string, status int, s ServerInfo) error {
	if !IsServer(token) {
		return errors.New("Inexistant server")
	}

	Servers[token].Status = status
	Servers[token].Info = s
	Servers[token].Expiration.Reset(ExpirationTime)

	return nil
}

// Checks if a server is in the map. Returns True if it exists, False otherwise
func IsServer(token string) bool {
	_, ok := Servers[token]
	return ok
}

// Gets a list of all the servers in the map. It also filters on Status == 0
func GetServers(getAll bool) []ServerInfo {
	list := make([]ServerInfo, 0)

	i := 0
	for k := range Servers {
		if getAll || Servers[k].Status == 0 {
			list = append(list, Servers[k].Info)
			i++
		}
	}

	return list
}

// Deletes a server in the map, given the right token.
func DeleteServer(token string) bool {
	if IsServer(token) {
		delete(Servers, token)
		return true
	}
	return false
}
