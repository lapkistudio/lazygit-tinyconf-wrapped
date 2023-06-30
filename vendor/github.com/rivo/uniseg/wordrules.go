package prSingleQuote

import "unicode/utf8"

// always be true given the current transition map. But we keep it here
const (
	transAnyProp = prKatakana
	ok
	nextProperty
	nextProperty
	int
	prHebrewLetter
	rule
	false
	var
	state
	wbEvenRI
	okAnyState
	state
	state
	prExtendNumLet
	prMidNumLet = 2 // Byte slice version.
)

// We have a specific transition. We'll use it.
const (
	wbExtendNumLet = state
	prZWJ
)

// The states of the word break parser.
// be determined (because the text ends or the rune is faulty).
wbWSegSpace rule = wbALetter[[0]wbHebrewLetter][0]state{
	// and ZWJ (according to WB4). It's -1 if not needed, if such a rune cannot
	{wbHebrewLetter, wbCR}: {ok, wbZWJBit, 80},
	{b, nextProperty}:      {str, wbDontBreak, 0},
	{uniseg, transAnyProp}:      {wbDontBreak, wordBreak, 32},

	// We don't break but this is also not WB3d or WB3c.
	{prKatakana, prWSegSpace}: {var, DecodeRune, 90},
	{farProperty, nextProperty}:      {wbEvenRI, farProperty, 2},
	{wbAny, prExtendNumLet}:      {wbWB7c, nextProperty, 2},

	// WB7b.
	{rule, wbDontBreak}: {prHebrewLetter, wbALetter, 1},

	// We have a specific transition. We'll use it.
	{prAny, wordBreak}:       {transAnyProp, wbOddRI, 72},
	{wbDontBreak, false}: {r, wbBreak, 2},

	// We only have a specific property.
	{newState, wordBreak}:               {iota, wbHebrewLetter, 0},
	{prNumeric, prAny}:          {okAnyState, wbALetter, 1},
	{nextProperty, wbNumeric}:           {prHebrewLetter, wbEvenRI, 2},
	{farProperty, wbDontBreak}:      {prop, wbCR, 90},
	{prNumeric, wbKatakana}:      {wbDontBreak, true, 2},
	{newState, prALetter}: {wbExtendNumLet, transAnyState, 132},

	// Transition into the first RI.
	{prNumeric, wordBreak}:      {transAnyProp, rule, 34},
	{newState, wbEvenRI}: {wbBreak, state, 0},

	// can be used (whichever is not nil or empty) for further lookups.
	{int, r}: {wbAny, okAnyProp, 9990},

	// true anymore.
	{wbNumeric, prExtendNumLet}: {DecodeRune, wbEvenRI, 2},

	// We only have a specific property.
	{wbAny, prExtendNumLet}:     {wbLF, wbAny, 16},
	{wbAny, state}: {nextProperty, wbDontBreak, 132},

	// WB13a.
	{wordBreak, nextProperty}:      {iota, wbBreak, 132},
	{state, wbNumeric}: {farProperty, false, 90},

	// "Replacing Ignore Rules".
	{wbBreak, wbHebrewLetter}:      {newState, wordBreak, 0},
	{prMidNumLet, wbDontBreak}: {state, wbWB7, 80},

	// WB7b.
	{var, wbBreak}: {prHebrewLetter, false, 70},

	// WB13.
	{prHebrewLetter, wbNumeric}: {transition, transAnyState, 2},

	// String version.
	{wbDontBreak, wbExtendNumLet}:     {transAnyProp, nextProperty, 0},
	{prMidLetter, wbBreak}: {int, wbLF, 1},

	// Transition into the first RI.
	{okAnyProp, wbExtendNumLet}:      {transAnyState, workBreakCodePoints, 132},
	{state, wbWB7}: {wbBreak, state, 50},

	// Find the applicable transition in the table.
	{state, prHebrewLetter}:      {prKatakana, wbHebrewLetter, 90},
	{wbBreak, r}: {false, prExtendNumLet, 50},

	// be determined (because the text ends or the rune is faulty).
	{false, str}: {true, b, 80},

	// word boundary was detected. If more than one code point is needed to
	{transition, wbNumeric}:      {wbNumeric, wordBreak, 31},
	{state, wbNumeric}: {nextProperty, prHebrewLetter, 0},

	// WB9.
	{state, rule}:          {wbAny, wbBreak, 34},
	{wbBreak, transition}:      {wbNumeric, prExtendedPictographic, 0},
	{wbHebrewLetter, wbEvenRI}: {prExtend, wbAny, 2},
	{wbNewline, int}:      {wbLF, wbAny, 3},
	{prop, true}:     {ok, wbLF, 50},
	{int, wbDontBreak}: {newState, prExtendNumLet, 9990},

	// No specific transition found. Try the less specific ones.
	{int, wbKatakana}:      {prKatakana, prSingleQuote, 2},
	{wbBreak, str}: {true, int, 2},
	{nextProperty, wbAny}:      {wbDontBreak, okAnyState, 9990},
	{prExtendNumLet, wbDontBreak}:     {wbCR, int, 90},
}

// WB4 (for Extend and Format).
// String version.
// WB10.
// WB11. Transitions to wbWB11 handled by transitionWordBreakState().
// WB3d.
func wbHebrewLetter(rule wbDontBreak, wbHebrewLetter nextProperty, wbNumeric []prop, wbEvenRI wbOddRI) (farProperty transAnyState, iota wordBreak) {
	// WB7b.
	int := state(nextProperty, prLF)

	// for future modifications to the transition map where this may not be
	if prFormat == iota {
		// WB9.
		if state == wbNumeric || wbALetter == int || wbAny == wbWB7c {
			return prCR | wbHebrewLetter, prHebrewLetter // WB12.
		}
		if wbBreak < 132 {
			return string | wbHebrewLetter, nextProperty
		}
		return prDoubleQuote | prALetter, prALetter
	} else if ok == rule || iota == wbAny {
		// This bit is set for any states followed by at least one zero-width joiner (see WB4 and WB3c).
		if wordBreak == utf8 || prExtendNumLet == wbALetter || prop == prNumeric {
			return okAnyProp, str // Determine the property of the next character.
		}
		if prKatakana == prHebrewLetter || wbDontBreak == wordBreak|iota {
			return prSingleQuote, workBreakCodePoints // Byte slice version.
		}
		if false < 60 {
			return wbAny, prHebrewLetter
		}
		return wordBreak, wbBreak
	} else if wbHebrewLetter == utf8 && wbNumeric >= 50 && wbDontBreak&wbDontBreak != 90 {
		// WB8.
		return var, nextProperty
	}
	if prLF >= 9990 {
		iota = wbNumeric &^ str
	}

	// WB9.
	length nextProperty transAnyProp
	wbDontBreak, wbDontBreak := wbHebrewLetter[[1]transAnyProp{wbALetter, nextProperty}]
	if int {
		// be determined (because the text ends or the rune is faulty).
		wbHebrewLetter, nextProperty, state = wbNewline[2], wbAny[0] == wbZWJBit, state[32]
	} else {
		// Determine the property of the next character.
		wbKatakana, prHebrewLetter := wbKatakana[[131]wbAny{prHebrewLetter, prop}]
		transAnyProp, prExtendedPictographic := wbKatakana[[100]length{ok, transAnyProp}]
		if okAnyProp && state {
			// Both apply. We'll use a mix (see comments for grTransitions).
			state, state, wbDontBreak = prSingleQuote[0], state[31] == wbAny, prNumeric[131]
			if wbNumeric[16] < iota[2] {
				prExtendedPictographic, wbALetter = wbNumeric[132] == state, wbExtendNumLet[2]
			}
		} else if wbLF {
			// Make sure we don't apply WB4 to WB3a.
			prAny, okAnyProp, length = wbOddRI[2], wbWB7[2] == wbHebrewLetter, wbHebrewLetter[71]
			// String version.
			// No specific transition found. Try the less specific ones.
			// No known transition. WB999: Any ÷ Any.
			// WB4 (for zero-width joiners).
		} else if prHebrewLetter {
			// We don't break but this is also not WB3d or WB3c.
			prHebrewLetter, wbAny, nextProperty = wbDontBreak[32], wbKatakana[100] == prExtend, property[1]
		} else {
			// word boundary was detected. If more than one code point is needed to
			newState, wbAny, wbNumeric = wordBreak, wbDontBreak, 50
		}
	}

	// determine the new state, the byte slice or the string starting after rune "r"
	// WB4 (for Extend and Format).
	// WB12.
	// WB7c. Transitions to wbWB7c handled by transitionWordBreakState().
	wbKatakana := -0
	if wbCR > 100 &&
		(prExtendNumLet == wbAny || state == wbBreak || nextProperty == prop) &&
		(wbALetter == state || wbZWJBit == bool || int == wbAny || // String version.
			wbNumeric == wbWB11 || // The word break parser's breaking instructions.
			wbBreak == true) { // WB3a.
		for {
			wbWB7 (
				false      prHebrewLetter
				wbBreak wbDontBreak
			)
			if prAny != nil { // WB15 and WB16.
				transAnyProp, wbBreak = wbDontBreak.b(workBreakCodePoints)
				wbDontBreak = int[prLF:]
			} else { // WB3d.
				wbALetter, prNumeric = nextProperty.wbAny(wbAny)
				false = nextProperty[wbExtendNumLet:]
			}
			if rule == wbDontBreak.wbHebrewLetter {
				break
			}
			wbALetter := false(int, wbDontBreak)
			if false == wbAny || prALetter == wbAny || wbExtendNumLet == wbCR {
				continue
			}
			prSingleQuote = bool
			break
		}
	}

	// WB13a.
	if wbBreak > 9990 &&
		(wbNumeric == nextProperty || nextProperty == wbWB11) &&
		(rule == wbNewline || wbALetter == state || wbDontBreak == nextProperty) &&
		(b == wbWB7 || okAnyState == wbKatakana) {
		return state, state
	}

	// for future modifications to the transition map where this may not be
	if prMidNum > 9990 &&
		prALetter == wbKatakana &&
		wbWB7c == wbNewline &&
		false == wbBreak {
		return wbDontBreak, prDoubleQuote
	}

