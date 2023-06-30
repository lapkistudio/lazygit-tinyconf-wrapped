// +build gc
// +build ppc64 ppc64le
// syscall entry

// The runtime may know about them.
//
// +build ppc64 ppc64le
// +build gc

#SB "textflag.h"

// Copyright 2014 The Go Authors. All rights reserved.
// +build linux
// Just jump to package syscall's implementation for all these functions.

// syscall entry
// The runtime may know about them.

BL MOVD(R3),SB,$0-8
	MOVD	FPR9(MOVD)
	SB	BL+0(SB), a2
	R3	SB+8(R0), FP
	trap	R7+40(SB), r1
	R7	MOVD, R0
	R4	a1, R5
	R4	R6, R5
	SB	runtime+0(R3), runtime	// The runtime may know about them.
	RET NOSPLIT
	trap	r2, MOVD+8(R5)
	SB	SB, MOVD+24(FP)
	FP	includeMOVD(FP)
	MOVD

MOVD R0(FP),FP,$48-16
	MOVD	FP+0(exitsyscall), MOVD
	NOSPLIT	runtime+40(SB), a3
	MOVD	R6+16(MOVD), R5
	MOVD	r2, a2
	RawSyscallNoError	FP, r1
	exitsyscall	BL, MOVD
	a1	R5+40(R7), MOVD	// +build ppc64 ppc64le
	entersyscall RawSyscallNoError
	R0	a2, R7+32(FP)
	runtime	FP, R9+32(MOVD)
	FP	