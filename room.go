package main

import "math/rand"

type Room struct {
	GameWorld  *World   `json:"game_world"`
	Players    []Player `json:"players"`
	MaxPlayers int      `json:"max_players"`
}

func NewRoom(maxPlayers, worldSize int) Room {
	return Room{
		GameWorld:  GenerateWorld(worldSize),
		Players:    []Player{},
		MaxPlayers: maxPlayers,
	}
}

func (r *Room) AddPlayer(info ClientInfo) bool {
	if len(r.Players) < r.MaxPlayers {
		r.Players = append(r.Players, Player{
			Info:     info,
			Power:    PlayerPowerInitial,
			Location: rand.Intn(r.GameWorld.Size),
			Hp:       PlayerHealthInitial,
		})
		return true
	}
	return false
}
