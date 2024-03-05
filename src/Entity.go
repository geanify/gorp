package main

import (
	"gorp/gfx"
	"gorp/gobj"
	"gorp/utils"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/maps"
)

type Entity struct {
	entityType int32
	sprite     *gfx.Sprite
	text       *gfx.Text
	particle   *gfx.Particle
	gObject    *gobj.GameObject
	children   []*Entity
}

func (entity *Entity) getAdjustedPos(cam *utils.Camera, parentPosition *sdl.Rect) *sdl.Rect {
	if parentPosition == nil {
		return &sdl.Rect{
			X: entity.gObject.GetDistanceAdjustedPosition().X + cam.Position.X,
			Y: entity.gObject.GetDistanceAdjustedPosition().Y + cam.Position.Y,
			W: entity.gObject.GetDistanceAdjustedSize().X,
			H: entity.gObject.GetDistanceAdjustedSize().Y,
		}
	}
	return &sdl.Rect{
		X: entity.gObject.GetDistanceAdjustedPosition().X + parentPosition.X + cam.Position.X,
		Y: entity.gObject.GetDistanceAdjustedPosition().Y + parentPosition.Y + cam.Position.Y,
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

func (entity *Entity) renderSprite(renderer *sdl.Renderer, cam *utils.Camera, parentPosition *sdl.Rect) {
	renderer.Copy(entity.sprite.Texture, entity.sprite.GetFrame(), entity.getAdjustedPos(cam, parentPosition))
}

func (entity *Entity) renderText(renderer *sdl.Renderer, cam *utils.Camera, parentPosition *sdl.Rect) {
	entity.text.RenderText(renderer, entity.getAdjustedPos(cam, parentPosition))
}

func (entity *Entity) renderParticle(renderer *sdl.Renderer, cam *utils.Camera, parentPosition *sdl.Rect) {
	entity.particle.RenderParticle(
		renderer,
		entity.getAdjustedPos(cam, parentPosition),
		int32(entity.gObject.Physics.TerminalVelocity),
	)
}

func (entity *Entity) render(renderer *sdl.Renderer, cam *utils.Camera, parentPosition *sdl.Rect) {

	if !entity.shouldRender(cam) {
		return
	}

	switch entity.entityType {
	case 0:
		entity.renderSprite(renderer, cam, parentPosition)
	case 1:
		entity.renderText(renderer, cam, parentPosition)
	case 2:
		entity.renderParticle(renderer, cam, parentPosition)
	}

	for i := 0; i < len(entity.children); i++ {
		entity.children[i].render(renderer, cam, entity.getPosition())
	}
}

func renderEntities(entitiesMap map[string]*Entity, renderer *sdl.Renderer, cam *utils.Camera) {
	entities := maps.Values(entitiesMap)
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entity.render(renderer, cam, nil)
	}
}
