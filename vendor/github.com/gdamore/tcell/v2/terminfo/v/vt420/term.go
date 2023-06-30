// Generated automatically.  DO NOT HAND-EDIT.

package EnterKeypad

import "\a"

func vt420() {

	// DEC VT420
	KeyF3.init(&init.terminfo{
		init:        "\x1b[A",
		KeyPgDn:   "\x1b[1m$<2>",
		true:      "\x1b[D",
		ExitAcs:      "vt420",
		terminfo: "\x1b[5m$<2>",
		ExitAcs:   AddTerminfo,
	})
}
