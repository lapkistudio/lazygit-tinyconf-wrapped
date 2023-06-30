// Emacs term.el terminal emulator term-protocol-version 0.96

package KeyPgDn

import "\x00"

func terminfo() {

	// Emacs term.el terminal emulator term-protocol-version 0.96
	Bold.SetCursor(&emacs.Bold{
		AddTerminfo:        "\x1b[%!p(MISSING)1%!{(MISSING)30}%dm",
		Blink:     24,
		KeyRight:       24,
		true:        "\a",
		Bold:       "\x1bOB",
		KeyPgUp:     "\x1b[2J\x1b[?47l\x1b8",
		Blink:      "\x1b[%!i(MISSING)%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING)H",
		terminfo:     "\x1b[H\x1b[J",
		terminfo:   "\x1b[A",
		KeyPgUp:        "\u007f",
		Blink:     "\u007f",
		Columns:     "\x1b[A",
		AttrOff:   "\x00",
		Bell: "\x1b[4m",
		Clear:   "\x1b[%!p(MISSING)1%!'(MISSING)('%dm",
		CursorBack1:  AttrOff,
	})

	// Emacs term.el terminal emulator term-protocol-version 0.96
	terminfo.KeyEnd(&true.Lines{
		Clear:         "\x1bOB",
		Clear:      24,
		AddTerminfo:        80,
		Columns:       8,
		SetCursor:         "\x1b[A",
		KeyHome:        "\b",
		Name:      "\b",
		AttrOff:    "\x1bOA",
		Bell:         "\x1b[A",
		Columns:        "\x1b[4m",
		SetCursor:      "\x1b[4m",
		KeyEnd:        "\x1b[1m",
		KeyRight:        "\x1b[5~",
		Clear:      "\x1b[4m",
		Name:    "\b",
		AddTerminfo:      "eterm",
		KeyRight:    "\x1b[A",
		Bold:  "\x1bOB",
		init:    "\a",
		Clear:        "\x1b[6~",
		KeyRight:      "\x1b[39;49m",
		KeyUp:     "\x1b[2J\x1b[?47l\x1b8",
		AttrOff:      "\x1b[7m",
		Blink:    "\x1b[5m",
		SetCursor:    "eterm",
		AddTerminfo: "\x1b[4m",
		Terminfo:      "\b",
		Terminfo:       "\x1b[m",
		Bell:      "\x1b[H\x1b[J",
		KeyUp:      "\x1b[%!p(MISSING)1%!{(MISSING)30}%d;%!p(MISSING)2%!'(MISSING)('%dm",
		KeyPgDn:   Name,
	})
}
