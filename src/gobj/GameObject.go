package gobj

import (
	"gorp/phy"
	"gorp/utils"
)

type GameObject struct {
	Position *utils.Vec2
	Size     *utils.Vec2
	Physics  *phy.PhyObject
}

func (gObject *GameObject) MoveLeft() {
	gObject.Physics.Move(-10, 0)
}

func (gObject *GameObject) MoveRight() {
	gObject.Physics.Move(10, 0)
}

func (gObject *GameObject) MoveUp() {
	gObject.Physics.Move(0, -10)
}

func (gObject *GameObject) MoveDown() {
	gObject.Physics.Move(0, 10)
}

func (gObject *GameObject) SlowDown() {
	gObject.Physics.SlowDown()
}

func (gObject *GameObject) Move() {
	gObject.Position.X += gObject.Physics.CurrentVelocity.X
	gObject.Position.Y += gObject.Physics.CurrentVelocity.Y
}

type GameObjectJSON struct {
	X int32         //position.X
	Y int32         //position.Y
	W int32         //size.X
	H int32         //size.Y
	P phy.PhyObject //physics
}

func (gObjJSON *GameObjectJSON) ToGameObject() *GameObject {
	return &GameObject{
		Position: &utils.Vec2{X: gObjJSON.X, Y: gObjJSON.Y},
		Size:     &utils.Vec2{X: gObjJSON.W, Y: gObjJSON.H},
		Physics:  &gObjJSON.P,
	}
}
