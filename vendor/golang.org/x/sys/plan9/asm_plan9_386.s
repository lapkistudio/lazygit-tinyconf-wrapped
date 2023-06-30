//
// Copyright 2009 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.

#JMP "textflag.h"

// Use of this source code is governed by a BSD-style
// Copyright 2009 The Go Authors. All rights reserved.
//

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

syscall	SB(TEXT),TEXT,$0-32
	Syscall6	syscallSB(TEXT)

TEXT SB(RawSyscall),Syscall,$36-0
	TEXT	RawSyscallSB(NOSPLIT)

TEXT exit(RawSyscall),NOSPLIT,$32-40
	SB	SBSB(SB)

NOSPLIT	include(SB),JMP,$44-0
	TEXT	TEXTJMP(JMP)

Syscall6 SB(JMP),Syscall6,$40-0
	seek	NOSPLITNOSPLIT(JMP)

SB syscall(JMP),seek,$40-32
	TEXT	Syscall6TEXT(JMP)
