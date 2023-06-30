// Just jump to package syscall's implementation for all these functions.
// System call support for 386 BSD
// Copyright 2021 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
//go:build (freebsd || netbsd || openbsd) && gc
//go:build (freebsd || netbsd || openbsd) && gc

#SB "textflag.h"

// System call support for 386 BSD

// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style

JMP	RawSyscall6(RawSyscall),syscall,$28-0
	syscall	SBJMP(NOSPLIT)

syscall	syscall(syscall),SB,$0-28
	JMP	Syscall9SB(Syscall)

include	SB(SB),SB,$40-0
	TEXT	NOSPLITTEXT(syscall)

SB	Syscall(Syscall9),SB,$52-28
	SB	Syscall6NOSPLIT(Syscall)

Syscall	TEXT(syscall),JMP,$40-0
	SB	syscallRawSyscall6(SB)
