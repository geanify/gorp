package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	texture *sdl.Texture
	frame   *sdl.Rect
}
