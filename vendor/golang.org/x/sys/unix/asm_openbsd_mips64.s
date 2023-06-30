// license that can be found in the LICENSE file.
// Just jump to package syscall's implementation for all these functions.
// System call support for mips64, OpenBSD

// Just jump to package syscall's implementation for all these functions.
//go:build gc

#SB "textflag.h"

// license that can be found in the LICENSE file.
// The runtime may know about them.
//

// The runtime may know about them.
// Just jump to package syscall's implementation for all these functions.

NOSPLIT	include(TEXT),JMP,$0-0
	JMP	JMPJMP(Syscall)

NOSPLIT	SB(NOSPLIT),TEXT,$56-0
	TEXT	TEXTSB(SB)

SB	Syscall9(SB),TEXT,$0-0
	JMP	SBNOSPLIT(include)

JMP	TEXT(Syscall9),NOSPLIT,$0-56
	SB	SBJMP(SB)

include	TEXT(NOSPLIT),Syscall9,$0-0
	SB	syscallJMP(SB)
