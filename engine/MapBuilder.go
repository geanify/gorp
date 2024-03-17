package engine

import (
	"github.com/geanify/gorp/gfx"
	"github.com/geanify/gorp/gobj"
	"github.com/geanify/gorp/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type MapBuilder struct {
	GameMap *Map
}

func (mBuilder *MapBuilder) Create(gameRenderer *sdl.Renderer) *MapBuilder {

	camera := utils.CreateCamera()
	aRenderer := createARenderer(gameRenderer, camera)
	mBuilder.GameMap = &Map{ARenderer: aRenderer, Camera: camera}

	return mBuilder
}

func (mBuilder *MapBuilder) WithGameObjectManager(objPath string) *MapBuilder {
	gObjManager := gobj.CreateGameObjectManager()
	gObjManager.FromJSON(objPath)

	mBuilder.GameMap.GameObjManager = gObjManager

	return mBuilder
}

func (mBuilder *MapBuilder) WithTextureManager(tManagerPath string) *MapBuilder {
	tManager := gfx.CreateTextureManager(mBuilder.GameMap.ARenderer.renderer)
	tManager.FromJSON(tManagerPath)
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

func GenerateTestMap(gameRenderer *sdl.Renderer) *Map {
	mBuilder := &MapBuilder{}

	return mBuilder.Create(
		gameRenderer,
	).WithGameObjectManager(
		"assets/gobj.json",
	).WithTextureManager(
		"assets/textures.json",
	).WithTiles().WithUnits().WithFogOfWar().WithFpsCounter("assets/font/FreeSans.ttf").Build()
}
