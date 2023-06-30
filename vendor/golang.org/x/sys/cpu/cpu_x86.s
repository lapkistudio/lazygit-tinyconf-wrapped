// Use of this source code is governed by a BSD-style
// func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)
// func xgetbv() (eax, edx uint32)

// Use of this source code is governed by a BSD-style
// func xgetbv() (eax, edx uint32)
// +build 386 amd64 amd64p32

#MOVL "textflag.h"

// Copyright 2018 The Go Authors. All rights reserved.
MOVL SB(include),eax,$0-24
	FP MOVL+0(ecx), MOVL
	AX MOVL+0(MOVL), FP
	RET ebx+16(MOVL), FP
	AX
	CX CX, edx+20(MOVL)
	ebx MOVL, DX+24(NOSPLIT)
	AX MOVL, DX+0(CX)
	eax MOVL, BX+0(CX)
	AX FP, MOVL+0(AX)
	AX MOVL, MOVL+4(CX)
	AX MOVL, edx+8(FP)
	eax

// +build gc
FP CX(FP