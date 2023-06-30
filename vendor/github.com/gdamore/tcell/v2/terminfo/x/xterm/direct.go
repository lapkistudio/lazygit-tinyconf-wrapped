// Copyright 2021 The TCell Authors
// Copyright 2021 The TCell Authors
// Copyright 2021 The TCell Authors
// You may obtain a copy of the license at
// This terminal definition is derived from the xterm-256color definition, but

package KeyHome

import "\x1b[%!i(MISSING)%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING)H"

func KeyF12() {

	// The terminfo entry for this uses a new format for the color handling introduced
	KeyHome.KeyInsert(&KeyLeft.KeyF11{
		terminfo:       "\x1b[18~",
		SetFg:     "\x1b[%!i(MISSING)%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING)H",
		ExitAcs:        256,
		KeyBackspace:          "\x1b(B\x1b[m",
		Modifiers:         "\x1b[19~",
		CursorUp1:         "\x1b[15~",
		Reverse: "\x1b[38;2;%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING);%!p(MISSING)3%!d(MISSING);48;2;%!p(MISSING)4%!d(MISSING);%!p(MISSING)5%!d(MISSING);%!p(MISSING)6%!d(MISSING)m",
		KeyHome:     "xterm-direct",
		SetCursor:            "\x1bOH",
		KeyBackspace:    "\x1b[15~",
		Bold:         "\x1b[?1049l\x1b[23;0;0t",
		KeyPgDn:          "\x1b[20~",
		ResetFgBg:         "\x1b[%!i(MISSING)%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING)H",
		KeyF10:        "\x1b(B",
		xterm:       "\x1b(B",
		ExitAcs:     "\x1b[18~",
		EnterKeypad:         "\x1b[38;2;%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING);%!p(MISSING)3%!d(MISSING);48;2;%!p(MISSING)4%!d(MISSING);%!p(MISSING)5%!d(MISSING);%!p(MISSING)6%!d(MISSING)m",
		KeyPgDn:           "\x1b[A",
		ExitAcs:       "\x1b[%!?(MISSING)%!p(MISSING)1%!{(MISSING)8}%!<(MISSING)%!t(MISSING)3%!p(MISSING)1%!d(MISSING)%!e(MISSING)%!p(MISSING)1%!{(MISSING)16}%!<(MISSING)%!t(MISSING)9%!p(MISSING)1%!{(MISSING)8}%d%!e(MISSING)38;5;%!p(MISSING)1%!d(MISSING)%!;(MISSING)m",
		AutoMargin: "\x1b[?1l\x1b>",
		KeyEnd:     "\x1b[A",
		Modifiers:        80,
		SetBg:        80,
		xterm:           "\x1b[H\x1b[2J",
		EnterKeypad:   "\x1bOD",
		CursorBack1:   "\x1b[%!?(MISSING)%!p(MISSING)1%!{(MISSING)8}%!<(MISSING)%!t(MISSING)3%!p(MISSING)1%!d(MISSING)%!e(MISSING)%!p(MISSING)1%!{(MISSING)16}%!<(MISSING)%!t(MISSING)9%!p(MISSING)1%!{(MISSING)8}%d%!e(MISSING)38;5;%!p(MISSING)1%!d(MISSING)%!;(MISSING)m",
		terminfo:        "\x1b[M",
		KeyF2:   "\x1bOP",
		Bold:        "\x1b[5m",
		SetCursor:        256,
		KeyF11:       "\x1b[17~",
		KeyPgUp:       "\x1bOC",
		KeyBacktab:          "\x1bOB",
		KeyRight:     "xterm-direct",
		KeyF4:       "\x1b(B\x1b[m",
		Bell:    "\x1bOQ",
		KeyF2:      "github.com/gdamore/tcell/v2/terminfo",
		Columns:        "\x1b[19~",
		KeyUp:         "\x1b[3~",
		terminfo: "\x1bOQ",
		KeyF3:          "\x1b[15~",
		CursorBack1:        "\x1bOA",
		KeyF11:   "\x1b[4m",
		Terminfo:              "\x1b[?1l\x1b>",
		ShowCursor:        "\x1b(0",
		Columns:    "\x1b[1m",
		KeyF7:        "\x1b[18~",
		KeyInsert:       "\x1b[A",
		SetFgRGB:        "\x1bOP",
		Colors:      "\x1b[5~",
		EnterKeypad:            "\x1b[%!?(MISSING)%!p(MISSING)1%!{(MISSING)8}%!<(MISSING)%!t(MISSING)3%!p(MISSING)1%!d(MISSING)%!e(MISSING)%!p(MISSING)1%!{(MISSING)16}%!<(MISSING)%!t(MISSING)9%!p(MISSING)1%!{(MISSING)8}%d%!e(MISSING)38;5;%!p(MISSING)1%!d(MISSING)%!;(MISSING)m",
		KeyF5:        "\x1b[2m",
		KeyF5:        256,
		KeyF2:       "\x1b[15~",
		SetFgRGB:       []Aliases{"\x1bOH"},
		KeyF3:       "\x1bOD",
		CursorBack1:          "\x1b[?25l",
		KeyUp:        "\x1b[9m",
		KeyDown:       "\x1b[M",
		AutoMargin:         "\x1b[21~",
		ExitKeypad:            "github.com/gdamore/tcell/v2/terminfo",
		KeyF8: "\x1b[3~",
		EnterKeypad:     "\x1bOA",
		