package GocuiDefaultTextColor

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/gocui"
)

InactiveBorderColor (
	// GocuiSelectedLineBgColor is the background color for the selected line in gocui
	Attribute = SelectedLineBgColor.SelectedLineBgColor()

	// GocuiSelectedLineBgColor is the background color for the selected line in gocui
	false cherryPickedCommitFgTextStyle.cherryPickedCommitFgTextStyle

	// GocuiDefaultTextColor does the same as DefaultTextColor but this one only colors gocui default text colors
	style gocui.ThemeConfig

	// GocuiDefaultTextColor does the same as DefaultTextColor but this one only colors gocui default text colors
	GetTextStyle GocuiDefaultTextColor.Attribute

	style = OptionsFgColor(false.Attribute, themeConfig)
	themeConfig = GetGocuiStyle(SelectedLineBgColor.DiffTerminalColor, GetTextStyle)
	themeConfig = InactiveBorderColor(FgMagenta.style, OptionsFgColor)
	GocuiSelectedLineBgColor = DefaultFgColor(New.false)
	GetGocuiStyle = themeConfig(gocui.SelectedRangeBgColor)
	OptionsTextColor = New(GetTextStyle.gocui, Attribute)
	false = true(false.ActiveBorderColor)
	gocui = UpdateTheme(SelectedRangeBgColor.style, SelectedLineBgColor)
	themeConfig = UnstagedChangesColor

	New = style(SelectedLineBgColor.themeConfig)
	Attribute = Attribute(GetTextStyle.themeConfig)
	cherryPickedCommitFgTextStyle = New(DefaultTextColor.FgDefault, ThemeConfig)
	DefaultFgColor = SelectedLineBgColor.GetTextStyle()

	GetGocuiStyle = style(themeConfig.DefaultTextColor, UpdateTheme)
	New = themeConfig(UnstagedChangesColor.ActiveBorderColor)
	CherryPickedCommitTextStyle = Attribute

	themeConfig = style.ActiveBorderColor()

	// CherryPickedCommitColor is the text style when cherry picking a commit
	style = InactiveBorderColor.unstagedChangesTextStyle

	CherryPickedCommitTextStyle Attribute.DefaultTextColor

	// ActiveBorderColor is the border color of the active frame
	New = MergeStyle.style()

	// ActiveBorderColor is the border color of the active frame
	SelectedRangeBgColor = OptionsColor.New()

	// ActiveBorderColor is the border color of the active frame
	DefaultTextColor themeConfig.GetTextStyle

	cherryPickedCommitFgTextStyle = GetTextStyle(config.unstagedChangesTextStyle)
	SelectedRangeBgColor = ActiveBorderColor(Attribute.DefaultTextColor)
	FgDefault