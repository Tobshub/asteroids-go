package main

import rl "github.com/gen2brain/raylib-go/raylib"

type AsteroidSize int32

const (
	Small AsteroidSize = iota
	Medium
	Large
)

type Asteroid struct {
	Size AsteroidSize

	Position     rl.Vector2
	Velocity     int32
	Rotation     int32
	Acceleration uint32
}

func (a *Asteroid) Update() {
}

func (a *Asteroid) Draw() {
}
