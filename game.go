package main

import "github.com/veandco/go-sdl2/sdl"

func gameLoop(renderer *sdl.Renderer, texture *sdl.Texture) {
	running := true
	for running {

		rect := &sdl.Rect{X: 0, Y: 0, W: 64, H: 64}
		sprite := Sprite{texture: texture, frame: rect}
		pos := &sdl.Rect{X: 100, Y: 100, W: 64, H: 64}
		entity := Entity{sprite: &sprite, position: pos}
		entity2 := Entity{sprite: &sprite, position: rect}
		entity.render(renderer)
		entity2.render(renderer)
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
