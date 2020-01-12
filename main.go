package main

import (
	"fmt"
)

var s *Server
var api *Api

func main() {
	s = new(Server)
	api = new(Api)
	s.Room = NewRoom(3, 64)

	go s.LaunchPaytimeTimer()
	fmt.Println("Started server at port 34000.")

	api.Init(s)
}
