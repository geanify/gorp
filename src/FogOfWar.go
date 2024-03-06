package main

import (
	"fmt"
	"gorp/gfx"
	"gorp/gobj"
	"gorp/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func generateFowFromTiles(i int32, j int32) *Entity {

	childSprite := gfx.Sprite{
		MaxFrames:  0,
		FrameIndex: 0,
		Color:      &sdl.Color{R: 0, G: 0, B: 0, A: 125},
		Animations: map[string]*gfx.Animation{
			"red": {
				StartFrame:     &sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
				AmountOfFrames: 1,
				FrameIndex:     0,
			},
		},
		CurrentAnimation: "red",
	}

	gobj2 := &gobj.GameObject{
		Position: &utils.Vec2{X: tileSize * i, Y: tileSize * j},
		Size:     &utils.Vec2{X: 64, Y: 64},
	}
	entity := &Entity{
		sprite:  &childSprite,
		gObject: gobj2,
	}
	return entity
}

func generateFogOfWar() map[string]*Entity {
	entities := make(map[string]*Entity)
	for i := int32(0); i < 50; i++ {
		for j := int32(0); j < 50; j++ {
			textureName := fmt.Sprintf("z-fow-%d-%d", i, j)
			entity := generateFowFromTiles(i, j)
			entities[textureName] = entity

		}
	}
	return entities
}
