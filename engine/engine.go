package engine

import "github.com/veandco/go-sdl2/sdl"

func RunEngine() {
	initSDL()
	window := initWindow()
	defer window.Destroy()
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	renderer := createRenderer(window)

	// createSurface(window)

	// window.UpdateSurface()

	gameLoop(renderer)

}
