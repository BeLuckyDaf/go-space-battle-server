package main

import (
	"fmt"
	"strings"
	"time"
)

// PaytimeInterval is the time interval between the payments
const PaytimeInterval = time.Second * 5

// Server is used as a general representation of a server
type Server struct {
	Room           Room `json:"room"`
	PaytimeEnabled bool `json:"paytime_enabled"`
	timer          *time.Timer
	timerRunning   bool
}

// DisablePaytime turns off the payments
func (s *Server) DisablePaytime() {
	s.PaytimeEnabled = false
}

// EnablePaytime turns on the payments
func (s *Server) EnablePaytime() {
	s.PaytimeEnabled = true
}

// handlePaytime gives power to players and
// reduces player HP if staying on someone else's station
func (s *Server) handlePaytime() {
	for i, p := range s.Room.Players {
		pname := p.Username
		loc := s.Room.Players[pname].Location
		point := s.Room.GameWorld.Points[loc]
		s.Room.Players[pname].Power++
		if point.LocType == LoctypeStation && strings.Compare(pname, point.OwnedBy) != 0 {
			p.Hp--
		}
		fmt.Println(s.Room.Players[i])
	}

	for _, l := range s.Room.GameWorld.Points {
		if l.LocType != LoctypeStation && strings.Compare(l.OwnedBy, "") != 0 {
			p := s.Room.Players[l.OwnedBy]
			if p == nil {
				continue
			}
			switch l.LocType {
			case LoctypePlanet:
				p.Power += 2
			case LoctypeAsteroid:
				p.Power++
			}
		}
	}
}

// LaunchPaytimeTimer resets and turns on the payments
func (s *Server) LaunchPaytimeTimer() {
	if s.timerRunning {
		fmt.Println("ANOTHER TIMER IS ALREADY RUNNING")
		return
	}

	s.EnablePaytime()

	if s.timer == nil {
		s.timer = time.NewTimer(PaytimeInterval)
	} else {
		s.timer.Reset(PaytimeInterval)
	}

	for {
		s.timerRunning = true
		a := <-s.timer.C

		// PAYTIME HERE
		fmt.Println("PAYTIME", a)
		s.handlePaytime()

		// RESET TIMER
		if s.PaytimeEnabled {
			s.timer.Reset(PaytimeInterval)
		} else {
			s.timerRunning = false
			break
		}
	}

	fmt.Println("PAYTIME TIMER STOPPED")
}
