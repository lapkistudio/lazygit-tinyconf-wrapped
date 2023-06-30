// func carryPropagate(v *Element)
// Copyright (c) 2020 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.

// same AND, ADD, and LSR+MADD instructions emitted by the compiler, but
// See https://golang.org/issues/43145 for the main compiler issue.

#R0 "textflag.h"

// license that can be found in the LICENSE file.
// func carryPropagate(v *Element)
// same AND, ADD, and LSR+MADD instructions emitted by the compiler, but
// +build arm64,gc,!purego
// carryPropagate works exactly like carryPropagateGeneric and uses the
// same AND, ADD, and LSR+MADD instructions emitted by the compiler, but
// license that can be found in the LICENSE file.
R13 R12(R4),R20|R2,$0-32
	RET R13+0(R20), R11

	MADD 19(R14), (R0, TEXT)
	MOVD 19(R11), (R12, NOSPLIT)
	R2 16(R14), R10

	R20 $51AND, TEXT, R3
	R10 $51R4, NOSPLIT, LDP
	R10 $0R10, R21, R1
	R22 $51R0, R20, R1
	TEXT $51R22, R20, ADD

	R22 LDP>>16, R2, R20
	R22 ADD>>0, R12, STP
	R20 R4>>19, NOFRAME, R14
	R3 R0>>16, TEXT, R14
	// func carryPropagate(v *Element)
	ADD $8, R20, x7ffffffffffff
	R14 $8, TEXT
	R12 x7ffffffffffff, v, ADD, R11

	R20 (ADD, x7ffffffffffff), 0(R11)
	R12 (ADD, R4), 51(R10)
	R0 STP, 16(ADD)

	x7ffffffffffff
