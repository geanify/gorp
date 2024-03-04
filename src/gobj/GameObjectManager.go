package gobj

import (
	"encoding/json"
	"fmt"
	"gorp/utils"
	"log"
	"os"
)

type GameObjectManager struct {
	gameObjects map[string]*GameObject
	Collisions  map[string]map[string]bool
}

func CreateGameObjectManager() *GameObjectManager {
	return &GameObjectManager{
		gameObjects: make(map[string]*GameObject),
		Collisions:  make(map[string]map[string]bool),
	}
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

func (gObjManager *GameObjectManager) CheckCollision(gObjOne *GameObject, gObjTwo *GameObject) bool {
	return utils.HasIntersection(gObjOne.Position, gObjOne.Size, gObjTwo.Position, gObjTwo.Size)
}

func (gObjManager *GameObjectManager) GenerateCollisionMatrix() {
	collisions := make(map[string]map[string]bool)

	for key, val := range gObjManager.gameObjects {
		for key2, val2 := range gObjManager.gameObjects {
			if key != key2 && val2.Physics.Solid && val.Physics.Solid {
				result := gObjManager.CheckCollision(val, val2)

				if collisions[key] == nil {
					collisions[key] = make(map[string]bool)
				}
				if collisions[key2] == nil {
					collisions[key2] = make(map[string]bool)
				}
				collisions[key][key2] = result
				collisions[key2][key] = result
			}
		}
	}

	gObjManager.Collisions = collisions
}

func (gObjManager *GameObjectManager) HasCollision(id string) bool {
	for _, val := range gObjManager.Collisions[id] {
		if val {
			return true
		}
	}
	return false
}
