package main

import "github.com/veandco/go-sdl2/sdl"

type AsyncRenderer struct {
	timeControl  *TimeControl
	renderer     *sdl.Renderer
	cam          *Camera
	shouldRender bool
}

func (aRenderer *AsyncRenderer) clearRenderer() {
	if !aRenderer.timeControl.shouldExecute() {
		return
	}
	aRenderer.renderer.Clear()
	aRenderer.shouldRender = true
}
func (aRenderer *AsyncRenderer) handleRendering(entities map[string]*Entity) {
	if !aRenderer.shouldRender {
		return
	}
	renderEntities(entities, aRenderer.renderer, aRenderer.cam)
}

func (aRenderer *AsyncRenderer) present() {
	if !aRenderer.shouldRender {
		return
	}
	aRenderer.renderer.Present()
	aRenderer.shouldRender = false
}

func createARenderer(renderer *sdl.Renderer, cam *Camera) *AsyncRenderer {
	return &AsyncRenderer{timeControl: createTimeControl(), renderer: renderer, cam: cam}
}
