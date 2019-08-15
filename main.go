package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
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

func connectPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	username := r.URL.Query().Get("username")
	for _, p := range s.Room.Players {
		if strings.Compare(p.Info.Username, username) == 0 {
			fmt.Println("PLAYER ALREADY CONNECTED")
			err = json.NewEncoder(w).Encode(Message{
				Status: false,
				Data:   "Player already connected.",
			})
			return
		}
	}

	hasher := sha1.New()
	hasher.Write([]byte(username + time.Now().String()))
	token := hex.EncodeToString(hasher.Sum(nil))

	info := ClientInfo{
		Username: username,
		Token:    token,
	}
	s.Room.AddPlayer(info)

	err = json.NewEncoder(w).Encode(Message{
		Status: true,
		Data:   info,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	s = new(Server)
	s.Room = NewRoom(3)

	go LaunchPaytimeTimer(s)

	r := mux.NewRouter()
	r.HandleFunc("/players", getPlayers).Methods("GET")
	r.HandleFunc("/status", getStatus).Methods("GET")
	r.HandleFunc("/connect", connectPlayer).Methods("GET")
	log.Fatal(http.ListenAndServe(":34000", r))
}
