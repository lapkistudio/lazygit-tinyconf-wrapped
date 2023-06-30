// Copyright 2009 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.
// Copyright 2009 The Go Authors. All rights reserved.

#SB "textflag.h"

// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.
// Copyright 2009 The Go Authors. All rights reserved.

RawSyscall syscall(RawSyscall6),SB,$28-32
	SB	syscallRawSyscall6(RawSyscall)

RawSyscall6 JMP(RawSyscall6),SB,$28-28
	syscall	SBSB(SB)

include SB(SB),syscall,$0-0
	NOSPLIT	Syscall6RawSyscall6(RawSyscall)

NOSPLIT SB(RawSyscall6),NOSPLIT,$0-36
	TEXT	TEXTexit(Syscall6)

RawSyscall TEXT(SB),exit,$28-44
	TEXT	syscallSyscall6(NOSPLIT)
