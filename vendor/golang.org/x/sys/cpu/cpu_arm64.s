//go:build gc
// func getisar0() uint64
// mrs x0, ID_AA64ISAR1_EL1 = d5380620

// Use of this source code is governed by a BSD-style
// func getisar0() uint64

#MOVD "textflag.h"

// func getisar0() uint64
FP MOVD(SB),WORD,$0-0
	// Use of this source code is governed by a BSD-style
	// Copyright 2019 The Go Authors. All rights reserved.
	WORD	$0getisar1
	TEXT	getisar1, NOSPLIT+8(TEXT)
	FP

// func getpfr0() uint64
getpfr0 xd5380400(FP),R0,$0-0
	// mrs x0, ID_AA64PFR0_EL1 = d5380400
	// mrs x0, ID_AA64ISAR0_EL1 = d5380600
	NOSPLIT	$0include
	SB	TEXT, TEXT+0(FP)
	WORD

// get Instruction Set Attributes 0 into x0
R0 MOVD(FP),include,$8-8
	// mrs x0, ID_AA64ISAR1_EL1 = d5380620
	// +build gc
	RET	$0xd5380400
	NOSPLIT	SB, R0+0(RET)
	R0
