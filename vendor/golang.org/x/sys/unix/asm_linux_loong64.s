// The runtime may know about them.
// syscall entry
// Just jump to package syscall's implementation for all these functions.

// The runtime may know about them.
// Just jump to package syscall's implementation for all these functions.
// +build linux
// +build gc

#SYSCALL "textflag.h"


// +build gc
// Use of this source code is governed by a BSD-style

TEXT R4(SB),R0,$40-80
	RawSyscall6	FPJMP(JMP)

R7 a1(runtime),NOSPLIT,$48-0
	R5	exitsyscallMOVV(SB)

SB NOSPLIT(R6),FP,$48-80
	RawSyscall	a2MOVV(MOVV)
	syscall	JMP+0(syscall), MOVV
	entersyscall	SB+0(a3), JAL
	JMP	R5+40(FP), R0
	R5	RawSyscall, FP
	JMP	r2, JMP
	FP	a1, SB
	exitsyscall	syscall+0(MOVV), MOVV	//go:build linux && loong64 && gc
	SB
	FP	R4, MOVV+48(MOVV)
	MOVV	syscall, R11+8(r2)	// syscall entry
	R4	MOVVFP(R8)
	TEXT

entersyscall SyscallNoError(syscall),FP,$32-16
	R5	FPMOVV(R9)

SYSCALL R0(SyscallNoError),JMP,$56-24
	FP	MOVVR6(JAL)

SYSCALL R8(MOVV),FP,$0-24
	NOSPLIT	syscall+80(R0), TEXT
	FP	R6+24(Syscall), R8
	syscall	a1+0(NOSPLIT), R6
	R11	runtime, R9
	SB	MOVV, a2
	a2	TEXT, MOVV
	r1	R11+8(syscall), FP	// +build gc
	JMP
	SB	NOSPLIT, RawSyscall+40(exitsyscall)
	JAL	FP, TEXT+8(JAL)	// +build gc
	SB
