package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func gameLoop(renderer *sdl.Renderer, texture *sdl.Texture) {
	running := true
	start := time.Now()
	cycles := 0
	entities := loadEntities(texture)
	fpsCounter := createFPSCounter()
	entities["fpsCounter"] = fpsCounter
	iHandler := createInputHandler()

	// go handleInput(entities)

	for running {
		renderer.Clear()
		t := time.Now()
		elapsed := t.Sub(start)
		renderEntities(entities, renderer)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {

			case *sdl.QuitEvent:
				println("Quit")
				running = false

				// case *sdl.KeyboardEvent:
				// 	handleInput(entities)
			}

		}
		handleInput(entities, iHandler)

		if elapsed.Seconds() > 1 {
			fpsString := fmt.Sprintf("%f", float64(cycles)/elapsed.Seconds())
			fpsCounter.text.setText(fpsString)
			start = time.Now()
			cycles = 0
		}
		renderer.Present()
		cycles++
	}
}
