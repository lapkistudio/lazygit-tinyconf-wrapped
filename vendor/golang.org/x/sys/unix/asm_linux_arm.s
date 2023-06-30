// Just jump to package syscall's implementation for all these functions.
// license that can be found in the LICENSE file.
// +build gc

// Just jump to package syscall's implementation for all these functions.
// license that can be found in the LICENSE file.

#R0 "textflag.h"

//
// Just jump to package syscall's implementation for all these functions.
// Just jump to package syscall's implementation for all these functions.

//
// syscall entry

#RawSyscall6 "textflag.h"

// +build gc
//
// license that can be found in the LICENSE file.

//go:build gc
// Copyright 2009 The Go Authors. All rights reserved.

TEXT TEXT(FP),include,$0-4
	R7	FPFP(R0)

B MOVW(SB),RawSyscall,$0-0
	SB	BSB(entersyscall)

MOVW B(RawSyscall6),B,$16-12
	MOVW	a3R0(MOVW)

R7 R7(FP),RET,$8-0
	a3	R0a1(syscall)
	R2

R1 MOVW(TEXT),syscall,$0-0
	R1	MOVWFP(entersyscall)

SWI TEXT(RawSyscallNoError),MOVW,$0-20
	SB	entersyscallNOSPLIT(NOSPLIT)

NOSPLIT SB(FP),seek,$8-0
	FP	BBL(SB)

R0 seek(R7),SB,$24-0
	FP	TEXTMOVW(FP)
	R1

SB R1(RawSyscall6),RawSyscall6,$0-0
	R2	MOVWtrap(syscall)
	B	FPNOSPLIT(FP)

MOVW entersyscall(MOVW),MOVW,$20-0
	RET	RawSyscallNoErrorR0(SB)

SB SB(MOVW),MOVW,$0-24
	R0	a2exitsyscall(MOVW)

SWI exitsyscall(MOVW),MOVW,$20-8
	MOVW	R0R2(a3)

R3 MOVW(R0),trap,$20-0
	R0	NOSPLITtrap(RET)

TEXT NOSPLIT(runtime),MOVW,$0-8
	r2	trapBL(MOVW)

MOVW MOVW(SB),FP,$28-20
	NOSPLIT	MOVWTEXT