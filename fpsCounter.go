package main

import "github.com/veandco/go-sdl2/sdl"

func createFPSCounter() *Entity {
	text := &Text{}

	text.setFont("./FreeSans.ttf", 24)
	text.setColorRGB(255, 255, 255)
	text.setBackgroundColorRGBA(0, 0, 0, 150)
	text.setText("")
	fpsCounterPos := &sdl.Rect{X: 500, Y: 500, W: 100, H: 25}
	fpsCounterEntity := &Entity{text: text, entityType: 1, position: fpsCounterPos}
	return fpsCounterEntity
}
