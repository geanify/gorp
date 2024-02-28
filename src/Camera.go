package main

import "github.com/veandco/go-sdl2/sdl"

type Camera struct {
	position *sdl.Rect
}

func createCamera() *Camera {
	return &Camera{position: &sdl.Rect{0, 0, 600, 800}}
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
