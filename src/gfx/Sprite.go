package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	Texture          *sdl.Texture
	Frame            *sdl.Rect
	MaxFrames        uint
	Animations       map[string]*Animation
	FrameIndex       uint
	CurrentAnimation string
	Color            *sdl.Color
}

func (sprite *Sprite) NextFrame() {
	if sprite.Animations == nil {
		return
	}
	currentAnimation := sprite.Animations[sprite.CurrentAnimation]
	currentAnimation.nextFrame()
}

func (sprite *Sprite) SetAnimation(animationName string) {
	sprite.CurrentAnimation = animationName
}

func (sprite *Sprite) GetFrame() *sdl.Rect {
	currentAnimation := sprite.Animations[sprite.CurrentAnimation]
	return currentAnimation.getFrame()
}

func (sprite *Sprite) SetTextureColorMode(r uint8, g uint8, b uint8) {
	sprite.Texture.SetColorMod(r, g, b)
}

func (sprite *Sprite) SetBlendModeAdd() {
	sprite.Texture.SetBlendMode(sdl.BLENDMODE_ADD)
}

func (sprite *Sprite) SetBlendModeBlend() {
	sprite.Texture.SetBlendMode(sdl.BLENDMODE_BLEND)
}

func (sprite *Sprite) SetBlendModeMod() {
	sprite.Texture.SetBlendMode(sdl.BLENDMODE_MOD)
}

func (sprite *Sprite) RenderColor(renderer *sdl.Renderer, pos *sdl.Rect) {
	if sprite.Color == nil {
		return
	}
	renderer.SetDrawColor(sprite.Color.R, sprite.Color.G, sprite.Color.B, sprite.Color.A)
	renderer.FillRect(pos)
	renderer.SetDrawColor(0, 0, 0, 255)
}
