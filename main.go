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

	err := json.NewEncoder(w).Encode(Message{
		Status: true,
		Data:   s.Room.Players,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	if s != nil {
		err = json.NewEncoder(w).Encode(Message{
			Status: true,
			Data:   s,
		})
	} else {
		err = json.NewEncoder(w).Encode(Message{
			Status: false,
			Data:   nil,
		})
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

func connectMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO connect player here, add to list

	err := json.NewEncoder(w).Encode(Message{
		Status: false,
		Data:   nil,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	s = new (Server)
	s.Room = NewRoom(3)
	s.Room.AddPlayer(ClientInfo{"BeLuckyDaf", "nil"})
	s.Room.AddPlayer(ClientInfo{"Ababwa", "nil"})

	go LaunchPaytimeTimer(s)

	r := mux.NewRouter()
	r.HandleFunc("/players", getPlayers).Methods("GET")
	r.HandleFunc("/status", getStatus).Methods("GET")
	r.HandleFunc("/connect", connectMe).Methods("GET")
	log.Fatal(http.ListenAndServe(":34000", r))
}
