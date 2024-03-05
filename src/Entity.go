package main

import (
	"gorp/gobj"
	"gorp/utils"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/maps"
)

type Entity struct {
	entityType int32
	sprite     *Sprite
	text       *Text
	gObject    *gobj.GameObject
}

func (entity *Entity) getAdjustedPos(cam *utils.Camera) *sdl.Rect {
	return &sdl.Rect{
		X: entity.gObject.GetDistanceAdjustedPosition().X + cam.Position.X,
		Y: entity.gObject.GetDistanceAdjustedPosition().Y + cam.Position.Y,
		W: entity.gObject.GetDistanceAdjustedSize().X,
		H: entity.gObject.GetDistanceAdjustedSize().Y,
	}
}

func (entity *Entity) getPosition() *sdl.Rect {
	return &sdl.Rect{
		X: entity.gObject.GetDistanceAdjustedPosition().X,
		Y: entity.gObject.GetDistanceAdjustedPosition().Y,
		W: entity.gObject.GetDistanceAdjustedSize().X,
		H: entity.gObject.GetDistanceAdjustedSize().Y,
	}
}

func (entity *Entity) shouldRender(cam *utils.Camera) bool {
	return entity.getPosition().HasIntersection(cam.InvertedPosition())
}

func (entity *Entity) renderSprite(renderer *sdl.Renderer, cam *utils.Camera) {
	renderer.Copy(entity.sprite.texture, entity.sprite.getFrame(), entity.getAdjustedPos(cam))
}

func (entity *Entity) renderText(renderer *sdl.Renderer, cam *utils.Camera) {
	entity.text.renderText(renderer, entity.getAdjustedPos(cam))
}

func (entity *Entity) render(renderer *sdl.Renderer, cam *utils.Camera) {

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

func renderEntities(entitiesMap map[string]*Entity, renderer *sdl.Renderer, cam *utils.Camera) {
	entities := maps.Values(entitiesMap)
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entity.render(renderer, cam)
	}
}
