package main

import (
	"fmt"
	"gorp/gfx"
	"gorp/gobj"
	"gorp/phy"
	"gorp/utils"

	"github.com/veandco/go-sdl2/sdl"
)

const tileSize = 64

func generateTileFromCoords(i int32, j int32, texture *sdl.Texture) *Entity {
	sprite := gfx.Sprite{
		Texture:    texture,
		MaxFrames:  0,
		FrameIndex: 0,
		Animations: map[string]*gfx.Animation{
			"red": {
				StartFrame:     &sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
				AmountOfFrames: 1,
				FrameIndex:     0,
			},
			"blue": {
				StartFrame:     &sdl.Rect{X: 0, Y: 64, W: 64, H: 64},
				AmountOfFrames: 1,
				FrameIndex:     0,
			},
		},
		CurrentAnimation: "red",
	}
	if (i+j)%2 == 0 {
		sprite.CurrentAnimation = "blue"
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

func generateTileMap(tManager *gfx.TextureManager) map[string]*Entity {
	entities := make(map[string]*Entity)
	texture := tManager.Textures["grass"]
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

func loadEntities(tManager *gfx.TextureManager, gObjManager *gobj.GameObjectManager) map[string]*Entity {
	entities := make(map[string]*Entity)

	rect := &sdl.Rect{X: 0, Y: 0, W: 64, H: 64}
	sprite := gfx.Sprite{
		Texture:    tManager.Textures["player"],
		Frame:      rect,
		MaxFrames:  4,
		FrameIndex: 0,
		Animations: map[string]*gfx.Animation{
			"down": {
				StartFrame:     &sdl.Rect{X: 64, Y: 0, W: 64, H: 64},
				AmountOfFrames: 4,
				FrameIndex:     0,
			},
			"left": {
				StartFrame:     &sdl.Rect{X: 64, Y: 64, W: 64, H: 64},
				AmountOfFrames: 4,
				FrameIndex:     0,
			},
			"right": {
				StartFrame:     &sdl.Rect{X: 64, Y: 128, W: 64, H: 64},
				AmountOfFrames: 4,
				FrameIndex:     0,
			},
			"up": {
				StartFrame:     &sdl.Rect{X: 64, Y: 192, W: 64, H: 64},
				AmountOfFrames: 4,
				FrameIndex:     0,
			},
		},
		CurrentAnimation: "down",
	}
	gObj := gObjManager.Get("player")
	entity := Entity{sprite: &sprite, gObject: gObj}
	entities["player"] = &entity

	return entities
}

func loadParticle(entities map[string]*Entity, tManager *gfx.TextureManager, gObjManager *gobj.GameObjectManager) {
	gObj := gObjManager.Get("test")

	particle := &gfx.Particle{
		MaxFrames:  10,
		MaxSpeed:   int32(10),
		FrameIndex: 0,
		InitialPos: &utils.Vec2{},
		CurrentPos: utils.Vec2{},
	}
	rect := &sdl.Rect{X: 0, Y: 0, W: 64, H: 64}
	sprite := &gfx.Sprite{
		Frame:      rect,
		MaxFrames:  4,
		FrameIndex: 0,
		Animations: map[string]*gfx.Animation{
			"red": {
				StartFrame:     &sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
				AmountOfFrames: 1,
				FrameIndex:     0,
			},
		},
		CurrentAnimation: "red",
		Color:            &sdl.Color{R: 255, G: 0, B: 0, A: 255},
	}

	entity := Entity{entityType: 0, particle: particle, gObject: gObj, sprite: sprite}
	entities["test"] = &entity
}
