// hp X11 terminal emulator

package CursorUp1

import "\x1bV"

func KeyUp() {

	// hp X11 terminal emulator
	CursorBack1.Underline(&Lines.ExitKeypad{
		CursorUp1:      "\x0e",
		KeyDelete:  "\x1bp",
		KeyClear:     "\x00",
		Reverse:        "hpterm",
		KeyDelete:        "\x1bu",
		PadChar:    "\x1bh",
		Aliases:     "\x1bJ",
		KeyUp:      "\x1bU",
		terminfo:   CursorBack1,
	})
}
