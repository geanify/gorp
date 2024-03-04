package main

import (
	"fmt"
	"gorp/gobj"
	"gorp/phy"
	"gorp/utils"

	"github.com/veandco/go-sdl2/sdl"
)

const tileSize = 64

func generateTileFromCoords(i int32, j int32, texture *sdl.Texture) *Entity {
	sprite := Sprite{
		texture:    texture,
		maxFrames:  0,
		frameIndex: 0,
		animations: map[string]*Animation{
			"red": {
				startFrame:     &sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
				amountOfFrames: 1,
				frameIndex:     0,
			},
			"blue": {
				startFrame:     &sdl.Rect{X: 0, Y: 64, W: 64, H: 64},
				amountOfFrames: 1,
				frameIndex:     0,
			},
		},
		currentAnimation: "red",
	}
	if (i+j)%2 == 0 {
		sprite.currentAnimation = "blue"
	}
	physics := phy.CreatePhyObject()
	if i >= 10 || j >= 10 {
		physics.Solid = true
	}
	gobj := &gobj.GameObject{
		Position: &utils.Vec2{X: tileSize * i, Y: tileSize * j},
		Size:     &utils.Vec2{X: 64, Y: 64},
		Physics:  physics,
	}
	entity := &Entity{sprite: &sprite, gObject: gobj}
	return entity
}

func generateTileMap(tManager *TextureManager) map[string]*Entity {
	entities := make(map[string]*Entity)
	texture := tManager.textures["grass"]
	for i := int32(0); i < 50; i++ {
		for j := int32(0); j < 50; j++ {
			textureName := fmt.Sprintf("z-texture-%d-%d", i, j)
			entity := generateTileFromCoords(i, j, texture)
			// entity.sprite.setTextureColorMode(255, 0, 0)
			entities[textureName] = entity
			// if i == 0 || j == 0 || i == 13 || j == 13 {
			// 	entities[textureName].physics = 1
			// }

		}
	}
	return entities
}

func loadEntities(tManager *TextureManager, gObjManager *gobj.GameObjectManager) map[string]*Entity {
	entities := make(map[string]*Entity)

	rect := &sdl.Rect{X: 0, Y: 0, W: 64, H: 64}
	sprite := Sprite{
		texture:    tManager.textures["player"],
		frame:      rect,
		maxFrames:  4,
		frameIndex: 0,
		animations: map[string]*Animation{
			"down": {
				startFrame:     &sdl.Rect{X: 64, Y: 0, W: 64, H: 64},
				amountOfFrames: 4,
				frameIndex:     0,
			},
			"left": {
				startFrame:     &sdl.Rect{X: 64, Y: 64, W: 64, H: 64},
				amountOfFrames: 4,
				frameIndex:     0,
			},
			"right": {
				startFrame:     &sdl.Rect{X: 64, Y: 128, W: 64, H: 64},
				amountOfFrames: 4,
				frameIndex:     0,
			},
			"up": {
				startFrame:     &sdl.Rect{X: 64, Y: 192, W: 64, H: 64},
				amountOfFrames: 4,
				frameIndex:     0,
			},
		},
		currentAnimation: "down",
	}
	gObj := gObjManager.Get("player")
	entity := Entity{sprite: &sprite, gObject: gObj}
	entities["player"] = &entity

	return entities
}
