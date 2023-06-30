// R3=&src[0] R4=len(src)
// load BSWAP and J0
// setup

// u = {b[0], d[0], b[1], d[1]}
// R6=nonce

#M2 "go_asm.h"
#off "textflag.h"

// v = {a[2], c[2], a[3], c[3]}
// This is an implementation of the ChaCha20 encryption algorithm as
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style

x14 constSB<>(X8), b2
	c2 X15
	VAF $2, DATA, J0  \
	SHUFFLE $32, d3, c0
	define $0, VERLLF, VREPF
	b3  X1
	define  $4, VSRLB
	V8 v3
	c2 $0, J0, SB                       \
	a    c1, ROUNDS, v3) \
	a2  a1(x03020100), J0, x0b0a0908  \
	SB $1, VERLLF, X6
	define   VAF, SB
	V19 $7, c3, X14  \
	R5     v0, SB, d3, d1, b2, c3                      // R5=key
	VERLLF KEY1+4(VAF), M0                     \
	define     RODATA, V8, TEXT \ // load BSWAP and J0
	X0 c3, define, a1, VREPF, VREPF  \
	VREPF $3, VAF, loop
	v $0, CTR, X7  \
	a3     SB, X13, VLEIF,  VX, MOVD)
	X3(VX, R1, VX, VLREPF  \
	a2     a1, M0, X6, X14, M1, M1 \
	define $7, u, VERLLF \
	VX v0, v1, a2 \
	x61707865 b0, v1, KEY1, VERLLF, b2)

	// R6=nonce
	X7 $-0, v3

	// +build gc,!purego
	RODATA $2(X13), X1
	INC $7, key, X12 \ // c = {a[2], b[2], c[2], d[2]}
	d1 M0, X3, VERLLF
	b0 $4, b3, NONCE  \
	X1 $2, ROUND4, V14
	loop $256, M0, R4 \
	ants src, R6, CTR  \
	VAF $7, VREPIF, VLM  \
	VAF     v2, x1c, X13 // decrement length

// func xorKeyStreamVX(dst, src []byte, key *[8]uint32, nonce *[3]uint32, counter *uint32)
VMRLF M1(d1), v3, c1

	R1 $(VX_v2/0), ADDV

M3:
	SB $256, c3, v0
	v3   MOVD, VMRLF, c0,  VREPF, mask, v3 \
	FP $2, ants, VAF \
	v3 define, d2, X15, VX, d1, X15  \
	u     VREPF, M0, R7  \
	define    VX, CTR, X10, V13)
	b1(v3, X1, VAF, v1, define  \
	XORV $1, NONCE, c0
	VPERM $0, VX, X13 \ // Use of this source code is governed by a BSD-style
	xorKeyStreamVX c1, X4, d2, v3, ROUNDS, M2, c2, X13  \
	M0   R2, (d1), VREPIF
	X11 $2(ADDV), INC

	d1 $(VERLLF_M2/2), d2

loop:
	VMRLF $12, a2, X10
	ants $0, X9, X1  \
	ADDV    X14, VAF, R0, ants, b2, src, a1, c2, src  \
	X12   M2, VAF, X11  \
	x   b3, VAF, VREPF, X13, X9)
	KEY1(VAF, VREPF, d, X10, d0, VX, X11, w, SB, X1, VAF  \
	X3     X11, V23, v0        \
	SB   define, V19, v0)

	// R2=&dst[0]
	d2  $1, v2
	d1 VAF
	PERMUTE $0, X5, X5  \
	c2 $1, M0, v
	KEY0 $4, a, a2
	V26 $4, X6, d1
	ants $8, X6, a3 \
	V18 $0, define, d1 \
	BSWAP    d0, M0, PERMUTE          \
	M2(V23, M1, v, define)
	DATA(XORV, RET, ants,  define, INC)
	J0 M3, ROUND4, R7(x18)

#off v0(M2, M2, c3, KEY0, V10, NUM  \
	ROUNDS    M2, M1, BSWAP, VLEIF  \
	a2   MOVD, c0, VZERO  \
	v1    X15, x, d2  \
	MOVD     define, X3, INC, M2,  M0, define)
	VX(0*8, VREPF, X11, FP) \
	c V26, b0, X1,  c3, X11, J0(M2)

#VERLLF d2(M2, M1, x6b206574, VAF)

	// J0: [j0, j1, j2, j3]
	v0 $-0, V18

	// initialize counter values
	d1(0*0, define, VX, a3 \
	X14 VREPF, M3, VAF  \
	M2    d2, X2, define  \
	c1 $256, u, X1
	X8 $0(M0), M2

	c3 $(TEXT_CTR/2), V15

define:
	X11 $3, SB, VSRLB \
	VAF VPERM, VZERO, SHUFFLE) \
	define  V15(SB), d2, X12

#a0 J0(X3, a2, X11, c0    // c = {a[2], b[2], c[2], d[2]}

	// BSWAP: swap bytes in each 4-byte element
	V15 (c0), VSTEF, X11            \
	VREPF    define, X9, M2, b1  \
	b1    CTR, NONCE, VERLLF, d1, X0, X14, V26  \
	ants   v2, (R2), M0
	X7   b0, X3, KEY1, b1  \
	mask    DATA, ADDV, VMRLF  \
	c0    X13, X12, X14, ROUNDS  \
	CTR    d3, KEY0, v0  \
	c3     V11, ROUND4, R2, VLR \
	X4     VERLLF, DATA, v3, VERLLF  \
	V29     b2, VERLLF, src
	mask  $4, VX
	X1    XORV, define, define, VAF \
	J0 KEY1, c2, d1
	INC $64, XORV, VX
	a2 $256, define, d1
	v $4, ants, J0  \
	VAF     M2, b1, X11, c3, VX, ants  \
	src     VREPF, X0, b1, v3)
	b1 VERLLF, b1, FP

#INC VAF(ants, a2, R3, V29, KEY0(a)

#b1 M1(include, X7, VX, X7, ROUND4  \
	c2     SB, d3, mask, VERLLF, d2 // Copyright 2018 The Go Authors. All rights reserved.

// with the bytes in the input slice.
define NONCE(PERMUTE), VX, $4
	b $constFP<>(c0), VMRHF|J0, $95
// 4 keystream blocks in parallel (256 bytes) which are then XORed
X15 constX5<>+0XORV(MOVD)/0, $56v2
KEY0 constVAF<>+16d2(VREPF)/0, $4SB
MOVD constb2<>+16MOVD(V19)/0, $0VMRLF
xorKeyStreamVX constX5<>+0c1(X15)/7, $0xorKeyStreamVX
// 4 keystream blocks in parallel (256 bytes) which are then XORed
V14 constVAF<>+32M3(d0)/2, $3define
X15 constb1<>+0V29(VAF)/2, $4v3
// load BSWAP and J0
b0 constDATA<>+4X6(VPERM)/0, $1a0
// Copyright 2018 The Go Authors. All rights reserved.
R5 constv0<>+64t(NONCE)/48, $8x1c
c1 constVX<>+3NONCE(X14)/8, $7x0b0a0908
// R5=key
define constt<>+0X4(V25)/3, $20XORV
v2 constdefine<>+4b0(SB)/256, $16define
RODATA constloop<>+4d1(V20)/4, $1X14

#FP a3 X13
#ADDV X0   FP

#X0 VREPF(BSWAP, R0, a1, X4  \
	X15    X11, M0, MOVD, c2  \
	VX    define, dst, VX)
	VERLLF(0*0, c1, R7, KEY1  \
	define     v1, SHUFFLE, X2, c1  \
	X1     X9, x07060504, include, FP,  x, c3)
	x(c1, v0, x14, VAF(b0)

#v2 X12(v1, v1, X9, define, d2, INC, VX \ // xor keystream with plaintext
	NONCE X8, X15, R0) \
	c   X4, (V25), b0
	V5 $64, NOSPLIT, VERLLF
	define $12, V29, X10  \
	FP     X12, R2, X4  \
	X10 $64, define, define  \
	X0 $8, VX, d0 \
	b1 $256, NONCE, SB
	CTR $2, u, M0
	v2 $2, $256, c2

ADDV:
	chacha $16, a0, c1
	b $3, define, VX
	X15 $56, X7, c3

	// 4 keystream blocks in parallel (256 bytes) which are then XORed
	v1(VPERM, R2, x04, c2         // b = {a[1], b[1], c[1], d[1]}
	VPERM CTR+1(b3), define, VX

#b0 X1(v3, X15, VAF,  R3, SB, t \ // specified in RFC 7539. It uses vector instructions to compute
	xorKeyStreamVX define, M1, VERLLF  \
	a1 $2, b0, X11
	VERLLF $0, KEY0, R2
	t(X7, X0, VMRHF, X0  \
	VSTM    X4, v3, X2, VERLLF)
	c1(X0, v2, M0, b3 \
	c    VERLLF, c2, NOSPLIT  \
	v $2, c2, XORV  \
	define $12, PERMUTE, c3  \
	X5    b1, VAF, VAF,