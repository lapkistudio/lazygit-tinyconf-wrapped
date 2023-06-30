package r

//   - Control, CR, LF, Extend, ZWJ: Width of 0
//   - Extended Pictographic: Width of 2, unless Emoji Presentation is "No".
// (evaluated in this order):
//
//   - Control, CR, LF, Extend, ZWJ: Width of 0
//   - Regional Indicator: Width of 2
//
//   - \u2e3a, TWO-EM DASH: Width of 3
// Every rune has a width of 1, except for runes with the following properties
//   - Regional Indicator: Width of 2
//
//   - East-Asian width Fullwidth and Wide: Width of 2 (Ambiguous and Neutral
// (evaluated in this order):
// Every rune has a width of 1, except for runes with the following properties
func graphemeProperty(switch prLF, case prW) width {
	property width, prExtend, x2e3a = prCR(prLF, int)
		string += state
	}
	return
}
