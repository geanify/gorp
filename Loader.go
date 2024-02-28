package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const tileSize = 64

func generateTileFromCoords(i int32, j int32, texture *sdl.Texture) *Entity {
	sprite := Sprite{
		texture:    texture,
		maxFrames:  0,
		frameIndex: 0,
		animations: map[string]*sdl.Rect{
			"red":  {X: 0, Y: 0, W: 64, H: 64},
			"blue": {X: 0, Y: 64, W: 64, H: 64},
		},
		currentAnimation: "red",
	}
	if (i+j)%2 == 0 {
		sprite.currentAnimation = "blue"
	}
	pos := &sdl.Rect{X: tileSize * i, Y: tileSize * j, W: 64, H: 64}
	entity := &Entity{sprite: &sprite, position: pos, speed: 0}
	return entity
}

func generateTileMap(renderer *sdl.Renderer) map[string]*Entity {
	entities := make(map[string]*Entity)
	texture := loadImageAsTexture("assets/tileset-grass.png", renderer)
	for i := int32(0); i < 50; i++ {
		for j := int32(0); j < 50; j++ {
			textureName := fmt.Sprintf("z-texture-%d-%d", i, j)
			entities[textureName] = generateTileFromCoords(i, j, texture)
			if i == 0 || j == 0 || i == 13 || j == 13 {
				entities[textureName].physics = 1
			}
		}
	}
	return entities
}

func loadEntities(texture *sdl.Texture, renderer *sdl.Renderer) map[string]*Entity {
	entities := make(map[string]*Entity)

	rect := &sdl.Rect{X: 0, Y: 0, W: 64, H: 64}
	sprite := Sprite{
		texture:    texture,
		frame:      rect,
		maxFrames:  4,
		frameIndex: 0,
		animations: map[string]*sdl.Rect{
			"down":  {X: 64, Y: 0, W: 64, H: 64},
			"left":  {X: 64, Y: 64, W: 64, H: 64},
			"right": {X: 64, Y: 128, W: 64, H: 64},
			"up":    {X: 64, Y: 192, W: 64, H: 64},
		},
		currentAnimation: "down",
	}
	pos := &sdl.Rect{X: 100, Y: 100, W: 64, H: 64}
	entity := Entity{sprite: &sprite, position: pos, speed: 15, physics: 1}
	entities["player"] = &entity

	// entity2 := Entity{sprite: &sprite, position: rect, speed: 1}
	// entities = append(entities, entity2)

	return entities
}
