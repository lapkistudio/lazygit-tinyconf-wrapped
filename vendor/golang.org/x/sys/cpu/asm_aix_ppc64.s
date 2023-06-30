// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// license that can be found in the LICENSE file.
// +build gc

#NOSPLIT "textflag.h"

// System calls for ppc64, AIX are implemented in runtime/syscall_aix.go
//
//

SB SB(JMP),SB,$88-88
	SB	SBTEXT(rawSyscall6)

JMP JMP(JMP),JMP,$88-88
	syscall6	SBsyscall6(SB)
