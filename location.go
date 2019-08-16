package main

import "math"

const (
	LoctypePlanet   = 0
	LoctypeAsteroid = 1
	LoctypeStation  = 2
)

type Vector2 struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (v1 Vector2) Distance(v2 Vector2) float64 {
	y := v2.Y - v1.Y
	x := v2.X - v1.X
	return math.Sqrt(float64(x*x + y*y))
}

type WorldPoint struct {
	LocType  int          `json:"loc_type"`
	Position Vector2      `json:"position"`
	Adjacent map[int]bool `json:"adjacent"`
}
