// Copyright 2020 Vladislav Smirnov

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// World is used as a general structure of a world
type World struct {
	Size   int                 `json:"size"`
	Points map[int]*WorldPoint `json:"points"`
}

/* --- WORLD GENERATION --- */

// TODO: optimize or rewrite world generation

// MinimalDistance is the distance between the nodes for a path
const MinimalDistance = 5.0

// EdgeDistance is a minimal distance of an edge (maybe not)
const EdgeDistance = 50.0

// GenerateWorld create a world of s points
func GenerateWorld(s int) *World {
	wp := make(map[int]*WorldPoint)

	fmt.Println("Generating world... 0%")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < s; i++ {
		wp[i] = &WorldPoint{
			LocType:  rand.Intn(3),
			Position: generatePosition(wp, i),
			Adjacent: make([]int, 0),
		}
	}

	fmt.Println("Generating world... 100%")

	for i := 0; i < s-1; i++ {
		fmt.Printf("Generating edges... %d%%\n", 100*(i+1)/s)
		nearestID := -1
		nearestDist := 0.0
		for j := i + 1; j < s; j++ {
			dist := wp[i].Position.Distance(wp[j].Position)
			if dist < EdgeDistance {
				wp[i].Adjacent = append(wp[i].Adjacent, j)
				wp[j].Adjacent = append(wp[j].Adjacent, i)
			}
			if nearestID == -1 || dist < nearestDist {
				nearestID = j
				nearestDist = dist
			}
		}
		// if no edges added, find the nearest
		if len(wp[i].Adjacent) == 0 {
			wp[i].Adjacent = append(wp[i].Adjacent, nearestID)
			wp[nearestID].Adjacent = append(wp[nearestID].Adjacent, i)
		}
	}

	fmt.Println("Generating edges... 100%")

	w := World{
		Size:   s,
		Points: wp,
	}

	return &w
}

// TODO: this may generate the same coordinates for different points
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
