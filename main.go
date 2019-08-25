package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
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
	ok, p, _ := getPlayerDataFromQuery(w, q)
	if !ok {
		return
	}
	target, err := strconv.Atoi(q.Get("target"))
	if err != nil {
		writeError(w, "Invalid target, NaN.")
	}

	if p.Location == target {
		writeError(w, "Cannot move to current position.")
	} else if p != nil && s.Room.GameWorld.Points[p.Location].Adjacent[target] {
		p.Location = target
		writeSuccess(w, p)
	} else {
		writeError(w, "Target is not an adjacent point.")
	}
}

func buyLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	ok, p, u := getPlayerDataFromQuery(w, q)
	if !ok {
		return
	}
	if strings.Compare(s.Room.GameWorld.Points[p.Location].OwnedBy, "") > 0 {
		writeError(w, "Point already owned.")
		return
	}
	if p.Power < 1 {
		writeError(w, "Not enough power.")
		return
	}

	s.Room.GameWorld.Points[p.Location].OwnedBy = u
	p.Power -= 1
	writeSuccess(w, s.Room.GameWorld.Points[p.Location])
}

func destroyLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	ok, p, _ := getPlayerDataFromQuery(w, q)
	if !ok {
		return
	}
	if strings.Compare(s.Room.GameWorld.Points[p.Location].OwnedBy, "") == 0 {
		writeError(w, "Point is not owned by anyone.")
		return
	}
	if p.Power < 1 {
		writeError(w, "Not enough power.")
		return
	}

	s.Room.GameWorld.Points[p.Location].OwnedBy = ""
	p.Power -= 1
	writeSuccess(w, s.Room.GameWorld.Points[p.Location])
}

func getPlayerDataFromQuery(w http.ResponseWriter, q url.Values) (bool, *Player, string) {
	username := q.Get("username")
	token := q.Get("token")
	p := s.Room.Players[username]
	if strings.Compare(token, p.Info.Token) != 0 {
		writeError(w, "Invalid token.")
		return false, nil, ""
	}
	return true, p, username
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
	r.HandleFunc("/buy", buyLocation).Methods("GET")
	r.HandleFunc("/destroy", destroyLocation).Methods("GET")
	log.Fatal(http.ListenAndServe(":34000", r))
}
