package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var s *Server

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(s.r.Players)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	s = new (Server)
	s.r = NewRoom(3)
	s.r.AddPlayer(ClientInfo{"BeLuckyDaf", "nil"})
	s.r.AddPlayer(ClientInfo{"Ababwa", "nil"})

	go LaunchPaytimeTimer(s)

	r := mux.NewRouter()
	r.HandleFunc("/players", getPlayers).Methods("GET")
	log.Fatal(http.ListenAndServe(":34000", r))
}
