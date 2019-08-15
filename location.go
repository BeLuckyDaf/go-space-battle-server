package main

const (
	LoctypePlanet   = 0
	LoctypeAsteroid = 1
	LoctypeStation  = 2
)

type Node struct {
	Id       int     `json:"id"`
	LocType  int     `json:"loc_type"`
	Adjacent []*Node `json:"adjacent"`
}
