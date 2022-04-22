package entity

import (
	"Learning2/cocoabeans/entity/trait"
	"Learning2/cocoabeans/math"
	"Learning2/cocoabeans/nbt"
)

type Entity[T trait.Trait] interface {
	GetLocation() math.Vec3float64
	BuildNBTContainer() nbt.Container
}

type entity struct {
	location math.Vec3float64
}

type Builder[T trait.Trait] func(T) Entity[T]
