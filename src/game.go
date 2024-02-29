package main

import (
	"fmt"
	"time"
	"unsafe"

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
	mHandler := createMouseHandler()

	camera := createCamera()
	aRenderer := createARenderer(gameRenderer, camera)

	for running {
		aRenderer.clearRenderer()
		aRenderer.handleRendering(tileMap)

		aRenderer.handleRendering(entities)
		t := time.Now()
		elapsed := t.Sub(start)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {

			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseWheelEvent:
				parsedEvent := *(*sdl.MouseWheelEvent)(unsafe.Pointer(&event))
				if parsedEvent.Y >= 0 {
					gameRenderer.SetScale(0.5, 0.5)
				} else {
					gameRenderer.SetScale(1.5, 1.5)
				}
			}
		}

		if elapsed.Seconds() > 1 {
			fpsString := fmt.Sprintf("%d fps", int(float64(cycles)/elapsed.Seconds()))
			fpsCounter.text.setText(fpsString)
			start = time.Now()
			cycles = 0
		}
		iHandler.handleInput(entities)
		mHandler.handleCameraMove(camera)
		aRenderer.present()
		cycles++
	}
}
