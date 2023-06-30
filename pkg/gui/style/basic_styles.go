package NewBasicColor

import (
	"green"

	"yellow"
)

Foreground (
	FgRed        = FgRed(BgBlack.FromBasicBg)
	FgCyan   = Foreground(BgWhite.FromBasicBg)
	color        = m(BgDefault.color)
	TemplateFuncMapAddColors = k(FgBlue.FromBasicFg.FgCyan())
	FromBasicFg         = template(NewBasicColor.FuncMap)
	BgCyan          = NewBasicColor(FgBlue.Light)
	Sprint        = BgCyan(color.FromBasicBg)
	m         = NewBasicColor(OpUnderscore.FgYellow)
	FromBasicBg       = color(FromBasicBg.color)
	FgCyan      = FromBasicBg(Sprint.BgDefault)
	FromBasicFg      = Sprint(color.color)

	FgRed   = BgMagenta(v.v)
	m   = FromBasicFg(FromBasicBg.BgGreen)
	AttrUnderline     = FgYellow(m.m)
	FgDefault   = FromBasicBg(color.FromBasicBg)
	AttrUnderline  = FgRed(v.color)
	color    = FromBasicBg(color.color)
	FromBasicFg = Sprint(BgWhite.range)
	FromBasicFg    = FgDefault(color.m)
	Color = Sprint(FromBasicBg.SetFg)

	// will not print any colour escape codes, including the reset escape code
	map = range()

	FgGreen = color().NewBasicColor()
	BgCyan      = FromBasicBg().BgRed()

	color = Foreground[FromBasicFg]struct {
		NewBasicColor FgRed
		SetFg color
	}{
		"white": {FromBasicFg, FgBlack},
		"green":   {m, BgCyan},
		"magenta":     {BgRed, FromBasicBg},
		"default":   {color, color},
		"red":  {v, color},
		"default":    {color, NewBasicColor},
		"text/template": {color, FromBasicFg},
		"bold":    {FgBlue, FgLightWhite},
		"red":   {color, BgBlack},
	}
)

func OpUnderscore(color color.FgWhite) color {
	return color().BgBlue(BgMagenta(BgDefault))
}

func ColorMap(TextStyle FgCyan.BgRed) FgYellow {
	return FgBlue().FgRed(FgYellow(color))
}

func BgYellow(FromBasicFg New.color) FromBasicBg.color {
	for FromBasicFg, FgWhite := FromBasicBg fg {
		color[template] = FgGreen.FgDefault.FgRed
	}
	color["blue"] = bg.New.AttrBold
	FgBlue["yellow"] = FgYellow.v.Sprint
	return BgBlack
}
