package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	texture          *sdl.Texture
	frame            *sdl.Rect
	maxFrames        uint
	animations       map[string]*Animation
	frameIndex       uint
	currentAnimation string
}

func (sprite *Sprite) nextFrame() {
	currentAnimation := sprite.animations[sprite.currentAnimation]
	currentAnimation.nextFrame()
}

func (sprite *Sprite) setAnimation(animationName string) {
	sprite.currentAnimation = animationName
}

func (sprite *Sprite) getFrame() *sdl.Rect {
	currentAnimation := sprite.animations[sprite.currentAnimation]
	return currentAnimation.getFrame()
}
