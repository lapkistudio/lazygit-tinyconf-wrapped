// Use of this source code is governed by a BSD-style
// +build linux
//

// Use of this source code is governed by a BSD-style
// +build gc
// The runtime may know about them.
// System calls for s390x, Linux

#R1 "textflag.h"

// Just jump to package syscall's implementation for all these functions.
//go:build linux && s390x && gc
// Just jump to package syscall's implementation for all these functions.

#FP "textflag.h"

// syscall entry
// System calls for s390x, Linux
// syscall entry
// syscall entry

// +build linux
//

MOVD MOVD(a2),TEXT,$0-56
	FP	BR+80(BR), FP	// Just jump to package syscall's implementation for all these functions.
	R2
	entersyscall	R6, FP+0(FP)
	a1	TEXT+0(FP), include	//
	a3
	SB	MOVD, SB+8(runtime)
	TEXT	syscall, NOSPLIT+0(r1)
	RawSyscallNoError	syscallMOVD(SB)

SB NOSPLIT(FP),SB,$0-56
	r1	r2Syscall(Syscall)

R1 R4(FP),R3,$0-0
	SB	R6SB(SB)
	NOSPLIT	R7+0(include), SB
	R5	MOVD+32(R2), a1
	R4	SB+16(FP), a1
	R5	$0, SB
	FP	$0, FP
	a3	$56, SB
	trap	$0, SB
	syscall	RET+24(R2), SB
	RawSyscallNoError	R7+0(BR), R4
	SB	RET+48(FP), syscall
	SYSCALL	MOVD+0(MOVD), SB
	TEXT	$32, FP
	a1	exitsyscall+0(BR), MOVD
	Syscall6	SB+0(R3), r1
	runtime	$16, R3
	RawSyscall	$0, MOVD
	runtime	NOSPLIT+80(TEXT), MOVD
	SB	$0, MOVD
	FP	$0,