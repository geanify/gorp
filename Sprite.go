package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	texture          *sdl.Texture
	frame            *sdl.Rect
	maxFrames        uint
	animations       map[string]*sdl.Rect
	frameIndex       uint
	currentAnimation string
}

func (sprite *Sprite) nextFrame() {
	sprite.frameIndex++

	if sprite.frameIndex >= sprite.maxFrames {
		sprite.frameIndex = 0
	}
}

func (sprite *Sprite) setAnimation(animationName string) {
	sprite.currentAnimation = animationName
}

func (sprite *Sprite) getFrame() *sdl.Rect {
	currentAnimation := sprite.animations[sprite.currentAnimation]
	animationFrame := &sdl.Rect{X: currentAnimation.X * int32(sprite.frameIndex), Y: currentAnimation.Y, W: 64, H: 64}
	return animationFrame
}
