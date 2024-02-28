package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/maps"
)

type Entity struct {
	entityType int32
	sprite     *Sprite
	text       *Text
	position   *sdl.Rect
	speed      int32 //per tickrate
	physics    int
}

func (entity *Entity) getAdjustedPos(cam *Camera) *sdl.Rect {
	return &sdl.Rect{
		X: entity.position.X + cam.position.X,
		Y: entity.position.Y + cam.position.Y,
		W: entity.position.W,
		H: entity.position.H,
	}
}

func (entity *Entity) renderSprite(renderer *sdl.Renderer, cam *Camera) {
	renderer.Copy(entity.sprite.texture, entity.sprite.getFrame(), entity.getAdjustedPos(cam))
}

func (entity *Entity) renderText(renderer *sdl.Renderer, cam *Camera) {
	entity.text.renderText(renderer, entity.getAdjustedPos(cam))
}

func (entity *Entity) render(renderer *sdl.Renderer, cam *Camera) {
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
	entity.position.X -= entity.speed
}

func (entity *Entity) moveRight(elapsed time.Duration) {
	entity.position.X += entity.speed
}

func (entity *Entity) moveUp(elapsed time.Duration) {
	entity.position.Y -= entity.speed
}

func (entity *Entity) moveDown(elapsed time.Duration) {
	entity.position.Y += entity.speed
}
