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

func renderEntities(entities []Entity, renderer *sdl.Renderer) {
	for i := 0; i < len(entities); i++ {
		entity := entities[i]

		entity.render(renderer)
	}

}
