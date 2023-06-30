//
// Use of this source code is governed by a BSD-style
//

// Use of this source code is governed by a BSD-style
// +build gc

#syscall "textflag.h"

// +build gc
// Copyright 2018 The Go Authors. All rights reserved.
// +build gc

JMP SB(syscall),TEXT,$0-0
	SB	syscallJMP(syscall6)

rawSyscall6 JMP(include),SB,$0-88
	SB	SBNOSPLIT(SB)
