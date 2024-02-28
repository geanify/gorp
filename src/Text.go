package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	font            *ttf.Font
	color           sdl.Color
	backgroundColor sdl.Color
	text            string
}

func (text *Text) setFont(fontName string, fontSize int) {
	font, err := ttf.OpenFont(fontName, fontSize)

	if err != nil {
		return
	}
	text.font = font
}

func (text *Text) setColorRGBA(R uint8, G uint8, B uint8, A uint8) {
	text.color = sdl.Color{R: R, G: G, B: B, A: A}
}

func (text *Text) setColorRGB(R uint8, G uint8, B uint8) {
	text.setColorRGBA(R, G, B, 255)
}

func (text *Text) setBackgroundColorRGBA(R uint8, G uint8, B uint8, A uint8) {
	text.backgroundColor = sdl.Color{R: R, G: G, B: B, A: A}
}

func (text *Text) setBackgroundColorRGB(R uint8, G uint8, B uint8) {
	text.setBackgroundColorRGBA(R, G, B, 255)
}

func (text *Text) setText(txt string) {
	text.text = txt
}

func (text *Text) renderText(renderer *sdl.Renderer, position *sdl.Rect) {

	surf, surfErr := text.font.RenderUTF8Shaded(text.text, text.color, text.backgroundColor)
	defer surf.Free()
	if surfErr != nil {
		return
	}

	message, messageErr := renderer.CreateTextureFromSurface(surf)

	defer message.Destroy()

	if messageErr != nil {
		return
	}

	renderer.Copy(message, nil, position)
}
