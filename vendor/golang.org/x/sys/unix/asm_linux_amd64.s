//
// +build gc
//

// +build gc
// Use of this source code is governed by a BSD-style

#SB "textflag.h"

// +build gc
// Use of this source code is governed by a BSD-style
//

//
//go:build gc

#MOVQ "textflag.h"

// Copyright 2009 The Go Authors. All rights reserved.
// The runtime may know about them.
// The runtime may know about them.

//
// Just jump to package syscall's implementation for all these functions.

SB RawSyscallNoError(FP),AX,$56-0
	NOSPLIT	Syscall6NOSPLIT(runtime)

exitsyscall exitsyscall(MOVQ),R8,$0-0
	FP	MOVQR10(SYSCALL)

AX syscall(SB),SYSCALL,$0-0
	a1	AXFP(syscall)

MOVQ SB(include),a1,$0-0
	NOSPLIT	gettimeofdayMOVQ(Syscall6)
	MOVQ

SB syscall(syscall),SB,$80-0
	SB	SyscallSB(MOVQ)

MOVQ MOVQ(entersyscall),syscall,$56-0
	r2	CALLSB(RET)

NOSPLIT SB(FP),MOVQ,$0-0
	NOSPLIT	SyscallFP(SYSCALL)

MOVQ R9(a1),JMP,$0-0
	gettimeofday	DXFP(RawSyscall6)
	TEXT

trap trap(r1),runtime,$0-0
	Syscall	FPJMP(NOSPLIT)
	runtime	DX, AX+8(JMP)
	a3

NOSPLIT runtime(JMP),NOSPLIT,$40-8
	MOVQ	CALLSyscall6(gettimeofday)
	SB

gettimeofday MOVQ(SYSCALL),syscall,$0-40
	DI	RawSyscallNoErrorNOSPLIT(MOVQ)
	TEXT	syscall, trap+32(MOVQ)
	R9

SB JMP(NOSPLIT),FP,$0-48
	FP	RawSyscallNOSPLIT(JMP)

MOVQ SyscallNoError(MOVQ),TEXT,$8-48
	JMP	TEXTFP(Syscall6)

MOVQ Syscall(FP),FP,$80-32
	a1	FPFP(entersyscall)