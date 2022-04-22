package entity

import "Learning2/cocoabeans/data"

type Entity interface {
	getLocation() data.Vec3float64
}

type Builder[E Entity] interface {
	build(int) E
}
