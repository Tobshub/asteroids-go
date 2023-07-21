package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MAX_ASTEROID_VELOCITY int32 = 8
	MIN_ASTEROID_VELOCITY int32 = 4
)

type AsteroidSize int32

const (
	Small  AsteroidSize = 15
	Medium AsteroidSize = 30
	Large  AsteroidSize = 60
)

type Asteroid struct {
	Size AsteroidSize

	Position  rl.Vector2
	Velocity  float32
	Direction int32
}

func (a *Asteroid) Update() {
	displace_x, displace_y := DisplacementComponents(a.Velocity, a.Direction)

	a.Position.X += displace_x
	a.Position.Y += displace_y
}

func (a *Asteroid) Draw() {
	rl.DrawCircleV(a.Position, float32(a.Size), rl.DarkGray)
}

func RandomAsteroidPosition(_size AsteroidSize, exclusion_area rl.Rectangle) rl.Vector2 {
	x := exclusion_area.X
	y := exclusion_area.Y

	var size float32 = float32(_size)

	for (x >= exclusion_area.X-size && x <= exclusion_area.X+exclusion_area.Width+size) ||
		(y >= exclusion_area.Y-size && y <= exclusion_area.Y+exclusion_area.Height+size) {
		x = rand.Float32() * float32(SCREEN_WIDTH)
		y = rand.Float32() * float32(SCREEN_HEIGHT)
	}

	return rl.Vector2{X: x, Y: y}
}

func RandomAsteroidVelocity() float32 {
	return float32(MIN_ASTEROID_VELOCITY) + float32(MAX_ASTEROID_VELOCITY-MIN_ASTEROID_VELOCITY)*rand.Float32()
}

func RandomAsteroidDirection() int32 {
	return int32(rand.Intn(360))
}
