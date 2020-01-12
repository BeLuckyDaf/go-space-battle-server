package main

import "math/rand"

// Room is used as a general representation of a room is the world
type Room struct {
	GameWorld  *World             `json:"game_world"`
	Players    map[string]*Player `json:"players"`
	MaxPlayers int                `json:"max_players"`
}

// NewRoom creates a new room in the world
func NewRoom(maxPlayers, worldSize int) Room {
	return Room{
		GameWorld:  GenerateWorld(worldSize),
		Players:    make(map[string]*Player),
		MaxPlayers: maxPlayers,
	}
}

// AddPlayer adds the client to the room
func (r *Room) AddPlayer(username string, token string) bool {
	if len(r.Players) < r.MaxPlayers {
		r.Players[username] = &Player{
			Username: username,
			Token:    token,
			Power:    PlayerPowerInitial,
			Location: rand.Intn(r.GameWorld.Size),
			Hp:       PlayerHealthInitial,
		}
		return true
	}
	return false
}
