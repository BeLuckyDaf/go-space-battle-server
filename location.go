package main

const (
	LoctypePlanet   = 0
	LoctypeAsteroid = 1
	LoctypeStation  = 2
)

type Node struct {
	adjacent []*Node
	LocType int `json:"loc_type"`
}
