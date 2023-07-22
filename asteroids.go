package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MAX_ASTEROID_VELOCITY int32 = 6
	MIN_ASTEROID_VELOCITY int32 = 2
)

type AsteroidSize int32

const (
	Small  AsteroidSize = 15
	Medium AsteroidSize = 30
	Large  AsteroidSize = 60
)

type Asteroid struct {
	Size   AsteroidSize
	Origin AsteroidSpawn

	Position  rl.Vector2
	Velocity  float32
	Direction int32
}

func (a *Asteroid) Update() bool {
	displace_x, displace_y := DisplacementComponents(a.Velocity, a.Direction)

	a.Position.X += displace_x
	a.Position.Y += displace_y
	size := float32(a.Size)

	if (a.Position.X+size < 0 && a.Origin != Left) || (a.Position.X-size > float32(SCREEN_WIDTH) && a.Origin != Right) ||
		(a.Position.Y+size < 0 && a.Origin != Top) || (a.Position.Y-size > float32(SCREEN_HEIGHT) && a.Origin != Bottom) {
		return true
	} else {
		return false
	}
}

func (a *Asteroid) Draw() {
	rl.DrawCircleV(a.Position, float32(a.Size), rl.DarkGray)
}

func RandomAsteroidVelocity() float32 {
	return float32(MIN_ASTEROID_VELOCITY) + float32(MAX_ASTEROID_VELOCITY-MIN_ASTEROID_VELOCITY)*rand.Float32()
}

func RemoveAsteroid(index int) {
	ASTEROIDS = append(ASTEROIDS[:index], ASTEROIDS[index+1:]...)
}

func SpawnAsteroid(origin AsteroidSpawn, x float32, y float32, direction int32) {
	asteroid := Asteroid{
		Size:      Large,
		Origin:    origin,
		Position:  rl.Vector2{X: x, Y: y},
		Velocity:  RandomAsteroidVelocity(),
		Direction: direction,
	}
	ASTEROIDS = append(ASTEROIDS, asteroid)
}
