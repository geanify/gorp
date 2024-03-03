package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager struct {
	textures map[string]*sdl.Texture
	renderer *sdl.Renderer
}

func createTextureManager(renderer *sdl.Renderer) *TextureManager {
	return &TextureManager{renderer: renderer, textures: make(map[string]*sdl.Texture)}
}

func (tManager *TextureManager) get(id string) *sdl.Texture {
	return tManager.textures[id]
}

func (tManager *TextureManager) set(id string, texture *sdl.Texture) {
	tManager.textures[id] = texture
}

func (tManager *TextureManager) loadImageAsTexture(fileLocation string) (texture *sdl.Texture) {
	texture, err := img.LoadTexture(tManager.renderer, fileLocation)

	if err != nil {
		panic(err)
	}

	return texture
}

func (tManager *TextureManager) fromJSON(path string) {
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
		tManager.textures[key] = tManager.loadImageAsTexture(val)
	}

}
