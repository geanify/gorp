package utils

import "github.com/veandco/go-sdl2/sdl"

type Camera struct {
	Position *sdl.Rect
}

func CreateCamera() *Camera {
	return &Camera{Position: &sdl.Rect{X: 0, Y: 0, W: 800, H: 600}}
}

func (cam *Camera) InvertedPosition() *sdl.Rect {
	return &sdl.Rect{X: (-1) * cam.Position.X, Y: (-1) * cam.Position.Y, W: 800, H: 600}
}

func (cam *Camera) MoveRight() {
	cam.Position.X += 10
}

func (cam *Camera) MoveLeft() {
	cam.Position.X -= 10
}

func (cam *Camera) MoveDown() {
	cam.Position.Y += 10
}

func (cam *Camera) MoveUp() {
	cam.Position.Y -= 10
}
