package main

// Slogger is a global logger object
var Slogger *Logger

var s *Server
var api *API

func main() {
	s = new(Server)
	api = new(API)
	s.Room = NewRoom(3, 64)
	Slogger = NewLogger("logs.txt")

	go s.LaunchPaytimeTimer()
	Slogger.Log("Started server at port 34000.")

	api.Init(s)
}
