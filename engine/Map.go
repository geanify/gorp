package engine

import (
	"github.com/geanify/gorp/gfx"
	"github.com/geanify/gorp/gobj"
	"github.com/geanify/gorp/utils"
)

type Map struct {
	MapSize        *utils.Vec2
	TileSize       *utils.Vec2
	Tiles          [][]*Entity
	GameObjManager *gobj.GameObjectManager
	TextureManager *gfx.TextureManager
	FogOfWar       *FogOfWar
	Units          map[string]*Entity
	ARenderer      *AsyncRenderer
	Camera         *utils.Camera
}

func (_map *Map) RenderMap() {
	_map.ARenderer.clearRenderer()
	_map.ARenderer.handleTileRendering(_map.Tiles)

	_map.ARenderer.handleRendering(_map.Units)
	_map.FogOfWar.UpdateFogOfWar(_map.Units)
	_map.ARenderer.handleTileRendering(_map.FogOfWar.fog)

	_map.ARenderer.present()
}
