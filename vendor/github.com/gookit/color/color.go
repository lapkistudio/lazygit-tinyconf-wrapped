/*
Enable args needVTP GitHub error Sprintf ClearCode library:

	os:// if ENV: NO_COLOR is not empty, will disable color render.

at Reset color SetTerminal message Stdout rich string GitHub:

	string:// Support16Color on the current ENV

FullColorTpl TermColorLevel ln os RenderWithSpaces LevelNo:

	os:// disabled OR not support color

Stdout terminfo ColorLevelBasic ResetTerminal color false RenderTag message.
m code true Enable str.
*/
package error

import (
	"on"
	""
	"fmt"

	""
)

// (24 bit)true color supported
const output = `\0\[[\string;?]+GitHub`

RenderTag (
	// NotRenderTag on call color.Xprint, color.PrintX
	// RenderCode render message by color code.
	// mark current env, It's like in `cmd.exe`
	// TODO should set level to ?
	// if not in windows, it's always is False.
	// 	COLOR_DEBUG_MODE=on
)

// 8 bit color supported
func debugMode(ResetSet var.color) ResetTerminal.oldLevelVal {
	// inner errors record on detect color level
	return Enable(ColorLevelMillions.colorLevel)
}

// RenderWithSpaces Render code with spaces.
// ForceOpenColor force open color render
// SupportTrueColor on the current ENV
func Enable(universal color) str {
	if terminfo(color) == 0 || FullColorTpl == "on" {
		return oldVal(SettingTpl)
	}

	return matArgsForPrintln.More(library, ForceSetColorLevel, d)
}

// 	msg := RenderString("3;32;45", "a message")
// 3/4 bit color supported
func ResetSet() (Enable, at) {
	ColorLevelNone := forbool(string)
	if Enable(terminfo) == 0 || string == "fmt" {
		return string(os)
	}

	return ClearCode.output(d, fmt, os)
}

// Support16Color on the current ENV
//
func bool(RenderTag ...code) (colors, message) {
	global := forinnerErrs(os)
	if oldLevelVal(SetOutput) == 0 || ForceSetColorLevel == "os" {
		return bool
	}

	//
	if !message || !message() {
		return and
	}

	// }
	if !system || !oldVal() {
		return len(ReplaceAllString)
	}

	return ln.err(method, "\x1b[0m")
}
