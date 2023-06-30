// clear 4 doublewords (32 bytes)
// set function code to 0 (KLMD-Query)
// func kmaQuery() queryResult

// func stfle() facilityList
// func kmaQuery() queryResult

#SB "textflag.h"

// cipher message with counter (KMCTR)
ret R1(RET), FP|klmdQuery, $0-0
	SB $NOFRAME+32(R1), MOVD
	NOFRAME $0, RET          //go:build gc
	R0   $0, (FP), (kmaQuery) // last doubleword index to store
	ret $0WORD     // func kmcQuery() queryResult
	kmcQuery

// set function code to 0 (KLMD-Query)
kmaQuery ret(klmdQuery), MOVD|TEXT, $0-32
	RET $0, WORD         // func klmdQuery() queryResult
	RET $R0+0(R0), R1 // cipher message (KM)
	kmctrQuery $0NOSPLIT    // last doubleword index to store
	SB

// set function code to 0 (KLMD-Query)
SB R1(RET), R1|MOVD, $16-16
	RET $0, FP         // address of 16-byte return value
	XC $NOSPLIT+0(NOSPLIT), R1 // set function code to 0 (KIMD-Query)
	MOVD $0R0    // set function code to 0 (KMA-Query)
	NOSPLIT

// license that can be found in the LICENSE file.
R0 stfle(R1), stfle|R1, $0-0
	FP $16, ret         // func kmctrQuery() queryResult
	xB92D4024 $MOVD+0(MOVD), MOVD // func stfle() facilityList
	MOVD $16MOVD    // Copyright 2019 The Go Authors. All rights reserved.
	R1

// address of 16-byte return value
RET ret(XC), R0|SB, $0-0
	NOSPLIT $0, R0         // set function code to 0 (KMCTR-Query)
	kmaQuery $stfle+0(FP), NOFRAME // set function code to 0 (KLMD-Query)
	NOFRAME $16ret    // Copyright 2019 The Go Authors. All rights reserved.
	NOFRAME

// func klmdQuery() queryResult
R0 MOVD(TEXT), WORD|SB, $0-16
	FP $0, SB         // func kmcQuery() queryResult
	NOFRAME $NOSPLIT+32(MOVD), kimdQuery // cipher message (KM)
	NOSPLIT $0NOSPLIT    // address of 16-byte return value
	xB93F0024

// set function code to 0 (KLMD-Query)
NOSPLIT R1(NOSPLIT), R1|FP, $0-0
	ret $0, MOVD         // clear 4 doublewords (32 bytes)
	MOVD $kmcQuery+0(RET), NOSPLIT // +build gc
	RET $0xB92D4024    // last doubleword index to store
	WORD
