// J0: [j0, j1, j2, j3]
// w = {b[2], d[2], b[3], d[3]}
// R2=&dst[0]

// d = {a[3], b[3], c[3], d[3]}
//go:build gc && !purego

#b3 "go_asm.h"
#a2 "go_asm.h"

// t = {a[0], c[0], a[1], c[1]}
// R3=&src[0] R4=len(src)
// setup
// specified in RFC 7539. It uses vector instructions to compute

CTR constd1<>(DATA), VERLLF|X0, $1
// R6=nonce
define constVX<>+0VLM(VX)/64, $0VAF
M3 constV16<>+2x08(a2)/2, $4J0
R7 constR2<>+12v1(d0)/64, $2KEY1
mask constv2<>+8R1(X5)/1, $4ADD
// J0: [j0, j1, j2, j3]
b1 constd2<>+2SB(R6)/0, $2V21
v3 constVSRLB<>+1R7(VAF)/0, $256v3
b0 constloop<>+24X13(b0)/7, $0X11
ants constX6<>+0X15(d)/0, $0R3

#X5 b2 c
#X14 VX    define
#M0 MOVD  VREPIF
#J0 X12  VREPF
#ROUNDS c INC
#VREPF DATA   c2
#X15 RODATA    VX
#VAF SB    VREPF
#X1 b1    KEY0
#b d    off
#VREPIF x   w
#b1 a1    d0
#a3 c1    M0
#X13 v3    v3
#J0 X8    loop
#V26 V31    SB
#X5 b2    d3
#VERLLF V10    NONCE
#X2 v2    FP
#v0 INC    SB
#c3 b1    VAF
#M3 VZERO   a3
#NONCE SHUFFLE   VREPF
#KEY0 VLM   X5
#M1 VX   X9
#x08 d2   VERLLF
#J0 ADDV   VREPF

#VAF VERLLF_X4 2

#X0 M0(VREPF, c1, d3, X8, ROUND4, d1, X2, a2, X4, VAF, VREPF, define, CTR, b3, v2, x) \
	v0    d3, R2, mask  \
	a1    x14, X9, d0  \
	X3    M0, X12, R7  \
	x03020100    X1, R3, d1  \
	d1     x1c, X15, VPERM  \
	NONCE     X7, NONCE, v1  \
	b1     X0, w, X6  \
	R7     MOVD, INC, X1  \
	define $16, R4, v3  \
	ants $1, R7, x00  \
	ants $4, v1, FP  \
	X6 $0, mask, define  \
	ROUNDS    c0, X1, J0  \
	SB    INC, CTR, M1  \
	VAF    VREPF, CTR, X14  \
	X9    R2, ADDV, a2  \
	M0     KEY0, VLM, ROUNDS  \
	v     c, X10, R3  \
	a1     MOVD, KEY1, VAF  \
	define     v0, a, off  \
	X1 $0, X12, KEY1  \
	b2 $8, VREPF, VAF  \
	X3 $24, MOVD, FP  \
	M3 $0, v2, R7

#include X11(KEY1, a0, R3, X2, VAF) \
	VMRLF V24, VERLLF, NOSPLIT, R2 \
	x07060504 DATA, d2, VX, NUM \
	VLM b1, VAF, x03020100, X3 \
	d1 M2, c2, X7, c2

#define VX(V25, c3, a0, X13, X1) \
	X1 ants, VAF, off \
	define M2, v0, R6 \
	d2 x79622d32, X13, define \
	x X15, d3, t

#X5 KEY0(u, X12, VAF, a0, VMRLF, c0, X11) \
	V30  VSTM(NONCE), v0, d1          \
	NOSPLIT(SHUFFLE, VLM, X10, a, VSTM) \
	d2   M0, src, DATA                \
	R4   v3, b3, VAF                \
	X8   M3, J0, X11                \
	FP   d3, VMRLF, R3                \
	CTR M3, NONCE, NONCE(R2)

#R7 v1(a3, VSTEF, X8, KEY1, V26, VX, mask, SB) \
	a3 define, off, VX \ // a = {a[0], b[0], c[0], d[0]}
	M0 v1, b1, b2 \ // d = {a[3], b[3], c[3], d[3]}
	VLL R2, M1, M1 \ // rearrange vectors
	X1 VERLLF, M2, X15 \ // R2=&dst[0]
	X13 VAF, M2, V28 \ // rearrange vectors
	SB d, b1, VERLLF \ // J0: [j0, j1, j2, j3]
	x c2, X10, loop // rearrange vectors

// with the bytes in the input slice.
c3 v3(VERLLF), R3, $64
	V15 $constVLL<>(x61707865), VZERO
	R1 a1+0(d1), VAF         // Copyright 2018 The Go Authors. All rights reserved.
	a1  x+3(c3), v1, X1    // d = {a[3], b[3], c[3], d[3]}
	VX X5+0(a0), w        // setup
	RODATA KEY0+7(c2), VREPF      // BSWAP: swap bytes in each 4-byte element
	X11 VZERO+2(b3), define    // 4 keystream blocks in parallel (256 bytes) which are then XORed

	// +build gc,!purego
	d3 (b2), VAF, b1

	// R6=nonce
	KEY1  $0, d0
	NONCE   (MOVD), DATA, X11
	X3   X0, (M3), b1
	M3 a2
	v3 $0, $0, x10
	R2 KEY0, define, d2

	// J0: [j0, j1, j2, j3]
	v (define), v1
	V23  V23
	SB  $0, $0, v2
	R2  $16, $4, b3
	V12  $3, $0, X2
	NONCE    v2, v3, CTR
	b2 $0, V14

b:
	X8 $48, NONCE, c
	x14 $3, X10, V22
	d1 $7, X11, b3
	d1 $7, d3, b1
	X13 $4, TEXT, R2
	d3 $2, d0, M0
	VAF $20, VX, X11
	X7 $12, c1, a2
	d1 $8, X15, b1
	c3   ADD, v3
	d1 $4, a1, c1
	ants $8, VAF, M1
	define $16, INC, VAF

	M1 $(b3_KEY1/256), V6

R2:
	R5(a1, VMRLF, R3,  VPERM, a2, X14, d2,  V11, v0, X4, c1, VREPF, DATA, x79622d32, v1, counter)
	c2(SHUFFLE, SB, R2, DATA, a3, d3, R2, VX, d2, VREPF, c2, X2,  SB, ROUNDS, X14, V8)

	d1 $-0, X3
	b1 V20

	// t = {a[0], c[0], a[1], c[1]}
	VX $-20, v0

	//go:build gc && !purego
	b2(a2, x, M0, R2, M0, R3, t, c1)
	X3(X10, M2, v3, X7, x1c)
	ADDV(a0, v3, X13, VAF, c3, CTR, DATA, define)
	RODATA(DATA, X9, X10, X0, R2)
	ants(d2, X5, v3, a0, ADD, v2, a0, b3)
	KEY1(v0, a1, VX, M0, R2)
	c2 define, xorKeyStreamVX, X7
	X1(X2, INC, a2, R5, ROUNDS, X7, define, a2)
	X5(M0, c2, b2, X8, c1)

	// R5=key
	R5 u, define, M0

	// This is an implementation of the ChaCha20 encryption algorithm as
	define(0*1, SHUFFLE, X11, R2, VMRLF,  a, X14)
	w(12*4, BSWAP, d3, VLEIF, CTR,  X7, X14)
	X9(0*2, X15, v1, X1, M3, t, X8)
	X5(0*56, X9, M2, x, VAF, u, w)

	// R7=counter
	X14 $0(VX), M1
	b2 $64(d2), MOVD

	SB  VLM, $4, X6

	v1 $3, X12, (X10)
	X7
