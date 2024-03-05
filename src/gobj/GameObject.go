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

func (gObject *GameObject) GetDistanceAdjustedPosition() *utils.Vec2 {

	if gObject.Physics == nil {
		return gObject.Position
	}
	if gObject.Physics.DistanceFromCam == 0 {
		return gObject.Position
	}

	distanceSq := gObject.Physics.DistanceFromCam * gObject.Physics.DistanceFromCam

	return &utils.Vec2{
		X: gObject.Position.X / int32(distanceSq),
		Y: gObject.Position.Y / int32(distanceSq),
	}
}

func (gObject *GameObject) GetDistanceAdjustedSize() *utils.Vec2 {
	if gObject.Physics == nil {
		return gObject.Size
	}
	if gObject.Physics.DistanceFromCam == 0 {
		return gObject.Size
	}

	distanceSq := gObject.Physics.DistanceFromCam * gObject.Physics.DistanceFromCam

	return &utils.Vec2{
		X: gObject.Size.X / int32(distanceSq),
		Y: gObject.Size.Y / int32(distanceSq),
	}
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

func (gObject *GameObject) InvertMovement() {
	gObject.Physics.InvertMovement()
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
