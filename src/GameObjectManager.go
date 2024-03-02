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

func (gObjManager *GameObjectManager) get(id string) *GameObject {
	return gObjManager.gameObjects[id]
}

func (gObjManager *GameObjectManager) set(id string, gObj *GameObject) {
	gObjManager.gameObjects[id] = gObj
}

func (gObjManager *GameObjectManager) fromJSON(path string) {
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

	for key, val := range payload {
		gObjManager.set(key, val.toGameObject())
	}

	// tManager.gameObjects = payload

}
