package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH      int32 = 800
	SCREEN_HEIGHT     int32 = 600
	INITIAL_ASTERIODS int   = 3
)

var (
	PLAYER    Player
	ASTEROIDS []Asteroid
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Asteroids")

	rl.SetTargetFPS(60)

	InitGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		UpdateGame()
		DrawGame()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func InitGame() {
	// init player
	PLAYER = Player{
		Width:     28,
		Height:    32,
		Center:    rl.Vector2{X: float32(SCREEN_WIDTH/2 - PLAYER.Width), Y: float32(SCREEN_HEIGHT/2 - PLAYER.Height)},
		Velocity:  0,
		Direction: 270,
	}

	ASTEROIDS = []Asteroid{}

	for i := 0; i < INITIAL_ASTERIODS; i++ {
		asteroid := Asteroid{
			Size: Large,
			Position: RandomAsteroidPosition(Large, rl.Rectangle{
				X:      PLAYER.Center.X - float32(PLAYER.Width/2),
				Y:      PLAYER.Center.Y - float32(PLAYER.Height/2),
				Width:  float32(PLAYER.Width),
				Height: float32(PLAYER.Height),
			}),
			Velocity:  RandomAsteroidVelocity(),
			Direction: RandomAsteroidDirection(),
		}

		ASTEROIDS = append(ASTEROIDS, asteroid)
	}
}

func DrawGame() {
	rl.ClearBackground(rl.RayWhite)
	PLAYER.Draw()

	for _, asteroid := range ASTEROIDS {
		asteroid.Draw()
	}
}

func UpdateGame() {
	for i := 0; i < len(ASTEROIDS); i++ {
		ASTEROIDS[i].Update()
	}

	PLAYER.Update()
}
