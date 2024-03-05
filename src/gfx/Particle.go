package gfx

import (
	"gorp/utils"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Particle struct {
	Respawn    bool
	stopRender bool
	MaxFrames  uint
	MaxSpeed   int32
	FrameIndex uint
	InitialPos *utils.Vec2
	CurrentPos utils.Vec2
}

func getRandomized(n int32) int32 {
	a := rand.Int31n(n)
	neg := rand.Int31n(2)

	if neg == 1 {
		return (-1) * a
	}
	return a
}

func (particle *Particle) GetNextPos() {
	if particle.FrameIndex >= particle.MaxFrames {
		if particle.Respawn {
			particle.CurrentPos = *particle.InitialPos
			particle.FrameIndex = 0
			return
		} else {
			particle.stopRender = true
			return
		}
	}

	x := getRandomized(particle.MaxSpeed)
	y := getRandomized(particle.MaxSpeed)

	particle.CurrentPos.X += x
	particle.CurrentPos.Y += y

	particle.FrameIndex++
}

func (particle *Particle) GetNextFrame() {
	particle.GetNextPos()
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
	if particle.stopRender {
		return
	}
	particle.MaxSpeed = maxSpeed
	particle.GetNextFrame()
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.FillRect(particle.getAdjustedPos(pos))
	renderer.SetDrawColor(0, 0, 0, 255)
}
