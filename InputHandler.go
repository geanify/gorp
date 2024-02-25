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

	if iHandler.isKeyPressed(sdl.SCANCODE_A) {
		entitiesMap["player"].moveLeft(elapsed)
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_D) {
		entitiesMap["player"].moveRight(elapsed)
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_W) {
		entitiesMap["player"].moveUp(elapsed)
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_S) {
		entitiesMap["player"].moveDown(elapsed)
	}
}

func createInputHandler() *InputHandler {
	return &InputHandler{keyboardState: sdl.GetKeyboardState(), start: time.Now()}
}

func handleInput(entities map[string]*Entity, iHandler *InputHandler) {
	iHandler.handleInput(entities)
}
