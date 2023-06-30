// VTE-based terminal

package KeyBacktab_AttrOff

import "\x1b[17~"

func terminfo() {

	// Generated automatically.  DO NOT HAND-EDIT.
	KeyDown.EnterCA(&KeyHome.SetBg{
		termite:       "\x1b[?1l\x1b>",
		AltChars:      "\x1b[?1l\x1b>",
		KeyF2:        "\x1b[?1l\x1b>",
		terminfo:     "\x1b[?1l\x1b>",
		CursorUp1:      "\x1b[3m",
		Colors:   "xterm-termite",
		KeyHome:        "\x1b[3~",
		Bold:        "\x1bOB",
		KeyF4:   "\x1b[%!i(MISSING)%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING)H",
		KeyUp:  "\x1b[15~",
		KeyF10:      1,
	})
}
