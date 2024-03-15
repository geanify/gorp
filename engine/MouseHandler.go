package engine

import (
	"github.com/geanify/gorp/utils"

	"github.com/veandco/go-sdl2/sdl"
)

type MouseHandler struct {
	timeControl utils.TimeControl
}

func (mHandler *MouseHandler) handleCameraMove(cam *utils.Camera) {

	if !mHandler.timeControl.ShouldExecute() {
		return
	}

	x, y, _ := sdl.GetMouseState()

	if x < 20 {
		cam.MoveRight()
	}
	if x > 780 {
		cam.MoveLeft()
	}
	if y < 20 {
		cam.MoveDown()
	}
	if y > 580 {
		cam.MoveUp()
	}

}

func handleMouse(mHandler *MouseHandler, cam *utils.Camera) {
	for {
		mHandler.handleCameraMove(cam)
	}
}

func createMouseHandler() *MouseHandler {
	return &MouseHandler{timeControl: *utils.CreateTimeControl()}
}
