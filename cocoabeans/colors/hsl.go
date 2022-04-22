package colors

import (
	"Learning2/cocoabeans/math/math_extra"
	"math"
)

type HSL interface {
	Hue() uint16 // expected to be bounded to [0..359]
	Saturation() float64
	Lightness() float64

	// ToRGB Converter
	ToRGB() RGB
	// ToHSL Converter
	ToHSL() HSL
	// ToHSV Converter
	ToHSV() HSV
}

type hsl struct {
	hue        uint16 // expected to be bounded to [0..359]
	saturation float64
	lightness  float64
}

func FromHSL(hue uint16, saturation float64, lightness float64) HSL {
	return hsl{math_extra.LoopClamp(hue, 0, 359), math_extra.Clamp(saturation, 0, 1), math_extra.Clamp(lightness, 0, 1)}
}

func (color hsl) Hue() uint16 {
	return color.hue
}

func (color hsl) Saturation() float64 {
	return color.saturation
}

func (color hsl) Lightness() float64 {
	return color.lightness
}

func (color hsl) ToRGB() RGB {
	c := (1 - math.Abs(2*color.lightness-1)) * color.saturation
	x := c * (1 - math.Abs(math.Mod(float64(color.hue)/60, 2)-1))
	m := color.lightness - c/2

	if color.hue < 60 {
		return rgb{ClampRGB((c + m) * 255), ClampRGB((x + m) * 255), ClampRGB((0 + m) * 255)}
	} else if color.hue < 120 {
		return rgb{ClampRGB((x + m) * 255), ClampRGB((c + m) * 255), ClampRGB((0 + m) * 255)}
	} else if color.hue < 180 {
		return rgb{ClampRGB((0 + m) * 255), ClampRGB((c + m) * 255), ClampRGB((x + m) * 255)}
	} else if color.hue < 240 {
		return rgb{ClampRGB((0 + m) * 255), ClampRGB((x + m) * 255), ClampRGB((c + m) * 255)}
	} else if color.hue < 300 {
		return rgb{ClampRGB((x + m) * 255), ClampRGB((0 + m) * 255), ClampRGB((c + m) * 255)}
	} else {
		return rgb{ClampRGB((c + m) * 255), ClampRGB((0 + m) * 255), ClampRGB((x + m) * 255)}
	}
}

func (color hsl) ToHSL() HSL {
	return color
}

func (color hsl) ToHSV() HSV {
	value := color.lightness + color.saturation*math.Min(color.lightness, 1-color.lightness)
	if value == 0 {
		return hsv{color.hue, 0, value}
	} else {
		return hsv{color.hue, 2 * (1 - color.lightness/value), value}
	}
}
