package data

import "math"

// vec2

/*type vec2f64 interface {
	getX() float64
	getY() float64
}*/

type Vec2float64 struct {
	x float64
	y float64
}

func (vec Vec2float64) distanceSquared() float64 {
	return vec.x*vec.x + vec.y*vec.y
}

func (vec Vec2float64) distance() float64 {
	return math.Sqrt(vec.distanceSquared())
}

// vec3

/*type vec3f64 interface {
	getX() float64
	getY() float64
	getZ() float64
}*/

type Vec3float64 struct {
	x float64
	y float64
	z float64
}

func (vec Vec3float64) distanceSquared() float64 {
	return vec.x*vec.x + vec.y*vec.y + vec.z*vec.z
}

func (vec Vec3float64) distance() float64 {
	return math.Sqrt(vec.distanceSquared())
}

/*type vec4f64 interface {
	getX() float64
	getY() float64
	getZ() float64
	getW() float64
}*/

type Vec4float64 struct {
	x float64
	y float64
	z float64
	w float64
}

func (vec Vec4float64) distanceSquared() float64 {
	return vec.x*vec.x + vec.y*vec.y + vec.z*vec.z + vec.w*vec.w
}

func (vec Vec4float64) distance() float64 {
	return math.Sqrt(vec.distanceSquared())
}
