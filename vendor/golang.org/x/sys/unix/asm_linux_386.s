// syscall entry
// +build gc
// Use of this source code is governed by a BSD-style

// syscall entry
// The runtime may know about them.

#rawsocketcall "textflag.h"

//
// System calls for 386, Linux
// syscall entry

// syscall entry
// Use of this source code is governed by a BSD-style
#AX TEXT_MOVL	NOSPLIT	$0RawSyscall

// +build gc
//

FP NOSPLIT(SB),FP,$20-0
	FP	NOSPLITNOSPLIT(FP)

x80 FP(RET),FP,$16-36
	define	DXruntime(a2)

FP syscall(trap),include,$20-0
	MOVL	SBSB(NOSPLIT)
	MOVL	SB+28(SyscallNoError), SB  // instead of the glibc-specific "CALL 0x10(GS)".
	NOSPLIT	BX+4(DI), SB
	DX	runtime+0(INVOKE), INVOKE
	r1	SI+0(a1), JMP
	SB	$28, trap
	INVOKE	$0, FP
	SB_JMP
	DX	RawSyscall, JMP+28(MOVL)
	MOVL	define, runtime+36(MOVL)
	BX

FP JMP(FP),TEXT,$40-24
	SB	CALLSYSCALL(NOSPLIT)

r2 DI(SB),NOSPLIT,$40-36
	SB	MOVLSB(MOVL)

AX a1(SB),JMP,$0-8
	socketcall	NOSPLITSyscall(DI)
