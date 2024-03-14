package main

import (
	"gorp/gfx"
	"gorp/gobj"
	"gorp/utils"
	"math"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/maps"
)

type FogOfWar struct {
	timeControl *utils.TimeControl
	size        int32
	fog         [][]*Entity
}

func CreateFogOfWar(size int32) *FogOfWar {
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

func (fow *FogOfWar) GetColor(pos *utils.Vec2, entity *Entity) *sdl.Color {
	distance := fow.CalculateDistance(pos, entity.gObject.Position)
	if entity.gObject.Physics == nil {
		return nil
	}
	lightCastDistanceUnits := entity.gObject.Physics.LightCastDistance
	distanceUnits := int32(distance) / fow.size
	if int32(distance) < lightCastDistanceUnits*fow.size {

		alpha := uint8(distanceUnits * (125) / lightCastDistanceUnits)
		return &sdl.Color{R: 0, G: 0, B: 0, A: alpha}
	}

	return nil
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

func getBrightestColor(colorA *sdl.Color, colorB *sdl.Color) *sdl.Color {
	if colorA.A < colorB.A {
		return colorA
	}
	return colorB
}

func (fow *FogOfWar) getTileEntity(pos *utils.Vec2) *Entity {
	i := pos.X / int32(fow.size)
	j := pos.Y / int32(fow.size)

	return fow.fog[i][j]
}

func (fow *FogOfWar) GenerateLightAroundPosition(entity *Entity) {

	if entity.gObject.Physics == nil {
		return
	}
	lightCastDistance := entity.gObject.Physics.LightCastDistance
	for i := -lightCastDistance; i < lightCastDistance; i++ {
		for j := -lightCastDistance; j < lightCastDistance; j++ {

			posX := entity.gObject.Position.X + i*fow.size
			posY := entity.gObject.Position.Y + j*fow.size
			pos := &utils.Vec2{X: posX, Y: posY}

			if !pos.IsPositiveOrZero() {
				continue
			}

			tileEntity := fow.getTileEntity(pos)

			newColor := fow.GetColor(tileEntity.gObject.Position, entity)
			if newColor != nil {
				tileEntity.sprite.Color = getBrightestColor(tileEntity.sprite.Color, newColor)
			}
		}
	}
}

func (fow *FogOfWar) UpdateFogOfWar(entitiesMap map[string]*Entity) {
	if !fow.timeControl.ShouldExecute() {
		return
	}

	for i := 0; i < len(fow.fog); i++ {
		for j := 0; j < len(fow.fog[i]); j++ {
			entity := fow.fog[i][j]
			entity.sprite.Color = &sdl.Color{R: 0, G: 0, B: 0, A: 125}
		}
	}

	entities := maps.Values(entitiesMap)
	for _, entity := range entities {
		fow.GenerateLightAroundPosition(entity)
	}
}

func (fow *FogOfWar) GenerateFogOfWar() {
	for i := int32(0); i < 50*(64/int32(fow.size)); i++ {
		entityRow := make([]*Entity, 0)
		fow.fog = append(fow.fog, entityRow)
		for j := int32(0); j < 50*(64/int32(fow.size)); j++ {
			entity := fow.GenerateFowFromTiles(i, j)
			fow.fog[i] = append(fow.fog[i], entity)
		}
	}
}
