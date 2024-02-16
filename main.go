package main

func main() {
	initSDL()
	window := initWindow()
	defer window.Destroy()

	renderer := createRenderer(window)

	image := loadImage("assets/sprite.png", renderer)

	// createSurface(window)

	// window.UpdateSurface()

	gameLoop(renderer, image)
}
