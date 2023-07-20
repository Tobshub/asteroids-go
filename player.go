package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MAX_PLAYER_SPEED int32 = 10
	ROTATION_SPEED   int32 = 2
)

type Player struct {
	Width  int32
	Height int32

	Center       rl.Vector2
	Velocity     int32
	Rotation     int32
	Acceleration int32
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyRight) {
		p.Rotation += ROTATION_SPEED
		if p.Rotation > 360 {
			p.Rotation = 0
		}
	} else if rl.IsKeyDown(rl.KeyLeft) {
		p.Rotation -= ROTATION_SPEED
		if p.Rotation < 0 {
			p.Rotation = 360
		}
	}
}

func RotationComponents(rotation int32) (float32, float32) {
	rotation_rad := DegToRad(rotation)
	return float32(math.Cos(rotation_rad)), float32(math.Sin(rotation_rad))
}

func DerivePoints(distance_x float32, distance_y float32, rotation int32) (float32, float32) {
	rot_sin, rot_cos := RotationComponents(rotation)

	x := distance_x*rot_cos - distance_y*rot_sin
	y := distance_x*rot_sin + distance_y*rot_cos

	return x, y
}

func (p *Player) Draw() {
	v1_x, v1_y := DerivePoints(0, float32(p.Height/2), p.Rotation)
	v1 := rl.Vector2{
		X: p.Center.X + v1_x,
		Y: p.Center.Y + v1_y,
	}

	v2_x, v2_y := DerivePoints(float32(p.Width/2), -1*float32(p.Height/2), p.Rotation)
	v2 := rl.Vector2{
		X: p.Center.X + v2_x,
		Y: p.Center.Y + v2_y,
	}

	v3_x, v3_y := DerivePoints(-1*float32(p.Width/2), -1*float32(p.Height/2), p.Rotation)
	v3 := rl.Vector2{
		X: p.Center.X + v3_x,
		Y: p.Center.Y + v3_y,
	}

	rl.DrawTriangle(v1, v2, v3, rl.Red)
}

func DegToRad(deg int32) float64 {
	return float64(deg) * math.Pi / 180
}
