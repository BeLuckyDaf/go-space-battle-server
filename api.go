package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Api struct {
	r *mux.Router
	s *Server
}

func (a *Api) getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	writeSuccess(w, a.s.Room.Players)
}

func (a *Api) getWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if s != nil {
		writeSuccess(w, a.s.Room.GameWorld)
	} else {
		writeError(w, "Server is nil.")
	}
}

func (a *Api) connectPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(a.s.Room.Players) >= a.s.Room.MaxPlayers {
		writeError(w, "Max players reached.")
		return
	}

	username := r.URL.Query().Get("username")

	if len(username) < 3 {
		writeError(w, "Username too short.")
		return
	}

	for _, p := range a.s.Room.Players {
		if strings.Compare(p.Username, username) == 0 {
			fmt.Println("PLAYER ALREADY CONNECTED")
			writeError(w, "Player already connected.")
			return
		}
	}

	hasher := sha1.New()
	hasher.Write([]byte(username + time.Now().String()))
	token := hex.EncodeToString(hasher.Sum(nil))

	a.s.Room.AddPlayer(username, token)

	// Only show token here, nowhere else
	data := make(map[string]interface{})
	data["player"] = a.s.Room.Players[username]
	data["token"] = a.s.Room.Players[username].Token

	writeSuccess(w, data)
}

func (a *Api) movePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	ok, p, _ := a.getPlayerDataFromQuery(w, q)
	if !ok {
		return
	}
	target, err := strconv.Atoi(q.Get("target"))
	if err != nil {
		writeError(w, "Invalid target, NaN.")
	}

	if p.Location == target {
		writeError(w, "Cannot move to current position.")
	} else if p != nil && a.s.Room.GameWorld.Points[p.Location].Adjacent[target] {
		p.Location = target
		writeSuccess(w, p)
	} else {
		writeError(w, "Target is not an adjacent point.")
	}
}

func (a *Api) buyLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	ok, p, u := a.getPlayerDataFromQuery(w, q)
	if !ok {
		return
	}
	if strings.Compare(a.s.Room.GameWorld.Points[p.Location].OwnedBy, "") > 0 {
		writeError(w, "Point already owned.")
		return
	}
	if p.Power < 1 {
		writeError(w, "Not enough power.")
		return
	}

	a.s.Room.GameWorld.Points[p.Location].OwnedBy = u
	p.Power--
	writeSuccess(w, a.s.Room.GameWorld.Points[p.Location])
}

func (a *Api) destroyLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	ok, p, _ := a.getPlayerDataFromQuery(w, q)
	if !ok {
		return
	}
	if strings.Compare(a.s.Room.GameWorld.Points[p.Location].OwnedBy, "") == 0 {
		writeError(w, "Point is not owned by anyone.")
		return
	}
	if p.Power < 1 {
		writeError(w, "Not enough power.")
		return
	}

	a.s.Room.GameWorld.Points[p.Location].OwnedBy = ""
	p.Power--
	writeSuccess(w, a.s.Room.GameWorld.Points[p.Location])
}

func (a *Api) attackPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	ok, p, _ := a.getPlayerDataFromQuery(w, q)
	if !ok {
		return
	}
	target := a.s.Room.Players[q.Get("target")]
	if target == nil {
		writeError(w, "Target player not found.")
		return
	}
	if p.Power < 1 {
		writeError(w, "Not enough power.")
		return
	}
	if target.Hp < 1 {
		writeError(w, "Target player already dead.")
		return
	}
	if target.Location != p.Location {
		writeError(w, "Target player is not in range.")
		return
	}

	target.Hp--
	p.Power--
	writeSuccess(w, target)
}

func (a *Api) getPlayerDataFromQuery(w http.ResponseWriter, q url.Values) (bool, *Player, string) {
	username := q.Get("username")
	token := q.Get("token")
	p := a.s.Room.Players[username]
	if p == nil {
		writeError(w, "Player not found.")
		return false, nil, ""
	}
	if strings.Compare(token, p.Token) != 0 {
		writeError(w, "Invalid token.")
		return false, nil, ""
	}
	if p.Hp < 1 {
		writeError(w, "Player dead.")
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

func (a *Api) Init(s *Server) {
	a.s = s
	a.r = mux.NewRouter()
	a.r.HandleFunc("/players", a.getPlayers).Methods("GET")
	a.r.HandleFunc("/world", a.getWorld).Methods("GET")
	a.r.HandleFunc("/connect", a.connectPlayer).Methods("GET")
	a.r.HandleFunc("/move", a.movePlayer).Methods("GET")
	a.r.HandleFunc("/buy", a.buyLocation).Methods("GET")
	a.r.HandleFunc("/destroy", a.destroyLocation).Methods("GET")
	a.r.HandleFunc("/attack", a.attackPlayer).Methods("GET")
	log.Fatal(http.ListenAndServe(":34000", a.r))
}