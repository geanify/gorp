package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type MouseHandler struct {
	start time.Time
}

func (mHandler *MouseHandler) handleCameraMove(cam *Camera) {
	now := time.Now()
	elapsed := now.Sub(mHandler.start)

	if elapsed.Milliseconds() < tickRateMS {
		return
	}

	mHandler.start = now

	x, y, _ := sdl.GetMouseState()

	if x < 20 {
		cam.moveRight()
	}
	if x > 780 {
		cam.moveLeft()
	}
	if y < 20 {
		cam.moveDown()
	}
	if y > 580 {
		cam.moveUp()
	}

}

func handleMouse(mHandler *MouseHandler, cam *Camera) {
	for {
		mHandler.handleCameraMove(cam)
	}
}

func createMouseHandler() *MouseHandler {
	return &MouseHandler{start: time.Now()}
}
