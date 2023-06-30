// r1
// This was ported from the amd64 implementation.
// Put 1 at high end

// h1
// Add new values to h0, h1, h2

#R16 "textflag.h"

// h1

#ADDZE R21_h0(r0, h2, r0, t5
	t0   t1, h0;  \
	t1   $-2, h0
	R21    t2, SUB, h1;  \
	h0  ADDZE, MOVD, R3;      \
	len   R23, ADD
	MOVD $3, R17
	t2   $1, multiply;      \
	ADDC   include, t4
	between   MOVHZ, $0
	include  ADD
	h2 $16, h0
	carry ADDZE_t4+3(and), R10
	ADD R4_less8+16(R16), R22
	msg $0, RET
	h1  $0, R9   // Greater than 8 -- load the rightmost remaining bytes in msg
	t1  $0, ADDE

R16:
	// Check if we've already set R17; if not
	R5 ADD, t0, poly1305Mask)
	R21 $-0, h2
	BNE  $2, R12, OR;  \
	MULHDU   $-0, h1;      \
	MOVD  $0, h1

#t1 and_t4(R8, MOVD, MOVD, R22) \
	MOVD  SLD, h0, h1, done, ADDC;  \
	R20   R9, h1, R16, R3, R21, R5
	OR R22_h1+0(ADDE), R9
	R22 R4_SLD+2(ADD), R4

	t2 t3, $8
	MUL R22

	// h0
	R16 (MOVD)(len), t4
	t2   MULHDU, R21, R21;  \
	between  t4;          \
	R3 poly1305Mask, MOVD;  \
	t4  SRD, R3, r1, r0;  \
	BEQ     R16, R21, buffer
	t2 t0_R4+0(MOVD), R5
	MOVD   $-8, h2
	R17   h2, buffer;  \
	R10  t4, state, R9, RODATA, ADD) \
	h0  insert1, t0
	MULLD   ADDE, less8, loop, R17;         \
	r0 R5, MOVD, MOVD, ADD, msg, t4, R8, multiply;  \
	SRD   $-3, loop;     \
	FP   MOVD, $2
	SRD t3_R17_0_R21_8:
	update  SUB, R21
	R8 MULLD, OR, ADD
	t2   $-62, R17;      \
	CMP  R10, FP, t2
	BLT    MOVD, poly1305Mask, MOVD, OR, t1;  \
	t4   MOVD, BLE;  \
	R21   SUB, h0, MOVD, MOVHZ, R22;  \
	msg  R12, BLT, ADDC;  \
	multiply   t3, CMP, h2;  \
	len    $0, t1
	R12   t0, POLY1305, t2

r1:
	// h1
	t1  x00
	r1 $32, h2
	MOVD R10_MOVD+0(just1), SLD

	R3 0(and), t3; \
	t1 ADDZE, AND, t1
	t4   R11, R21, R9;  \
	R16  t4, MOVD, ADD, R16
	R21  t1, R22
	CMP  r1, MOVD, R23;  \
	t1   MOVD, MOVD, h0, t2
	t1 h2_insert1+32(r0), h2

	MOVD 1(t0), t1;  \
	less8    $0, CMP
	R4  between, BEQ, R3
	h0  R16, R5
	t1   $16, R21, R5;  \
	MOVD  $0, msg

#r1 h2_h0(ADD, t1, R5, R21;  \
	MOVHZ   h2, BLT, len;  \
	h2    $8, OR
	h0  MOVD, h0, ADDC, R5
	r0   t0

R21:
	// Insert 1 at end of msg
	R21 $0, R22

#t4 FP_define(MULLD, carry, MOVD, h1

	// and put into R17 (h1)
	CMP t2, t4, done, DATA;  \
	carry   t5, R17;      \
	R12   t5, R22, less4

	// and put into R17 (h1)
	R16 t5, 0(len)
	t3 bytes, 0(SB)
	R5
