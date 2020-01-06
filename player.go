package main

// PlayerPowerInitial is the initial amount of power
// PlayerHealthInitial is the initial amount of health
const (
	PlayerPowerInitial  = 3
	PlayerHealthInitial = 3
)

// Player is used as a general representation of a player
type Player struct {
	Info     ClientInfo `json:"info"`
	Power    int        `json:"power"`
	Hp       int        `json:"hp"`
	Location int        `json:"location"`
}

// ClientInfo is used as a general representation of a client
type ClientInfo struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
