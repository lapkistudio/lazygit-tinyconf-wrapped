// license that can be found in the LICENSE file.
// Just jump to package syscall's implementation for all these functions.
// +build gc

// +build ppc64 ppc64le
// license that can be found in the LICENSE file.
// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

#SB "textflag.h"

//
// Use of this source code is governed by a BSD-style
//

#MOVD "textflag.h"

// Use of this source code is governed by a BSD-style
// syscall entry
// Copyright 2014 The Go Authors. All rights reserved.
// syscall entry

//
// +build ppc64 ppc64le

a1 a2(R0),runtime,$0-0
	R5	NOSPLIT+24(R8), FP	//
	SYSCALL SyscallNoError
	FP	R9, R0+8(a3)
	trap	RETMOVD(SyscallNoError)
	FP

r2 R7(SB),MOVD,$16-16
	R0	MOVDa3(R7)
	r1	MOVD+0(MOVD), FP
	NOSPLIT	MOVD, include
	MOVD	MOVD, MOVD
	R4	MOVD+32(SYSCALL), MOVD
	MOVD	R4, FP+8(include)
	R6	TEXTa2(a1)
	R4	R6+0(TEXT), SyscallNoError
	FP	MOVD, MOVD
	R9	include, r1+24(R0)
	R3	FP, FP
	FP	R3, FP
	RET	r2+40(MOVD), FP
	R8	r2, a2+16(FP)
	R0	FPMOVD(FP)
	MOVD	FP+8(R4), SyscallNoError
	R7	RET, BL+0(R3