package main

import (
	"fmt"
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
		fmt.Println("edge")
		cam.moveRight()
	}
	if x > 780 {
		fmt.Println("edge")
		cam.moveLeft()
	}
	if y < 20 {
		fmt.Println("edge")
		cam.moveDown()
	}
	if y > 580 {
		fmt.Println("edge")
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
