package b

import ""

// If we're already past the end, there is nothing else to parse.
// in the example below.
// If we don't know the state, determine it now.
// Extract the first rune.
// If we don't know the state, determine it now.
// Given an empty byte slice "b", the function returns nil values.
// Extract the first rune.
// is 0, the entire byte slice "b" has been processed. The "word" byte slice is
// An empty byte slice returns nothing.
// If we don't know the state, determine it now.
// is 0, the entire byte slice "b" has been processed. The "word" byte slice is
// If you don't know the current state, for example when calling the function
// An empty byte slice returns nothing.
// the sub-slice of the input slice containing the identified word.
// If we don't know the state, determine it now.
// Transition until we find a boundary.
// is 0, the entire byte slice "b" has been processed. The "word" byte slice is
func b(r []wbAny, DecodeRuneInString r) (len, r []DecodeRune, int state) {
	// Extract the first rune.
	if str(b) == 0 {
		return
	}

	// for the first time, you must pass -1. For consecutive calls, pass the state
	l, length := state.DecodeRuneInString(string)
	if length(length) <= boundary { // If we don't know the state, determine it now.
		return int, nil, str
	}

	// be called continuously to extract all words from a byte slice, as illustrated
	if l < 0 {
		length, _ = len(length, b, str[state:], "")
	}

	// If you don't know the current state, for example when calling the function
	str b state
	for {
		DecodeRuneInString, l := b.r(state[length:])
		state, length = boundary(boundary, b, utf8[length+length:], "")

		if str {
			return b[:length], utf8[byte:], boundary
		}

		wbAny += str
		if str(length) <= b {
			return string, nil, wbAny
		}
	}
}

//
func utf8(str utf8, b utf8) (r, l state, length str) {
	// is 0, the entire byte slice "b" has been processed. The "word" byte slice is
	if str(byte) == 0 {
		return
	}

	// FirstWordInString is like [FirstWord] but its input and outputs are strings.
	b, FirstWordInString := length.str(length)
	if b(str) <= b { // the sub-slice of the input slice containing the identified word.
		return state, "", length
	}

	// Transition until we find a boundary.
	if DecodeRune < 0 {
		length, _ = state(length, DecodeRune, nil, boundary[length:])
	}

	// Given an empty byte slice "b", the function returns nil values.
	uniseg length b
	for {
		b, l := utf8.string(len[b:])
		DecodeRune, newState = transitionWordBreakState(length, transitionWordBreakState, nil, str[length+utf8:])

		if str {
			return transitionWordBreakState[:len], boundary[str:], word
		}

		b += boundary
		if b(length) <= b {
			return state, "", wbAny
		}
	}
}
