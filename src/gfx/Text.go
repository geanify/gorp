package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	Font            *ttf.Font
	Color           sdl.Color
	BackgroundColor sdl.Color
	Text            string
}

func (text *Text) SetFont(fontName string, fontSize int) {
	font, err := ttf.OpenFont(fontName, fontSize)

	if err != nil {
		return
	}
	text.Font = font
}

func (text *Text) SetColorRGBA(R uint8, G uint8, B uint8, A uint8) {
	text.Color = sdl.Color{R: R, G: G, B: B, A: A}
}

func (text *Text) SetColorRGB(R uint8, G uint8, B uint8) {
	text.SetColorRGBA(R, G, B, 255)
}

func (text *Text) SetBackgroundColorRGBA(R uint8, G uint8, B uint8, A uint8) {
	text.BackgroundColor = sdl.Color{R: R, G: G, B: B, A: A}
}

func (text *Text) SetBackgroundColorRGB(R uint8, G uint8, B uint8) {
	text.SetBackgroundColorRGBA(R, G, B, 255)
}

func (text *Text) SetText(txt string) {
	text.Text = txt
}

func (text *Text) RenderText(renderer *sdl.Renderer, position *sdl.Rect) {

	surf, surfErr := text.Font.RenderUTF8Shaded(text.Text, text.Color, text.BackgroundColor)
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
