package gfx

import "github.com/veandco/go-sdl2/sdl"

type Animation struct {
	StartFrame     *sdl.Rect
	AmountOfFrames int
	FrameIndex     int
}

func (animation *Animation) nextFrame() {
	animation.FrameIndex++

	if animation.FrameIndex >= animation.AmountOfFrames {
		animation.FrameIndex = 0
	}
}

func (animation *Animation) getFrame() *sdl.Rect {
	return &sdl.Rect{
		X: animation.StartFrame.X * int32(animation.FrameIndex),
		Y: animation.StartFrame.Y,
		W: animation.StartFrame.W,
		H: animation.StartFrame.H,
	}
}
