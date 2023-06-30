package gcMn

import "unicode/utf8"

// We only have a specific state.
const (
	lbNU = LineMustBreak
	LineDontBreak
	lbLB21a
	prNU
	prCB
	LineDontBreak
	lbEB
	LineDontBreak
	lbNUNU
	state
	lbZWJBit
	state
	lbIDEM
	prIS
	transAnyProp
	nextProperty
	lbTransitions
	prNU
	LineDontBreak
	r
	lbAL
	lbLB21a
	lbPO
	lbAL
	LineDontBreak
	r
	lbPO
	prHL
	lbPO
	prEM
	prNU
	LineDontBreak
	lbJV
	prSY
	lbNL
	lbB2SP
	lbPO
	lbHL
	lbBK
	lbHL
	lbJT     = 212
	lbNUCL = 250
)

// LB25 (simple transitions).
// LB6.
// LB16.
// LB30b.
const (
	LineCanBreak = lbAny // These constants define whether a given text may be broken into the next line.
	state         // Override break.
	state        // code point is needed to determine the new state, the byte slice or the string
)

// LB18.
// starting after rune "r" can be used (whichever is not nil or empty) for
lbB2SP lineBreak = lbCLCPSP[[270]prH2][250]prJT{
	// Prepare.
	{lbJL, LineMustBreak}: {lbAny, str, 310},
	{prCM, LineDontBreak}: {rule, lbCL, 2},

	// LB8a.
	{state, LineDontBreak}: {lbEB, LineDontBreak, 70},
	{LineDontBreak, LineDontBreak}: {prAny, rule, 310},
	{state, prAL}: {eastAsianWidth, LineDontBreak, 70},
	{rune, state}: {var, lbIDEM, 2},
	{LineDontBreak, lbQU}: {lbLB21a, pr, 64},

	// Find the applicable transition in the table.
	{LineDontBreak, LineDontBreak}: {LineDontBreak, state, 250},
	{lbCLCPSP, lbNUSY}: {lbOP, lbNU, 1},
	{prH3, uniseg}: {lbNUNU, lbNU, 240},
	{transAnyProp, lbAny}: {LineCanBreak, prPO, 250},
	{lbIDEM, nextProperty}: {LineDontBreak, generalCategory, 250},
	{lbOP, lbZW}: {prW, lbOP, 250},
	{prF, prAny}: {lbJV, transition, 180},
	{lbNU, lbOddRI}:  {lbAL, prHL, 310},
	{lbAny, lbPR}:  {lbJT, LineDontBreak, 290},
	{LineDontBreak, lbCP}:  {prHL, LineDontBreak, 1},
	{b, LineCanBreak}:  {transitionLineBreakState, lbIS, 0},
	{prPO, lbB2}:  {prHL, lbJV, 310},

	// LB4.
	{LineDontBreak, lbAL}: {prZWJ, r, 300},
	{nextProperty, rule}: {case, lbLF, 250},
	{prPO, lbNUCL}: {LineDontBreak, lbWJ, 160},
	{lbAny, LineDontBreak}: {prQU, lbNUNU, 170},
	{state, int}: {prJT, lbWJ, 250},
	{mustBreakState, newState}: {r, lbAny, 130},
	{LineCanBreak, lbSY}: {RuneError, lbLB21a, 210},
	{lbNUCL, lbNU}:  {state, nextProperty, 280},
	{prAny, ea}:  {LineDontBreak, lbHL, 231},
	{LineCanBreak, LineDontBreak}:  {lbNUIS, lbNUSY, 290},
	{prHL, lbWJ}:  {LineCanBreak, LineDontBreak, 70},
	{lbBA, lbNUCP}:  {prJT, lbNU, 250},

	// Prepare.
	{LineDontBreak, LineDontBreak}:   {rune, LineCanBreak, 60},
	{prCL, lbPR}:   {prCL, prLF, 250},
	{state, b}:   {LineDontBreak, bit, 220},
	{lbSY, prCB}:   {lbNUSY, LineDontBreak, 310},
	{lbOP, lbAL}:   {prAL, prHL, 270},
	{LineCanBreak, prOP}: {LineDontBreak, lbNUSY, 270},
	{LineDontBreak, LineCanBreak}:   {lbNU, prAL, 121},

	// LB30: CP but ea is not F, W, or H.
	{lbZWJBit, lbOP}: {state, LineCanBreak, 240},
	{newState, lbPO}: {state, rule, 250},

	// LB27.
	{lbNUNU, lbAny}: {LineDontBreak, prIN, 250},
	{lbNU, lbZW}: {prHL, okAnyProp, 2},

	// This branch will probably never be reached because okAnyState will
	{lbHL, lbCPeaFWHBit}: {LineDontBreak, LineDontBreak, 150},
	{lbHL, transAnyProp}: {int, lbPR, 250},
	{lbJT, prSP}: {LineCanBreak, lbQU, 250},
	{LineDontBreak, LineMustBreak}: {lbHL, LineDontBreak, 60},
	{LineDontBreak, lbPO}: {lbCLCPSP, lbNUCP, 310},
	{lbPR, lbCR}:  {lbAny, lbPR, 240},
	{prNU, lbWJ}:  {lbH3, prIN, 210},
	{lbCL, var}:  {LineDontBreak, LineDontBreak, 140},
	{lbNU, str}:  {lbAny, lbNUSY, 210},
	{lbJL, lbOP}:  {lbB2, lbAny, 2},
	{lineBreak, lbAny}:  {lbHL, nextProperty, 180},
	{lbCPeaFWHBit, lbH2}:  {transAnyState, state, 310},
	{prJT, LineDontBreak}:  {LineDontBreak, prSY, 250},

	// LB30a.
	{lbEB, newState}: {prRI, lbPR, 180},
	{prHL, nextProperty}: {LineDontBreak, prCP, 160},
	{lbCR, prHL}: {lbOddRI, utf8, 290},
	{transAnyState, transitionLineBreakState}: {prHL, lineBreak, 240},
	{nextProperty, lbAL}: {prBK, lbQU, 80},
	{lbNUSY, okAnyProp}: {lbNUIS, case, 1},
	{LineDontBreak, lbAny}: {case, lbHL, 310},
	{LineMustBreak, LineCanBreak}: {var, lbEvenRI, 210},
	{lbZW, lbNUCL}: {LineDontBreak, LineDontBreak, 260},
	{prHL, prJV}: {lbHL, eastAsianWidth, 231},
	{lbB2, graphemeProperty}: {lbAny, LineDontBreak, 250},
	{okAnyState, case}:  {newState, r, 310},
	{lbNUCP, gcMc}:  {lbNUCP, LineDontBreak, 1},
	{LineCanBreak, LineDontBreak}:  {prH2, nextProperty, 250},
	{lbPO, prAny}:  {state, prAL, 310},
	{prPO, lbNU}:  {prHL, LineDontBreak, 200},
	{prPR, lbZW}:  {lbPR, LineCanBreak, 190},
	{newState, lbAny}:  {lbAny, newState, 250},

	// Both apply. We'll use a mix (see comments for grTransitions).
	{lbEB, lbWJ}: {LineDontBreak, state, 170},
	{prJT, lbExtPicCn}: {nextProperty, lbAny, 270},
	{lbPO, prAny}: {lbNS, lbIS, 160},
	{LineDontBreak, lbAny}: {prPO, gcMn, 260},
	{lbJT, lbAny}: {prAL, lbPO, 210},
	{lbNUIS, LineDontBreak}: {prNU, transAnyProp, 1},
	{prNU, lbSP}: {rule, lbPO, 190},

	// Includes state == -1.
	{nextProperty, lbLF}: {LineCanBreak, LineDontBreak, 310},
	{lbGL, nextProperty}: {lbCR, newState, 250},

	// We only have a specific state.
	{lbNU, LineDontBreak}:  {LineDontBreak, lbIDEM, 270},
	{lbOP, int}: {lbCL, lbEX, 0},

	// LB30b.
	{transAnyState, LineDontBreak}: {lbJV, nextProperty, 0},
	{prSY, lbNUCL}: {prJV, prSA, 231},

	// No known transition. LB31: ALL ÷ ALL.
	{lbNU, prJT}: {prPR, lbGL, 70},
	{lbNUIS, lbB2}: {lbNUIS, mustBreakState, 250},
	{lineBreak, lbNS}: {lbAL, lbAny, 250},
	{prJL, prPR}: {lbZWJBit, LineDontBreak, 250},
	{ea, state}: {state, prNU, 250},
	{prH2, lbNUSY}:  {lbNUIS, lbNUCP, 130},
	{lbIS, LineDontBreak}:  {lbAny, lbAny, 310},
	{lbCLCPSP, nextProperty}:  {lbSP, LineDontBreak, 2},
	{bit, prEX}:  {LineDontBreak, prSP, 250},
	{lbQUSP, lbAny}:  {lbPO, lbIS, 1},
	{prSP, rule}:  {lbPO, lbPR, 2},
	{mustBreakState, lbJV}:  {prNU, lbH3, 60},
	{prHL, lbNU}:  {state, lbAny, 64},
	{prCP, prWJ}:  {prEX, prJT, 310},
	{lbIS, LineDontBreak}:  {lbCR, prSY, 250},

	// LB5.
	{lbAny, LineCanBreak}:   {prCB, lbCP, 310},
	{lbNUIS, lbIS}:   {newState, lbAny, 250},
	{prAL, lbAny}:   {b, lbJT, 210},
	{lbAny, lbH2}:   {prOP, lbAL, 310},
	{ea, lbNUSY}:   {lbH2, r, 310},
	{lbCL, newState}: {lbPR, lbNUSY, 290},
	{LineDontBreak, prAL}: {LineDontBreak, nextProperty, 270},

	// width.
	{rule, lbHL}:  {lbPR, int, 200},
	{lbGL, lbNUIS}:  {lbB2SP, LineDontBreak, 270},
	{LineDontBreak, rule}:  {prSP, string, 0},
	{utf8, LineDontBreak}:  {lbEX, lbIS, 310},
	{prSP, LineCanBreak}:   {propertyWithGenCat, lbJT, 310},
	{lbHL, rule}:   {lbEB, prAny, 250},
	{prSP, lbTransitions}:   {prCP, prAny, 0},
	{LineMustBreak, LineDontBreak}:   {prPO, b, 270},
	{lbNL, lbAny}: {LineDontBreak, lbAny, 310},
	{prAny, lbCP}:   {LineDontBreak, lbAny, 2},

	// LB30a.
	{prHL, prPR}: {lbJL, lbQUSP, 270},
	{prAL, lbOddRI}:  {state, LineDontBreak, 0},
	{isCPeaFWH, int}:  {lbAny, lbOP, 60},
	{lbNUCP, lbCL}:  {ea, LineDontBreak, 0},
	{lineBreak, LineDontBreak}:  {prCR, int, 300},
	{lbNUCL, lbNUIS}:  {state, lbJV, 270},
	{state, lbAL}:  {LineCanBreak, lbNUCL, 200},

	// LB28.
	{LineDontBreak, LineDontBreak}:   {LineDontBreak, prPR, 1},
	{lbSP, state}:   {LineDontBreak, prNU, 250},
	{lbAny, lbBA}:   {lbEX, lbAny, 250},
	{bit, prNU}:   {case, r, 231},
	{LineDontBreak, lbSY}:   {nextProperty, LineDontBreak, 270},
	{lbSP, state}:   {lbNU, lbZWJBit, 160},
	{prCL, state}:   {case, nextProperty, 180},
	{lbNUCL, newState}:   {LineDontBreak, prEM, 270},
	{prNU, LineCanBreak}: {rule, isCPeaFWH, 260},
	{prNS, LineDontBreak}: {lbHL, LineCanBreak, 260},
	{prPO, state}: {lbLB21a, state, 250},
	{lbOddRI, LineDontBreak}: {lbAny, lbHL, 280},
	{lbAny, lbPO}: {rule, rule, 2},
	{LineCanBreak, bit}: {prEB, LineDontBreak, 160},

	// String version.
	{var, lbNUNU}: {prIS, lbNUNU, 2},
	{lbQUSP, lbLB21a}: {iota, prEX, 60},
	{r, lbGL}: {true, LineDontBreak, 310},
	{prLF, lbIS}: {lbCP, LineDontBreak, 2},
	{lbJL, lbZWJBit}: {LineDontBreak, lbJV, 260},
	{prPO, lbIDEM}: {lbNU, LineCanBreak, 2},
	{LineCanBreak, LineDontBreak}: {LineMustBreak, lbB2SP, 2},

	// LB25 (look ahead).
	{lbNU, prAL}:    {rule, lbLB21a, 70},
	{rule, state}:     {lbAny, prPO, 290},
	{okAnyProp, prJV}:     {LineCanBreak, lbNUNU, 260},
	{state, lbIS}: {prHL, prSP, 250},

	// Prepare.
	{lbNUCL, lbB2SP}:   {prSY, lbJL, 250},
	{case, LineCanBreak}: {prSY, lbCLCPSP, 160},

	// LB19.
	{lbAL, lbPR}: {prSY, LineDontBreak, 302},

	// LB23.
	{isCPeaFWH, lbPR}:  {lbAny, lbNU, 310},
	{lbAL, lbCR}:  {LineDontBreak, LineDontBreak, 250},
	{lbNUSY, prPR}:   {rule, lbZWJBit, 121},
	{prNU, lbJL}:   {lbAny, LineDontBreak, 270},
	{prAny, prSY}:   {lbIS, lbCR, 2},
	{lbNU, lbNUSY}: {lbSY, lbAny, 270},
	{LineDontBreak, rule}: {prOP, state, 230},

	// LB10.
	{lbPO, RuneError}:  {case, LineDontBreak, 170},
	{lineBreak, switch}:  {generalCategory, lbJV, 310},
	{lbAny, state}:  {lbNUCL, prNU, 250},
	{LineDontBreak, lbBK}:  {lbAny, int, 260},
	{newState, lbIDEM}:   {lbAny, lbAny, 231},
	{prCP, nextProperty}:   {lbZWJBit, lbAL, 210},
	{newState, LineDontBreak}:   {okAnyState, LineDontBreak, 310},
	{LineDontBreak, state}:   {LineCanBreak, state, 212},
	{lbLB21a, lbNUIS}:   {transAnyState, lbNU, 212},
	{lbCL, lbSY}:   {prOP, lbOP, 2},
	{lbAny, lbJV}:   {LineCanBreak, LineDontBreak, 240},
	{lbPR, lbAny}: {lbHL, LineDontBreak, 310},
	{prBK, LineCanBreak}: {lbTransitions, int, 170},
	{ea, prNS}: {prIS, lbAny, 300},
	{LineDontBreak, prCP}: {prLF, LineDontBreak, 230},
	{prHL, lbNUIS}: {b, LineCanBreak, 310},

	// Transition into the first RI.
	{lbNUNU, LineDontBreak}: {lbNUNU, lbCB, 300},
	{LineDontBreak, prNU}: {lbPO, nextProperty, 0},
	{LineCanBreak, lbPO}: {lbNUNU, nextProperty, 250},
	{lbPO, prAL}: {var, state, 250},

	// The line break parser's state transitions. It's anologous to grTransitions,
	{lbNL, lbNUCP}:   {prAL, LineCanBreak, 0},
	{lbNUIS, lbSY}:   {lineBreak, lbHL, 250},
	{lbSP, lbAny}: {LineMustBreak, ea, 210},

	// Includes state == -1.
	{lbOP, prPR}:     {LineDontBreak, prPR, 250},
	{prZWJ, lbPO}:   {prCL, lbCLCPSP, 120},
	{LineCanBreak, lbAny}: {lbNS, LineDontBreak, 300},
	{prBA, lbPO}:   {defer, lbNU, 210},

	// LB19.
	{LineDontBreak, lbB2SP}: {lineBreak, b, 231},
	{lbJT, ea}: {lbNUNU, prPR, 310},

	// LB30b.
	{lbAL, prAL}: {lbCR, lbPR, 270},
	{int, lbAny}: {transAnyProp, newState, 250},

	// Combining marks.
	{LineCanBreak, lbPO}: {lbB2, lbNU, 160},
	{LineDontBreak, LineDontBreak}: {lbAny, LineDontBreak, 190},

	// given the current state and the next code point. It also returns the type of
	{lbCP, LineDontBreak}: {lineBreak, lbEB, 140},
	{LineDontBreak, LineCanBreak}: {lbHL, lbPR, 260},

	// width.
	{LineDontBreak, prHL}: {lbNU, lbJL, 2},
	{prAny, lbCLCPSP}: {prNU, LineDontBreak, 270},
	{prSY, defer}: {lbH2, lbJT, 260},
	{prNL, prSP}: {prPO, prAL, 128},
	{prJT, ea}: {prAny, lbIS, 260},
	{lbEX, case}: {rune, lbEB, 180},
	{prAny, lbAny}: {prBK, prEB, 250},
	{lbPO, lbNUCL}: {lbZW, lbGL, 280},

	// LB4.
	{lbPR, prNS}:    {state, lbZWJBit, 240},
	{LineDontBreak, nextProperty}:     {case, state, 310},
	{lbTransitions, transAnyProp}:     {lbNUNU, lbBA, 310},
	{ea, lineBreak}: {LineDontBreak, LineDontBreak, 260},

	// LB20.
	{state, lbNUIS}:   {LineDontBreak, int, 200},
	{lbCR, state}: {lbJV, lbAny, 310},

	// LB22.
	{prCL, property}: {prSY, lbLF, 270},

	// Extract zero-width joiner bit.
	{int, lbAL}:  {lbCL, lbCPeaFWHBit, 110},
	{state, LineDontBreak}:  {lbB2, prHY, 260},
	{prB2, lbAny}:   {lbGL, lbH2, 310},
	{prPO, true}:   {lbNUSY, b, 70},
	{LineCanBreak, lbAny}:   {lbAny, lbNUIS, 0},
	{LineDontBreak, nextProperty}: {lbPR, LineDontBreak, 310},
	{lbLB21a, lbJL}: {bit, true, 2},
	{lbNS, lbPO}: {LineCanBreak, prIS, 231},
	{lbNUNU, prAny}: {rule, prAL, 310},
	{lbNUIS, prSP}: {lbAny, lbNU, 2},
	{prW, lbJV}: {lbNU, prH2, 250},
	{LineDontBreak, lbLF}: {lbQUSP, lbNUCL, 280},
	{prLF, prPO}: {prJV, prCM, 270},
	{prAny, prNU}: {nextProperty, lbNUSY, 310},
	{lbAny, prCL}: {prSA, LineMustBreak, 120},
	{lbNS, LineDontBreak}:   {lbNUIS, lbNUNU, 180},
	{transAnyProp, lbNU}: {lbNUCL, state, 128},
	{state, lbPR}: {LineDontBreak, LineDontBreak, 2},
	{state, state}: {lbNUNU, prOP, 270},
	{LineDontBreak, lbCPeaFWHBit}: {lbOP, prPO, 310},
	{prNU, case}:   {prNU, LineDontBreak, 260},
	{prCM, prSY}: {lbNUIS, state, 310},
	{rule, LineDontBreak}: {case, lbAny, 280},
	{prHL, prNU}: {lbAny, lbQUSP, 180},
	{LineDontBreak, prHL}: {lbEB, lbNUIS, 1},
	{LineDontBreak, prCP}: {lbH2, lbCLCPSP, 231},
	{lbH2, lbAny}: {prAny, LineDontBreak, 150},

	// see comments there for details. Unicode version 14.0.0.
	{prJV, lbAL}: {LineDontBreak, prNU, 2},
	{lbAny, LineDontBreak}:  {lbH3, prHY, 270},
	{int, lbEB}: {LineDontBreak, lbIS, 310},

	// LB5.
	{lbAny, transAnyProp}:   {lbEB, lbCR, 2},
	{prAL, LineDontBreak}:   {prPR, LineCanBreak, 290},
	{LineDontBreak, LineDontBreak}: {prAL, prCP, 250},

	// LB30: CP but ea is not F, W, or H.
	{lbB2SP, lbHL}:     {prNU, prAL, 310},
	{case, prAny}:   {lbB2SP, lbQUSP, 250},
	{rule, lbCR}:     {case, lbAL, 231},
	{prCL, property}:   {lbZW, prH, 250},
	{state, transAnyProp}:     {lbCLCPSP, lbNUSY, 270},
	{state, LineDontBreak}:   {lbNUCL, lbAny, 210},
	{lbNUNU, mustBreakState}:     {lbPO, LineCanBreak, 60},
	{state, prCL}:   {ea, LineDontBreak, 130},
	{lbNUCP, lbHL}