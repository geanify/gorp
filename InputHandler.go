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

func (iHandler *InputHandler) handleInput(entities []*Entity) {
	now := time.Now()
	elapsed := time.Now().Sub(iHandler.start)

	if elapsed.Milliseconds() < tickRateMS {
		return
	}

	iHandler.keyboardState = sdl.GetKeyboardState()
	iHandler.start = now

	if iHandler.isKeyPressed(sdl.SCANCODE_A) {
		for i := 0; i < len(entities); i++ {
			entities[i].moveLeft(elapsed)
		}
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_D) {
		for i := 0; i < len(entities); i++ {
			entities[i].moveRight(elapsed)
		}
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_W) {
		for i := 0; i < len(entities); i++ {
			entities[i].moveUp(elapsed)
		}
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_S) {
		for i := 0; i < len(entities); i++ {
			entities[i].moveDown(elapsed)
		}
	}
}

func createInputHandler() *InputHandler {
	return &InputHandler{keyboardState: sdl.GetKeyboardState(), start: time.Now()}
}

func handleInput(entities []*Entity, iHandler *InputHandler) {
	iHandler.handleInput(entities)
}
