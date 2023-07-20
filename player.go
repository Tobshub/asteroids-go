package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MAX_PLAYER_VELOCITY float32 = 10
	PLAYER_ACCELERATION         = .2
	ROTATION_SPEED      int32   = 5
)

type Player struct {
	Width  int32
	Height int32

	Center    rl.Vector2
	Direction int32 // direction in degrees

	Velocity float32
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyRight) {
		p.Direction -= ROTATION_SPEED
		if p.Direction < 0 {
			p.Direction = 360
		}
	} else if rl.IsKeyDown(rl.KeyLeft) {
		p.Direction += ROTATION_SPEED
		if p.Direction > 360 {
			p.Direction = 0
		}
	}

	if rl.IsKeyDown(rl.KeyUp) {
		p.Velocity = float32(math.Min(float64(MAX_PLAYER_VELOCITY), float64(p.Velocity+PLAYER_ACCELERATION)))
	} else {
		p.Velocity = float32(math.Max(0, float64(p.Velocity-PLAYER_ACCELERATION*2)))
	}

	if p.Velocity > 0 {
		displace_x, displace_y := DisplacementComponents(p.Velocity, p.Direction)
		new_x := p.Center.X + float32(displace_x)
		new_y := p.Center.Y + float32(displace_y)

		if new_x < 0 {
			p.Center.X = 0
		} else {
			p.Center.X = float32(math.Min(float64(SCREEN_WIDTH), float64(new_x)))
		}

		if new_y < 0 {
			p.Center.Y = 0
		} else {
			p.Center.Y = float32(math.Min(float64(SCREEN_HEIGHT), float64(new_y)))
		}
	}
}

func DisplacementComponents(v float32, angle int32) (int32, int32) {
	x := int32(float32(math.Cos(DegToRad(angle))) * v)
	y := int32(float32(math.Sin(DegToRad(angle))) * v)

	return x * -1, y
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
	v1_x, v1_y := DerivePoints(0, float32(p.Height/2), p.Direction)
	v1 := rl.Vector2{
		X: p.Center.X + v1_x,
		Y: p.Center.Y + v1_y,
	}

	v2_x, v2_y := DerivePoints(float32(p.Width/2), -1*float32(p.Height/2), p.Direction)
	v2 := rl.Vector2{
		X: p.Center.X + v2_x,
		Y: p.Center.Y + v2_y,
	}

	v3_x, v3_y := DerivePoints(-1*float32(p.Width/2), -1*float32(p.Height/2), p.Direction)
	v3 := rl.Vector2{
		X: p.Center.X + v3_x,
		Y: p.Center.Y + v3_y,
	}

	rl.DrawTriangle(v1, v2, v3, rl.Red)
}

func DegToRad(deg int32) float64 {
	return float64(deg) * math.Pi / 180
}
