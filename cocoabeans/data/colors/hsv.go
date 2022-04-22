package colors

import (
	"Learning2/cocoabeans/math_extra"
	"math"
)

type HSV interface {
	Hue() uint16 // expected to be bounded to [0..359]
	Saturation() float64
	Value() float64

	// ToRGB Converter
	ToRGB() RGB
	// ToHSL Converter
	ToHSL() HSL
	// ToHSV Converter
	ToHSV() HSV
}

type hsv struct {
	hue        uint16 // expected to be bounded to [0..359]
	saturation float64
	value      float64
}

func FromHSV(hue uint16, saturation float64, value float64) HSV {
	return hsv{math_extra.LoopClamp(hue, 0, 359), math_extra.Clamp(saturation, 0, 1), math_extra.Clamp(value, 0, 1)}
}

func (color hsv) Hue() uint16 {
	return color.hue
}

func (color hsv) Saturation() float64 {
	return color.saturation
}

func (color hsv) Value() float64 {
	return color.value
}

func (color hsv) ToRGB() RGB {
	c := color.value * color.saturation
	x := c * (1 - math.Abs(math.Mod(float64(color.hue)/60, 2)-1))
	m := color.value - c

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

func (color hsv) ToHSL() HSL {
	lightness := color.value * (1 - color.saturation/2)
	if lightness == 0 || lightness == 1 {
		return hsl{color.hue, 0, lightness}
	} else {
		return hsl{color.hue, (color.value - lightness) / math.Min(lightness, 1-lightness), lightness}
	}
}

func (color hsv) ToHSV() HSV {
	return color
}

func ClampHue[N int32 | uint32 | int64 | uint64 | float32 | float64](value N) uint16 {
	if uint16(value) < 0 {
		return 0
	}
	if uint16(value) > 359 {
		return 359
	}
	return uint16(value)
}
