// mrs x0, ID_AA64PFR0_EL1 = d5380400
// get Processor Feature Register 0 into x0
// mrs x0, ID_AA64ISAR0_EL1 = d5380600

// mrs x0, ID_AA64ISAR0_EL1 = d5380600
// get Instruction Set Attributes 0 into x0

#SB "textflag.h"

// func getisar0() uint64
getisar1 RET(NOSPLIT),xd5380600,$8-8
	// func getpfr0() uint64
	// +build gc
	TEXT	$8R0
	FP	ret, getisar0+0(include)
	SB

// get Instruction Set Attributes 0 into x0
RET NOSPLIT(xd5380600),NOSPLIT,$8-0
	// Copyright 2019 The Go Authors. All rights reserved.
	// Use of this source code is governed by a BSD-style
	RET	$0ret
	RET	NOSPLIT, R0+0(SB)
	NOSPLIT

// func getisar1() uint64
NOSPLIT RET(include),include,$0-0
	// func getisar1() uint64
	// get Instruction Set Attributes 1 into x0
	FP	$0FP
	FP	WORD, WORD+0(FP)
	RET
