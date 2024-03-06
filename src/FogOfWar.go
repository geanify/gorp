package main

import (
	"fmt"
	"gorp/gfx"
	"gorp/gobj"
	"gorp/utils"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func generatePosition(i int32, j int32) *utils.Vec2 {
	return &utils.Vec2{X: tileSize * i, Y: tileSize * j}
}

func calculateDistance(pos1 *utils.Vec2, pos2 *utils.Vec2) int {

	return int(math.Abs(float64(pos1.X-pos2.X)) + math.Abs(float64(pos1.Y-pos2.Y)))
}

func getColor(pos *utils.Vec2, entities map[string]*Entity) *sdl.Color {
	alpha := uint8(125)

	for _, entity := range entities {
		distance := calculateDistance(pos, entity.gObject.Position)
		if entity.gObject.Physics == nil {
			continue
		}
		if distance < entity.gObject.Physics.LightCastDistance*64 {
			alpha = 0
		}
	}

	return &sdl.Color{R: 0, G: 0, B: 0, A: alpha}
}

func generateFowFromTiles(i int32, j int32, entities map[string]*Entity) *Entity {

	position := generatePosition(i, j)
	childSprite := gfx.Sprite{
		MaxFrames:  0,
		FrameIndex: 0,
		Color:      getColor(position, entities),
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
		Position: position,
		Size:     &utils.Vec2{X: 64, Y: 64},
	}
	entity := &Entity{
		sprite:  &childSprite,
		gObject: gobj2,
	}
	return entity
}

func updateFogOfWar(fowEntities map[string]*Entity, entities map[string]*Entity) {
	for _, entity := range fowEntities {
		entity.sprite.Color = getColor(entity.gObject.Position, entities)
	}
}

func generateFogOfWar(entities map[string]*Entity) map[string]*Entity {
	fowEntities := make(map[string]*Entity)
	for i := int32(0); i < 50; i++ {
		for j := int32(0); j < 50; j++ {
			textureName := fmt.Sprintf("z-fow-%d-%d", i, j)
			entity := generateFowFromTiles(i, j, entities)
			fowEntities[textureName] = entity

		}
	}
	return fowEntities
}
