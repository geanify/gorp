package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type MouseHandler struct {
	start time.Time
}

func (mHandler *MouseHandler) handleCameraMove() {
	now := time.Now()
	elapsed := now.Sub(mHandler.start)

	if elapsed.Milliseconds() < tickRateMS {
		return
	}

	mHandler.start = now

	x, y, _ := sdl.GetMouseState()

	if x < 20 {
		fmt.Println("edge")
	}
	if y < 20 {
		fmt.Println("edge")
	}

}

func handleMouse(mHandler *MouseHandler) {
	for {
		mHandler.handleCameraMove()
	}
}

func createMouseHandler() *MouseHandler {
	return &MouseHandler{start: time.Now()}
}
