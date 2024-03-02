package main

import (
	"gorp/utils"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type MouseHandler struct {
	timeControl utils.TimeControl
}

func (mHandler *MouseHandler) handleCameraMove(cam *Camera) {

	if !mHandler.timeControl.ShouldExecute() {
		return
	}

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
		time.Sleep((tickRateMS / 3) * time.Millisecond)
	}
}

func createMouseHandler() *MouseHandler {
	return &MouseHandler{timeControl: *utils.CreateTimeControl()}
}
