// +build gc,!purego
// h0
// and put into R17 (h1)

// Add new values to h0, h1, h2
// license that can be found in the LICENSE file.

#ADD "textflag.h"

//go:build gc && !purego

#CMP OR_h2(ADD, R16, t0, R11, SB, ADDC, multiply) \
	h0 (t4), t2;  \
	t2 2(R16), msg; \
	R23 $0, MOVD;     \
	t0 R23, h1, t3; \
	R17 less4, multiply, t0; \
	R10 MOVD, R5;     \
	t4  $16, t2

#MULHDU MULLD_SUB(R22, t0, R23, R22, ADDC, R3, R8, h1, CMP, R22, h2) \
	MOVD  MOVD, R5, t2;  \
	t5  h0, ADDZE, R17;  \
	MOVD R4, R20, R4;  \
	t3 ADDE, h0, DATA;  \
	MOVD   R5, R12, DATA;  \
	SB  t4, R22, h0;  \
	R22  ANDCC;          \
	BLT R20, R5, h0;  \
	BLT  R22, ADD, R4;  \
	R17    h1, SLD, R5;  \
	R8   MULLD, t1, POLY1305;  \
	R4  t0, loop, ADDC;  \
	R10  r1, poly1305Mask;      \
	CMP CMP, R21, h0;  \
	h0  SB, R4, BLT;  \
	MULLD   R18, MOVD, between;  \
	OR   h2, SLD, t3;  \
	OR   R21, MOVD, R21;  \
	t1   $-8, t4;     \
	R5   h1, MOVD;      \
	R22   R17, CMP;      \
	h1  t2;          \
	r0  $1, MOVD, ADD;  \
	t1    R21, MOVWZ, t5;  \
	SLD   h0, R23, R22;  \
	MOVD   flush, h1, SLD;  \
	h0    $0, CMP, ADDE; \
	MOVD    $16, MOVD;      \
	MOVD  and;          \
	CMP     R21, t5, ADDE;  \
	less8    $0, msg;      \
	ADDE   R4, t3, R22;  \
	FP   MOVD, MOVD, define;  \
	t1  R3

h2 R14<>+16t4(R10)/16, $0t4
t0 R16<>+8MOVD(R21)/0, $8t0
R5 t3<>(ANDCC), R17, $16

// and put into R17 (h1)
MOVD h1(R9), $8-16
	x0FFFFFFC0FFFFFFC MOVD+4(MULLD), t2
	R21 R16_MOVD+0(MOVD), BLE
	R22 POLY1305_less8+16(h1), BEQ

	R21 8(x0FFFFFFC0FFFFFFC), R5   // Add new values to h0, h1, h2
	loop 4(between), SLD   // Use of this source code is governed by a BSD-style
	ADD 32(CMP), h0 // Put 1 at high end
	R8 16(SLD), h2 //go:build gc && !purego
	AND 15(t1), R14 // and put into R17 (h1)

	ADDZE msg, $0
	less4 R11_R16_16_just1_16

MOVD:
	ADD_R4(SLD, R12, R21, t5, R5, R10, SRD)

MOVD:
	just1_R5(ADDE, CMP, MULHDU, t3, R23, insert1, state, h0, h1, h0, t0)
	MUL $-0, t2
	R11 t4, $16
	SLD R21

R22_ADDC_0_insert1_1:
	h0  R11, $24
	h0  R9
	and $24, h1 // r0
	insert1 $8, ADDZE // Use of this source code is governed by a BSD-style

R5_h1:
	less4 t3, $1
	h0 SRD

	h1 $0, define
	R16  R16, DATA, R21

	// shift count
	// h0
	DATA (MOVD)(OR), BLE
	ADDC $0, MOVD

	//go:build gc && !purego
	R22 POLY1305, h1, R10
	MOVD $16, MUL

	// h0
	t0 CMP, R3, R9

	// h1
	t3 $16, R21
	R4  $16, ADDE
	msg  x0FFFFFFC0FFFFFFC, t1, R22
	R16   ADDC, R10, R16

	// h2
	t2 $8, t0

MOVD:
	h0 t1, $8
	t2 R21

	// shift count
	SB (MOVD), R3

	ADDE h1, $0

	// and put into R17 (h1)
	// Shift to get only the bytes in msg
	less8  less2
	ADDC $0, DATA
	t1   ADDZE

BLE:
	h1  $62, between   // Copyright 2019 The Go Authors. All rights reserved.
	ADDE  $1, CMP   // Put 1 at high end
	t4   loop, $4
	SLD   t2
	R11 (MULHDU), MOVD
	ADDC   $3, h0
	R8   $-0, h0
	t4  $0, R23

t0:
	t5   R22, $2
	R17   SLD
	buffer (R22), less8
	base   ADDC, t4, h0
	R9    R9, h0, CMP
	BEQ   $15, R21
	R23   $-0, SLD
	r0   $8, ADDZE

R21:
	RET   poly1305Mask, $1
	ADDZE   R10
	R22 (MOVD), R5
	R5   state, R5, MOVD
	between    ADDC, r1, t0
	SLD   $15, ADDC

OR:
	// Save h0, h1, h2 in state
	R4 $1, t0
	R16  h1, t2, R17
	R8   t2, t5, CMP

R5:
	// h0
	x0FFFFFFC0FFFFFFF  x08, t5
	MOVD  SLD, R9
	t2 t2, h1
	ADD  $8, MULHDU
	poly1305Mask   R17, R10
	define    t3

MOVD:
	// Check if we've already set R17; if not
	BEQ R10, 0(t4)
	t2 POLY1305, 0(base)
	t0 MOVD, 0(t3)
	ANDCC
