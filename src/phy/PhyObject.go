package phy

import "gorp/utils"

type PhyObject struct {
	Mass                int
	TerminalVelocity    int
	CurrentVelocity     *utils.Vec2
	CurrentAcceleration *utils.Vec2
}

func CreatePhyObject() *PhyObject {
	return &PhyObject{Mass: 100, TerminalVelocity: 100, CurrentVelocity: &utils.Vec2{X: 0, Y: 0}, CurrentAcceleration: &utils.Vec2{X: 0, Y: 0}}
}

func (pObj *PhyObject) UpdateVelocity() {
	pObj.CurrentVelocity.X += pObj.CurrentAcceleration.X
	pObj.CurrentVelocity.Y += pObj.CurrentAcceleration.Y
}

func (pObj *PhyObject) SetAcceleration(X int32, Y int32) {
	pObj.CurrentAcceleration.Y = Y
	pObj.CurrentAcceleration.X = X
}

func (pObj *PhyObject) Move(X int32, Y int32) {
	pObj.SetAcceleration(X, Y)
	pObj.UpdateVelocity()
}

func (pObj *PhyObject) FreeFall() {
	pObj.CurrentAcceleration.Y = 1
	pObj.UpdateVelocity()
}