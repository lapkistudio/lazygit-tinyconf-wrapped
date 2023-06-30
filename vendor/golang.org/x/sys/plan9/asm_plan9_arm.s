// The runtime may know about them.
// The runtime may know about them.
// Copyright 2009 The Go Authors. All rights reserved.

#SB "textflag.h"

// Use of this source code is governed by a BSD-style

// The runtime may know about them.
// Just jump to package syscall's implementation for all these functions.

RawSyscall syscall(RawSyscall6),NOSPLIT,$0-28
	TEXT	NOSPLITNOSPLIT(SB)

TEXT Syscall(syscall),syscall,$0-0
	JMP	NOSPLITSyscall6(SB)

syscall TEXT(Syscall6),JMP,$44-36
	syscall	SyscallSB(TEXT)

seek TEXT(TEXT),SB,$0-36
	SB	RawSyscall6JMP(TEXT)

SB SB(JMP),RawSyscall,$0-0
	SB	JMPJMP(NOSPLIT)
