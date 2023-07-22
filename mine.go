package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	MINE_SIZE = 10

	MAX_MINE_COUNT    = 3
	NEW_MINE_LIFETIME = 500
)

type Mine struct {
	Center   rl.Vector2
	LifeTime int
}

func PlaceMine() {
	MINES = append(MINES, Mine{
		Center:   rl.Vector2{X: PLAYER.Center.X, Y: PLAYER.Center.Y},
		LifeTime: NEW_MINE_LIFETIME,
	})
}

func (mine *Mine) Update() bool {
	if mine.LifeTime > 0 {
		mine.LifeTime--
		return false
	} else {
		return true
	}
}

func (mine *Mine) Draw() {
	rl.DrawPoly(mine.Center, 4, MINE_SIZE, 45, rl.Red)
}

func DetonateMine(index int) {
	MINES = append(MINES[:index], MINES[index+1:]...)
}
