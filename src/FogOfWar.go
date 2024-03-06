package main

import (
	"fmt"
	"gorp/gfx"
	"gorp/gobj"
	"gorp/utils"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type FogOfWar struct {
	timeControl *utils.TimeControl
	size        int
	tiles       map[string]*Entity
	fog         map[string]*Entity
}

func CreateFogOfWar(size int, tileMap map[string]*Entity) *FogOfWar {
	fow := &FogOfWar{size: size, timeControl: utils.CreateTimeControl()}
	fow.GenerateFogOfWar(tileMap)
	return fow
}

func (fow *FogOfWar) GeneratePosition(i int32, j int32) *utils.Vec2 {
	return &utils.Vec2{X: int32(fow.size) * i, Y: int32(fow.size) * j}
}

func (fow *FogOfWar) CalculateDistance(pos1 *utils.Vec2, pos2 *utils.Vec2) int {
	deltaX := int(math.Abs(float64(pos1.X - pos2.X)))
	deltaY := int(math.Abs(float64(pos1.Y - pos2.Y)))
	return deltaX + deltaY
}

func (fow *FogOfWar) GetColor(pos *utils.Vec2, entities map[string]*Entity) *sdl.Color {
	alpha := uint8(125)

	for _, entity := range entities {
		distance := fow.CalculateDistance(pos, entity.gObject.Position)
		if entity.gObject.Physics == nil {
			continue
		}
		if distance < entity.gObject.Physics.LightCastDistance*fow.size {
			alpha = 0
		}
	}

	return &sdl.Color{R: 0, G: 0, B: 0, A: alpha}
}

func (fow *FogOfWar) GenerateFowFromTiles(i int32, j int32) *Entity {

	position := fow.GeneratePosition(i, j)
	childSprite := gfx.Sprite{
		MaxFrames:  0,
		FrameIndex: 0,
		Color:      fow.GetColor(position, fow.tiles),
		Animations: map[string]*gfx.Animation{
			"red": {
				StartFrame:     &sdl.Rect{X: 0, Y: 0, W: int32(fow.size), H: int32(fow.size)},
				AmountOfFrames: 1,
				FrameIndex:     0,
			},
		},
		CurrentAnimation: "red",
	}

	gobj2 := &gobj.GameObject{
		Position: position,
		Size:     &utils.Vec2{X: int32(fow.size), Y: int32(fow.size)},
	}
	entity := &Entity{
		sprite:  &childSprite,
		gObject: gobj2,
	}
	return entity
}

func (fow *FogOfWar) UpdateFogOfWar(entities map[string]*Entity) {
	if !fow.timeControl.ShouldExecute() {
		return
	}
	for _, entity := range fow.fog {
		entity.sprite.Color = fow.GetColor(entity.gObject.Position, entities)
	}
}

func (fow *FogOfWar) GenerateFogOfWar(tiles map[string]*Entity) {
	fow.tiles = tiles
	fowEntities := make(map[string]*Entity)
	for i := int32(0); i < (50*64)/int32(fow.size); i++ {
		for j := int32(0); j < (50*64)/int32(fow.size); j++ {
			textureName := fmt.Sprintf("z-fow-%d-%d", i, j)
			entity := fow.GenerateFowFromTiles(i, j)
			fowEntities[textureName] = entity

		}
	}
	fow.fog = fowEntities
}
