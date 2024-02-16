package main

import "github.com/veandco/go-sdl2/sdl"

func gameLoop(renderer *sdl.Renderer, texture *sdl.Texture) {
	running := true
	for running {

		renderer.Copy(texture, nil, nil)
		renderer.Present()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {

			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break

			default:
				break
			}

		}

	}
}
