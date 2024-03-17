package engine

import "github.com/geanify/gorp/sfx"

type InputHandler struct {
	kHandlerAnimation *KeyHandler
	kHandlerMovement  *KeyHandler
	mHandler          *MouseHandler
	_map              *Map
	audio             *sfx.Audio
}

func CreateInputHandler(_map *Map) *InputHandler {
	kHandlerAnimation := createKeyHandler()
	kHandlerMovement := createKeyHandler()
	mHandler := createMouseHandler()
	audio := sfx.CreateAudio()
	audio.GenerateChunks()

	return &InputHandler{
		kHandlerAnimation: kHandlerAnimation,
		kHandlerMovement:  kHandlerMovement,
		mHandler:          mHandler,
		_map:              _map,
		audio:             audio,
	}
}

func (iHandler *InputHandler) HandleInput() {
	iHandler.kHandlerAnimation.animationHandler(iHandler._map.Units, iHandler.audio)
	iHandler.kHandlerMovement.handleMovement(iHandler._map.GameObjManager)
	iHandler.mHandler.handleCameraMove(iHandler._map.Camera)
}
