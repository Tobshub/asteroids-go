package main

import rl "github.com/gen2brain/raylib-go/raylib"

const MAX_PLAYER_SPEED int32 = 10
const ROTATION_SPEED int32 = 1

type Player struct {
	Width  int32
	Height int32

	Position     rl.Vector2
	Velocity     int32
	Rotation     int32
	Acceleration int32
}

func (p *Player) Update() {
	switch key := rl.GetKeyPressed(); key {
	case rl.KeyRight:
		p.Rotation += ROTATION_SPEED
		if p.Rotation > 360 {
			p.Rotation = 0
		}
	case rl.KeyLeft:
		p.Rotation -= ROTATION_SPEED
		if p.Rotation < 0 {
			p.Rotation = 360
		}
	}
}

func (p *Player) Draw() {
	v2 := rl.Vector2{
		X: p.Position.X - float32(p.Width/2),
		Y: p.Position.Y + float32(p.Height),
	}
	v3 := rl.Vector2{
		X: p.Position.X + float32(p.Width/2),
		Y: p.Position.Y + float32(p.Height),
	}

	rl.DrawTriangle(p.Position, v2, v3, rl.Red)
}
