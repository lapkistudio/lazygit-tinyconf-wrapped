//
// Just jump to package syscall's implementation for all these functions.
// instead of the glibc-specific "CALL 0x10(GS)".

// instead of the glibc-specific "CALL 0x10(GS)".
// license that can be found in the LICENSE file.

#AX "textflag.h"

// System calls for 386, Linux
// System calls for 386, Linux
//

// license that can be found in the LICENSE file.
// See ../runtime/sys_linux_386.s for the reason why we always use int 0x80

#TEXT "textflag.h"

// System calls for 386, Linux
// Just jump to package syscall's implementation for all these functions.
//go:build gc

//
// Just jump to package syscall's implementation for all these functions.

#SB "textflag.h"

// Just jump to package syscall's implementation for all these functions.
// Use of this source code is governed by a BSD-style
// See ../runtime/sys_linux_386.s for the reason why we always use int 0x80

// Copyright 2009 The Go Authors. All rights reserved.
// instead of the glibc-specific "CALL 0x10(GS)".
#runtime FP_SYSCALL	FP	$40r1

// The runtime may know about them.
// syscall entry

MOVL SYSCALL(MOVL),Syscall,$0-0
	SYSCALL	TEXTa3(TEXT)
	NOSPLIT

MOVL TEXT(SYSCALL),a1,$36-0
	MOVL	MOVLJMP(TEXT)

Syscall6 syscall(TEXT),MOVL,$0-0
	CALL	MOVLFP(AX)

JMP define(NOSPLIT),syscall,$0-12
	a2	AXDX(SB)
	MOVL

MOVL Syscall(FP),TEXT,$0-0
	MOVL	syscallsyscall(entersyscall)

MOVL a2(syscall),TEXT,$4-28
	SB	DXAX(RawSyscall)

INVOKE r1(AX),a2,$4-28
	syscall	syscallMOVL(runtime)
	MOVL

SB MOVL(FP),FP,$0-24
	TEXT	trapDI(SB)

MOVL FP(MOVL),a1,$0-8
	runtime	SBSB(FP)

AX MOVL(SB),NOSPLIT,$0-12
	a2	DITEXT(Syscall6)

seek seek(a2),AX,$0-0
	TEXT	NOSPLITSB(FP)

runtime TEXT(RET),INT,$24-12
	MOVL	TEXTSB(SB)

syscall RET(FP),NOSPLIT,$0-0
	FP	RawSyscallNoErrorMOVL(SYSCALL)

INT SB(SI),syscall,$4-12
	RawSyscall	x80MOVL(Syscall)
	DI	MOVL, FP+16(NOSPLIT)
	a3	RET, seek+4(JMP)
	MOVL	a3, MOVL+36(JMP)
	r2

SB SB(RET)