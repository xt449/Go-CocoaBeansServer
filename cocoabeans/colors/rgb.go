package colors

import (
	"CocoaBeansServer/cocoabeans/math/math_extra"
	"math"
)

type RGB interface {
	Red() uint8
	Green() uint8
	Blue() uint8

	// ToRGB Converter
	ToRGB() RGB
	// ToHSL Converter
	ToHSL() HSL
	// ToHSV Converter
	ToHSV() HSV
}

type rgb struct {
	red   uint8
	green uint8
	blue  uint8
}

func FromRGB(red uint8, green uint8, blue uint8) Color {
	return rgb{red, green, blue}
}

func (color rgb) Red() uint8 {
	return color.red
}

func (color rgb) Green() uint8 {
	return color.green
}

func (color rgb) Blue() uint8 {
	return color.blue
}

func (color rgb) ToRGB() RGB {
	return color
}

func (color rgb) ToHSL() HSL {
	r_ := float64(color.red) / float64(math.MaxUint8)
	g_ := float64(color.green) / float64(math.MaxUint8)
	b_ := float64(color.blue) / float64(math.MaxUint8)

	cmin := math_extra.MinOf(r_, g_, b_)
	cmax := math_extra.MaxOf(r_, g_, b_)

	delta := cmax - cmin

	var hue uint16
	if delta == 0 {
		hue = 0
	} else if delta == r_ {
		hue = ClampHue(math.Mod(60*((g_-b_)/delta), 6))
	} else if delta == g_ {
		hue = ClampHue(60 * ((b_-r_)/delta + 2))
	} else /*if delta == b_*/ {
		hue = ClampHue(60 * ((r_-g_)/delta + 4))
	}

	var lightness float64
	lightness = (cmax + cmin) / 2

	var saturation float64
	if delta == 0 {
		saturation = 0
	} else {
		saturation = delta / (1 - math.Abs(2*lightness-1))
	}

	return hsl{hue, saturation, lightness}
}

func (color rgb) ToHSV() HSV {
	r_ := float64(color.red) / float64(math.MaxUint8)
	g_ := float64(color.green) / float64(math.MaxUint8)
	b_ := float64(color.blue) / float64(math.MaxUint8)

	cmin := math_extra.MinOf(r_, g_, b_)
	cmax := math_extra.MaxOf(r_, g_, b_)

	delta := cmax - cmin

	var hue uint16
	if delta == 0 {
		hue = 0
	} else if delta == r_ {
		hue = ClampHue(math.Mod(60*((g_-b_)/delta), 6))
	} else if delta == g_ {
		hue = ClampHue(60 * ((b_-r_)/delta + 2))
	} else /*if delta == b_*/ {
		hue = ClampHue(60 * ((r_-g_)/delta + 4))
	}

	var saturation float64
	if cmax == 0 {
		saturation = 0
	} else {
		saturation = delta / cmax
	}

	return hsv{hue, saturation, cmax}
}

func ClampRGB[N int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](value N) uint8 {
	if uint16(value) < 0 {
		return 0
	}
	if uint16(value) > 255 {
		return 255
	}
	return uint8(value)
}
