package main

import (
	"gorp/gobj"
	"gorp/utils"
)

func createFPSCounter() *Entity {
	text := &Text{}

	text.setFont("../assets/font/FreeSans.ttf", 24)
	text.setColorRGB(255, 255, 255)
	text.setBackgroundColorRGBA(0, 0, 0, 150)
	text.setText("")
	gobj := &gobj.GameObject{
		Position: &utils.Vec2{X: 500, Y: 500},
		Size:     &utils.Vec2{X: 100, Y: 25},
	}
	fpsCounterEntity := &Entity{text: text, entityType: 1, gObject: gobj}
	return fpsCounterEntity
}
