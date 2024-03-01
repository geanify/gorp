package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func handleQuit() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch event.(type) {

		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}

func handleFpsCounter(fpsCounter *Entity, start *time.Time, cycles *int) {
	t := time.Now()
	elapsed := t.Sub(*start)
	*cycles++
	if elapsed.Seconds() < 1 {
		return
	}

	fpsString := fmt.Sprintf("%d fps", int(float64(*cycles)/elapsed.Seconds()))
	fpsCounter.text.setText(fpsString)
	*start = time.Now()
	*cycles = 0
}

func gameLoop(gameRenderer *sdl.Renderer, texture *sdl.Texture) {
	start := time.Now()
	cycles := 0

	tileMap := generateTileMap(gameRenderer)
	entities := loadEntities(texture, gameRenderer)
	fpsCounter := createFPSCounter()
	entities["fpsCounter"] = fpsCounter
	iHandler := createInputHandler()
	mHandler := createMouseHandler()

	camera := createCamera()
	aRenderer := createARenderer(gameRenderer, camera)

	for {
		aRenderer.clearRenderer()
		aRenderer.handleRendering(tileMap)

		aRenderer.handleRendering(entities)

		aRenderer.present()
		iHandler.handleInput(entities)
		mHandler.handleCameraMove(camera)

		handleFpsCounter(fpsCounter, &start, &cycles)

		handleQuit()
	}
}
