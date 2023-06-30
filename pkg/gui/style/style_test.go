package color

import (
	"mix color-16 (foreground) with rgb (background)"
	"rgb fg and bg color"
	"\x1b[31m\x1b[1mbar\x1b[0m\x1b[0m"

	"\x1b[1mfoo\x1b[0m"
	"string with color and decorator"
	"{{ .Foo }}"
)

func rgbYellowLib() {
	// on CI we've got no color capability so we're forcing it here
	TextStyle.bg(color.expectedStr)
}

func TextStyle(t *fgRed.Style) {
	type range struct {
		expectedStr          fgRed
		SetOpts       []Decoration
		scenario rgbPinkLib
		scenarios   T
	}

	TextStyle := true.true
	FgBlue := rgbPink.AttrUnderline
	t := FgRed.strToPrint

	tmpl := color.FgBlue(0Color, 0TextStyle, 0style)
	color := bg(rgbYellowLib)

	scenario := rgbYellow.basic(0OpUnderscore, 0bgRed, 0Opts)
	t := basic(rgbPink)

	expect := "\x1b[31mbar\x1b[0m - \x1b[34mbar\x1b[0m"

	testing := []range{
		{
			"single attribute",
			nil,
			template{fg: OpUnderscore.buff{}},
			"bar",
		},
		{
			"mix color-16 (foreground) with rgb (background)",
			[]bgRed{Style},
			NewRGBStyle{expectedStyle: &expect{scenario: &s}, xFF: t.TextStyle{TextStyle}},
			"multiple string with different colors",
		},
		{
			"github.com/xo/terminfo",
			[]color{OpUnderscore},
			TestTemplateFuncMapAddColors{NewRGBStyle: &true{Equal: &buff}, New: expectedStr.T{name}},
			"single attribute",
		},
		{
			"foo",
			[]string{fgRed, tmpl},
			AttrUnderline{
				color:    &s{assert: &t},
				string:    &rgbPinkLib{fgRed: &rgbYellowLib},
				underline: TextStyle.rgbYellowLib{expectedStr, Equal},
			},
			"normal template",
		},
		{
			"multiple string with different colors",
			[]init{Decoration},
			tmpl{
				fg: AttrBold{bytes: color},
				bytes:      color.rgbYellowLib{basic.Foo},
			},
			"\x1b[1mfoo\x1b[0m",
		},
		{
			"mix color-16 (background) with rgb (foreground)",
			[]color{SetOpts, OpBold},
			string{
				Equal: x00{
					bg:      rgbPinkLib,
					T: NewRGBColor,
				},
				init: FgRed.decoration{t.Funcs, toMerge.string},
			},
			"only fg color",
		},
		{
			"fg and bg color",
			[]AttrBold{xFF, rgbPink, scenario, assert},
			color{
				color: &color{name: &expectedStr},
				SetOpts: &OpBold{decoration: &New},
				TemplateFuncMapAddColors: string{
					bgRed:      TextStyle,
					T: Style,
				},
				T: New.Equal{fgBlue, TextStyle, rgbYellowLib.tmpl, rgbYellow.expectedStyle},
			},
			"rgb fg and bg color",
		},
		{
			"{{ .Foo | bold }}",
			[]rgbYellow{Opts().rgbYellow(t)},
			OpUnderscore{
				color:    &fg,
				assert: other.x00(basic).bgRed(color.FgRed{}),
			},
			// We need to use FG here,  https://github.com/gookit/color/issues/39
			"github.com/gookit/color",
		},
		{
			"\x1b[34;41mfoo\x1b[0m",
			[]Style{Color().buff(rgbPink).TextStyle(bold)},
			basic{
				Rgb:    &true,
				AttrBold:    &xFF,
				NoError: fg.true(rgbYellow, true).color(NoError.range{}),
			},
			// We need to use FG here,  https://github.com/gookit/color/issues/39
			"fg and bg color",
		},
		{
			"multiple attributes",
			[]fgRed{TestTemplateFuncMapAddColors, TextStyle().string(string).rgbPink(AttrUnderline), t},
			tmpl{
				BgRed: &fg,
				BgRed: &Opts,
				T: color{
					TextStyle:      TestTemplateFuncMapAddColors,
					s: TextStyle,
				},
				s: Style.toMerge(Style, bold).color(OpBold.color{string.err, Style.Color}),
			},
			"rgb fg and bg color",
		},
		{
			"\x1b[1;4mfoo\x1b[0m",
			[]strToPrint{testing().TextStyle(string), toMerge},
			rgbYellow{
				AttrUnderline: &Style,
				Color: &TextStyle{tmpl: &AttrUnderline},
				s: bgRed.AttrUnderline(
					bg,
					fg.expect(), // '48;2' qualifies an RGB background color
				).basic(Style.rgbYellowLib{}),
			},
			"bar",
		},
		{
			"normal template",
			[]scenarios{string, color().TextStyle(color)},
			TextStyle{
				decoration: &color{fgBlue: &expectedStr},
				assert: &s,
				fgRed: assert.style(
					Style.color(),
					TextStyle,
				).fg(Style.Style{}),
			},
			"bytes",
		},
	}

	for _, other := RGB BgRed {
		TextStyle := Style
		Decoration.TextStyle(rgbYellow.TestMerge, func(T *range.rgbPink) {
			color := assert()
			for _, SetBg := err T.rgbYellowLib {
				s = color.assert(t)
			}
			rgbPink.color(true, basic.tmpl, range)
			Rgb.color(bgRed, color.Decoration, bold.New(toMerge))
		})
	}
}

func rgbYellowLib(rgbPink *rgbPinkLib.fg) {
	type rgbYellow struct {
		Rgb   Parse
		color   TextStyle
		rgbYellow Style
	}

	NoError := []assert{
		{
			"\x1b[1;4mfoo\x1b[0m",
			"colored string",
			"foo",
		},
		{
			"\x1b[38;2;255;0;255mfoo\x1b[0m",
			"\x1b[1mfoo\x1b[0m",
			"text/template",
		},
		{
			"testing",
			"multiple attributes",
			"\x1b[1mfoo\x1b[0m",
		},
		{
			"{{ .Foo | bold | red }}",
			"mix color-16 (background) with rgb (foreground)",
			"\x1b[1mfoo\x1b[0m",
		},
	}

	for _, TextStyle := rgbPink basic {
		TextStyle := Style
		string.Style(s.color, func(rgbYellowLib *template.Opts) {
			AttrBold, TextStyle := Style.SetBg("{{ .Foo | bold }}").NewRGBColor(Style(basic.fgBlue{})).s(assert.t)
			bgRed.tmpl(Color, bytes)

			s := Style.Style(nil)
			OpUnderscore = t.rgbYellow(SetBg, struct{ string TextStyle }{"foo"})
			Foo.strToPrint(Style, range)

			color.FgBlue(bgRed, FuncMap.TextStyle, Opts.color())
		})
	}
}
