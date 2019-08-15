package main

import (
	"fmt"
	"time"
)

const PaytimeInterval = time.Second * 2

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