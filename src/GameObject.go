package main

type GameObject struct {
	position *Vec2
	size     *Vec2
	speed    int32 //per tickrate
	physics  int32
}

func (gObject *GameObject) moveLeft() {
	gObject.position.X -= gObject.speed
}

func (gObject *GameObject) moveRight() {
	gObject.position.X += gObject.speed
}

func (gObject *GameObject) moveUp() {
	gObject.position.Y -= gObject.speed
}

func (gObject *GameObject) moveDown() {
	gObject.position.Y += gObject.speed
}

type GameObjectJSON struct {
	X int32 //position.X
	Y int32 //position.Y
	W int32 //size.X
	H int32 //size.Y
	S int32 //speed
	P int32 //physics
}

func (gObjJSON *GameObjectJSON) toGameObject() *GameObject {
	return &GameObject{
		position: &Vec2{X: gObjJSON.X, Y: gObjJSON.Y},
		size:     &Vec2{X: gObjJSON.W, Y: gObjJSON.H},
		speed:    gObjJSON.S,
		physics:  gObjJSON.P,
	}
}
