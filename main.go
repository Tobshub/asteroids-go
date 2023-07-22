package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH      int32 = 800
	SCREEN_HEIGHT     int32 = 600
	INITIAL_ASTERIODS int   = 5
)

var (
	PLAYER    Player
	ASTEROIDS []Asteroid
	has_lost  bool = false
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
	PLAYER = Player{
		Width:     28,
		Height:    32,
		Center:    rl.Vector2{X: float32(SCREEN_WIDTH/2 - PLAYER.Width), Y: float32(SCREEN_HEIGHT/2 - PLAYER.Height)},
		Velocity:  0,
		Direction: 270,
	}

	ASTEROIDS = []Asteroid{}

	for i := 0; i < INITIAL_ASTERIODS; i++ {
		SpawnAsteroidFromOrigin(RandomAsteroidOrigin())
	}
}

func DrawGame() {
	rl.ClearBackground(rl.RayWhite)
	PLAYER.Draw()

	for _, asteroid := range ASTEROIDS {
		asteroid.Draw()
	}
}

var frame_count int = 0

func UpdateGame() {
	frame_count++
	if frame_count%60 == 0 {
		SpawnAsteroidFromOrigin(RandomAsteroidOrigin())
	}

	for i := 0; i < len(ASTEROIDS); i++ {
		var asteroid *Asteroid = &ASTEROIDS[i]

		if (rl.CheckCollisionCircleRec(
			asteroid.Position,
			float32(asteroid.Size),
			rl.Rectangle{
				X:      PLAYER.Center.X - float32(PLAYER.Width/2),
				Y:      PLAYER.Center.Y - float32(PLAYER.Height/2),
				Width:  float32(PLAYER.Width),
				Height: float32(PLAYER.Height),
			},
		)) {
			has_lost = true
			break
		}

		/** is_out :=  */
		asteroid.Update()

		// if is_out {
		// 	RemoveAsteroid(i)
		// }
	}

	if has_lost {
		EndGame()
		return
	}

	PLAYER.Update()
}

func EndGame() {
	has_lost = false
	InitGame()
}
