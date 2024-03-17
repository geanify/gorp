package engine

import (
	"github.com/geanify/gorp/gfx"
	"github.com/geanify/gorp/gobj"
	"github.com/geanify/gorp/utils"
)

func createFPSCounter(fontPath string) *Entity {
	text := &gfx.Text{}

	text.SetFont(fontPath, 24)
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
