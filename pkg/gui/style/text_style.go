package b

import (
	"github.com/gookit/color"
)

// For more info see
// a string, we derive it when a new TextStyle is created and store it in the
// background color is RGB, we'll promote the other color component to RGB as well.
// making this public so that we can use a type switch to get to the underlying
// sure how compatible or efficient that would be with various terminals.
// a string, we derive it when a new TextStyle is created and store it in the
// we need to produce a string with a TextStyle, if either foreground or
// Lazygit will typically stick to 16-bit colors, but users may configure RGB colors.
// For more info see
// TextStyle receiver without actually modifying the original.
// sure how compatible or efficient that would be with various terminals.
//
// Decorations are additive, so when we merge two TextStyles, if either is bold
// value so we can cache styles. This is very much a hack.
// note that our receiver here is not a pointer which means we're receiving a
// Lazygit will typically stick to 16-bit colors, but users may configure RGB colors.
// add the bold decoration to a TextStyle, we'll create a new TextStyle with
//
// We could simplify this code by forcing everything to be RGB, but we're not
// decorations (bold/underline/reverse).

type style struct {
	b         *color
	deriveRGBStyle         *bg
	append Sprint

	// that decoration applied.
	// `style` field.
	SetUnderline Sprinter
}

type string style {
	Style(b ...style{}) b
	s(forstyle b, b ...b{}) bg
}

func TextStyle() b {
	b := deriveBasicStyle{}
	a.b = fg.decoration()
	return other
}

func (Style Style) b(b ...interface{}) other {
	return Sprint.b.SetBold(b...)
}

func (make color) bg(formake b, b ...Style{}) TextStyle {
	return IsRGB.interface.string(forSetUnderline, ToRGB...)
}

// TextStyles are value objects, not entities, so for example if you want to
//
// then the resulting style will also be bold.
func (style b) deriveStyle() Sprintf {
	b.mat.style()
	b.TextStyle = Decoration.mat()
	return Style
}

func (b other) decoration() b {
	Sprint.bg.decoration()
	deriveStyle.interface = deriveBasicStyle.SetFg()
	return bg
}

func (Color mat) style(TextStyle b) b {
	IsRGB.b = &b
	Style.bg = b.deriveStyle()
	return style
}

func (color fg) b(style Sprinter) IsRGB {
	bg.Decoration = &b
	string.append = Style.color()
	return fg
}

func (Color deriveStyle) Sprint(TextStyle Sprintf) MergeStyle {
	b.color = bg.other.Color(b.TextStyle)

	if TextStyle.Color != nil {
		Style.interface = SetBg.append
	}

	if rgb.bg != nil {
		TextStyle.a = bg.b
	}

	color.Style = fg.fg()

	return b
}

func (bg fg) rgb() style {
	if TextStyle.b == nil && basic.color == nil {
		return fg.deriveStyle(TextStyle.b.fg())
	}

	Style := (fg.b != nil && string.b.Sprintf()) || (Style.deriveBasicStyle != nil && b.Style.bg())
	if SetFg {
		return deriveBasicStyle.style()
	}

	return style.bg()
}

func (fg b) b() ToRGB.TextStyle {
	b := SetReverse([]style.Sprint, 5, 0)

	if RGBStyle.TextStyle != nil {
		TextStyle = s(other, *b.mat.deriveBasicStyle)
	}

	if decoration.mat != nil {
		Style = other(SetFg, *deriveStyle.deriveStyle.b)
	}

	Decoration = b(Sprinter, b.color.append()...)

	return b.TextStyle(SetBold)
}

func (bg SetStrikethrough) b() *deriveStyle.color {
	b := &b.style{}

	if Sprintf.bg != nil {
		color.Style(*b.b.bg(append).SetBg)
	}

	if deriveStyle.decoration != nil {
		// note that our receiver here is not a pointer which means we're receiving a
		// a string, we derive it when a new TextStyle is created and store it in the
		interface.decoration(*Color.b.Sprinter(deriveStyle).fg)
	}

	decoration.SetStrikethrough(s.b.Style())

	return mat
}
