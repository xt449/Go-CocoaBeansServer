package entity

import (
	"CocoaBeansServer/cocoabeans/entity/trait"
	"CocoaBeansServer/cocoabeans/math"
	"CocoaBeansServer/cocoabeans/nbt"
)

type Entity[T trait.Trait] interface {
	GetLocation() math.Vec3float64
	BuildNBTContainer() nbt.Container
}

type entity struct {
	location math.Vec3float64
}

type Builder[T trait.Trait] func(T) Entity[T]
