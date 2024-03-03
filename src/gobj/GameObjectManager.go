package gobj

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type GameObjectManager struct {
	gameObjects map[string]*GameObject
}

func CreateGameObjectManager() *GameObjectManager {
	return &GameObjectManager{gameObjects: make(map[string]*GameObject)}
}

func (gObjManager *GameObjectManager) Get(id string) *GameObject {
	return gObjManager.gameObjects[id]
}

func (gObjManager *GameObjectManager) Set(id string, gObj *GameObject) {
	gObjManager.gameObjects[id] = gObj
}

func (gObjManager *GameObjectManager) FromJSON(path string) {
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
		gObjManager.Set(key, val.ToGameObject())
	}
}
