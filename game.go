package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func gameLoop(gameRenderer *sdl.Renderer, texture *sdl.Texture) {
	running := true
	start := time.Now()
	cycles := 0
	tileMap := generateTileMap(gameRenderer)
	entities := loadEntities(texture, gameRenderer)
	fpsCounter := createFPSCounter()
	entities["fpsCounter"] = fpsCounter
	iHandler := createInputHandler()

	go handleInput(entities, iHandler)
	for running {
		gameRenderer.Clear()
		t := time.Now()
		elapsed := t.Sub(start)
		renderEntities(tileMap, gameRenderer)
		renderEntities(entities, gameRenderer)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {

			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}

		}

		if elapsed.Seconds() > 1 {
			fpsString := fmt.Sprintf("%f", float64(cycles)/elapsed.Seconds())
			fpsCounter.text.setText(fpsString)
			start = time.Now()
			cycles = 0
		}
		gameRenderer.Present()
		cycles++
	}
}
