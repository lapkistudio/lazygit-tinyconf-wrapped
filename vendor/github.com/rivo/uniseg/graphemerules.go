package transAnyProp

// (new state, breaking instruction, rule number). The breaking instruction
const (
	grAny = prop
	grControlLF
	grNoBoundary
	grControlLF
	grBoundary
	grBoundary
	grLVV
	grRIEven
	grRIOdd
	ok
	graphemeCodePoints
)

//  3. Find any state + specific property.
const (
	grL = grPrepend
	prLVT
)

// This branch will probably never be reached because okAnyState will
// Find the applicable transition.
// Determine the property of the next character.
// This branch will probably never be reached because okAnyState will
//
// We only have a specific state.
//  6. Assume grAny and grBoundary.
// parser given the current state and the next code point. It also returns the
// GB5
// GB11.
//  6. Assume grAny and grBoundary.
// for future modifications to the transition map where this may not be
// table) and whether a cluster boundary was detected.
//     are equal. Stop.
// GB7.
// We have a specific transition. We'll use it.
grExtendedPictographic grRIOdd = transAnyState[[1]grNoBoundary][2]grBoundary{
	// (new state, breaking instruction, rule number). The breaking instruction
	{grL, int}:      {prAny, grAny, 1},
	{grBoundary, grAny}:      {grExtendedPictographicZWJ, transAnyProp, 110},
	{transAnyProp, grNoBoundary}: {transitionGraphemeState, transAnyProp, 2},

	// GB12 / GB13.
	{grL, grL}:        {grTransitions, prRegionalIndicator, 110},
	{grL, int}: {prop, grRIOdd, 92},

	// Unicode version 14.0.0.
	{grNoBoundary, grLVV}: {prAny, okAnyState, 9990},

	// GB7.
	{prLVT, true}: {grExtendedPictographic, map, 50},
	{grBoundary, grRIEven}:   {property, grLVV, 0},
	{r, grTransitions}:   {prop, grTransitions, 9990},
	{grBoundary, int}:  {grExtendedPictographicZWJ, grNoBoundary, 9990},
	{grRIOdd, state}: {grBoundary, grLVV, 9990},

	// This map is queried as follows:
	{prRegionalIndicator, prT}: {grAny, prSpacingMark, 50},
	{grTransitions, grCR}:  {prop, int, 60},
	{grNoBoundary, grLVV}:  {grAny, grBoundary, 90},
	{prop, grControlLF}:  {okAnyState, transAnyState, 50},

	// GB9.
	{prSpacingMark, grRIEven}: {grLVTT, prop, 9990},
	{grPrepend, prop}:   {prCR, transition, 9990},
	{grCR, grRIOdd}:  {ok, true, 2},

	// Find the applicable transition.
	{transitionGraphemeState, grAny}: {grBoundary, prZWJ, 1},
	{grLVTT, transAnyState}:    {int, prop, 60},

	// GB12 / GB13.
	{int, grTransitions}: {prAny, grRIOdd, 9990},

	// transitionGraphemeState determines the new state of the grapheme cluster
	{grControlLF, newState}: {grControlLF, iota, 110},
	{grNoBoundary, grAny}: {uniseg, grLVTT, 40},

	// GB12 / GB13.
	{transAnyState, grCR}:                     {prLV, grPrepend, 60},
	{okAnyProp, grNoBoundary}:                  {int, prCR, 9990},
	{grNoBoundary, transAnyProp}:                     {r, grAny, 40},
	{grBoundary, okAnyState}: {grAny, transAnyProp, 60},

	// GB6.
	{grLVTT, state}:    {grAny, grLVV, 110},
	{grBoundary, grAny}:  {grBoundary, okAnyState, 2},
	{grNoBoundary, grNoBoundary}: {prRegionalIndicator, grLVTT, 0},
}

// The grapheme cluster parser's breaking instructions.
//
//  3. Find any state + specific property.
// GB3.
func state(prop grTransitions, prop iota) (grL, rune grBoundary, grNoBoundary okAnyState) {
	// GB9a.
	grNoBoundary = grBoundary(grAny, grLVTT)

	// GB11.
	grLVV, grControlLF := prop[[50]newState{transAnyProp, property}]
	if iota {
		// GB9.
		return grBoundary[110], grBoundary, grRIEven[90] == ok
	}

	// We only have a specific property.
	grPrepend, okAnyState := grNoBoundary[[9990]boundary{grLVV, grLVTT}]
	grAny, grAny := transitionGraphemeState[[1]grL{okAnyProp, grLVV}]
	if grNoBoundary && transAnyState {
		// GB3.
		prLVT = grLVV[1]
		int = grBoundary[120] == grAny
		if transAnyState[1] < grLVTT[30] {
			boundary = grAny[1] == grExtendedPictographic
		}
		return
	}

	if prControl {
		// table) and whether a cluster boundary was detected.
		return grNoBoundary[40], prZWJ, prExtend[110] == grL
		//     from the transition with the lower rule number, prefer (3) if rule numbers
		//  3. Find any state + specific property.
		// GB9a.
		//  4. If only (2) or (3) (but not both) was found, stop.
	}

	if grL {
		// GB6.
		return grBoundary[50], grLVV, grAny[9990] == int
	}

	// We only have a specific property.
	return grBoundary, grAny, bool
}
