package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const tickRateMS = 100 //miliseconds

type InputHandler struct {
	keyboardState []uint8
	start         time.Time
}

func (iHandler *InputHandler) isKeyPressed(key int) bool {
	return iHandler.keyboardState[key] == 1
}

func (iHandler *InputHandler) handleInput(entitiesMap map[string]*Entity) {
	now := time.Now()
	elapsed := time.Now().Sub(iHandler.start)

	if elapsed.Milliseconds() < tickRateMS {
		return
	}

	iHandler.keyboardState = sdl.GetKeyboardState()
	iHandler.start = now

	player := entitiesMap["player"]

	if iHandler.isKeyPressed(sdl.SCANCODE_A) {
		player.sprite.nextFrame()
		player.moveLeft(elapsed)
		player.sprite.setAnimation("left")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_D) {
		player.moveRight(elapsed)
		player.sprite.nextFrame()
		player.sprite.setAnimation("right")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_W) {
		player.moveUp(elapsed)
		player.sprite.nextFrame()
		player.sprite.setAnimation("up")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_S) {
		player.moveDown(elapsed)
		player.sprite.nextFrame()
		player.sprite.setAnimation("down")
	}
}

func createInputHandler() *InputHandler {
	return &InputHandler{keyboardState: sdl.GetKeyboardState(), start: time.Now()}
}

func handleInput(entities map[string]*Entity, iHandler *InputHandler) {
	iHandler.handleInput(entities)
}
