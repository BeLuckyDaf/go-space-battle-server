package main

import (
	"fmt"
	"math/rand"
)

const MinimalDistance = 5.0

type World struct {
	Size   int                `json:"size"`
	Points map[int]WorldPoint `json:"points"`
}

func GenerateWorld(s int) *World {
	wp := make(map[int]WorldPoint)

	for i := 0; i < s; i++ {
		wp[i] = WorldPoint{
			LocType:  rand.Intn(3),
			Position: generatePosition(wp, i),
			Adjacent: nil,
		}
	}

	w := World{
		Size:   s,
		Points: wp,
	}

	return &w
}

func generatePosition(wp map[int]WorldPoint, s int) Vector2 {
	v := Vector2{
		X: rand.Intn(1000),
		Y: rand.Intn(1000),
	}

	for !checkDistance(v, wp, s) {
		v = Vector2{
			X: rand.Intn(1000),
			Y: rand.Intn(1000),
		}
	}

	return v
}

func checkDistance(v Vector2, wp map[int]WorldPoint, s int) bool {
	if s == 0 {
		return true
	}

	for i := 0; i < s; i++ {
		p, ok := wp[i]
		if !ok {
			fmt.Println("Invalid map access. Perhaps checkDistance size argument is wrong.")
		}

		if p.Position.Distance(v) < MinimalDistance {
			return false
		}
	}

	return true
}
