package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	LASER_SPEED  = 3
	LASER_WIDTH  = 10
	LASER_HEIGHT = 20
)

type Laser struct {
	Center    rl.Vector2
	Direction int32
}

func (laser *Laser) Update() bool {
	displace_x, displace_y := DisplacementComponents(LASER_SPEED, laser.Direction)

	laser.Center.X += displace_x
	laser.Center.Y += displace_y

	if laser.Center.X-LASER_WIDTH/2 < 0 || laser.Center.X+LASER_WIDTH/2 > float32(SCREEN_WIDTH) {
		return true
	} else if laser.Center.Y-LASER_HEIGHT/2 < 0 || laser.Center.Y+LASER_HEIGHT/2 > float32(SCREEN_HEIGHT) {
		return true
	} else {
		return false
	}
}

func (laser *Laser) Draw() {
	laser_rec := rl.Rectangle{
		X:      laser.Center.X - LASER_WIDTH/2,
		Y:      laser.Center.Y - LASER_HEIGHT/2,
		Width:  LASER_WIDTH,
		Height: LASER_HEIGHT,
	}
	direction := float32(laser.Direction - 270 + 45)

	rl.DrawRectanglePro(laser_rec, rl.NewVector2(0, 0), direction, rl.Red)
}

func ShootLaser() {
	new_laser := Laser{
		Center:    rl.Vector2{X: PLAYER.Center.X, Y: PLAYER.Center.Y},
		Direction: PLAYER.Direction,
	}
	LASERS = append(LASERS, new_laser)
}

func VaporizeLaser(index int) {
	LASERS = append(LASERS[:index], LASERS[index+1:]...)
}
