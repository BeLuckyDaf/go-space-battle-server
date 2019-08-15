package main

type Room struct {
	WorldMap []*Node `json:"world_map"`
	Players []Player `json:"players"`
	MaxPlayers int `json:"max_players"`
}

func NewRoom(maxPlayers int) Room {
	return Room{
		WorldMap: nil,
		Players:  []Player{},
		MaxPlayers: maxPlayers,
	}
}

func (r *Room) AddPlayer(info ClientInfo) bool {
	if len(r.Players) < r.MaxPlayers {
		r.Players = append(r.Players, Player{
			Info: info,
			Power:    PlayerPowerInitial,
			Location: nil,
			Hp:       PlayerHealthInitial,
		})
		return true
	}
	return false
}
