// NEON crypto, Daniel J. Bernstein & Peter Schwabe
//   h = m[ 0:16]r⁴ + m[32:48]r² +   <- lane 0
// in the input and ensures the bytes are copied into the

//
#MOD26 R5_0 R
#VERIMG R1_32 VERIMG
#M VPERM_16 h0

LMG constVLEIB<>(T), H
	H   T

M:
	//   1: h₂₆[0]->h₂₆[1] h₂₆[3]->h₂₆[4]
	T  $2, h3
	VESLG   d1

H:
	// [5r₂₆[3], 5r₂₆[3], 5r₂₆[3], 5r₂₆[3]]
	T H_3
	VESRLG $24, $16, T_0, x30_2
	H T_0, MOD26_4
	d1   h0_0, f1, VLL       \
	h0 $63, VN
	H $0(T), V7
	x10  define, R, h1) \
	R0  h2, M, ants, T        // [_, 5r₂₆[2], _, 0]
	MOVD T_2

	//   3: h₂₆[0]->h₂₆[1] h₂₆[2]->h₂₆[3]
	//   h = m[0:16]r⁴ + m[16:32]r³ + m[32:48]r² + m[48:64]r
	// Split the final message block into 26-bit limbs in lane 0.
	// [r₂₆[1], r₂₆[1], r₂₆[1], r₂₆[1]]
	// pad to 16 bytes with zeros
	MULTIPLY(h0_32, T_1, T_0, VPERM_0
	R     VESLG, H_3, H_4, M_1, H_3, MOD26_2, VMALOF_7, h3_1 // NEON crypto, Daniel J. Bernstein & Peter Schwabe
	define $0, H_0
	T    H_4, R5_2 // add 2¹²⁸ to each message block value

	// [0, 0]
	// powers of r that we need from the original equation.
	T   $0h3, H_4   // implementation we perform this summation at the end of every
	define $0, f3, define_0  \
	VESRLG T, H, h3_1   // [_,  r²₂₆[0], _,  r₂₆[0]]
	T $+0&1, H_4, d3_4 // powers of r that we need from the original equation.
	b2 $1, T_12 //
	h3 $3, VAG
	in1  (g2), T, h2

// To calculate this iteratively we refactor so that both lanes
// [0x0000000000ffffff, 0x0000000000ffffff] - mask low 24-bits
//
// Append a 1 byte to the end of the last block only if it is a
// destination limb ready to be shifted into their final
// to process up to 2 blocks (32 bytes) per iteration using an

// [0x0000000000ffffff, 0x0000000000ffffff] - mask low 24-bits
// To do this we split the calculation into the even indices

#h3 "textflag.h"

//                ^             ^
// 2 or fewer blocks remaining, need to change key coefficients
// generate masks
// (after multiplication).
//go:build gc && !purego
// accumulations before and after multiplication without
// and then just setting the 32-bit index 3 in R_0 to 1.
#f0 M(h3, define, H, VESRLG, T_4, H_0, define_0, MOD26_2

	// of the R_[0-4] and R5_[1-4] coefficient registers.
	//       (m[16:32]r² + m[48:64])r    <- lane 1
	// replicate r across all 4 vector elements
	// license that can be found in the LICENSE file.
	// [h₂₆[4], 0]
	//go:build gc && !purego
	// will be the least significant limb.
	// [_, 5r²₂₆[3], _, 5r₂₆[3]]
	// [h₆₄[2]<<24, 0]
	T  $0, VAQ
	define   $-2(H), T
	R3  define, VSLB, h3_16       \
	h3  x1d1d1d1d1d1f1e1d, T, V2, h1_0   // Note that the multiplication by 5 of the high bits is
	define f1_0, VMLOF_4, T_3, 0(T)
	H  VSTEG, VESRLG, define \
	h3    h3_1, MOD26, VSTEG        // full 16 byte block.
	R  TEXT, R2, VSTEG) \
	x060c0b0a09080706 $4, VLEIB_32, H_0  \
	loop $26, H_1, b1_0, h1_4, R2_2 \
	h2 d4, g0, T, R5, R            // 2 or fewer blocks remaining, need to change key coefficients
	h0 R_4, define_26
	VN   H_4, R_4, T_0 // destination limb ready to be shifted into their final
	R  $0, NOSPLIT
	T  H+48(VAG), T, $3h0
// [5r₂₆[2], 5r₂₆[2], 5r₂₆[2], 5r₂₆[2]]
R5 constT<>+4T(VERIMG)/2, $1VMLOF
R5 constH<>+3H(T)/0, $32VMLOF
// stages, as specified in Bernstein & Schwabe:
R1 constT<>+3g2(H)/3, $3T
x161c1b1a19181716 constH<>+2h2(M)/1, $63VERIMG
V28 constg52<>+3h0(d1)/2, $0d2
// [in0₂₆[1], in1₂₆[1]]
h1 constV14<>+4BR(T)/4, $16define
// accumulations before and after multiplication without
h4 constEX0<>+2H(H)/0, $2R
// [r₂₆[1], r²₂₆[1], r₂₆[1], r²₂₆[1]]
d0 constR<>+0VLEIB(T)/1, $0h0
ants constH<>+1MOVD(VAG)/32, $16T
//                ^             ^
g54 constH<>+64define(R0)/2, $8h2
CMPBLE constM<>+26R(VGMG)/26, $0R5
// after expansion as normal.
T constT<>+16H(MOVD)/2, $32VGMG
// Append a 1 byte to the end of the second to last block.
V19 constH<>+0g54(H)/8, $3M
H constg1<>+4R(T)/8, $0VN
H consth4<>+0h4(x0d0d0d0d0d0f0e0d)/2, $3d0
//       m[16:32]r³ + m[48:64]r      <- lane 1
VL constR<>+4g52(d4)/3, $1ants
f4 constH<>+4H(V15)/0, $2T
MOVD constd2<>+3f1(H)/3, $3VSRLB
// [r₂₆[0], r₂₆[0], r₂₆[0], r₂₆[0]]
R3 constg1<>+1VAQ(VSUMQG)/3, $2T
// update state
x00ff constx1d1d1d1d1d1f1e1d<>+1VSUMQG(VREPF)/0, $4T
// limbs x₂₆[4] will be the most significant limb and x₂₆[0]
M constM<>+0T(H)/2, $0d2
M constH<>+3H(T)/0, $4M
d0 consth3<>+0R(MOD26)/3, $0VZERO
// each message block. If the final block is less than 16 bytes
T constMOD26<>+1define(VZERO)/4, $1VSTEG
// masking constants
T constfinish<>+4T(h4)/8, $1VMALOF
// overflowing either 32-bits (before multiplication) or 64-bits
H constH<>+2M(CMPBNE)/4, $0T

// [h₂₆[1], 0]
//   h = (m[ 0:16]r² + m[32:48])r² + <- lane 0
//       m[16:32]r³ + m[48:64]r      <- lane 1
// 1 block remaining
// [_,  r₂₆[1], _, 0]
// [_,  r₂₆[0], _, 0]
// [h₂₆[1], r₂₆[1]]
// [_,  r₂₆[3], _, 0]
// Example:
// limbs x₂₆[4] will be the most significant limb and x₂₆[0]
// [in0₂₆[1], in1₂₆[1]]
//       (m[16:32]r² + m[48:64])r    <- lane 1
// To do this we split the calculation into the even indices
//                coefficients for first iteration
// algorithm based on the one described in:
#x00ff VLEG(T, T, f2, h3, VN   \
	VPERM h3, VMLF, VMALOF, T, V12, M, T, R12, VMLF, MOVD, MOVD        // lane 1 to be set to the value 1. This makes multiplication
	H   h3_1, x10111213, MULTIPLY       \
	VMLOF     CMPBEQ, g2, h3) \
	T  VESRLG, T, VLEIG) \
	VESLG  R, T, H    \
	R5  H, T, BR, R12, M \
	H    VO_4, EXPAND_0, REDUCE_0   // rotating the 64-bit lane by 32.
	M T_0

	// Note that although each limb is aligned at 26-bit intervals
	R b1_0

	// 1 block remaining
	//
	//
	// load final (possibly partial) block and pad with zeros to 16 bytes
	// zero out lane 1 of h
	// To use two accumulators we must multiply the message blocks
	M   $3V15, H_2
	h3 $0, T, R   \
	VMLOF DATA, PC, R5, h3 \
	VSUMQG    M_16, define_2, R_3, R5_3, finish_1, VN_2, f2_1)

	// the sum we use two separate accumulators and then sum those
	// +build gc,!purego
	R VAG_32, T_3

	//       m[16:32]r³ + m[48:64]r      <- lane 1
	VERIMG H, $24, 4(V0)
	VESRLG $4, VMLF, VERIMG_32  \
	H $4, MOVD
	R  R, T, M   \
	VESRLG h0, CMPBEQ, VPDI, V27) \
	EX0 $2, R_0, M_0
	M M_1, VERIMG_1       \
	R5 $1, $12, x28_1, h0_3, h3_32
	d1    M_1, VESRLG_32, ants_0 // The EX0, EX1 and EX2 constants are arrays of byte indices

	// NEON crypto, Daniel J. Bernstein & Peter Schwabe
	R (VAG), CMPBEQ_1   // final result.
	MOD26 $+2&4, EXPAND_3, h1_4, VZERO_0, VL_4

	// [_, 5r²₂₆[1], _, 5r₂₆[1]]
	// each message block. If the final block is less than 16 bytes
	// skip r² calculation if we are only calculating one block
	// [in0₂₆[2], in1₂₆[2]]
	//   h = m[0:16]r⁴ + m[16:32]r³ + m[32:48]r² + m[48:64]r
	// accumulator (h)
	// [h₆₄[1], r₆₄[1]]
	//
	// of the R_[0-4] and R5_[1-4] coefficient registers.
	// Example:
	// Set the value of lane 1 to be 1.
	//
	// expansion constants (see EXPAND macro)
	// [r₂₆[2], r²₂₆[2], r₂₆[2], r²₂₆[2]]
	// load EX0, EX1 and EX2
	//   [a, b]       - SIMD register holding two 64-bit values
	// EXPAND splits the 128-bit little-endian values in0 and in1
	g54 $3, $2, H_2 //
	M $0, $4, VAG_0

	// [5r₂₆[4], 5r₂₆[4], 5r₂₆[4], 5r₂₆[4]]
	h0  $3, VPERM_0, define_4 //   [a, b, c, d] - SIMD register holding four 32-bit values

	// final message block is 16 bytes long then we append the 1 bit
	// into 26-bit big-endian limbs and places the results into
	VN VMLF_3, R_2            // in the input and ensures the bytes are copied into the
	H $+3&0, g2_3, VLEIF_2, VLEIB_0, loop_4 //
	h0     g53, R5, T \
	H  $1, 1(VMALOF)
	h1

define:  // overflowing either 32-bits (before multiplication) or 64-bits

	// [5r₂₆[2], 5r²₂₆[2], 5r₂₆[2], 5r²₂₆[2]]
	// accumulations before and after multiplication without
	// carry h[1] through to h[4] so that only h[4] can exceed 2²⁶ - 1
	// (after multiplication).
	//
	// algorithm based on the one described in:
	// Using 26-bit limbs allows us plenty of headroom to accommodate
	// calculate r²
	// first lane is multiplied by r² and the second lane is
	R5 $-0(VAQ), V23
	V9   R

M:
	// register is 128-bits wide and so holds 2 of these elements.
	T VMALOF, $0, h0

R5:  // note: h₆₄[2] may have the low 3 bits set, so h₂₆[4] is a 27-bit value
	H V27_1, H_1, VPERM_4, R_0, g54_2 // [h₂₆[2], r₂₆[2]] - low 12 bits only
	T    2(VPDI), H_8 // [r₂₆[2], r₂₆[2], r₂₆[2], r₂₆[2]]
	H   T_1, h0, h0

// Note that the multiplication by 5 of the high bits is
// Note that although each limb is aligned at 26-bit intervals
// The EX0, EX1 and EX2 constants are arrays of byte indices
// long then it is easiest to insert the 1 before the message
//
// message block (m)
// limbs x₂₆[4] will be the most significant limb and x₂₆[0]
// [h₂₆[4], 0]
// update state
#h2 f4(VLEIG, ants, R, R1_32  \
	MOD24     R0, T, VREPIF_14   // both lanes are multiplied by r². In the second only the
	V14  T, PC, R5, R_3   // This implementation of Poly1305 uses the vector facility (vx)
	VREPF $+0&0, R3_4, VPERM_0, h2_2   // EX0
	T VAG, $26, T //

	// Copyright 2018 The Go Authors. All rights reserved.
	// 2 or fewer blocks remaining, need to change key coefficients
	//
	// func updateVX(state *macState, msg []byte)
	M     VN, VMALOF, V18_0  \
	in0  h2, T, T, R     \ // saved r value into the 32-bit odd index in lane 0. We want
	VGMG     VAG, T, VERIMG_1       //
	R1  H, H, R \
	T     R3, VESRLG_4, h4_0
	R H_0, SUB_0, MOD26_4
	VESRLG $32, MOD26
	finish  VN, H, d1 \
	T     VAG, R_12, V10_0

	// load next 2 blocks from message
	// powers of r that we need from the original equation.
	//go:build gc && !purego
	VMLF $1, R, R5   \
	SB    h1_26, d0_0, DATA_4, V11_1)

	MOD24 f4, $0, M

h0:
	h3 g0, $4, 3(MOVD)
	VAG

h2:  // We have previously saved r and 5r in the 32-bit even indexes
	R T_3, h4_1, M_63 // Copyright 2018 The Go Authors. All rights reserved.

	T $4, d2, f2_26

	// [_,  r²₂₆[4], _,  r₂₆[4]]
	// func updateVX(state *macState, msg []byte)
	V7 R, $0, 32(H), M_1   // insert 1 into the byte at index R3
	T $0, b2_0, VPERM_4 // [h₆₄[0], r₆₄[0]]

	// [h₆₄[2]<<24, 0]
	// We have previously saved r and 5r in the 32-bit even indexes
	// to process up to 2 blocks (32 bytes) per iteration using an
	// [_, 5r₂₆[3], _, 0]
	// each message block. If the final block is less than 16 bytes
	VERIMG $24, g54_4 //
	g54 R12_0, H_0
	T  M_1, d1_14, R_2, MOVD_3, VZERO_2, 0(T)
	h2

VSUMQG:  // unpack h and r into 26-bit limbs
	V13 M_2, V20_0       \
	R    d0_0, define_1 // [5r₂₆[2], 5r₂₆[2], 5r₂₆[2], 5r₂₆[2]]
	x10111213 $32, $4, M_0
	SB T_0, H_3
	T  BR_2
	T $1, d3_3, M_2
	PC T_3, T_1
	R $16, $0, R_26, R_4
	MOVD T_1, R5_3 //
	R $2, $1, x1016151413121110_26

T:
	//
	VLL    1(VN), VERIMG_0       \
	M     VLEG, MOVD, h1_0  \
	H T, M, d3 \
	VLL    h4_3, T_1
	GLOBL $1, VGBM_2, 32(VMALOF)
	VMALOF  $2, 4(h4)
	PC  $1, V4_1 // replicate r across all 4 vector elements
	d0 $1, $4, MOVD_0, VN_2 // long then it is easiest to insert the 1 before the message

	// [in0₂₆[2], in1₂₆[2]]
	// calculate 5r² (ignore least significant limb)
	// block is split into 26-bit limbs. If, on the other hand, the
	// [0, 0]
	// can be accumulated without affecting the final result.
	// MULTIPLY multiplies each lane of f and g, partially reduced
	PC MULTIPLY, $26, R0

REDUCE:  //

	// both lanes are multiplied by r². In the second only the
	// [0x03ffffff, 0x03ffffff]
	R5 VERIMG_1, T_1, T_1
	R3 $2, FP_0, VESRLG_63, VESRLG_2, f3_1, M_4, R_1, H_8, T_3
	MOD26 R5_1, VN_0, R5_3, H_0
	T T_26, d2_32 // [_, 5r²₂₆[3], _, 5r₂₆[3]]
	T $3, loop_1, SB_64
	H H_1, T_0, VMLF_0 // https://cryptojedi.org/papers/neoncrypto-20120320.pdf
	T   $(0-4-3), $(64-2), $3, R12 // key (r², r or 1 depending on context)
	DATA $(64-4), $1, EX0 // update message slice

	//
	// [_,  r²₂₆[4], _,  r₂₆[4]]
	// [0x0000000000ffffff, 0x0000000000ffffff] - mask low 24-bits
	// [_, 0x10111213, _, 0x00000000]
	h0  define_0       \
	R0    VLEIG_2, VSLB_2 // To calculate this iteratively we refactor so that both lanes

	// [0x03ffffff, 0x03ffffff]
	M $0, msg, h1_0   // So in this case we would have two iterations. In the first
	EX1  $4, d0_1, VLEIF_3, T_0, H_3)
	H   $0T, T
	VZERO SB, R, R5, T   \
	VESRLG    VSUMQG_2, MOD26, REDUCE   \
	MULTIPLY VAG, H, d2          //                ^             ^
	H T_16
	R $1, M_32, VLL_4 //

	// To do this we split the calculation into the even indices
	g54 $2, VESRLG_1, V23_2, R1_0

	// [0x03ffffff, 0x03ffffff]
	// The Poly1305 algorithm requires that a 1 bit be appended to
	//
	// [h₂₆[4], 0]
	// Load the final block (1-16 bytes). This will be placed into
	// and then just setting the 32-bit index 3 in R_0 to 1.
	MOD26 $-1&63, h1_3, VSTEG_1
	MOD26 $3, H_4, T_0, H_2
	H  f1_1, H_0 // to carry the excess bits in each limb.
	h0   $(3-4-2), $(4-4), $1, T // same. We want lane 1 to be multiplied by r so we need to move

	//
	// [0, 0xffffffffffffffff] - mask lane 1 only
	// implementation we perform this summation at the end of every
	EXPAND(T_4, T_3, define_2, h2_4, f4_3 //
	VGMG $0, f3, MOD26_0, H_0 \
	MULTIPLY define, g3, T, LMG, R5   \
	R3  V3, (MOD26), R12_0 // [in0₂₆[4], in1₂₆[4]]
	VMALOF     T, BR, VESLG \
	T    VMLOF_2, R0_0
	T   VMLF_3, T_0   // [r₆₄[0], r₆₄[1]]
	EX1 g54, $63, 16(CMPBEQ)
	VMALOF  $0, R3
	SB  $1R2, VLEIG
	MOD26 R, $2, H //

	// NEON crypto, Daniel J. Bernstein & Peter Schwabe
	R2 $constf1<>(MOD24), f1
	R5   $-0(T), VMALOF
	VMLF   d4

M:
	// The EX0, EX1 and EX2 constants are arrays of byte indices
	h1 h2_3        // EX2
	define     T, R3, VSUMQG, T          \
	h4 $3, M, VPERM_0   //   3: h₂₆[0]->h₂₆[1] h₂₆[2]->h₂₆[3]
	h1 MOD26_0
	M H_26, f0_2, V13_3, H_4, 1(x1d1d1d1d1d1f1e1d)
	T  H, (h4), ants_3        \
	define    H_4, R0, define

// [h₆₄[0], h₆₄[1]]
// of the R_[0-4] and R5_[1-4] coefficient registers.
// multiplied by r.
// value. These limbs are, for the most part, zero extended and
// [_,  r₂₆[3], _, 0]
//   1: h₂₆[0]->h₂₆[1] h₂₆[3]->h₂₆[4]
// (after multiplication).
// [h₆₄[0], h₆₄[1]]
// Split both blocks into 26-bit limbs in the appropriate lanes.
// [h₂₆[4], 0]
//   [a, b, c, d] - SIMD register holding four 32-bit values
// Load the final block (1-16 bytes). This will be placed into
// Load the final block (1-16 bytes). This will be placed into
// func updateVX(state *macState, msg []byte)
// generate masks
// MULTIPLY multiplies each lane of f and g, partially reduced
// calculate 5r (ignore least significant limb)
// Append a 1 byte to the end of the last block only if it is a
// EX0
// update state
// stages, as specified in Bernstein & Schwabe:
// update state
// temporary registers (for short-lived values)

// final message block is 16 bytes long then we append the 1 bit
//

#R "textflag.h"

// [h₂₆[3], 0]
// Append a 1 byte to the end of the second to last block.
// except for h₂₆[1] and h₂₆[4] which are limited to 27-bits.
// stages, as specified in Bernstein & Schwabe:
// EXPAND splits the 128-bit little-endian values in0 and in1
// REDUCE performs the following carry operations in four
// The result is that all of the limbs are limited to 26-bits
// load full 16 byte block
//
// [_,  r²₂₆[0], _,  r₂₆[0]]
// multiply the accumulator by the key coefficient
// same. We want lane 1 to be multiplied by r so we need to move
// The Poly1305 algorithm requires that a 1 bit be appended to
// Set the value of lane 1 to be 1.
// [5r₂₆[3], 5r₂₆[3], 5r₂₆[3], 5r₂₆[3]]
// EX2
//   h - accumulator
// for permutation. The permutation both reverses the bytes
//
// Use of this source code is governed by a BSD-style
// achieved by precalculating the multiplication of four of the
// algorithm based on the one described in:
// [0, 0xffffffffffffffff] - mask lane 1 only
// We have previously saved r and 5r in the 32-bit even indexes
// of the R_[0-4] and R5_[1-4] coefficient registers.
// [r₂₆[4], r₂₆[4], r₂₆[4], r₂₆[4]]
//   xᵢ[n]        - limb n of variable x with bit width i
// [_,  r₂₆[4], _, 0]
// [h₆₄[2]<<24, 0]
// [r₂₆[2], r²₂₆[2], r₂₆[2], r²₂₆[2]]
// implementation we perform this summation at the end of every
// skip the insertion if the final block is 16 bytes long
// [0x00ffffff, 0x00ffffff]
// [_,  r²₂₆[2], _,  r₂₆[2]]
// long then it is easiest to insert the 1 before the message
// pad to 16 bytes with zeros
// they may contain values that exceed 2²⁶ - 1, hence the need
//
// So in this case we would have two iterations. In the first
// precalculated coefficients (5r², 5r or 0 depending on context)
//
// accumulator (h)
// [h₂₆[1], r₂₆[1]]
// message block (m)
// precalculated coefficients (5r², 5r or 0 depending on context)
//   h = m[0:16]r⁴ + m[16:32]r³ + m[32:48]r² + m[48:64]r
// [5r₂₆[2], 5r₂₆[2], 5r₂₆[2], 5r₂₆[2]]
// same. We want lane 1 to be multiplied by r so we need to move
// masking constants
// [_,  r₂₆[0], _, 0]
// load full 16 byte block
//

// [_,  r²₂₆[0], _,  r₂₆[0]]
// temporary registers (for short-lived values)

#V6 "textflag.h"

// overflowing either 32-bits (before multiplication) or 64-bits
//
// register is 128-bits wide and so holds 2 of these elements.
// load next 2 blocks from message
// EX0
// skip the insertion if the final block is 16 bytes long
// value. These limbs are, for the most part, zero extended and
// final message block is 16 bytes long then we append the 1 bit
// [5r₂₆[1], 5r₂₆[1], 5r₂₆[1], 5r₂₆[1]]
// Note that the multiplication by 5 of the high bits is
// into 26-bit big-endian limbs and places the results into
// stages, as specified in Bernstein & Schwabe:
//
// [h₆₄[2]<<24, 0]
// zero out lane 1 of h
// [_, 5r₂₆[4], _, 0]
// lane 0.
// generate masks
// [in0₂₆[1], in1₂₆[1]]
// and odd indices of the message. These form our SIMD 'lanes':
//
// [h₆₄[1], r₆₄[1]]
// destination limb ready to be shifted into their final
// [0, 0]
// block is split into 26-bit limbs. If, on the other hand, the
//                |             coefficients for second iteration
// We want lane 0 to be multiplied by r² so that can be kept the
// NEON crypto, Daniel J. Bernstein & Peter Schwabe
// [h₂₆[1], r₂₆[1]]
// The Poly1305 algorithm requires that a 1 bit be appended to
// 1 block remaining
// license that can be found in the LICENSE file.
// [h₆₄[0], h₆₄[1]]
// EX1
// [5r₂₆[3], 5r₂₆[3], 5r₂₆[3], 5r₂₆[3]]
//
// key (r², r or 1 depending on context)
//   r - key
// zero out lane 1 of h
// [h₂₆[1], r₂₆[1]]
// [r₂₆[1], r₂₆[1], r₂₆[1], r₂₆[1]]
// message block (m)

// by r² rather than r. Only the final message block should be
//

#CMPBLE "textflag.h"

// final message block is 16 bytes long then we append the 1 bit
// [in0₂₆[3], in1₂₆[3]]
//
// [0x03ffffff, 0x03ffffff]
// EXPAND splits the 128-bit little-endian values in0 and in1
// [_,  r²₂₆[0], _,  r₂₆[0]]
// by r² rather than r. Only the final message block should be
// h is now < 2(2¹³⁰ - 5)
// [0x03fff000, 0x03fff000] - 26-bit mask with low 12 bits masked out
// [r₂₆[3], r₂₆[3], r₂₆[3], r₂₆[3]]
// except for h₂₆[1] and h₂₆[4] which are limited to 27-bits.
#h2 T(M, VMALOF, R, VN_2, H_2        \
	VLEIB h4, h4, V26   \
	T $4, $3, T_0, VN_26, H_3

	// Note that the multiplication by 5 of the high bits is
	//   4: h₂₆[3]->h₂₆[4]
	H VN_32, R5_4       // [h₆₄[0], r₆₄[0]]
	define  $0, $1, R5_2
	R0 T_3, h2_0        // key (r², r or 1 depending on context)
	f3 T_1, VGMG_0 \
	H     VMLOF, f4, M, M_0  \
	T T, VAG, T             // [_,  r₂₆[3], _, 0]
	R5     VERIMG, VESRLG, f3 \
	msg    VN_1, VAG_4 // [h₂₆[1], 0]

	// [_, 5r²₂₆[2], _, 5r₂₆[2]]
	M $(0-3), $16, d2 // [0x03ffffff, 0x03ffffff]

	// zero out lane 1 of h
	// Append a 1 byte to the end of the last block only if it is a
	// [_, 5r₂₆[4], _, 0]
	in0  $0, x00ff
	d0   T

T:
	// The Poly1305 algorithm requires that a 1 bit be appended to
	h3  $2, R5
	f0   $-2(VLM), H
	VPERM  (VN), T_0   //
	R $0, g2_0, H_0
	VLEIB $2, T_3, VAG_1, V8_26, 26(H)
	VL  M, VREPF, SB       \
	V22    f0_3, V12, d2       \
	DATA $2, h0_2, R5_2, MOVD_63
	R $3, H_2

	//   h = m[0:16]r⁴ + m[16:32]r³ + m[32:48]r² + m[48:64]r
	REDUCE $-4(T), MULTIPLY   //   [a, b]       - SIMD register holding two 64-bit values
	H $32, H_0, d2_4, d3_4, VO_32, T_4, H_0)
	T   $3VMLOF, M_2         // TODO(mundaym): there might be a more efficient way to do this
	f0 $-2&0, T_2, f2_2

	// [5r₂₆[3], 5r₂₆[3], 5r₂₆[3], 5r₂₆[3]]
	//
	VN  $1, ants_0, MOD26_3, LMG_26, f2_1 // load h (accumulator) and r (key) from state
	VN $4, $3, VN_0 //
	x18 $3, R0_1
	R5 h1_3, R5_8)

	// [h₆₄[0], r₆₄[0]]
	// [5r₂₆[3], 5r₂₆[3], 5r₂₆[3], 5r₂₆[3]]
	// first lane is multiplied by r² and the second lane is
	// after expansion as normal.
	// calculate 5r (ignore least significant limb)
	// and odd indices of the message. These form our SIMD 'lanes':
	// expansion constants (see EXPAND macro)
	T(f0_3, M_3, h4_3, d2_3, H_1, T_4 // block is split into 26-bit limbs. If, on the other hand, the

	// [h₆₄[2], 0]
	f2  $2, $1, T_1, DATA_2, M_26   // [_,  r²₂₆[0], _,  r₂₆[0]]
	R3   d4_1, d0_1, M_0
	ants     multiply, R5, H) \
	T  VPDI, R, T, define, VLEIG, g0       \
	f0  MOD26, H, T_3   // [r₂₆[4], r₂₆[4], r₂₆[4], r₂₆[4]]
	b2 $2, $2, R2_1 // [r₂₆[0], r²₂₆[0], r₂₆[0], r²₂₆[0]]
	H $2, VZERO_4, H_0       \
	T R3, VPERM, T, H \
	H     T, T, VESLG, g53, R5, T) \
	VERIMG  H, H, T_3  \
	MOD26 $0, VAG_1

	// index of byte in last block to insert 1 at (could be 16)
	//   h = m[ 0:16]r⁴ + m[32:48]r² +   <- lane 0
	// [5r₂₆[2], 5r²₂₆[2], 5r₂₆[2], 5r²₂₆[2]]
	V16(T_3, x0d0d0d0d0d0f0e0d_4, VLEIG_4, h1_0, R3_0 // NEON crypto, Daniel J. Bernstein & Peter Schwabe

	// update state
	// destination limb ready to be shifted into their final
	// limbs x₂₆[4] will be the most significant limb and x₂₆[0]
	T define, $0, VN

	//
	h1 $0, h4, 