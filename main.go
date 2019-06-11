package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type info struct {
	Version       string `json:"version"`
	LastCommitSHA string `json:"lastcommitsha"`
	Description   string `json:"description"`
}

type app struct {
	AppName []info `json:"myapplication"`
}

type health struct {
	Health string `json:"health"`
}

type welcome struct {
	Welcome string `json:welcome`
}

func main() {
	fmt.Println("starting http server ")
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/info", getInfo)
	r.HandleFunc("/health", healthCheck)

	s := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	myHealth := health{
		Health: "Still alive!",
	}

	healthJSON, err := json.Marshal(myHealth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(healthJSON)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	myWelcome := welcome {
		Welcome: "Hello World!",
	}

	welcomeJSON, err := json.Marshal(myWelcome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(welcomeJSON)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	var version = os.Getenv("VERSION")
	var lastCommitSHA = os.Getenv("LAST_COMMIT_SHA")

	myInfo := info{
		Version:       version,
		LastCommitSHA: lastCommitSHA,
		Description:   "pre-interview technical test",
	}

	myApp := app{
		AppName: []info{
			myInfo,
		},
	}

	infoJSON, err := json.Marshal(myApp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(infoJSON)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
