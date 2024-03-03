package main

import (
	"gorp/gobj"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/maps"
)

type Entity struct {
	entityType int32
	sprite     *Sprite
	text       *Text
	gObject    *gobj.GameObject
}

func (entity *Entity) getAdjustedPos(cam *Camera) *sdl.Rect {
	return &sdl.Rect{
		X: entity.gObject.Position.X + cam.position.X,
		Y: entity.gObject.Position.Y + cam.position.Y,
		W: entity.gObject.Size.X,
		H: entity.gObject.Size.Y,
	}
}

func (entity *Entity) getPosition() *sdl.Rect {
	return &sdl.Rect{
		X: entity.gObject.Position.X,
		Y: entity.gObject.Position.Y,
		W: entity.gObject.Size.X,
		H: entity.gObject.Size.Y,
	}
}

func (entity *Entity) shouldRender(cam *Camera) bool {
	return entity.getPosition().HasIntersection(cam.invertedPosition())
}

func (entity *Entity) renderSprite(renderer *sdl.Renderer, cam *Camera) {
	renderer.Copy(entity.sprite.texture, entity.sprite.getFrame(), entity.getAdjustedPos(cam))
}

func (entity *Entity) renderText(renderer *sdl.Renderer, cam *Camera) {
	entity.text.renderText(renderer, entity.getAdjustedPos(cam))
}

func (entity *Entity) render(renderer *sdl.Renderer, cam *Camera) {

	if !entity.shouldRender(cam) {
		return
	}

	switch entity.entityType {
	case 0:
		entity.renderSprite(renderer, cam)
	case 1:
		entity.renderText(renderer, cam)
	}

}

func renderEntities(entitiesMap map[string]*Entity, renderer *sdl.Renderer, cam *Camera) {
	entities := maps.Values(entitiesMap)
	for i := 0; i < len(entities); i++ {
		entity := entities[i]

		entity.render(renderer, cam)
	}
}

func (entity *Entity) moveLeft(elapsed time.Duration) {
	entity.gObject.MoveLeft()
}

func (entity *Entity) moveRight(elapsed time.Duration) {
	entity.gObject.MoveRight()
}

func (entity *Entity) moveUp(elapsed time.Duration) {
	entity.gObject.MoveUp()
}

func (entity *Entity) moveDown(elapsed time.Duration) {
	entity.gObject.MoveDown()
}
