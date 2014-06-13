package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Get command-line flags
	listen := flag.String("bind", "127.0.0.1:9902", "Address:Port or Socket where to listen to")
	flag.Parse()

	InitServers()

	r := mux.NewRouter()
	// GET  - Read stuff
	g := r.Methods("GET").Subrouter()
	g.HandleFunc("/", ServerListAPI)
	// POST - Actions
	p := r.Methods("POST").Subrouter()
	p.HandleFunc("/add", AddServerAPI)
	p.HandleFunc("/update", UpdateServerAPI)
	p.HandleFunc("/delete", DeleteServerAPI)

	http.Handle("/", r)

	fmt.Println("Listening on " + *listen)
	http.ListenAndServe(*listen, nil)
}
