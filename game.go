package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func gameLoop(renderer *sdl.Renderer, texture *sdl.Texture) {
	running := true
	start := time.Now()
	cycles := 0
	for running {
		t := time.Now()
		elapsed := t.Sub(start)
		if elapsed.Seconds() > 1 {
			fps := float64(cycles) / elapsed.Seconds()
			printFPS(renderer, fps)
			start = time.Now()
			cycles = 0
		}
		entities := loadEntities(texture)
		renderEntities(entities, renderer)
		renderer.Present()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {

			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}

		}
		cycles++
	}
}
