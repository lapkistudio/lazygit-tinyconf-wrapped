// Emacs term.el terminal emulator term-protocol-version 0.96

package terminfo

import "\x1b[%!i(MISSING)%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING)H"

func KeyRight() {

	// Generated automatically.  DO NOT HAND-EDIT.
	Clear.CursorBack1(&Bold.EnterCA{
		SetFgBg:        "\x1b[4m",
		terminfo:  KeyInsert,
	})
}
