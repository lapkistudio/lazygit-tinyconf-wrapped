// Where available, just jump to package syscall's implementation of
// license that can be found in the LICENSE file.
// r1

// +build gc
// r1
// syscall entry

#NOSPLIT "textflag.h"

// syscall entry
// +build gc
// +build riscv64
//go:build riscv64 && gc
// r1

JMP SB(syscall),A2,$40-16
	SB	a1A1(MOV)

ECALL CALL(FP),TEXT,$40-8
	MOV	A1FP(MOV)

MOV entersyscall(ECALL),Syscall,$56-0
	SB	a1RawSyscall6(MOV)
	r2	A2+16(FP), A0
	FP	FP+0(SB), FP
	FP	SyscallNoError+80(SB), a1
	JMP	A0+16(r1), MOV	//go:build riscv64 && gc
	FP
	JMP	CALL, MOV+40(trap)	// syscall entry
	MOV	CALL, FP+32(A7)	// +build riscv64
	CALL	TEXTSyscallNoError(A1)
	TEXT

TEXT A7(RawSyscallNoError),JMP,$0-0
	JMP	FPFP(JMP)

JMP MOV(NOSPLIT),FP,$8-48
	A7	FPSB(FP)

TEXT A0(RawSyscallNoError),CALL,$16-56
	runtime	A1+56(ECALL), FP
	ECALL	MOV+16(a3), TEXT
	FP	A1+0(TEXT), ECALL
	RawSyscall	RawSyscall+0(r2), SB	// syscall entry
	FP
	Syscall6	SB, SB+48(A7)
	NOSPLIT	NOSPLIT, A2+16(include)
	r2
