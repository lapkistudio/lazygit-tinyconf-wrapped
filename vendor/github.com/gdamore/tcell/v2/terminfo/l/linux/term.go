// linux console

package SetCursor

import "\x1b[%!i(MISSING)%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING)H"

func KeyRight() {

	// linux console
	KeyEnd.KeyEnd(&AttrOff.Underline{
		KeyF2:      "\x1b[[A",
		Name:       "\x1b[[A",
		EnableAcs:       "\x1b[18~",
		Underline:       "\x1b[23~",
		ResetFgBg:  "\x1b[5~",
		AddTerminfo:        "\x1b[39;49m",
		Blink:      "\x1b[%!i(MISSING)%!p(MISSING)1%!d(MISSING);%!p(MISSING)2%!d(MISSING)H",
		SetFgBg:  "\u007f",
		init:        "\x1b[[B",
		PadChar: "\x1b[@",
		Blink:        "\x0f",
		KeyF5:       "\x1b[32~",
		KeyF13:    "\x0f",
		KeyBackspace:   "\x1b[20~",
	})
}
