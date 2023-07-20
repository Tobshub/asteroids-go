package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH  int32 = 800
	SCREEN_HEIGHT int32 = 600
)

var PLAYER Player

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
		Width:        32,
		Height:       48,
		Center:       rl.Vector2{X: float32(SCREEN_WIDTH/2 - PLAYER.Width), Y: float32(SCREEN_HEIGHT/2 - PLAYER.Height)},
		Velocity:     0,
		Acceleration: 0,
		Rotation:     0,
	}
}

func DrawGame() {
	rl.ClearBackground(rl.RayWhite)
	PLAYER.Draw()
}

func UpdateGame() {
	PLAYER.Update()
}
