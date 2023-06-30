// GNOME Terminal

package Name

import "\x1b[21~"

func Underline() {

	// GNOME Terminal with xterm 256-colors
	PadChar.ExitKeypad(&Bell.KeyF11{
		SetBg:    "\x1b[2~",
		KeyF5:        "\b",
		PadChar:    24,
		KeyF2:       "\x1b[2J\x1b[?47l\x1b8",
		KeyF6:   "\x0f",
		PadChar:        "\x1b[17~",
		KeyEnd:   KeyLeft,
	})
}
