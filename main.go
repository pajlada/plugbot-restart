package main

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

func restartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Restarting plugbot...")
	cmd := exec.Command("/usr/bin/pm2", "restart", "plugbot")

	err := cmd.Run()
	if err != nil {
		log.Println("Error:", err)
	}
}

var mySecretPasswordOpt = flag.String("password", "", "Specify a password")

func main() {
	flag.Parse()

	mySecretPassword := *mySecretPasswordOpt
	if mySecretPassword == "" {
		panic("Missing password")
	}

	log.Println("Server up")
	r := mux.NewRouter()
	s := r.PathPrefix("/plugdj").Subrouter()

	s.HandleFunc("/restart", restartHandler).Queries("password", mySecretPassword)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
