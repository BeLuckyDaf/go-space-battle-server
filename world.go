package main

import (
	"fmt"
	"math/rand"
)

const MinimalDistance = 5.0
const EdgeDistance = 300.0

type World struct {
	Size   int                 `json:"size"`
	Points map[int]*WorldPoint `json:"points"`
}

func GenerateWorld(s int) *World {
	wp := make(map[int]*WorldPoint)

	fmt.Println("Generating world... 0%")

	for i := 0; i < s; i++ {
		wp[i] = &WorldPoint{
			LocType:  rand.Intn(3),
			Position: generatePosition(wp, i),
			Adjacent: make(map[int]bool),
		}
	}

	fmt.Println("Generating world... 100%")

	for i := 0; i < s-1; i++ {
		fmt.Printf("Generating edges... %d%%\n", 100*(i+1)/s)
		for j := i + 1; j < s; j++ {
			if wp[i].Position.Distance(wp[j].Position) < EdgeDistance {
				wp[i].Adjacent[j] = true
				wp[j].Adjacent[i] = true
			}
		}
	}

	fmt.Println("Generating edges... 100%")

	w := World{
		Size:   s,
		Points: wp,
	}

	return &w
}

func generatePosition(wp map[int]*WorldPoint, s int) Vector2 {
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

func checkDistance(v Vector2, wp map[int]*WorldPoint, s int) bool {
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