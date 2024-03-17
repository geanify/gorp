package engine

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
	fpsCounter.text.SetText(fpsString)
	*start = time.Now()
	*cycles = 0
}

func gameLoop(gameRenderer *sdl.Renderer) {
	start := time.Now()
	cycles := 0

	_map := GenerateTestMap(gameRenderer)
	iHandler := CreateInputHandler(_map)

	loadParticle(_map.Units, _map.GameObjManager, _map.TextureManager)

	for {
		_map.RenderMap()
		iHandler.HandleInput()

		handleFpsCounter(_map.Units["fpsCounter"], &start, &cycles)

		handleQuit()
	}
}
