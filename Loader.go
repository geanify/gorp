package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func loadEntities(texture *sdl.Texture) map[string]*Entity {
	entities := make(map[string]*Entity)

	rect := &sdl.Rect{X: 0, Y: 0, W: 64, H: 64}
	sprite := Sprite{
		texture:    texture,
		frame:      rect,
		maxFrames:  4,
		frameIndex: 0,
		animations: map[string]*sdl.Rect{
			"down":  {X: 64, Y: 0, W: 64, H: 64},
			"left":  {X: 64, Y: 64, W: 64, H: 64},
			"right": {X: 64, Y: 128, W: 64, H: 64},
			"up":    {X: 64, Y: 192, W: 64, H: 64},
		},
		currentAnimation: "down",
	}
	pos := &sdl.Rect{X: 100, Y: 100, W: 64, H: 64}
	entity := Entity{sprite: &sprite, position: pos, speed: 25}
	entities["player"] = &entity

	// entity2 := Entity{sprite: &sprite, position: rect, speed: 1}
	// entities = append(entities, entity2)

	return entities
}
