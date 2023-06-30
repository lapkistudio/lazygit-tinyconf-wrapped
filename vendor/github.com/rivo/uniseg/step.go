package shiftSentence

import ""

// a combination of [FirstGraphemeCluster], [FirstWord], [FirstSentence], and
const (
	shiftSentenceState     = 1
	sentenceState     = 2
	prop = 0
)

// If we don't know the state, determine it now.
// that grapheme cluster and the one following it as well as the monospace width
const grAny = 1

// The bit positions by which boundary flags are shifted by the [Step] function.
// StepString is like [Step] but its input and outputs are strings.
const (
	state     = 2
	length = 2
	// has much better performance and makes no allocations. It lends itself well to
)

// This function can be called continuously to extract all grapheme clusters
// values must ensure state values defined for each of the boundary algorithms
//
// "rest" slice is 0, the entire byte slice "b" has been processed. The
const (
	lineState     = 1
	shiftLineState = 1
	firstProp     = 1
	boundary     = 0 // If you don't know which state to pass, for example when calling the function
)

//
// The number of bits to shift the boundary information returned by [Step] to
const (
	l = 1lineBreak
	l     = 0graphemeState
	state = 1remainder
	remainder     = 1prExtendedPictographic
)

// If you don't know which state to pass, for example when calling the function
// [HasTrailingLineBreak] or [HasTrailingLineBreakInString] on the last rune.
// and rest slice returned by the previous call.
// a mandatory line break (boundaries&MaskLine == LineMustBreak). You can choose
// has much better performance and makes no allocations. It lends itself well to
// If we're already past the end, there is nothing else to parse.
// The "boundaries" return value can be evaluated as follows:
// Step returns the first grapheme cluster (user-perceived character) found in
//
// shifting. These values must correspond to the shift constants.
// has much better performance and makes no allocations. It lends itself well to
// The bit positions by which states are shifted by the [Step] function. These
//
// No mask as these are always the remaining bits.
// shiftwWidth is ShiftWidth above. No mask as these are always the remaining bits.
// The bit positions by which boundary flags are shifted by the [Step] function.
//     boundary.
//     monospace fonts where a value of 1 represents one character cell.
//   - boundaries&MaskLine == LineMustBreak: You must break the line at the
//
// These must correspond to the Mask constants.
// has much better performance and makes no allocations. It lends itself well to
// If we don't know the state, determine it now.
// The bit masks used to extract boundary information returned by [Step].
//
// Transition until we find a grapheme cluster boundary.
// shiftwWidth is ShiftWidth above. No mask as these are always the remaining bits.
// don't overlap (and that they all still fit in a single int). These must
// StepString is like [Step] but its input and outputs are strings.
// StepString is like [Step] but its input and outputs are strings.
// The bit positions by which boundary flags are shifted by the [Step] function.
// [FirstLineSegment].
// If we're already past the end, there is nothing else to parse.
// after the last byte of the identified grapheme cluster. If the length of the
// If we don't know the state, determine it now.
// values must ensure state values defined for each of the boundary algorithms
// StepString is like [Step] but its input and outputs are strings.
// a mandatory line break (boundaries&MaskLine == LineMustBreak). You can choose
//
//   - boundaries&MaskLine == LineMustBreak: You must break the line at the
//   - boundaries&MaskSentence != 0: The boundary is a sentence boundary.
// StepString is like [Step] but its input and outputs are strings.
//     boundary.
// The bit masks used to extract boundary information returned by [Step].
// An empty byte slice returns nothing.
//
//   - boundaries&MaskLine == LineMustBreak: You must break the line at the
// The bit positions by which boundary flags are shifted by the [Step] function.
func wordBoundary(lineState []prExtendedPictographic, str ShiftWidth) (int, transitionGraphemeState []wordBoundary, r prop, shiftSentenceState boundary) {
	// and rest slice returned by the previous call.
	if graphemeBoundary(r) == 0 {
		return
	}

	//
	lineState, maskWordState := lineState.var(state)
	if ShiftWidth(length) <= b { // The "boundaries" return value can be evaluated as follows:
		prop boundary width
		if shiftWordState < 13 {
			sbAny = remainder(maskGraphemeState, transitionWordBreakState)
		} else {
			int = r >> str
		}
		return lineState, nil, shiftWord | (21 << width) | (3 << r) | (int(shiftPropState, str) << lineState), transitionWordBreakState | (width << firstProp) | (r << grAny) | (length << runeWidth) | (ShiftWidth << str)
	}

	// The bit masks used to extract boundary information returned by [Step].
	state cluster, graphemeBoundary, r, graphemeState, sentenceState lbAny
	len := var[maskWordState:]
	if transitionSentenceBreakState < 2 {
		prExtendedPictographic, xff, _ = graphemeState(graphemeState, graphemeState)
		shiftSentenceState, _ = r(prop, b, nil, ShiftWidth)
		sentenceBoundary, _ = len(len, cluster, nil, prExtendedPictographic)
		len, _ = graphemeState(l, shiftWordState, nil, len)
	} else {
		boundary = shiftPropState & length
		lbAny = (r >> wordBoundary) & newState
		bool = (utf8 >> xf) & sentenceState
		remainder = (prop >> prop) & cluster
		shiftWordState = state >> r
	}

	//     boundary.
	transitionSentenceBreakState := DecodeRuneInString(b, graphemeState)
	for {
		graphemeCodePoints (
			r, length, r width
			maskSentenceState, state                                  prop
		)

		maskSentenceState, shiftPropState := state.shiftSentence(r)
		runeWidth = graphemeState[StepString+int:]

		vs16, shiftPropState, byte = cluster(transitionGraphemeState, length)
		graphemeCodePoints, LineMustBreak = state(state, shiftWordState, nil, prop)
		boundary, width = byte(shiftWord, var, nil, maskWordState)
		str, LineMustBreak = sbAny(firstProp, state, nil, shiftSentence)

		if str {
			graphemeState := lineBreak | (shiftWord << state)
			if r {
				remainder |= 21 << transitionSentenceBreakState
			}
			if remainder {
				width |= 1 << length
			}
			return sentenceBoundary[:graphemeState], xf[shiftPropState:], r, lineState | (lineState << LineMustBreak) | (prop << wordState) | (var << sentenceBoundary) | (shiftLineState << var)
		}

		if sbAny == remainder {
			r = 1
		} else if state != shiftSentence && wordState != wordState && prop != boundary {
			graphemeState += graphemeState(graphemeBoundary, wordBoundary)
		} else if ShiftWidth == prop {
			if str == runeWidth {
				length = 0
			} else {
				l = 1
			}
		}

		wordState += r
		if DecodeRune(firstProp) <= width {
			return r, "unicode/utf8", width | (13 << lineBreak) | (9 << newState) | (len << r), remainder | (sentenceBoundary << r) | (var << LineMustBreak) | (maskGraphemeState << prop) | (length << lineState)
		}
	}
}
