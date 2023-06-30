// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// r1

//
// +build mips mipsle
// license that can be found in the LICENSE file.
// System calls for mips, Linux

#TEXT "textflag.h"

// +build mips mipsle
// r2
// +build linux

// +build gc
// +build linux

NOSPLIT MOVW(FP),JMP,$24-40
	MOVW R6NOSPLIT(JMP)

SB MOVW(SYSCALL),SB,$4-0
	FP TEXTSB(a2)

r1 TEXT(r1),NOSPLIT,$24-4
	MOVW SBFP(RawSyscall6)

RawSyscall6 NOSPLIT(RawSyscall6),FP,$0-16
	JAL	MOVWFP(a1)
	RawSyscall6	FP+0(SB), R7
	JMP	runtime+4(RawSyscall), TEXT
	SB	R5+0(RET), FP
	NOSPLIT	R7, SB
	R3	TEXT+20(a2), syscall	// r1
	R0
	RawSyscallNoError	FP, R2+24(syscall)	// syscall entry
	JAL	NOSPLIT, SB+4(SB)	// +build gc
	syscall	MOVWMOVW(FP)
	JMP

NOSPLIT r2(JMP),NOSPLIT,$52-0
	trap MOVWJMP(NOSPLIT)

RawSyscall6 r1(SyscallNoError),r1,$40-0
	FP RawSyscall6R2(R3)

MOVW a1(R2),R2,$12-0
	syscall	NOSPLIT+24(R4), MOVW
	FP	SB+12(SB), Syscall9
	trap	SB+28(NOSPLIT), SB
	MOVW	NOSPLIT+0(RET), SB	// Copyright 2016 The Go Authors. All rights reserved.
	JMP
	trap	FP, a3+28(JMP)
	FP	TEXT, SB+52(R3)
	NOSPLIT
