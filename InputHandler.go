package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func isKeyPressed(keyboardState []uint8, key int) bool {
	return keyboardState[key] == 1
}

func handleGameInput(entities []Entity) {
	keyboardState := sdl.GetKeyboardState()

	if isKeyPressed(keyboardState, sdl.SCANCODE_A) {
		fmt.Println("A pressed")
		for i := 0; i < len(entities); i++ {
			entities[i].move()
		}
	}
}

func handleInput(entities []Entity) {
	handleGameInput(entities)
}
