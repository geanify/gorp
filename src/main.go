package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	initSDL()
	window := initWindow()
	defer window.Destroy()
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	renderer := createRenderer(window)

	texture := loadImageAsTexture("../assets/pokemon.png", renderer)

	// createSurface(window)

	// window.UpdateSurface()

	gameLoop(renderer, texture)

}
