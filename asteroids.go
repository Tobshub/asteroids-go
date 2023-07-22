package main

import rl "github.com/gen2brain/raylib-go/raylib"

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

func ExplodeAsteroid(index int) {
	asteroid := ASTEROIDS[index]

	if asteroid.Size != Small {
		var new_size AsteroidSize

		if asteroid.Size == Large {
			new_size = Medium
		} else if asteroid.Size == Medium {
			new_size = Small
		}

		CreateLeftSideDebris(&asteroid, new_size)
		CreateRightSideDebris(&asteroid, new_size)
	}

	ASTEROIDS = append(ASTEROIDS[:index], ASTEROIDS[index+1:]...)
}

func CreateLeftSideDebris(asteroid *Asteroid, size AsteroidSize) {
	// TODO: fix direction
	SpawnAsteroid(
		asteroid.Origin,
		asteroid.Position.X-float32(asteroid.Size),
		asteroid.Position.Y-float32(asteroid.Size),
		asteroid.Direction+90, size,
		asteroid.Velocity,
	)
}

func CreateRightSideDebris(asteroid *Asteroid, size AsteroidSize) {
	// TODO: fix direction
	SpawnAsteroid(
		asteroid.Origin,
		asteroid.Position.X+float32(asteroid.Size),
		asteroid.Position.Y+float32(asteroid.Size),
		asteroid.Direction-90, size,
		asteroid.Velocity,
	)
}

func SpawnAsteroid(origin AsteroidSpawn, x float32, y float32, direction int32, size AsteroidSize, v float32) {
	asteroid := Asteroid{
		Size:      size,
		Origin:    origin,
		Position:  rl.Vector2{X: x, Y: y},
		Velocity:  v,
		Direction: direction,
	}
	ASTEROIDS = append(ASTEROIDS, asteroid)
}

func DeSpawnAsteroid(index int) {
	ASTEROIDS = append(ASTEROIDS[:index], ASTEROIDS[index+1:]...)
}
