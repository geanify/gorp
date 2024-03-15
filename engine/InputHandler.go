package engine

import (
	"time"

	"github.com/geanify/gorp/gobj"
	"github.com/geanify/gorp/sfx"
	"github.com/geanify/gorp/utils"

	"github.com/veandco/go-sdl2/sdl"
)

type InputHandler struct {
	keyboardState []uint8
	start         time.Time
	timeControl   *utils.TimeControl
}

func (iHandler *InputHandler) isKeyPressed(key int) bool {
	return iHandler.keyboardState[key] == 1
}

func (iHandler *InputHandler) handleMovement(gameObjects *gobj.GameObjectManager) {
	if !iHandler.timeControl.ShouldExecute() {
		return
	}
	iHandler.keyboardState = sdl.GetKeyboardState()

	player := gameObjects.Get("player")

	gameObjects.GenerateCollisionMatrix()

	player.SlowDown()
	if iHandler.isKeyPressed(sdl.SCANCODE_A) {
		player.MoveLeft()
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_D) {
		player.MoveRight()
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_W) {
		player.MoveUp()
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_S) {
		player.MoveDown()
	}

	if gameObjects.HasCollision("player") {
		player.InvertMovement()
	}

	player.Move()
}

func (iHandler *InputHandler) animationHandler(entitiesMap map[string]*Entity, audio *sfx.Audio) {
	if !iHandler.timeControl.ShouldExecute() {
		return
	}

	iHandler.keyboardState = sdl.GetKeyboardState()

	player := entitiesMap["player"]
	if iHandler.isKeyPressed(sdl.SCANCODE_SPACE) {
		go audio.PlayTrack("test1")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_A) {
		player.sprite.NextFrame()
		player.sprite.SetAnimation("left")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_D) {
		player.sprite.NextFrame()
		player.sprite.SetAnimation("right")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_W) {
		player.sprite.NextFrame()
		player.sprite.SetAnimation("up")
	}
	if iHandler.isKeyPressed(sdl.SCANCODE_S) {
		player.sprite.NextFrame()
		player.sprite.SetAnimation("down")
	}
}

func createInputHandler() *InputHandler {
	return &InputHandler{keyboardState: sdl.GetKeyboardState(), start: time.Now(), timeControl: utils.CreateTimeControl()}
}

// func handleInput(entities map[string]*Entity, iHandler *InputHandler) {
// 	for {
// 		iHandler.handleInput(entities)
// 		time.Sleep((tickRateMS / 3) * time.Millisecond)
// 	}
// }
