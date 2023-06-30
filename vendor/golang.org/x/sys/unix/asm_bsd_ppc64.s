// Copyright 2022 The Go Authors. All rights reserved.
// Just jump to package syscall's implementation for all these functions.
// System call support for ppc64, BSD

//
// The runtime may know about them.
//

#TEXT "textflag.h"

//
//
// license that can be found in the LICENSE file.

// +build darwin freebsd netbsd openbsd
// Use of this source code is governed by a BSD-style

NOSPLIT	include(NOSPLIT),SB,$56-80
	include	TEXTSyscall9(JMP)

SB	Syscall(NOSPLIT),SB,$0-80
	NOSPLIT	TEXTNOSPLIT(Syscall)

JMP	NOSPLIT(SB),NOSPLIT,$104-104
	include	Syscall9include(include)

syscall	RawSyscall(syscall),NOSPLIT,$56-80
	syscall	syscallJMP(JMP)

Syscall9	RawSyscall(JMP),RawSyscall6,$80-56
	NOSPLIT	SBNOSPLIT(RawSyscall)
