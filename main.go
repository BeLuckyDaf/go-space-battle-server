package main

// Slogger is a global logger object
var Slogger *Logger

var s *Server
var api *API

func main() {
	Slogger = NewLogger("logs.txt")
	s = NewServer()
	api = NewAPI(s)
	s.Room = NewRoom(3, 64)

	go s.LaunchPaytimeTimer()
	Slogger.Log("Started server at port 34000.")

	api.Start()
}
