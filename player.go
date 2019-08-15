package main

const (
	PlayerPowerInitial = 3
	PlayerHealthInitial = 3
)

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
