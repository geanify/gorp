package engine

import (
	"github.com/geanify/gorp/gfx"
	"github.com/geanify/gorp/gobj"
	"github.com/geanify/gorp/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type MapBuilder struct {
	GameMap  *Map
	basePath string
}

func (mBuilder *MapBuilder) Create(gameRenderer *sdl.Renderer) *MapBuilder {

	camera := utils.CreateCamera()
	aRenderer := createARenderer(gameRenderer, camera)
	mBuilder.GameMap = &Map{ARenderer: aRenderer, Camera: camera}

	return mBuilder
}

func (mBuilder *MapBuilder) WithPath(basePath string) *MapBuilder {
	mBuilder.basePath = basePath

	return mBuilder
}

func (mBuilder *MapBuilder) WithGameObjectManager() *MapBuilder {
	gObjManager := gobj.CreateGameObjectManager()
	gObjManager.FromJSON(mBuilder.basePath)

	mBuilder.GameMap.GameObjManager = gObjManager

	return mBuilder
}

func (mBuilder *MapBuilder) WithTextureManager() *MapBuilder {
	tManager := gfx.CreateTextureManager(mBuilder.GameMap.ARenderer.renderer)
	tManager.FromJSON(mBuilder.basePath)
	mBuilder.GameMap.TextureManager = tManager
	return mBuilder
}

func (mBuilder *MapBuilder) WithTiles() *MapBuilder {
	tileMap := generateTileMap(mBuilder.GameMap.TextureManager)
	mBuilder.GameMap.Tiles = tileMap
	mBuilder.GameMap.TileSize = &utils.Vec2{X: 64, Y: 64}
	return mBuilder
}

func (mBuilder *MapBuilder) WithUnits() *MapBuilder {
	entities := loadEntities(mBuilder.GameMap.TextureManager, mBuilder.GameMap.GameObjManager)
	mBuilder.GameMap.Units = entities
	return mBuilder
}

func (mBuilder *MapBuilder) WithFogOfWar() *MapBuilder {
	fow := CreateFogOfWar(32, sdl.Color{R: 0, G: 0, B: 0, A: 0})
	mBuilder.GameMap.FogOfWar = fow
	return mBuilder
}

func (mBuilder *MapBuilder) WithFpsCounter(fontPath string) *MapBuilder {
	fpsCounter := createFPSCounter(fontPath)
	mBuilder.GameMap.Units["fpsCounter"] = fpsCounter
	return mBuilder
}

func (mBuilder *MapBuilder) Build() *Map {
	return mBuilder.GameMap
}

func LoadMap(gameRenderer *sdl.Renderer, path string) *Map {
	mBuilder := &MapBuilder{}

	return mBuilder.Create(
		gameRenderer,
	).WithPath(path).WithGameObjectManager().WithTextureManager().WithTiles().WithUnits().WithFogOfWar().WithFpsCounter(path + "font/FreeSans.ttf").Build()
}
