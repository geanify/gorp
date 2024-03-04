package utils

type Vec2 struct {
	X int32
	Y int32
}

func (vec *Vec2) Empty() bool {
	return vec == nil
}

func HasIntersection(a *Vec2, sizeA *Vec2, b *Vec2, sizeB *Vec2) bool {

	if a.Empty() || sizeA.Empty() || b.Empty() || sizeB.Empty() {
		return false
	}

	if a.X >= b.X+sizeB.X || a.X+sizeA.X <= b.X || a.Y >= b.Y+sizeB.Y || a.Y+sizeA.Y <= b.Y {
		return false
	}

	return true
}
