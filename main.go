package main

import (
	"github.com/geanify/gorp/editor"
	"github.com/geanify/gorp/engine"
)

func main() {
	go editor.StartServer()
	engine.RunEngine()
}
