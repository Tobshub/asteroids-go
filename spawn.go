package main

import "math/rand"

type AsteroidSpawn int32

// origin determines direction range
//
// direction -> range (not inclusive)
//
// Top -> 0 - 180
//
// Bottom -> 180 - 360
//
// Left -> 90 - 270
//
// Right -> 270 - 270+180
const (
	Top AsteroidSpawn = iota
	Bottom
	Left
	Right
)

func RandomInt32(min int32, max int32) int32 {
	return min + rand.Int31n(max-min)
}

func RandomAsteroidDirection(min int32, max int32) int32 {
	deg := RandomInt32(min+1, max)
	if deg > 360 {
		return deg - 360
	} else {
		return deg
	}
}

func RandomAsteroidVelocity() float32 {
	return float32(MIN_ASTEROID_VELOCITY) + float32(MAX_ASTEROID_VELOCITY-MIN_ASTEROID_VELOCITY)*rand.Float32()
}

func RandomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandomX() float32 {
	return RandomFloat32(0, float32(SCREEN_WIDTH))
}

func RandomY() float32 {
	return RandomFloat32(0, float32(SCREEN_HEIGHT))
}

func SpawnAsteroidFromOrigin(origin AsteroidSpawn) {
	var x, y float32
	var direction int32

	size := Large

	switch origin {
	case Top:
		x = RandomX()
		y = float32(0 - size)
		direction = RandomAsteroidDirection(0, 180)
	case Bottom:
		x = RandomX()
		y = float32(SCREEN_HEIGHT + int32(size))
		direction = RandomAsteroidDirection(180, 360)
	case Left:
		x = float32(0 - size)
		y = RandomY()
		direction = RandomAsteroidDirection(90, 270)
	case Right:
		x = float32(SCREEN_WIDTH + int32(size))
		y = RandomY()
		direction = RandomAsteroidDirection(270, 270+180)
	}
	SpawnAsteroid(origin, x, y, direction, Large, RandomAsteroidVelocity())
}

func RandomAsteroidOrigin() AsteroidSpawn {
	n := rand.Intn(4)
	switch n {
	case 0:
		return Top
	case 1:
		return Bottom
	case 2:
		return Left
	case 3:
		return Right
	}
	return Top
}
