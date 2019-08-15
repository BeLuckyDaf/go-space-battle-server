package main

import (
	"fmt"
	"time"
)

const (
	LoctypePlanet   = 0
	LoctypeAsteroid = 1
	LoctypeStation  = 2
)

const (
	PlayerPowerInitial = 3
	PlayerHealthInitial = 3
)

const PaytimeInterval = time.Second * 2

type Node struct {
	adjacent []*Node
	LocType int `json:"loc_type"`
}

type Player struct {
	Info ClientInfo `json:"info"`
	Power    int `json:"power"`
	Hp       int `json:"hp"`
	Location *Node `json:"location"`
}

type ClientInfo struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

type Room struct {
	WorldMap []*Node
	Players []*Player
	MaxPlayers int
}

func NewRoom(maxPlayers int) *Room {
	return &Room{
		WorldMap: nil,
		Players:  []*Player{},
		MaxPlayers: maxPlayers,
	}
}

func (r *Room) AddPlayer(info ClientInfo) bool {
	if len(r.Players) < r.MaxPlayers {
		r.Players = append(r.Players, &Player{
			Info: info,
			Power:    PlayerPowerInitial,
			Location: nil,
			Hp:       PlayerHealthInitial,
		})
		return true
	}
	return false
}

type Server struct {
	r *Room
	t *time.Timer
	paytimeEnabled bool
	timerRunning bool
}

func (s *Server) DisablePaytime() {
	s.paytimeEnabled = false
}

func (s *Server) EnablePaytime() {
	s.paytimeEnabled = true
}

func (s *Server)handlePaytime() {
	for _, p := range s.r.Players {
		p.Power++
		fmt.Println(p)
	}
}

func LaunchPaytimeTimer(s *Server) {
	if s.timerRunning {
		fmt.Println("ANOTHER TIMER IS ALREADY RUNNING")
		return
	}

	s.EnablePaytime()

	if s.t == nil {
		s.t = time.NewTimer(PaytimeInterval)
	} else {
		s.t.Reset(PaytimeInterval)
	}

	for {
		s.timerRunning = true
		a := <-s.t.C

		// PAYTIME HERE
		fmt.Println("PAYTIME", a)
		s.handlePaytime()

		// RESET TIMER
		if s.paytimeEnabled {
			s.t.Reset(PaytimeInterval)
		} else {
			s.timerRunning = false
			break
		}
	}

	fmt.Println("PAYTIME TIMER STOPPED")
}