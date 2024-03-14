package gfx

import (
	"encoding/json"
	"log"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager struct {
	Textures map[string]*sdl.Texture
	Renderer *sdl.Renderer
}

func CreateTextureManager(renderer *sdl.Renderer) *TextureManager {
	return &TextureManager{Renderer: renderer, Textures: make(map[string]*sdl.Texture)}
}

func (tManager *TextureManager) Get(id string) *sdl.Texture {
	return tManager.Textures[id]
}

func (tManager *TextureManager) Set(id string, texture *sdl.Texture) {
	tManager.Textures[id] = texture
}

func (tManager *TextureManager) LoadImageAsTexture(fileLocation string) (texture *sdl.Texture) {
	texture, err := img.LoadTexture(tManager.Renderer, fileLocation)

	if err != nil {
		panic(err)
	}

	return texture
}

func (tManager *TextureManager) FromJSON(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload map[string]string
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	for key, val := range payload {
		tManager.Textures[key] = tManager.LoadImageAsTexture(val)
	}

}
