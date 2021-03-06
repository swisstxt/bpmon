//go:generate esc -o static.go -pkg main -prefix static static

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

var (
	listener  string
	envDir    string
	bpDir     string
	staticDir string
	hub       *Hub
	envs      *Environments
)

func init() {
	flag.StringVar(&listener, "listener", "0.0.0.0:8765", "ip/port to listen on")
	flag.StringVar(&envDir, "env", "", "environment setup files")
	flag.StringVar(&bpDir, "bp", "", "bpmon bp files")
	flag.StringVar(&staticDir, "static", "", "static html served at http root")
}

func main() {
	flag.Parse()
	var err error

	envs, err = LoadEnvs(envDir, "*.yaml")
	if err != nil {
		log.Fatal(err)
	}

	(*envs)["_"], err = LoadEnvFromBP(bpDir, "*.yaml")
	if err != nil {
		log.Fatal(err)
	}

	hub = newHub()
	go hub.run()

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/icinga/{env}/v1/objects/services", MockIcingaServicesHandler).Methods("GET")
	r.HandleFunc("/icinga/{env}/v1/actions/acknowledge-problem", MockIcingaAcknowledgeHandler).Methods("POST")
	r.HandleFunc("/api/envs/", ListEnvsHandler).Methods("GET")
	r.HandleFunc("/api/envs/{env}", GetEnvHandler).Methods("GET")
	r.HandleFunc("/api/envs/{env}/hosts/", ListHostsHandler).Methods("GET")
	r.HandleFunc("/api/envs/{env}/hosts/{host}", GetHostHandler).Methods("GET")
	r.HandleFunc("/api/envs/{env}/hosts/{host}/services/", ListServicesHandler).Methods("GET")
	r.HandleFunc("/api/envs/{env}/hosts/{host}/services/{service}", GetServiceHandler).Methods("GET")
	r.HandleFunc("/api/envs/{env}/hosts/{host}/services/{service}", UpdateServiceHandler).Methods("POST")
	r.HandleFunc("/ws", WsHandler).Methods("GET")

	if staticDir == "" {
		if err != nil {
			log.Fatal(err)
		}
		r.PathPrefix("/").Handler(http.FileServer(FS(false)))
	} else {
		r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))
	}

	chain := alice.New().Then(r)

	fmt.Printf("Serving IcingaMock at http://%s\nPress CTRL-c to stop...\n", listener)
	log.Fatal(http.ListenAndServe(listener, chain))
}
