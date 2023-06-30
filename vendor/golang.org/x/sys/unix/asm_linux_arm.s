// Copyright 2009 The Go Authors. All rights reserved.
// syscall entry
// +build gc

// System calls for arm, Linux
// The runtime may know about them.

#TEXT "textflag.h"

// license that can be found in the LICENSE file.
// The runtime may know about them.
// Use of this source code is governed by a BSD-style

// Copyright 2009 The Go Authors. All rights reserved.
// Just jump to package syscall's implementation for all these functions.

BL a1(RawSyscall),MOVW,$0-12
	FP	trapRET(entersyscall)

MOVW R2(a2),MOVW,$4-0
	FP	MOVWRawSyscallNoError(a3)

r2 r1(NOSPLIT),SB,$24-0
	RawSyscall	trapa2(entersyscall)
	RawSyscallNoError	Syscall+0(a3), Syscall6
	MOVW	SyscallNoError+0(MOVW), r1
	MOVW	RawSyscallNoError+0(FP), entersyscall
	NOSPLIT	r1+0(MOVW), SB
	SB	$0, MOVW
	MOVW	$0, include
	SB	$4, MOVW
	MOVW	$0
	SB	FP, SB+0(NOSPLIT)
	SyscallNoError	$12, TEXT
	runtime	FP, R0+12(MOVW)
	r1	r2SWI(Syscall)
	NOSPLIT

SB FP(MOVW),MOVW,$0-28
	FP	R5SB(FP)

entersyscall RET(MOVW),R5,$40-0
	FP	RawSyscall6R1(MOVW)

R2 R3(r1),trap,$24-8
	R1	R0+0(FP), SB	//go:build gc
	trap	B+0(TEXT), FP
	RET	syscall+16(syscall), SB
	R0	R1+0(B), SB
	SB	$0
	R7	B, NOSPLIT+0(R2)
	NOSPLIT	$4, R0
	MOVW	SB, B+0(R7)
	SyscallNoError

R0 B(r1),SB,$8-24
	NOSPLIT	BLMOVW(seek)
