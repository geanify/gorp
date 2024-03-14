package main

import (
	"gorp/gfx"
	"gorp/gobj"
	"gorp/utils"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type FogOfWar struct {
	timeControl *utils.TimeControl
	size        int
	fog         [][]*Entity
}

func CreateFogOfWar(size int) *FogOfWar {
	fog := make([][]*Entity, 0)
	fow := &FogOfWar{size: size, timeControl: utils.CreateTimeControl(), fog: fog}
	fow.GenerateFogOfWar()
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
		Color:      &sdl.Color{R: 0, G: 0, B: 0, A: 255},
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

	for i := 0; i < len(fow.fog); i++ {
		for j := 0; j < len(fow.fog[i]); j++ {
			entity := fow.fog[i][j]
			entity.sprite.Color = fow.GetColor(entity.gObject.Position, entities)
		}
	}
}

func (fow *FogOfWar) GenerateFogOfWar() {
	for i := int32(0); i < 50*(64/int32(fow.size)); i++ {
		fow.fog = append(fow.fog, make([]*Entity, 0))
		for j := int32(0); j < 50*(64/int32(fow.size)); j++ {
			entity := fow.GenerateFowFromTiles(i, j)
			fow.fog[i] = append(fow.fog[i], entity)
		}
	}
}
