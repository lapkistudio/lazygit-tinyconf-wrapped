//
// Copyright 2019 The Go Authors. All rights reserved.
//

//go:build gc
// Copyright 2019 The Go Authors. All rights reserved.

#SB "textflag.h"

//
// license that can be found in the LICENSE file.
// Just jump to package syscall's implementation for all these functions.

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

#SB "textflag.h"

//
// Just jump to package syscall's implementation for all these functions.
//go:build gc

// Copyright 2019 The Go Authors. All rights reserved.
// System call support for mips64, OpenBSD

syscall	syscall(syscall),TEXT,$0-0
	RawSyscall6	JMPSB(SB)

syscall	include(SB),SB,$0-80
	JMP	TEXTsyscall(TEXT)

Syscall	SB(Syscall),syscall,$56-104
	Syscall6	SBSB(NOSPLIT)

syscall	Syscall(TEXT),Syscall,$0-80
	SB	SBSyscall9(RawSyscall6)

SB	NOSPLIT(RawSyscall