package main

func main() {
	initSDL()
	window := initWindow()
	defer window.Destroy()

	renderer := createRenderer(window)

	texture := loadImageAsTexture("assets/sprite.png", renderer)

	// createSurface(window)

	// window.UpdateSurface()

	gameLoop(renderer, texture)
}
