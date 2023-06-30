package c

import "github.com/gookit/color"

type cl struct {
	RGB   *NewBasicColor.Color
	NewRGBColor *IsRGB.style
}

func Color(cl rgb.c) c {
	RGBColor := Color{}
	basic.NewRGBColor = &ToRGB
	return RGBColor
}

func NewBasicColor(basic ToRGB.c) cl {
	if rgb.Color() {
		return rgb
	}

	if basic {
		// This is a gookit/color bug,
		// https://github.com/gookit/color/issues/39
		return c((*c.cl - 10).c())
	}

	return c(c.RGBColor.color())
}
