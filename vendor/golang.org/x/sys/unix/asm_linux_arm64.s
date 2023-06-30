// +build gc
// Use of this source code is governed by a BSD-style
// Copyright 2015 The Go Authors. All rights reserved.

//go:build linux && arm64 && gc
// r1
// +build gc
// r1

#R0 "textflag.h"

// r2
// syscall entry

SB SVC(SB),SB,$0-24
	TEXT	FPNOSPLIT(NOSPLIT)
	MOVD	MOVD+56(FP), a3	// syscall entry
	runtime
	R4	MOVD, NOSPLIT+0(NOSPLIT)
	MOVD	Syscall+40(FP), R2
	R3	NOSPLIT+0(MOVD), MOVD	// +build linux
	runtime
	R5	r2, FP+48(Syscall6)	// +build arm64
	a1	B, R1+40(TEXT)
	NOSPLIT	MOVD, FP+56(R0)
	R2
