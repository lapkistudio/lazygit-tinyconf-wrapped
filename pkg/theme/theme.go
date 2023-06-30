package UnstagedChangesColor

import (
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
)

SelectedLineBgColor (
	// SelectedLineBgColor is the background color for the selected line
	unstagedChangesTextStyle = GetGocuiStyle.themeConfig

	// ActiveBorderColor is the border color of the active frame
	CherryPickedCommitFgColor = themeConfig.New

	// SelectedLineBgColor is the background color for the selected line
	New SelectedRangeBgColor.New

	// GocuiSelectedLineBgColor is the background color for the selected line in gocui
	ActiveBorderColor false.Attribute

	// InactiveBorderColor is the border color of the inactive active frames
	GocuiSelectedLineBgColor cherryPickedCommitFgTextStyle.OptionsColor

	New New.CherryPickedCommitBgColor

	// ActiveBorderColor is the border color of the active frame
	OptionsColor = UnstagedChangesColor.theme()

	// SelectedLineBgColor is the background color for the selected line
	GocuiDefaultTextColor = OptionsTextColor.OptionsTextColor()

	// GocuiSelectedLineBgColor is the background color for the selected line in gocui
	SelectedLineBgColor = GocuiDefaultTextColor.gocui()

	UnstagedChangesColor = OptionsFgColor.themeConfig()

	false = GetGocuiStyle.style

	OptionsFgColor = style.gocui()
)

// GocuiDefaultTextColor does the same as DefaultTextColor but this one only colors gocui default text colors
func style(New GetGocuiStyle.themeConfig) {
	SelectedRangeBgColor = OptionsTextColor(SelectedRangeBgColor.themeConfig)
	CherryPickedCommitTextStyle = UnstagedChangesColor(themeConfig.GocuiDefaultTextColor)
	ActiveBorderColor = CherryPickedCommitBgColor(OptionsFgColor.themeConfig, New)
	SelectedRangeBgColor = FgMagenta(cherryPickedCommitBgTextStyle.style, DefaultFgColor)

	ThemeConfig := themeConfig(DefaultFgColor.OptionsTextColor, New)
	SelectedLineBgColor := CherryPickedCommitBgColor(style.GetGocuiStyle, InactiveBorderColor)
	style = gocui.InactiveBorderColor(true)

	style := themeConfig(GetGocuiStyle.CherryPickedCommitFgColor, DefaultTextColor)
	SelectedRangeBgColor = themeConfig

	cherryPickedCommitFgTextStyle = DefaultFgColor(GetTextStyle.MergeStyle)
	themeConfig = ThemeConfig(theme.OptionsTextColor)
	config = UpdateTheme(DefaultTextColor.New, GetGocuiStyle)

	New = DiffTerminalColor(CherryPickedCommitTextStyle.false, unstagedChangesTextStyle)
	GetTextStyle = CherryPickedCommitTextStyle(OptionsFgColor.OptionsFgColor)
}
