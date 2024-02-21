package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func printFPS(renderer *sdl.Renderer, fps float64) {
	font, err := ttf.OpenFont("./FreeSans.ttf", 24)

	if err != nil {
		return
	}

	white := sdl.Color{255, 255, 255, 255}
	black := sdl.Color{0, 0, 0, 255}

	fpsStr := fmt.Sprintf("fps %f", fps)

	surf, err2 := font.RenderUTF8Shaded(fpsStr, white, black)

	defer surf.Free()

	if err2 != nil {
		return
	}

	message, err3 := renderer.CreateTextureFromSurface(surf)

	defer message.Destroy()

	if err3 != nil {
		return
	}

	rect := &sdl.Rect{X: 500, Y: 500, W: 100, H: 25}

	renderer.Copy(message, nil, rect)

}
