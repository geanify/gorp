package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Entity struct {
	entityType int32
	sprite     *Sprite
	text       *Text
	position   *sdl.Rect
	speed      int32
}

func (entity *Entity) renderSprite(renderer *sdl.Renderer) {
	renderer.Copy(entity.sprite.texture, entity.sprite.frame, entity.position)
}

func (entity *Entity) renderText(renderer *sdl.Renderer) {
	entity.text.renderText(renderer, entity.position)
}

func (entity *Entity) render(renderer *sdl.Renderer) {
	switch entity.entityType {
	case 0:
		entity.renderSprite(renderer)
	case 1:
		entity.renderText(renderer)
	}

}

func renderEntities(entities []Entity, renderer *sdl.Renderer) {
	for i := 0; i < len(entities); i++ {
		entity := entities[i]

		entity.render(renderer)
	}
}

func (entity *Entity) move() {
	entity.position.X = entity.position.X + entity.speed
	entity.position.Y = entity.position.Y + entity.speed
}
