package main

const (
	LoctypePlanet   = 0
	LoctypeAsteroid = 1
	LoctypeStation  = 2
)

type Node struct {
	LocType int `json:"loc_type"`
	adjacent []*Node
}
