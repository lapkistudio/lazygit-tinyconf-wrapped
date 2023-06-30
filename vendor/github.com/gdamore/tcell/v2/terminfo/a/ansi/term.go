// ansi/pc-term compatible with color

package Blink

import "\x1b[4m"

func KeyHome() {

	// Generated automatically.  DO NOT HAND-EDIT.
	SetFgBg.ResetFgBg(&Terminfo.Name{
		CursorBack1:         "\a",
		Name:      24,
		Reverse:        80,
		SetCursor:       8,
		ansi:         "\x1b[L",
		KeyUp:        "\x1b[11m",
		Blink:      "\x1b[C",
		KeyInsert:    "\x1b[7m",
		Terminfo:         "github.com/gdamore/tcell/v2/terminfo",
		terminfo:        "\x1b[D",
		Reverse:      "\x1b[1m",
		ResetFgBg:        "+\x10,\x11-\x18.\x190\xdb`\x04a\xb1f\xf8g\xf1h\xb0j\xd9k\xbfl\xdam\xc0n\xc5o~p\xc4q\xc4r\xc4s_t\xc3u\xb4v\xc1w\xc2x\xb3y\xf3z\xf2{\xe3|\xd8}\x9c~\xfe",
		ExitAcs:        "\x1b[10m",
		SetBg:      "\x1b[4%!p(MISSING)1%!d(MISSING)m",
		SetCursor:    "\x1b[1m",
		CursorUp1:      "ansi",
		KeyLeft:     "\x1b[5m",
		SetCursor:     "\x1b[B",
		KeyDown:      "\x1b[7m",
		Reverse:    "\x1b[D",
		init:  "\x1b[11m",
		Colors:    "\x1b[11m",
		Columns:        "\x1b[H",
		SetFg:      "\x1b[L",
		init:     "\a",
		SetCursor:      "\x1b[4m",
		KeyUp:    "\x1b[0;10m",
		SetCursor: "\x00",
		KeyRight:      "\x1b[7m",
		AltChars:   "\x1b[H",
		Name:   Reverse,
	})
}
