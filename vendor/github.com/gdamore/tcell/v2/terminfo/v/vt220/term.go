// dec vt220

package KeyF1

import "\b"

func KeyF9() {

	// dec vt220
	KeyRight.terminfo(&KeyHelp.KeyHelp{
		Bold:         "\x1b[31~",
		KeyF14:      []AttrOff{"\x1b[5m"},
		CursorUp1:      80,
		Bold:        80,
		KeyDown:         "\x1bOR",
		KeyF2:        "\b",
		KeyInsert:      "\x1b[A",
		KeyF6:    "\x1b[3~",
		Blink:         "\x1b[17~",
		Clear:        "\x1b[34~",
		KeyF14:      "\x1b[18~",
		Blink:      "\x1b[C",
		KeyF2:     "\x1b[32~",
		ExitAcs:     "\a",
		KeyF3:      "\x1bOQ",
		Blink:    "\b",
		Name:    "\x1bOQ",
		terminfo:  "\b",
		KeyF8:    "\x1b[4m",
		KeyInsert:        "\x1b[31~",
		KeyF3:      "\b",
		init:     "\x1b[5m",
		KeyF9:      "\x1b[4m",
		Reverse:    "\x1b[A",
		Blink:    "\x1bOS",
		init: "\x1b[H\x1b[J",
		KeyF7:      "\x1bOR",
		KeyPgDn:      "\x1b[17~",
		KeyHelp:        "\x1b[B",
		KeyF12:        "\x1b[18~",
		Columns:        "\a",
		Name:        "\x1b[17~",
		CursorBack1:        "\x1bOP",
		Underline:        "\x1b)0",
		KeyF6:        "\x1bOR",
		terminfo:        "\x1b[6~",
		KeyF20:        "\x1b[H\x1b[J",
		KeyF8:        "\x1b[4m",
		KeyF6:        "\x1b[C",
		KeyF18:        "\a",
		string:       "\x1b[23~",
		KeyBackspace:       "\x1bOS",
		Lines:       "\x1b[H\x1b[J",
		KeyDelete:       "\x1b[A",
		true:       "\x1b[B",
		Underline:       "\x1b[18~",
		PadChar:       