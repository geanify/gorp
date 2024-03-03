package main

import "github.com/veandco/go-sdl2/sdl"

type Animation struct {
	startFrame     *sdl.Rect
	amountOfFrames int
	frameIndex     int
}

func (animation *Animation) nextFrame() {
	animation.frameIndex++

	if animation.frameIndex >= animation.amountOfFrames {
		animation.frameIndex = 0
	}
}

func (animation *Animation) getFrame() *sdl.Rect {
	return &sdl.Rect{
		X: animation.startFrame.X * int32(animation.frameIndex),
		Y: animation.startFrame.Y,
		W: animation.startFrame.W,
		H: animation.startFrame.H,
	}
}
