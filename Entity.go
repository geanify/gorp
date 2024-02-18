package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Entity struct {
	sprite   *Sprite
	position *sdl.Rect
}

func (entity *Entity) render(renderer *sdl.Renderer) {
	renderer.Copy(entity.sprite.texture, entity.sprite.frame, entity.position)
}
