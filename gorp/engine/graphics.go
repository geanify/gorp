package engine

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func initSDL() {
	if err := sdl.Init(sdl.INIT_EVERYTHING | img.INIT_PNG); err != nil {
		panic(err)
	}
	mix.Init(mix.INIT_MP3)
	ttf.Init()
	// initAudio()
	defer sdl.Quit()

}

func initWindow() (window *sdl.Window) {
	window, err := sdl.CreateWindow("HELLO GO-SDL", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	return window
}

func initAudio() {
	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
}

func createRenderer(window *sdl.Window) (renderer *sdl.Renderer) {
	renderer, err := sdl.CreateRenderer(window, 0, 0)

	if err != nil {
		panic(err)
	}

	return renderer
}

func createSurface(window *sdl.Window) (surface *sdl.Surface) {
	surface, err := window.GetSurface()

	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)

	return surface
}

func loadImageAsTexture(fileLocation string, renderer *sdl.Renderer) (texture *sdl.Texture) {
	texture, err := img.LoadTexture(renderer, fileLocation)

	if err != nil {
		panic(err)
	}

	return texture
}
