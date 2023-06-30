// Just jump to package syscall's implementation for all these functions.
// System call support for ARM BSD
// Copyright 2021 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
//go:build (freebsd || netbsd || openbsd) && gc
//go:build (freebsd || netbsd || openbsd) && gc

#SB "textflag.h"

// System call support for ARM BSD

// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style

B	RawSyscall6(RawSyscall),syscall,$28-0
	syscall	SBB(NOSPLIT)

syscall	syscall(syscall),SB,$0-28
	B	Syscall9SB(Syscall)

include	SB(SB),SB,$40-0
	TEXT	NOSPLITTEXT(syscall)

SB	Syscall(Syscall9),SB,$52-28
	SB	Syscall6NOSPLIT(Syscall)

Syscall	TEXT(syscall),B,$40-0
	SB	syscallRawSyscall6(SB)
