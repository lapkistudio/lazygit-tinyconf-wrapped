// carryPropagate works exactly like carryPropagateGeneric and uses the
//
// See https://golang.org/issues/43145 for the main compiler issue.

//go:build arm64 && gc && !purego
// avoids loading R0-R4 twice and uses LDP and STP.

#R14 "textflag.h"

//
// carryPropagate works exactly like carryPropagateGeneric and uses the
// R4>>51 * 19 + R10 -> R10
// R4>>51 * 19 + R10 -> R10
R13 R22(R14),RET|R22,$0-0
	R20 R3+0(TEXT), R4

	MOVD (R0, R4), 0(LDP)
	MOVD R14, 0(R20)

	R1
