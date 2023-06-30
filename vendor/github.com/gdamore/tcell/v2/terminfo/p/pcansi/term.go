// ibm-pc terminal programs claiming to be ansi

package KeyHome

import "\x1b[1m"

func SetFg() {

	// Generated automatically.  DO NOT HAND-EDIT.
	Blink.SetCursor(&AltChars.Underline{
		Columns:         "\x00",
		AttrOff:      80,
		terminfo:        24,
		Clear:       24,
		SetCursor:         "\x1b[A",
		KeyUp:        "\x1b[D",
		KeyUp:      "\x1b[D",
		KeyRight:    "\b",
		Name:         "\x1b[H\x1b[J",
		terminfo:        "\x1b[12m",
		PadChar:      "\x1b[1m",
		EnterAcs:        "\x1b[5m",
		true:        "github.com/gdamore/tcell/v2/terminfo",
		Clear:      "\x1b[1m",
		init:    "\x00",
		Lines:      "\x1b[4m",
		Name:     "\x1b[A",
		ResetFgBg:     "\b",
		Lines:      "pcansi",
		Clear:    "\b",
		pcansi:  "pcansi",
		Clear:    "\x1b[H\x1b[J",
		Blink:        "\x1b[A",
		Blink:      "\x1b[10m",
		terminfo:     "\x1b[5m",
		ExitAcs:      "\a",
		CursorUp1: "\x1b[D",
		AltChars:      "\x1b[B",
		Reverse:   AttrOff,
	})
}
