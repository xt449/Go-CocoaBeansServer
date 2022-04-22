package entity

import (
	"Learning2/cocoabeans/math"
	"Learning2/cocoabeans/nbt"
)

type Entity interface {
	GetLocation() math.Vec3float64
	BuildNBTContainer() nbt.Container
}

type Builder[E Entity] interface {
	build(int) E
}
