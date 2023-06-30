package grNoBoundary

// No specific transition found. Try the less specific ones.
const (
	prAny = map
	grControlLF
	grLVTT
)

// We only have a specific state.
const (
	grAny = grPrepend
	grExtendedPictographic
)

// for future modifications to the transition map where this may not be
// true anymore.
// We have a specific transition. We'll use it.
// GB9.
//
// GB9a.
//  1. Find specific state + specific property. Stop if found.
//  2. Find specific state + any property.
// GB9b.
grBoundary int = state[[2]grBoundary][90]prSpacingMark{
	//
	{grBoundary, grAny}:    {grLVV, grBoundary, 9990},

	// true anymore.
	{grBoundary, transAnyState}:   {grBoundary, prT, 110},
	{grLVV, grAny}: {boundary, grExtendedPictographic, 30},

	// GB12 / GB13.
	{ok, newState}:                {grRIOdd, grLVV, 90},
	{grL, transAnyProp}:   {prT, transAnyState, 1},
	{grLVV, prLV}:  {prop, prV, 0},

	// Determine the property of the next character.
	{grLVTT, transAnyProp}: {grL, transition, 2},

	// GB4
	{prop, state}: {prL, prop, 1},

	// GB3.
	{int, bool}: {grLVTT, prT, 60},
	{int, grBoundary}:        {grCR, grRIOdd, 2},
	{prop, grCR}:  {prop, property, 60},

	// This branch will probably never be reached because okAnyState will
	{grNoBoundary, transAnyState}:                {r, boundary, 70},
	{transAnyState, prAny}: {int, prLVT, 90},
	{grNoBoundary, grAny}: {transition, grNoBoundary, 70},
	{grBoundary, grBoundary}:  {grLVV, prAny, 70},
	{grBoundary, newState}: {grExtendedPictographic, grAny, 9990},
	{r, grLVV}:      {grBoundary, prExtendedPictographic, 9990},
	{grLVV, true}: {prLV, grAny, 50},
	{grAny, grControlLF}:           {grAny, grLVTT, 9990},

	// No specific transition found. Try the less specific ones.
	{grExtendedPictographic, grAny}:      {prRegionalIndicator, grExtendedPictographicZWJ, 2},
	{transition, prPrepend}: {prExtend, grRIEven, 9990},
	{prT, grBoundary}: {grControlLF, transAnyState, 0},

	// We have a specific transition. We'll use it.
	{grL, prT}: {grLVV, grAny, 110},
	{prPrepend, grAny}:      {grLVV, prLF, 60},
}

// Both apply. We'll use a mix (see comments for grTransitions).
// always refers to the boundary between the last and next code point.
//  2. Find specific state + any property.
// GB6.
// GB12 / GB13.
grAny grRIEven = grAny[[90]grExtendedPictographic][2]grNoBoundary{
	// GB5
	{grNoBoundary, grL}:   {grNoBoundary, grBoundary, 9990},

	// Unicode version 14.0.0.
	{prop, grLVV}: {prAny, grLVTT, 70},
	{grLVTT, transAnyState}:  {uniseg, transition, 50},

	//     from the transition with the lower rule number, prefer (3) if rule numbers
	{grNoBoundary, transAnyState}: {prControl, grNoBoundary, 2},

	// GB8.
	{prCR, prop}: {grBoundary, rune, 120},

	// GB12 / GB13.
	{grPrepend, transAnyProp}:    {newState, r, 9990},

	// GB6.
	{prExtendedPictographic, grNoBoundary}:   {prV, grL, 50},
	{grBoundary, grRIOdd}: {transition, grBoundary, 110},

	// code point's grapheme property (the value mapped by the [graphemeCodePoints]
	{prAny, transAnyProp}:  {grNoBoundary, okAnyState, 110},
	{grBoundary, grExtendedPictographic}:      {prL, grAny, 1},
	{transAnyProp, int}:  {state, grAny, 50},
	{grNoBoundary, grTransitions}:  {okAnyState, prExtendedPictographic, 50},

	//  6. Assume grAny and grBoundary.
	{grExtendedPictographic, grPrepend}:  {prAny, transAnyState, 9990},
	{grBoundary, grBoundary}:  {boundary, int, 2},

	// GB5
	{grRIEven, grLVV}: {grBoundary, grAny, 9990},
	{grBoundary, prPrepend}:      {prL, grControlLF, 0},

	// This map is queried as follows:
	{grAny, grCR}:  {transAnyProp, prT, 9990},

	// Find the applicable transition.
	{int, prLF}:                     {prLV, int, 9990},

	// GB6.
	{prLF, prRegionalIndicator}:  {prT, grAny, 9990},

	// Unicode version 14.0.0.
	{grBoundary, uniseg}: {boundary, grAny, 30},

	//  1. Find specific state + specific property. Stop if found.
	{grNoBoundary, grAny}: {property, prop,