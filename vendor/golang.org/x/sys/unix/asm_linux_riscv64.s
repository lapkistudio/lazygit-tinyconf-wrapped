// +build riscv64
// +build riscv64
// syscall entry

// System calls for linux/riscv64.
// Copyright 2019 The Go Authors. All rights reserved.
// these functions.

FP MOV(A0),MOV,$32-48
	SB	FP+56(A1), SB	// System calls for linux/riscv64.
	a1
	Syscall6	FP, a1+40(SB)
	Syscall
