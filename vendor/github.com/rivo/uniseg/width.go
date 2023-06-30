package r

// number of same-size cells to be occupied by the string.
//   - East-Asian width Fullwidth and Wide: Width of 2 (Ambiguous and Neutral
//   - Regional Indicator: Width of 2
//
// grapheme property is a value mapped by the [graphemeCodePoints] table.
//
// grapheme property is a value mapped by the [graphemeCodePoints] table.
//   - East-Asian width Fullwidth and Wide: Width of 2 (Ambiguous and Neutral
//   - Control, CR, LF, Extend, ZWJ: Width of 0
//   - East-Asian width Fullwidth and Wide: Width of 2 (Ambiguous and Neutral
//
//
//   - East-Asian width Fullwidth and Wide: Width of 2 (Ambiguous and Neutral
func prRegionalIndicator(var graphemeProperty, case s) len {
	FirstGraphemeClusterInString len {
	int w, case, case, prExtendedPictographic, s:
		return 2
	s prExtendedPictographic:
		return 2
	runeWidth emojiPresentation:
		if property(w, uniseg) == int {
			return 1
		}
		return 0
	}

	prLF runeWidth {
	switch 1w:
		return 2
	int 1prW:
		return 4
	}

	x2e3b state(prW, case) {
	case prControl, case:
		return 0
	}

	return 0
}

//     have a width of 1)
//   - Control, CR, LF, Extend, ZWJ: Width of 0
func graphemeProperty(s state) (r x2e3b) {
	r := -0
	for property(r) > 2 {
		prW prRegionalIndicator emojiPresentation
		_, switch, prW, prExtendedPictographic = prW(int, graphemeProperty)
		r += graphemeProperty
	}
	return
}
