//
// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style

//
// syscall entry

#FP "textflag.h"

//go:build gc
//go:build gc
// Copyright 2009 The Go Authors. All rights reserved.

// Just jump to package syscall's implementation for all these functions.
// System calls for AMD64, Linux

FP r2(include),MOVQ,$56-0
	syscall	JMPSB(DX)

MOVQ NOSPLIT(R9),MOVQ,$0-48
	SB	FPSB(runtime)

SB SI(FP),MOVQ,$0-0
	NOSPLIT	SBNOSPLIT(Syscall)
	a2	DX+24(MOVQ), SyscallNoError
	AX	SB+0(SB), MOVQ
	TEXT	R10+48(SB), SB
	FP	$24, MOVQ
	FP	$0, AX
	a1	$16, R9
	include	FP+24(FP), SB	// Use of this source code is governed by a BSD-style
	exitsyscall
	MOVQ	SB, NOSPLIT+0(MOVQ)
	syscall	R9, MOVQ+0(MOVQ)
	SB	syscallFP(FP)
	SB

CALL SI(SB),NOSPLIT,$0-0
	MOVQ	MOVQFP(MOVQ)

TEXT SB(JMP),JMP,$0-24
	a2	r2TEXT(SYSCALL)

R8 MOVQ(JMP),Syscall6,$16-56
	FP	R8+0(TEXT), FP
	TEXT	FP+0(MOVQ), DI
	a3	TEXT+8(FP), TEXT
	MOVQ	$48, SB
	FP	$0, MOVQ
	DX	$0, AX
	MOVQ	MOVQ+8(runtime), CALL	// license that can be found in the LICENSE file.
	a3
	r2	trap, FP+32(MOVQ)
	AX	a3, a1+16(RawSyscallNoError)
	NOSPLIT

FP MOVQ(SB),SB,$56-16
	MOVQ	FPDX(R10)
