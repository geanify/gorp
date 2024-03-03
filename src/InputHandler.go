package main

import (
	"gorp/utils"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const tickRateMS = 100 //miliseconds

type InputHandler struct {
	keyboardState []uint8
	start         time.Time
	timeControl   *utils.TimeControl
}

func (iHandler *InputHandler) isKeyPressed(key int) bool {
	return iHandler.keyboardState[key] == 1
}

func (iHandler *InputHandler) handleInput(entitiesMap map[string]*Entity) {
	if !iHandler.timeControl.ShouldExecute() {
		return
	}
	elapsed := iHandler.timeControl.GetElapsed()

	iHandler.keyboardState = sdl.GetKeyboardState()

	player := entitiesMap["player"]

	if iHandler.isKeyPressed(sdl.SCANCODE_A) {
		player.sprite.nextFrame()
		player.moveLeft(elapsed)
		player.sprite.setAnimation("left")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_D) {
		player.sprite.nextFrame()
		player.moveRight(elapsed)
		player.sprite.setAnimation("right")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_W) {
		player.sprite.nextFrame()
		player.moveUp(elapsed)
		player.sprite.setAnimation("up")
	} else {
		// player.sprite.nextFrame()
		// player.moveDown(elapsed)
		// player.sprite.setAnimation("down")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_S) {
		player.sprite.nextFrame()
		player.moveDown(elapsed)
		player.sprite.setAnimation("down")
	}
}

func createInputHandler() *InputHandler {
	return &InputHandler{keyboardState: sdl.GetKeyboardState(), start: time.Now(), timeControl: utils.CreateTimeControl()}
}

// func handleInput(entities map[string]*Entity, iHandler *InputHandler) {
// 	for {
// 		iHandler.handleInput(entities)
// 		time.Sleep((tickRateMS / 3) * time.Millisecond)
// 	}
// }
