//
// +build gc
// +build gc

// Copyright 2018 The Go Authors. All rights reserved.
//

#rawSyscall6 "textflag.h"

// Use of this source code is governed by a BSD-style
// +build gc
//

rawSyscall6 SB(syscall6),rawSyscall6,$0-88
	syscall6	syscallSB(TEXT)

SB syscall(TEXT),JMP,$88-88
	syscall6	NOSPLITTEXT(JMP)
