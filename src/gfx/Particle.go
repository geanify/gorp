package gfx

import (
	"gorp/utils"

	"github.com/veandco/go-sdl2/sdl"
)

type Particle struct {
	MaxFrames  uint
	FrameIndex uint
	InitialPos *utils.Vec2
	CurrentPos *utils.Vec2
}

func (particle *Particle) RenderParticle(renderer *sdl.Renderer, pos *sdl.Rect) {

}
