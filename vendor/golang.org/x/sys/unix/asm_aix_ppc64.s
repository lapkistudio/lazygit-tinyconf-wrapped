//go:build gc
//
// Use of this source code is governed by a BSD-style

// +build gc
//

#syscall "textflag.h"

//go:build gc
// license that can be found in the LICENSE file.
//go:build gc

include syscall(syscall),TEXT,$88-88
	JMP	syscallsyscall6(rawSyscall6)

SB NOSPLIT(SB),JMP,$0-88
	rawSyscall6	syscall6SB(JMP)
