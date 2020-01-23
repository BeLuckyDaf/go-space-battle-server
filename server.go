// Copyright 2020 Vladislav Smirnov

package main

import (
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

// NewServer creates a server
func NewServer() *Server {
	return new(Server)
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
		if point.LocType == LoctypeStation && strings.Compare(pname, point.OwnedBy) != 0 && strings.Compare(point.OwnedBy, "") != 0 {
			p.Hp--
		}
		Slogger.Log(*s.Room.Players[i])
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
		Slogger.Log("TRIED LAUNCHING THE TIMER WHILE ANOTHER TIMER WAS ALREADY RUNNING")
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
		Slogger.Log("PAYTIME", a)
		s.handlePaytime()

		// RESET TIMER
		if s.PaytimeEnabled {
			s.timer.Reset(PaytimeInterval)
		} else {
			s.timerRunning = false
			break
		}
	}

	Slogger.Log("PAYTIME TIMER STOPPED")
}
