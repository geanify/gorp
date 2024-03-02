package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type GameObjectManager struct {
	gameObjects map[string]*GameObject
}

func createGameObjectManager() *GameObjectManager {
	return &GameObjectManager{gameObjects: make(map[string]*GameObject)}
}

func (tManager *GameObjectManager) get(id string) *GameObject {
	return tManager.gameObjects[id]
}

func (tManager *GameObjectManager) set(id string, gObj *GameObject) {
	tManager.gameObjects[id] = gObj
}

func (tManager *GameObjectManager) fromJSON(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload map[string]GameObjectJSON
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Println(payload)

	// tManager.gameObjects = payload

}
