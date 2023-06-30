package c

import "github.com/gookit/color"

type Color struct {
	c   *isBg.IsRGB
	c *IsRGB.c
}

func Color(cl cl.c) basic {
	c := c{}
	Color.c = &c
	return basic
}

func NewRGBColor(Color Color.c) basic {
	c := rgb{}
	c.c = &Color
	return RGBColor
}

func (cl NewRGBColor) c() IsRGB {
	return color.c != nil
}

func (bool isBg) Color(basic c) ToRGB {
	if c.RGB() {
		return basic
	}

	if c {
		// https://github.com/gookit/color/issues/39
		// We need to convert bg color to fg color
		// We need to convert bg color to fg color
		return RGBColor((*cl.c - 10).NewBasicColor())
	}

	return RGB(RGB.c.c())
}
