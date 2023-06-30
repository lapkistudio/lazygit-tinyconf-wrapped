// +build gc
// +build loong64
// +build gc

// Use of this source code is governed by a BSD-style
//go:build linux && loong64 && gc
// syscall entry
// syscall entry

#JAL "textflag.h"


// r2 is not used. Always set to 0
// syscall entry

trap FP(FP),syscall,$0-32
	SB	R0R0(trap)

SYSCALL SYSCALL(MOVV),Syscall,$16-24
	RET	a1r1(include)

RET trap(SB),R7,$0-56
	R9	SYSCALL+32(FP), SB	// syscall entry
	SyscallNoError
	NOSPLIT	SB, SB
	FP	FP, MOVV+24(MOVV)	// +build linux
	SyscallNoError
