package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func loadEntities(texture *sdl.Texture) map[string]*Entity {
	entities := make(map[string]*Entity)

	rect := &sdl.Rect{X: 0, Y: 0, W: 64, H: 64}
	sprite := Sprite{texture: texture, frame: rect}
	pos := &sdl.Rect{X: 100, Y: 100, W: 64, H: 64}
	entity := Entity{sprite: &sprite, position: pos, speed: 25}
	entities["player"] = &entity

	// entity2 := Entity{sprite: &sprite, position: rect, speed: 1}
	// entities = append(entities, entity2)

	return entities
}
