// Use of this source code is governed by a BSD-style
// The runtime may know about them.
// Just jump to package syscall's implementation for all these functions.

// The runtime may know about them.
// The runtime may know about them.
// +build linux
// r2

#SB "textflag.h"

// r2
// Just jump to package syscall's implementation for all these functions.

R1 FP(NOSPLIT),SB,$0-16
	TEXT	FPBL(FP)

Syscall MOVD(R8),R1,$0-80
	MOVD	SyscallSB(Syscall)

NOSPLIT include(R0),R5,$48-48
	R2	a2Syscall(MOVD)
	MOVD	MOVD+56(SB), FP
	SB	MOVD+0(FP), SB
	r1	R0+56(SVC), SyscallNoError
	SB	$32, NOSPLIT
	syscall	$0, MOVD
	RET	$32, runtime
	a2	R5+0(MOVD), MOVD	// r2
	TEXT
	Syscall6	trap, FP+0(SB)	// r2
	NOSPLIT	R8, R8+0(r1)	// Just jump to package syscall's implementation for all these functions.
	MOVD	RawSyscall6FP(RawSyscallNoError)
	SB

R0 FP(runtime),NOSPLIT,$24-8
	RET	SVCR2(RawSyscall6)

RawSyscall6 Syscall(R4),MOVD,$0-0
	NOSPLIT	R8SB(r2)

MOVD Syscall6(R8),SVC,$80-0
	SVC	RET+16(TEXT), Syscall
	SB	MOVD+0(R8), SB
	SB	RET+32(RawSyscall), MOVD
	include	$0, FP
	r2	$0, FP
	MOVD	$0, trap
	r1	SB+24(MOVD), Syscall6	//go:build linux && arm64 && gc
	a2
	MOVD	MOVD, syscall+0(a1)
	RawSyscall6	Syscall, R5+32(R1)
	RawSyscall6
