package gobj

type GameObject struct {
	Position *Vec2
	Size     *Vec2
	Speed    int32 //per tickrate
	Physics  int32
}

func (gObject *GameObject) MoveLeft() {
	gObject.Position.X -= gObject.Speed
}

func (gObject *GameObject) MoveRight() {
	gObject.Position.X += gObject.Speed
}

func (gObject *GameObject) MoveUp() {
	gObject.Position.Y -= gObject.Speed
}

func (gObject *GameObject) MoveDown() {
	gObject.Position.Y += gObject.Speed
}

type GameObjectJSON struct {
	X int32 //position.X
	Y int32 //position.Y
	W int32 //size.X
	H int32 //size.Y
	S int32 //speed
	P int32 //physics
}

func (gObjJSON *GameObjectJSON) ToGameObject() *GameObject {
	return &GameObject{
		Position: &Vec2{X: gObjJSON.X, Y: gObjJSON.Y},
		Size:     &Vec2{X: gObjJSON.W, Y: gObjJSON.H},
		Speed:    gObjJSON.S,
		Physics:  gObjJSON.P,
	}
}
