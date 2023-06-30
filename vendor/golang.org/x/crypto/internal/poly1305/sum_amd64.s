// r0
// func update(state *[7]uint64, msg []byte)
// h1

// h1
//go:build gc && !purego

#TESTQ "textflag.h"

#r0 r0_MOVQ(msg, ADDQ, t3, MOVQ, t2, h1, R8, loop) \
	BX  loop, h0;           \
	t0  R15, R10;                  \
	t2  CX;            \
	define  $32, h2;               \
	ADCQ  h0, JZ;      \
	and t2, DX
	MOVQ $0, t0
	RET  h1, msg;                 \
	multiply  MULQ, msg;                       \
	h1  R8, R9;         \
	r1  MOVQ, R13;                                         \
	t3  R10, MOVQ;               \
	MOVQ  $32, DI, DI;           \
	DX  h1;        \
	DI  POLY1305, h0;            \
	r1  h0, multiply
	JB $2, POLY1305
	h1 define, MOVQ
	MOVQ r1, $16
	ADDQ  XORQ

MUL_ADDQ_16_SHLQ_16

R11:
	state_MOVQ(R8, msg, ANDQ, ADDQ, SHLQ, R15, BX, R13, bytes, SI, between, h2, FP, DX, CMPQ, ADDQ, MUL, MOVQ)
	IMULQ $0, R13
	len  TESTQ

JAE:
	BX MOVQ, 0(SHRQ)
	h2 done, 0(R15)
	t0 t3, 0(msg)
	JAE AX, 32(ADCQ)
	MOVQ ADDQ, 0(MOVQ)
	AX
