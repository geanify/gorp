package main

import "github.com/veandco/go-sdl2/sdl"

func createFPSCounter() Entity {
	text := &Text{}

	text.setFont("./FreeSans.ttf", 24)
	text.setColorRGB(255, 255, 255)
	text.setBackgroundColorRGB(0, 0, 0)
	text.setText("")
	fpsCounterPos := &sdl.Rect{500, 500, 100, 25}
	fpsCounterEntity := Entity{text: text, entityType: 1, position: fpsCounterPos}
	return fpsCounterEntity
}
