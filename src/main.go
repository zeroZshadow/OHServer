package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var debug *bool

func main() {
	// Get command-line flags
	listen := flag.String("bind", "127.0.0.1:9902", "Address:Port or Socket where to listen to")
	debug = flag.Bool("debug", false, "Enable debug messages")
	flag.Parse()

	InitServers()

	r := mux.NewRouter()
	// GET  - Read stuff
	g := r.Methods("GET").Subrouter()
	g.HandleFunc("/", FullServerListAPI)
	g.HandleFunc("/{version}", ServerListAPI)
	// POST - Actions
	p := r.Methods("POST").Subrouter()
	p.HandleFunc("/add", AddServerAPI)
	p.HandleFunc("/update", UpdateServerAPI)
	p.HandleFunc("/delete", DeleteServerAPI)

	http.Handle("/", r)

	fmt.Println("Listening on " + *listen)
	http.ListenAndServe(*listen, nil)
}
