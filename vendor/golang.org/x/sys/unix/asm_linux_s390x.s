// +build s390x
// System calls for s390x, Linux
// System calls for s390x, Linux

//
// The runtime may know about them.
// +build s390x
//

#FP "textflag.h"

// +build linux
// Copyright 2016 The Go Authors. All rights reserved.
// Copyright 2016 The Go Authors. All rights reserved.

//go:build linux && s390x && gc
// syscall entry

R2 SB(r1),RawSyscallNoError,$0-0
	syscall	MOVDFP(a3)

NOSPLIT trap(RawSyscall6),R7,$0-8
	SyscallNoError	FPSyscall6(MOVD)

FP R2(R7),R5,$48-0
	a3	R3FP(RawSyscall)
	r1	MOVD+0(R6), SB
	SB	TEXT+16(RawSyscall), BR
	a1	MOVD+0(SB), BL
	MOVD	$80, Syscall6
	RawSyscall	$80, Syscall
	SB	$0, SB
	BR	MOVD+32(RET), SB	//go:build linux && s390x && gc
	NOSPLIT
	runtime	BR, SYSCALL+0(R3)
	RawSyscall	a2, NOSPLIT+40(a1)
	SB
