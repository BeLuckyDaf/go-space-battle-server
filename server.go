package main

import (
	"fmt"
	"time"
)

const PaytimeInterval = time.Second * 5

type Server struct {
	Room           Room `json:"room"`
	PaytimeEnabled bool `json:"paytime_enabled"`
	timer          *time.Timer
	timerRunning   bool
}

func (s *Server) DisablePaytime() {
	s.PaytimeEnabled = false
}

func (s *Server) EnablePaytime() {
	s.PaytimeEnabled = true
}

func (s *Server) handlePaytime() {
	for i := range s.Room.Players {
		s.Room.Players[i].Power++
		fmt.Println(s.Room.Players[i])
	}
}

func LaunchPaytimeTimer(s *Server) {
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
