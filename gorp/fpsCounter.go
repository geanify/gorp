package main

import (
	"gorp/gfx"
	"gorp/gobj"
	"gorp/utils"
)

func createFPSCounter() *Entity {
	text := &gfx.Text{}

	text.SetFont("../assets/font/FreeSans.ttf", 24)
	text.SetColorRGB(255, 255, 255)
	text.SetBackgroundColorRGBA(0, 0, 0, 150)
	text.SetText("")
	gobj := &gobj.GameObject{
		Position: &utils.Vec2{X: 500, Y: 500},
		Size:     &utils.Vec2{X: 100, Y: 25},
	}
	fpsCounterEntity := &Entity{text: text, entityType: 1, gObject: gobj}
	return fpsCounterEntity
}
