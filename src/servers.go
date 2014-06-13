package main

import (
	"time"
)

var Servers map[string]ServerItem

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
	Players        PlayerInfo
}

type PlayerInfo struct {
	Nickname string
}

// We might change how we manage servers, so have some wrappers

func InitServers() {
	Servers = make(map[string]ServerItem)
}

func AddServer(status int, s ServerInfo) string {
	token := RandStr(10)
	SetServer(token, status, s)
	return token
}

func SetServer(token string, status int, s ServerInfo) {
	var timer *time.Timer
	if _, ok := Servers[token]; ok {
		timer = Servers[token].Expiration
		timer.Reset(ExpirationTime)
	} else {
		timer = time.NewTimer(ExpirationTime)
		go func() {
			<-timer.C
			DeleteServer(token)
		}()
	}

	Servers[token] = ServerItem{
		Status:     status,
		Info:       s,
		Expiration: timer,
	}
}

func IsServer(token string) bool {
	_, ok := Servers[token]
	return ok
}

func GetServers() []ServerInfo {
	list := make([]ServerInfo, 0)

	i := 0
	for k := range Servers {
		if Servers[k].Status == 0 {
			list = append(list, Servers[k].Info)
			i++
		}
	}

	return list
}

func DeleteServer(token string) bool {
	if _, ok := Servers[token]; ok {
		delete(Servers, token)
		return true
	}
	return false
}
