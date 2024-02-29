package main

import "github.com/veandco/go-sdl2/sdl"

type Camera struct {
	position *sdl.Rect
}

func createCamera() *Camera {
	return &Camera{position: &sdl.Rect{X: 0, Y: 0, W: 800, H: 600}}
}

func (cam *Camera) invertedPosition() *sdl.Rect {
	return &sdl.Rect{X: (-1) * cam.position.X, Y: (-1) * cam.position.Y, W: 800, H: 600}
}

func (cam *Camera) moveRight() {
	cam.position.X += 10
}

func (cam *Camera) moveLeft() {
	cam.position.X -= 10
}

func (cam *Camera) moveDown() {
	cam.position.Y += 10
}

func (cam *Camera) moveUp() {
	cam.position.Y -= 10
}
