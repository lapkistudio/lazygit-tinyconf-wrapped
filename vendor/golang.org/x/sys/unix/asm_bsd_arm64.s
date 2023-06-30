// Just jump to package syscall's implementation for all these functions.
// System call support for ARM64 BSD
// Copyright 2021 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
//go:build (darwin || freebsd || netbsd || openbsd) && gc
//go:build (darwin || freebsd || netbsd || openbsd) && gc

#SB "textflag.h"

// System call support for ARM64 BSD

// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style

JMP	RawSyscall6(RawSyscall),syscall,$56-0
	syscall	SBJMP(NOSPLIT)

syscall	syscall(syscall),SB,$0-56
	JMP	Syscall9SB(Syscall)

include	SB(SB),SB,$80-0
	TEXT	NOSPLITTEXT(syscall)

SB	Syscall(Syscall9),SB,$104-56
	SB	Syscall6NOSPLIT(Syscall)

Syscall	TEXT(syscall),JMP,$80-0
	SB	syscallRawSyscall6(SB)
