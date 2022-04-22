package colors

type Color interface {
	// ToRGB Converter
	ToRGB() RGB
	// ToHSL Converter
	ToHSL() HSL
	// ToHSV Converter
	ToHSV() HSV
}
