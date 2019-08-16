package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var s *Server

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	writeSuccess(w, s.Room.Players)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if s != nil {
		writeSuccess(w, s)
	} else {
		writeError(w, "Server is nil.")
	}
}

func connectPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(s.Room.Players) >= s.Room.MaxPlayers {
		writeError(w, "Max players reached.")
		return
	}

	username := r.URL.Query().Get("username")

	if len(username) < 3 {
		writeError(w, "Username too short.")
		return
	}

	for _, p := range s.Room.Players {
		if strings.Compare(p.Info.Username, username) == 0 {
			fmt.Println("PLAYER ALREADY CONNECTED")
			writeError(w, "Player already connected.")
			return
		}
	}

	hasher := sha1.New()
	hasher.Write([]byte(username + time.Now().String()))
	token := hex.EncodeToString(hasher.Sum(nil))

	s.Room.AddPlayer(ClientInfo{
		Username: username,
		Token:    token,
	})

	writeSuccess(w, s.Room.Players[username])
}

func movePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	username := q.Get("username")
	target, err := strconv.Atoi(q.Get("target"))
	if err != nil {
		writeError(w, "Invalid target, NaN.")
	}
	token := q.Get("token")

	p := s.Room.Players[username]
	if p.Location == target {
		writeError(w, "Cannot move to current position.")
	} else if p != nil && strings.Compare(token, p.Info.Token) == 0 &&
		s.Room.GameWorld.Points[p.Location].Adjacent[target] {
		p.Location = target
		writeSuccess(w, p)
	} else {
		writeError(w, "Target is not an adjacent point.")
	}
}

func writeError(w http.ResponseWriter, m interface{}) {
	_ = json.NewEncoder(w).Encode(Message{
		Status: false,
		Data:   m,
	})
}

func writeSuccess(w http.ResponseWriter, m interface{}) {
	_ = json.NewEncoder(w).Encode(Message{
		Status: true,
		Data:   m,
	})
}

func main() {
	s = new(Server)
	s.Room = NewRoom(3, 64)

	go LaunchPaytimeTimer(s)
	fmt.Println("Started server.")

	r := mux.NewRouter()
	r.HandleFunc("/players", getPlayers).Methods("GET")
	r.HandleFunc("/status", getStatus).Methods("GET")
	r.HandleFunc("/connect", connectPlayer).Methods("GET")
	r.HandleFunc("/move", movePlayer).Methods("GET")
	log.Fatal(http.ListenAndServe(":34000", r))
}
