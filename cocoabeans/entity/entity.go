package entity

import (
	"Learning2/cocoabeans/math"
)

type Entity interface {
	getLocation() math.Vec3float64
}

type Builder[E Entity] interface {
	build(int) E
}
