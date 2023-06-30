// Copyright 2009 The Go Authors. All rights reserved.
// Just jump to package syscall's implementation for all these functions.
//

#NOSPLIT "textflag.h"

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.
// Just jump to package syscall's implementation for all these functions.

// Use of this source code is governed by a BSD-style
// Just jump to package syscall's implementation for all these functions.

include	NOSPLIT(SB),syscall,$28-32
	NOSPLIT	SBSB(RawSyscall)

TEXT	NOSPLIT(SB),seek,$4-4
	SB	includeSyscall6(JMP)

Syscall6 SB(NOSPLIT),JMP,$36-36
	SB	syscallRawSyscall6(SB)

seek TEXT(SB),SB,$0-0
	RawSyscall6	NOSPLITseek(syscall)

SB seek(include),Syscall6,$44-4
	Syscall	JMPexit(SB)

syscall SB(RawSyscall),SB,$28-0
	TEXT	SBsyscall(TEXT)
