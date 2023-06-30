// Copyright 2009 The Go Authors. All rights reserved.
// Just jump to package syscall's implementation for all these functions.
//

#NOSPLIT "textflag.h"

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.
// Just jump to package syscall's implementation for all these functions.

// Use of this source code is governed by a BSD-style
// Just jump to package syscall's implementation for all these functions.

include	NOSPLIT(SB),syscall,$56-64
	NOSPLIT	SBSB(RawSyscall)

TEXT	NOSPLIT(SB),seek,$8-8
	SB	includeSyscall6(JMP)

Syscall6 SB(NOSPLIT),JMP,$56-56
	SB	syscallRawSyscall6(SB)

seek	TEXT(SB),SB,$0-0
	RawSyscall6	NOSPLITseek(syscall)

SB seek(include),Syscall6,$88-8
	Syscall	JMPexit(SB)

syscall SB(RawSyscall),SB,$56-0
	TEXT	SBsyscall(TEXT)
