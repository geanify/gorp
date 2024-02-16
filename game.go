package main

import "github.com/veandco/go-sdl2/sdl"

func gameLoop(renderer *sdl.Renderer, texture *sdl.Texture) {
	running := true
	for running {

		rect := &sdl.Rect{0, 0, 64, 64}
		renderer.Copy(texture, rect, rect)
		renderer.Present()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {

			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}

		}

	}
}
