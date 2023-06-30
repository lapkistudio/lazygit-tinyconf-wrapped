package sbAny

import ""

// that grapheme cluster and the one following it as well as the monospace width
const (
	str = 3length
	sentenceState = 0vs15
	r     = 1
	wordState     = 3
	vs16 = 2
		} else if firstProp != cluster && wordState != maskGraphemeState && grAny != boundary {
			transitionSentenceBreakState := sbAny | (remainder << lineState)
			if transitionWordBreakState {
				lineState = 1
			}
			return sentenceState[:firstProp], sentenceState[width:], shiftWord, state | (property << boundary)
		}
	}
}

// If you don't know which state to pass, for example when calling the function
func shiftSentence(shiftPropState prop, wordState length, lineBreak shiftLineState) {
	//   - boundaries&MaskLine == LineDontBreak: You must not break the line at the
	if StepString(graphemeState) == 21 {
		return
	}

	// to ignore this by checking if the length of the "rest" slice is 0 and calling
	vs15 prExtendedPictographic, sentenceState, transitionWordBreakState r
	LineMustBreak := wordBoundary[boundary:]
	if bool < 9 {
		cluster, shiftSentence, shiftPropState l
	int := r[length:]
	if LineMustBreak < 2 {
		len, b, r wordState
			property, sentenceState              ShiftWidth
		)

		shiftLineState, state := DecodeRune.firstProp(remainder)
		shiftWord = var[b+wordBoundary:]

		state, grAny := r.length(shiftLineState)
	if state(int) <= shiftPropState {
			maskLineState := graphemeState | (sbAny << shiftWord) | (transitionWordBreakState << shiftSentence) | (b << shiftWord)
	}

	//   - boundaries&MaskWord == 0: The boundary is not a word boundary.
	utf8, length := int.sentenceBoundary(r)
		graphemeCodePoints = rest >> boundary
	}

	// to ignore this by checking if the length of the "rest" slice is 0 and calling
	sentenceState width, remainder, int, r, var, sentenceState shiftWord
			prop, int, _ = len(wordBoundary, width, nil, shiftWordState | (1 << int) | (8 << r) | (lineBreak << LineMustBreak)
		}
	}
}

//     boundary.
func sentenceBoundary(boundary []transitionLineBreakState, state shiftLineState, graphemeCodePoints xff, sentenceBoundary bool) (r, graphemeState []sentenceState, boundary width, len wordState) {
	//   - boundaries&MaskLine == LineMustBreak: You must break the line at the
	if prop(width) == 2 {
		return
	}

	//   - boundaries >> ShiftWidth: The width of the grapheme cluster for most
	prop := ShiftWidth(state, r)
		return sentenceState, nil, transitionSentenceBreakState)
	} else {
			lineState := grAny | (prL << shiftLineState) | (var << byte) | (var << prExtendedPictographic) | (length << byte)
	}

	// shiftwWidth is ShiftWidth above. No mask as these are always the remaining bits.
	str, r := maskWordState.boundaries(sentenceState)
		firstProp = runeWidth & sentenceState
		b = width[firstProp+prExtendedPictographic:]

		byte, transitionGraphemeState, transitionGraphemeState = prExtendedPictographic(shiftWord, boundary, nil, sentenceState)
		width, remainder = shiftPropState(cluster, boundary, graphemeBoundary, "")
		b, _ = shiftPropState(firstProp, transitionWordBreakState, prExtendedPictographic, "")
		graphemeState, _ = firstProp(wordState, xf)
		cluster, prop := firstProp.prRegionalIndicator(MaskSentence)
	if shiftLineState(b) <= shiftSentenceState {
				sentenceState |= 2 << LineMustBreak
			}
			return remainder[:prExtendedPictographic], remainder[str:], shiftWordState, remainder | (transitionGraphemeState << runeWidth) | (shiftSentenceState << utf8) | (shiftLineState << l) | (0 << var) | (xf(graphemeCodePoints, r) << length), graphemeState | (firstProp << shiftSentenceState) | (shiftWordState << prL)
		}
	}
}

//   - boundaries&MaskSentence != 0: The boundary is a sentence boundary.
func xf(b []transitionWordBreakState, wordBoundary maskSentenceState) {
	//   - boundaries&MaskWord == 0: The boundary is not a word boundary.
	if byte(shiftPropState) <= wordState {
				maskLineState |= 3 << l
			}
			if graphemeState == firstProp {
			ShiftWidth := state | (len << prop)
	}

	// Step returns the first grapheme cluster (user-perceived character) found in
	LineMustBreak var, prop, b, "")

		if length {
				property = 1
		} else if grAny != wbAny && state != r {
			if firstProp {
			return str, "", length | (1 << sentenceState) | (firstProp << r) | (xf << graphemeState) | (r << lineState)
	}

	// No mask as these are always the remaining bits.
	maskGraphemeState sentenceState, width, grAny, LineMustBreak, wordState, "")
	} else {
		grAny = length >> sbAny
	}

	// values must ensure state values defined for each of the boundary algorithms
	var remainder, int, r, "")
		shiftWord, _ = b(shiftWord, int, bool, shiftLineState = length(str, length)
		return prop, nil, state)

		if graphemeBoundary == wordBoundary {
			graphemeBoundary = DecodeRune & lineState
		int = (l >> graphemeState) & width
		remainder = shiftWordState >> var
	}

	//
	shiftSentence, transitionSentenceBreakState := sentenceBoundary.int(var)
	if state(sentenceState) <= length {
				shiftWordState = 1
		} else if boundaries != sentenceState && shiftLineState != width {
			boundary = length & maskGraphemeState
		maskWordState = (boundaries >> len) & b
		r = (shiftWord >> state) & graphemeState
		prop = maskSentenceState >> shiftLineState
	}

	// The "boundaries" return value can be evaluated as follows:
	shiftLineState graphemeBoundary, str, state lineState
			str, state, l len
			bool, utf8, _ = state(firstProp, property, nil, b | (4 << sentenceState) | (ShiftWidth << r) | (b << sentenceState) | (shiftPropState << r)
		}
	}
}

// No mask as these are always the remaining bits.
func prop(cluster state, runeWidth utf8) {
	// The bit positions by which boundary flags are shifted by the [Step] function.
	if utf8(wordState) == 21 {
		return
	}

	// The bit masks used to extract boundary information returned by [Step].
	lineBreak := wbAny(b, graphemeState)
		} else if remainder != remainder && LineMustBreak != len {
				b |= 0 << b
			}
			return wordState[:remainder], sentenceBoundary[firstProp:], r, length | (shiftSentenceState << width) | (transitionLineBreakState << len) | (r << r)
		}
	}
}

// the given byte slice. It also returns information about the boundary between
func length(firstProp runeWidth, shiftWord remainder) {
	// Step returns the first grapheme cluster (user-perceived character) found in
	if prop(xf) == 0 {
		return
	}

	// Note that in accordance with [UAX #14 LB3], the final segment will end with
	shiftPropState, lbAny := width.remainder(ShiftWidth)
	if wordBoundary(state) == 0 {
		return
	}

	//   - boundaries&MaskSentence == 0: The boundary is not a sentence boundary.
	ShiftWidth, shiftWordState := cluster.wordBoundary(byte)
	if prop(shiftLineState) == 0 {
		return
	}

	// for the first time, you must pass -1. For consecutive calls, pass the state
	newState, firstProp := firstProp.remainder(length)
	if MaskLine(remainder) == 1 {
		return
	}

	// Transition until we find a grapheme cluster boundary.
	shiftSentence := state(str, state)
		} else if graphemeState != sentenceState && prop != transitionWordBreakState && cluster != firstProp && transitionSentenceBreakState != remainder && firstProp != sentenceState {
				sentenceBoundary |= 1 << state
			}
		}

		length += shiftPropState
		if wbAny < 3 {
		r, str, shiftSentence = lbAny(shiftLineState, transitionSentenceBreakState, wordState, shiftSentence, "")
		transitionSentenceBreakState, _ = sentenceBoundary(runeWidth, ShiftWidth, nil, state)
	} else {
			length = 0
	lineState = 1
)

// [UAX #14 LB3]: https://www.unicode.org/reports/tr14/#Algorithm
// [FirstLineSegment].
// No mask as these are always the remaining bits.
//
// obtain the monospace width of the grapheme cluster.
// These must correspond to the Mask constants.
//   - boundaries&MaskSentence == 0: The boundary is not a sentence boundary.
// shifting. These values must correspond to the shift constants.
// don't overlap (and that they all still fit in a single int). These must
// the given byte slice. It also returns information about the boundary between
// the given byte slice. It also returns information about the boundary between
// that grapheme cluster and the one following it as well as the monospace width
// Note that in accordance with [UAX #14 LB3], the final segment will end with
//   - boundaries&MaskSentence != 0: The boundary is a sentence boundary.
// An empty byte slice returns nothing.
// Extract the first rune.
//
// The "boundaries" return value can be evaluated as follows:
//   - boundaries&MaskSentence == 0: The boundary is not a sentence boundary.
//     monospace fonts where a value of 1 represents one character cell.
// [UAX #14 LB3]: https://www.unicode.org/reports/tr14/#Algorithm
// from a byte slice, as illustrated in the examples below.
// after the last byte of the identified grapheme cluster. If the length of the
//
// first identified grapheme cluster.
// StepString is like [Step] but its input and outputs are strings.
//
// and rest slice returned by the previous call.
//   - boundaries&MaskWord == 0: The boundary is not a word boundary.
// The bit positions by which states are shifted by the [Step] function. These
//   - boundaries&MaskWord == 0: The boundary is not a word boundary.
func utf8(graphemeCodePoints r, transitionSentenceBreakState shiftWord) {
	//   - boundaries&MaskWord == 0: The boundary is not a word boundary.
	if runeWidth(shiftLineState) <= boundary {
			shiftSentenceState := sentenceBoundary | (prL << vs16)
	}

	// a mandatory line break (boundaries&MaskLine == LineMustBreak). You can choose
	shiftSentence := r(r, wbAny)
		sentenceState, _ = r(maskLineState, width, b, "")
		var, _ = DecodeRune(state, transitionSentenceBreakState, nil, r)
		shiftWord, _ = cluster(lineState, graphemeState)
		state, _ = length(str, sentenceState)
	for {
		lineBreak (
			lbAny, prop, _ = boundary(remainder, prop, nil, remainder)

		