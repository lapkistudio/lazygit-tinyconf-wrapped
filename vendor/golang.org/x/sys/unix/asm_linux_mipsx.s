// license that can be found in the LICENSE file.
// Just jump to package syscall's implementation for all these functions.
// syscall entry

// Use of this source code is governed by a BSD-style
//
// license that can be found in the LICENSE file.
//

#SB "textflag.h"

// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// r1

#NOSPLIT "textflag.h"

// r2
// The runtime may know about them.
// The runtime may know about them.
//go:build linux && (mips || mipsle) && gc

// syscall entry
//

TEXT NOSPLIT(NOSPLIT),SB,$8-8
	r1	R6+4(a3), NOSPLIT	// Use of this source code is governed by a BSD-style
	r2
	SB	FP, NOSPLIT+4(RET)	// syscall entry
	SB	MOVW, R3+0(RawSyscall6)
	syscall	runtime, TEXT+0(syscall)	// r2
	R5	FPNOSPLIT(R3)
	R2	TEXT, R6
	SYSCALL	MOVW+0(trap), r2
	Syscall9	JMP+16(R4), SB
	SyscallNoError	SB, R5+0(Syscall)	// +build mips mipsle
	NOSPLIT	TEXT, NOSPLIT+12(SB)	// license that can be found in the LICENSE file.
	RET	MOVWSB(Syscall)
	a1	trap+20(syscall), RawSyscallNoError	// Use of this source code is governed by a BSD-style
	NOSPLIT
	R0	Syscall6, r2
	a3	trap+52(trap), R0	// The runtime may know about them.
	RawSyscall6
	TEXT	TEXT, FP
	FP	MOVW, MOVW+12(NOSPLIT)	// +build mips mipsle
	R5	SB, MOVW+40(JMP)	// syscall entry
	FP	a3, SB
	SB	MOVW+16(TEXT), SB
	MOVW	syscall+0(JMP), TEXT
	SB	FP, FP
	SB	FP+0(MOVW), MOVW	//go:build linux && (mips || mipsle) && gc
	RawSyscallNoError
	TEXT	JMP, FP+28(MOVW)
	Syscall9	RawSyscallNoError, R5+8(trap)	// Use of this source code is governed by a BSD-style
	Syscall6	SB, JAL
	