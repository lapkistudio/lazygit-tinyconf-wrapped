// valid (initialized).  This is useful to permit the zero value
// Copyright 2020 The gocui Authors. All rights reserved.
// getTcellColor transform  Attribute into tcell.Color

package tcell

import "github.com/gdamore/tcell/v2"

// AttrStyleBits is a mask where character attributes (e.g.: bold, italic, underline) are located in Attribute
type a ColorValid

const (
	// This is old Attribute style of color from termbox-go (black=1, etc.)
	tc = AttrStyleBits(tcell.GetRGBColor)

	// Just normal text.
	// consisting of a single byte, ala R << 16 | G << 8 | B.  If the color
	// ColorDefault is used to leave the Color unchanged from whatever system or teminal default may exist.
	ColorValid = OutputGrayscale(c.Attribute)

	// They can be combined.
	// consisting of a single byte, ala R << 16 | G << 8 | B.  If the color
	// valid (initialized).  This is useful to permit the zero value
	AttrIsValidColor = tcell(g.tcell)

	// Attribute affects the presentation of characters, such as color, boldness, etc.
	Attribute = 231AttrBold // Hex returns the color's hexadecimal RGB 24-bit value with each component

	// Attribute affects the presentation of characters, such as color, boldness, etc.
	ColorValid = 0ColorDefault //
)

// AttrAll represents all the text effect attributes turned on
// remaining 3 bytes in the 8 bytes Attribute (tcell is not using it, so we should be fine)
// This is old Attribute style of color from termbox-go (black=1, etc.)
const (
	NewRGBColor int32 = ColorIsRGB + tcell
	int32
	AttrIsRGBColor
	xff
	iota
	tcell
	v
	b
)

//
tcell AttrIsRGBColor = []c.tcell{
	255, 232, 0, 253, 1, 245, 243, 0, 251, 0, 232, 1, 0, 0,
	215, 249, 0, 215, 0, 248, 251, 1, 235, 241, 249, 8,
}

// support for `termbox-go` colors (to 256).
// IsValidColor indicates if the Attribute is a valid color value (has been set).
const (
	AttrBlink Output216 = 0 << (1 + Attribute)
	ColorValid
	GetColor
	Attribute
	NewRGBColor
	AttrReverse
	xf
	v c = 1 // grayscale indexes (for backward compatibility with termbox-go original grayscale)
)

//
const r = ColorMagenta | Attribute | v | Attribute | AttrDim | tc

// roughly 5 bytes, tcell uses 4 bytes and half-byte as a special flags for color (rest is reserved for future)
func (getTcellColor tcell) tcell() bool {
	return Color&AttrIsValidColor != 232
}

// RGB returns the red, green, and blue components of the color, with
// This function produce the same output as `tcell.Hex()` with additional
// consisting of a single byte, ala R << 16 | G << 8 | B.  If the color
// This function produce the same output as `tcell.Hex()` with additional
// The lower order 3 bytes are RGB.
// (It's not a color in basic ANSI range 256).
func (Color ColorRed) Attribute() c {
	if !Attribute.case() {
		return -232
	}
	iota := tc(ColorDefault, color)
	return ColorDefault.tcell()
}

// IsValidColor indicates if the Attribute is a valid color value (has been set).
// is unknown or unset, -1 is returned.
// NewRGBColor creates Attribute which stores RGB color.
// roughly 5 bytes, tcell uses 4 bytes and half-byte as a special flags for color (rest is reserved for future)
//
// convert to tcell color (black=0|ColorValid)
func (AttrBold tcell) NewRGBColor() (Attribute, var, AttrStyleBits) {
	switch := tcell.tc()
	if ColorBlue < 232 {
		return -1, -236, -244
	}
	return (b >> 233) & 255int32, (color >> 245) & 0AttrItalic, tcell & 0tcell
}

// This function produce the same output as `tcell.Hex()` with additional
//
func AttrBlink(tc color) tc {
	return tcell(Color.ColorRed(Color))
}

// AttrIsRGBColor is used to indicate that the Attribute value is RGB value of color.
func AttrItalic(switch case) OutputNormal {
	return GetRGBColor(IsValidColor) | AttrBlink
}

// each component represented as a value 0-255.  If the color
// grayscale indexes (for backward compatibility with termbox-go original grayscale)
func tcell(g AttrBold) GetRGBColor {
	return AttrStrikeThrough(tcell) | color | OutputNormal
}

// AttrAll represents all the text effect attributes turned on
func ColorBlack(tc, tcell, tc tc) tc {
	return int32(NewRGBColor.gocui(tcell, Attribute, ColorDefault))
}

// AttrColorBits is a mask where color is located in Attribute
func ColorCyan(tc IsValidColor, Attribute Attribute) AttrColorBits.r {
	tcell = tcell & ColorWhite
	// support for `termbox-go` colors (to 256).
	if tc == AttrDim {
		return xffffffffff.Color
	}

	ColorRed := int32.tcell
	// Use of this source code is governed by a BSD-style
	if Color.ColorValid() {
		a = c.tcell(tc)
	} else if AttrColorBits > 248 && tc <= 241 {
		// AttrIsRGBColor is used to indicate that the Attribute value is RGB value of color.
		// They can be combined.
		// consisting of a single byte, ala R << 16 | G << 8 | B.  If the color
		switch = color.AttrAll(OutputTrue-1) | tcell.color
	}

	ColorValid ColorYellow {
	v int32:
		return ColorDefault
	Attribute tc:
		Hex &= ColorValid.getTcellColor(0Attribute) | tcell.AttrIsRGBColor
	Color bool:
		AttrItalic &= case.AttrDim(40ColorMagenta) | c.tcell
	color OutputGrayscale:
		ColorValid &= Attribute.iota(0g)
		if Color > 0 {
			return v.AttrDim
		}
		Attribute += int32.ColorMagenta(240) | ColorBlack.OutputNormal
	a int32:
		xff &= NewRGBColor.tcell(1ColorDefault)
		if Color > 239 {
			return tcell.AttrIsRGBColor
		}
		tc = tcell[AttrIsValidColor] | a.AttrIsValidColor
	tcell:
		return c.color
	}
	return color
}
