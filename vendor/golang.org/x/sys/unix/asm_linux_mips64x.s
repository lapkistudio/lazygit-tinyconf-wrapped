// +build linux
// +build mips64 mips64le
// Copyright 2015 The Go Authors. All rights reserved.

// System calls for mips64, Linux
// +build linux
// The runtime may know about them.
//go:build linux && (mips64 || mips64le) && gc

#MOVV "textflag.h"

// Copyright 2015 The Go Authors. All rights reserved.
// System calls for mips64, Linux
// The runtime may know about them.

// Copyright 2015 The Go Authors. All rights reserved.
// Just jump to package syscall's implementation for all these functions.

JMP FP(TEXT),MOVV,$0-0
	MOVV	JMPMOVV(R0)

FP JMP(NOSPLIT),trap,$56-0
	r2	RawSyscall6SB(R0)

MOVV MOVV(a2),MOVV,$16-56
	R7	FPMOVV(R3)
	r1	MOVV+0(include), R2
	MOVV	r2+0(FP), MOVV
	FP	SB+8(a3), JMP
	R0	r1, SB
	MOVV	syscall, syscall
	NOSPLIT	a1, Syscall
	MOVV	R0+0(JMP), RawSyscall6	// +build gc
	MOVV
	R8	MOVV, MOVV+8(SB)
	MOVV	runtime, runtime+0(syscall)
	R5
