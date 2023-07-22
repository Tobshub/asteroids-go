package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH  int32 = 800
	SCREEN_HEIGHT int32 = 600

	INITIAL_ASTERIODS int = 5

	GAME_SPEED = 50
)

var (
	PLAYER    Player
	ASTEROIDS []Asteroid
	MINES     []Mine
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

	MINES = []Mine{}
	ASTEROIDS = []Asteroid{}

	for i := 0; i < INITIAL_ASTERIODS; i++ {
		SpawnAsteroidFromOrigin(RandomAsteroidOrigin())
	}
}

func DrawGame() {
	rl.ClearBackground(rl.RayWhite)
	if has_lost {
		rl.DrawText("Game Over", SCREEN_WIDTH/2-40, SCREEN_HEIGHT/2-10, 20, rl.Red)
		rl.DrawText("Press ENTER to restart", SCREEN_WIDTH/2-80, SCREEN_HEIGHT/2+25, 16, rl.Gray)
	} else {
		PLAYER.Draw()

		for _, laser := range MINES {
			laser.Draw()
		}

		for _, asteroid := range ASTEROIDS {
			asteroid.Draw()
		}
	}
}

var frame_count int = 0

func UpdateGame() {
	if has_lost {
		if rl.IsKeyPressed(rl.KeyEnter) {
			has_lost = false
			InitGame()
		}
	} else {
		frame_count++

		if frame_count%GAME_SPEED == 0 {
			SpawnAsteroidFromOrigin(RandomAsteroidOrigin())
		}

		if rl.IsKeyPressed(rl.KeySpace) && len(MINES) <= MAX_MINE_COUNT {
			PlaceMine()
		}

	asteroid_loop:
		for i := 0; i < len(ASTEROIDS); i++ {
			asteroid := &ASTEROIDS[i]

			for j := 0; j < len(MINES); j++ {
				mine := &MINES[j]

				if rl.CheckCollisionCircleRec(
					asteroid.Position, float32(asteroid.Size),
					rl.Rectangle{
						X:      mine.Center.X - float32(MINE_SIZE/2),
						Y:      mine.Center.Y - float32(MINE_SIZE/2),
						Width:  MINE_SIZE,
						Height: MINE_SIZE,
					},
				) {
					ExplodeAsteroid(i)
					DetonateMine(j)
					continue asteroid_loop
				}

				life_time_over := mine.Update()

				if life_time_over {
					DetonateMine(j)
				}
			}

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

			is_out := asteroid.Update()

			if is_out {
				DeSpawnAsteroid(i)
			}
		}

		if has_lost {
			return
		}

		PLAYER.Update()
	}
}
