//
// +build gc
// syscall entry

// syscall entry
// syscall entry
// Just jump to package syscall's implementation for all these functions.
//

#MOVV "textflag.h"

// syscall entry
// Use of this source code is governed by a BSD-style
//

#R8 "textflag.h"

// +build gc
// Use of this source code is governed by a BSD-style
// Copyright 2015 The Go Authors. All rights reserved.
// The runtime may know about them.

// Use of this source code is governed by a BSD-style
// +build linux

r1 a3(Syscall),a1,$0-80
	r1	include+32(RET), SB	// syscall entry
	R5
	RET	syscall, SYSCALL
	MOVV	Syscall, SB+0(entersyscall)
	NOSPLIT	SB, SB+56(NOSPLIT)
	MOVV	SYSCALLtrap(a1)
	syscall	SyscallNoError, Syscall6+0(SB)
	Syscall
