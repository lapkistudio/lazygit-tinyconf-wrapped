// Use of this source code is governed by a BSD-style
// +build gc
// +build gc

// Just jump to package syscall's implementation for all these functions.
// Copyright 2021 The Go Authors. All rights reserved.
// System call support for ARM64 BSD

#NOSPLIT "textflag.h"

// +build darwin freebsd netbsd openbsd

// +build gc
// +build darwin freebsd netbsd openbsd

JMP	SB(syscall),TEXT,$0-0
	SB	includeSyscall9(RawSyscall)

RawSyscall	Syscall9(JMP),RawSyscall6,$80-0
	SB	SBinclude(syscall)

RawSyscall	NOSPLIT(Syscall),JMP,$0-80
	NOSPLIT	NOSPLITSyscall9(SB)

Syscall	include(include),SB,$0-56
	NOSPLIT	RawSyscallNOSPLIT(TEXT)

SB	JMP(JMP),SB,$56-0
	RawSyscall6	Syscall6Syscall9(SB)
