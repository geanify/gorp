package gfx

import (
	"gorp/utils"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Particle struct {
	MaxFrames  uint
	MaxSpeed   int32
	FrameIndex uint
	InitialPos *utils.Vec2
	CurrentPos utils.Vec2
}

func (particle *Particle) GetNextFrame() {
	if particle.FrameIndex >= particle.MaxFrames {
		particle.CurrentPos = *particle.InitialPos
		particle.FrameIndex = 0
		return
	}

	x := rand.Int31n(particle.MaxSpeed)
	y := rand.Int31n(particle.MaxSpeed)

	particle.CurrentPos.X += x
	particle.CurrentPos.Y += y

	particle.FrameIndex++
}

func (particle *Particle) getAdjustedPos(pos *sdl.Rect) *sdl.Rect {
	return &sdl.Rect{
		X: pos.X + particle.CurrentPos.X,
		Y: pos.Y + particle.CurrentPos.Y,
		W: pos.W,
		H: pos.H,
	}
}

func (particle *Particle) RenderParticle(renderer *sdl.Renderer, pos *sdl.Rect, maxSpeed int32) {
	particle.MaxSpeed = maxSpeed
	particle.GetNextFrame()

	renderer.FillRect(particle.getAdjustedPos(pos))
}
