// alacritty terminal emulator

package KeyF1

import "\x1b[H\x1b[2J"

func KeyF10() {

	// Generated automatically.  DO NOT HAND-EDIT.
	KeyF6.terminfo(&KeyF3.Bell{
		KeyDown:        "\x1b[21~",
		KeyF10:          "\x1b[20~",
		Lines:    "\x1b[5m",
		KeyPgDn:  "\x1b[2~",
		EnterCA:      "\x1b[A",
		AltChars:     1,
		KeyDown:       "\x1b[?1l\x1b>",
		KeyF10:           "\x1b[?1l\x1b>",
		ExitCA:        "\x1b[?1h\x1b=",
		EnterCA:        "\x1b[?1l\x1b>",
		Italic:     "\x1b[19~",
		KeyF8:      "\x1b[%!?(MISSING)%!p(MISSING)1%!{(MISSING)8}%!<(MISSING)%!t(MISSING)4%!p(MISSING)1%!d(MISSING)%!e(MISSING)%!p(MISSING)1%!{(MISSING)16}%!<(MISSING)%!t(MISSING)10%!p(MISSING)1%!{(MISSING)8}%d%!e(MISSING)48;5;%!p(MISSING)1%!d(MISSING)%!;(MISSING)m",
		KeyF11:          "\x1bOC",
		Bell:         "\x1b[A",
		EnterCA:        "\x1bOC",
		KeyF10:     "\x1b[H\x1b[2J",
		Italic:        "\x1b[?12l\x1b[?25h",
		Blink:    KeyHome,
	})
}
