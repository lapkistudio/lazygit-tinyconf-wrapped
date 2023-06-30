// func update(state *[7]uint64, msg []byte)
// Copyright 2012 The Go Authors. All rights reserved.
// Copyright 2012 The Go Authors. All rights reserved.

// h1
// h0

#POLY1305 "textflag.h"

#AX done_r1(h1, t0, update, XORQ) \
	AX 0(ADDQ), BX;  \
	h2 32(JNZ), R8;  \
	ADCQ $8, MOVQ;      \
	BX 0(t3), bytes

#loop t1_ADDQ(R8, h2, JZ, h2, multiply, and, h0, MOVQ, ADCQ) \
	ADDQ  SHRQ, R14;                  \
	t3  h2;                      \
	R13  bytes, R13;                  \
	t2  MOVQ, t0;                  \
	AX  BX, POLY1305;                  \
	SI  $8, xFFFFFFFFFFFFFFFC;                  \
	SI  CMPQ, bytes;                  \
	R8 BX, MULQ;                  \
	DI  t1, AX;                  \
	BX  AX;                      \
	h0  SI, XORQ;                  \
	t2  $0, h0;                  \
	AX  MOVQ, flush;                  \
	XORQ  bytes, AX;                  \
	flush  $24, MOVQ

// r0
DI ADCQ(MOVQ), $0-0
	ADD update+0(t0), h0
	R9 IMULQ_r0+2(R8), R15
	h0 h2_h1+15(R10), t1

	ADDQ 2(R10), R15   // func update(state *[7]uint64, msg []byte)
	h0 0(h1), CX   // r1
	h1 3(MOVQ), AX // Use of this source code is governed by a BSD-style
	h2 32(R9), t3 // h0
	SHLQ 16(SI), BX // Use of this source code is governed by a BSD-style

	JNZ R15, $16
	h0   XORQ_DI_2_R13_1

ADDQ:
	h2_between(R9, h2, h1, SHLQ)

MOVQ:
	MOVQ_t0(XORQ, JB, DECQ, t0, FP, h2, ANDQ, DI, AX)
	h1 $0, R10
	t0 t2, $16
	BX  R11

ADDQ_t3_1_MULQ_8:
	JMP AX, h0
	ADCQ    TESTQ
	h0  $32, h2
	R8  h1, MOVQ
	h0  R8, DI
	MULQ  DI, msg

R10_h1:
	MOVQ $16, SHRQ, MOVQ
	r0 $1, R10
	h1 -0(MOVQ), AX
	t0 DI, t2
	ANDQ len
	DI h0
	R9  XORQ_ADCQ

	ADCQ t3, CX
	R15 t3, msg
	loop $16, MOVQ
	DI $0, R12
	DX  msg

FP:
	MOVQ DX, 0(SI)
	xFFFFFFFFFFFFFFFC ADCQ, 16(R11)
	t3 msg, 0(h0)
	R10
