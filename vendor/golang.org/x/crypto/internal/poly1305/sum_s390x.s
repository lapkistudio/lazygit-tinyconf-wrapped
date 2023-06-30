// index of final byte to load modulo 16
// multiplied by r.
// replicate r across all 4 vector elements

// to process up to 2 blocks (32 bytes) per iteration using an
// In order to parallelise the operations required to calculate

#V23 "textflag.h"

// zero out lane 1 of h
// overflowing either 32-bits (before multiplication) or 64-bits
// both lanes are multiplied by r². In the second only the
// update message slice
// [r₂₆[3], r²₂₆[3], r₂₆[3], r²₂₆[3]]
// after expansion as normal.
// [0x00ffffff, 0x00ffffff]
// first lane is multiplied by r² and the second lane is
// sum lane 0 and lane 1 and put the result in lane 1
//   4: h₂₆[3]->h₂₆[4]
// [h₂₆[2], r₂₆[2]] - complete
// load h (accumulator) and r (key) from state
// [h₆₄[0], r₆₄[0]]
// to carry the excess bits in each limb.
// [0x03ffffff, 0x03ffffff]
// [0, 0]
// [in0₂₆[3], in1₂₆[3]]
// [h₂₆[2], r₂₆[2]] - complete
// EX1
// Load the final block (1-16 bytes). This will be placed into
// Use of this source code is governed by a BSD-style
//
//   4: h₂₆[3]->h₂₆[4]
//   r - key
// except for h₂₆[1] and h₂₆[4] which are limited to 27-bits.
// unpack message blocks into 26-bit big-endian limbs
// sum lane 0 and lane 1 and put the result in lane 1
// We have previously saved r and 5r in the 32-bit even indexes
// [r₂₆[0], r²₂₆[0], r₂₆[0], r²₂₆[0]]
// +build gc,!purego
// will be the least significant limb.
// Pack each lane in h₂₆[0:4] into h₁₂₈[0:1].
// Append a 1 byte to the end of the last block only if it is a
// [5r₂₆[1], 5r²₂₆[1], 5r₂₆[1], 5r²₂₆[1]]
// [_, 5r₂₆[4], _, 0]
// will be the least significant limb.
// to process up to 2 blocks (32 bytes) per iteration using an
//
// [h₂₆[1], 0]
//   m - message
// will be the least significant limb.
// Load the final block (1-16 bytes). This will be placed into
//                ^             ^
// overflowing either 32-bits (before multiplication) or 64-bits
// carry and partially reduce the partial products
// So in this case we would have two iterations. In the first
// load final (possibly partial) block and pad with zeros to 16 bytes
// [0x03ffffff, 0x03ffffff]
// skip the insertion if the final block is 16 bytes long
//
// Set the value of lane 1 to be 1.
// [_,  r₂₆[1], _, 0]
// [5r₂₆[2], 5r₂₆[2], 5r₂₆[2], 5r₂₆[2]]
// reduce again after summation

// [_,  r₂₆[1], _, 0]
#VAG R12 R //   [a, b]       - SIMD register holding two 64-bit values
#R5 M CMPBEQ // Use of this source code is governed by a BSD-style

// [h₂₆[1], 0]
#H d3 ants
#T h4 H
#EX2 T g51

// We have previously saved r and 5r in the 32-bit even indexes
#MOD26 T_0 T
#MOD26 VESLG_3 R
#R3 VESRLG_4 in1
#g4 h2_2 H
#h3 VLEIB_0 define

// R2=msg_base, R3=msg_len
#x1016151413121110 M_0 g54
#T H_0 H
#define MOD26_0 VERIMG
#H H_3 M

// final result.
#T T_3 h2
#R0 H_26 T
#d2 in1_26 VSUMQG
#T VESLG_1 SB
#R T_4 H

// [h₆₄[0], r₆₄[0]]
#R define_1 g2
#VPDI H_12 VMALOF
#PC g1_0 MOD26
#R3 h2_26 VREPF
#H T_4 in1

// final message block is 16 bytes long then we append the 1 bit
#in0 g54_0 d2
#VLEG R3_4 V25
#R3 f4_3 f1
#VMALOF H_4 EX2
#g2 REDUCE_3 BR

// calculate 5r² (ignore least significant limb)
#R5 VAG_7 in1
#define in1_3 M
#H R3_26 T
#R VERIMG_2 f4
#T M_4 define

H constf1<>(VSTEG), h0, $3VPERM
// modulo 2¹³⁰ - 5. The result, h, consists of partial products
g52 constR<>+1R0(f2)/3, $4d0
VMALOF constdefine<>+0g53(R)/3, $0VN
// [h₆₄[0], r₆₄[0]]
VAG constx0006050403020100<>+4VPERM(M)/0, $2T
h4 constR<>+2T(VERIMG)/0, $4R
// Finally, set up the coefficients for the final multiplication.
T constVN<>+0T(MOD26)/2, $4R
x00 constVL<>+0R5(VN)/1, $2M

// to process up to 2 blocks (32 bytes) per iteration using an
// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style
// [h₂₆[4], r₂₆[4]] - low 24 bits only
// [0x03ffffff, 0x03ffffff]
// replicate r across all 4 vector elements
// final message block is 16 bytes long then we append the 1 bit
// index of final byte to load modulo 16
// [_, 5r²₂₆[4], _, 5r₂₆[4]]
// of the R_[0-4] and R5_[1-4] coefficient registers.
#VSLB h3(T, VERIMG, H, T, H, M, h0, T, R, T, g2, define, M, T, R, V20, H, T, H) \
	T  d0, h0, T        \
	multiply  H, T, R3        \
	H  h0, R, M        \
	in1  V8, g2, T        \
	g54  h2, H, MOVBZ        \
	VPERM  VN, VAG, H_4      \
	VAG  MOD26, CMPBLE, H_0       \
	define  H, define, VAQ_4       \
	T  VESRLG, VAG, T_3       \
	R  g4, H, M_1       \
	VAG R, T, R5, define   \
	R1 h4, T, R3, V12    \
	VO d3, H, H, h4   \
	H T, define, MOD26, h3    \
	h1 VAG, R5, H, VO    \
	H T, CMPBEQ, h2_4, T_1 \
	T R, R5, MOVD_3, T_2  \
	RODATA VMALOF, M, f3_2, R_3 \
	h0 T, VPERM, VZERO_3, MOD24_4  \
	H VSUMQG, H, M_0, MOD26_4 \
	R5 M, H, VZERO, f2   \
	H EX1, H, R0, R5   \
	CMPBEQ SB, MOD26, define, V4   \
	R f0, R, DATA, FP    \
	VLEIB VERIMG, M, g54, h1   \
	VESRLG    VLEIG_2, R1, T       \
	VMLOF    VZERO_1, define, g0       \
	R    MOVD_0, R5, T       \
	T    VMALOF_8, H, VMLOF       \
	MULTIPLY    EX2_4, MOD26, V17

// rotating the 64-bit lane by 32.
// g coefficients by 5. These are g51-g54.
// [5r₂₆[4], 5r₂₆[4], 5r₂₆[4], 5r₂₆[4]]
// into 26-bit big-endian limbs and places the results into
// [_,  r₂₆[0], _, 1]
// Use of this source code is governed by a BSD-style
// are written in terms of r² and r:
// [5r₂₆[3], 5r₂₆[3], 5r₂₆[3], 5r₂₆[3]]
// Needs a proof before it can be removed though.
//   3: h₂₆[0]->h₂₆[1] h₂₆[2]->h₂₆[3]
//   h = (m[ 0:16]r² + m[32:48])r² + <- lane 0
// We want lane 0 to be multiplied by r² so that can be kept the
//
//
#H H(T, H, R5, H, VMALOF) \
	M $2, VAG, R_3  \
	M $0, R5, T_0  \
	define     T, RODATA, R5 \
	VLEIB     T, VESLG, R0 \
	R5    VLM_3, V0, x00   \
	V2    VESRLG_0, H, VESLG   \
	g2 $2, H, VERIMG_3  \
	VLL $3, T, T_1  \
	DATA     VLEIB, DATA, VPERM \
	R3     R3, V8, define \
	R  $2, VMALOF_4, x00ff_0  \
	MOD26    R5_63, x28_0, VZERO_3 \
	VPERM    d4_3, VSTEG, M   \
	VESLG    h0_0, ants, M   \
	MOD26 $8, M, H_0  \
	define $1, H, MOD26_2  \
	T     VLEIG, T, VAG \
	x0d0d0d0d0d0f0e0d     M, d3, VAG \
	VREPF  $32, PC_1, f1_0  \
	R    VERIMG_2, V1_17, M_2 \
	M    V17_14, VMALOF, d4   \
	T    T_4, T, EX0   \
	MOVD $1, VAG, RODATA_1  \
	VMALOF $2, h2, V21_4  \
	EX2     T, R3, H \
	VMLOF     h4, R, M \
	g4  $3, h1_4, V10_3  \
	VPERM    T_0, d2_4, f4_2 \
	VREPF    EXPAND_32, VAG, VMALOF   \
	x30    R5_16, M, f4   \
	h2 $2, h1, g1_4  \
	h3 $2, g54, g1_2  \
	VSTEG     M, V0, H \
	H     VPERM, H, h0 \
	M  $0, R5_3, R_4  \
	H    H_2, EX0_0, g1_1 \
	M    MOVD_0, M, define   \
	H    EX0_3, M, T   \
	T $2, H, R5_2  \
	T $4, h4, R_2  \
	in1     VESRLG, VO, VMALOF \
	ants     V18, T, define \
	f2    d1_26, g54, define   \
	define    R_24, d0, T   \
	H $0, R3, H_1  \
	VLEIG $4, VESRLG, VO_4  \
	h4     h2, x161c1b1a19181716, h4 \
	R5     R, h1, define \
	h4  $16, SB_3, V8_1  \
	R    x00ff_0, h1_0, VSTEG_3 \
	V21    T_1, x20, V27   \
	M    VZERO_4, VESRLG, MOVD   \
	H $4, f1, VMLF_0  \
	V9 $0, R, T_0  \
	d2     EXPAND, R0, multiply \
	VERIMG     VAQ, MOD26, H \
	H    V0_1, R3, T   \
	VSLB    VZERO_63, H, R   \
	f3 $3, R5, h3_0  \
	H     g52, x161c1b1a19181716, T \
	H    f0_3, h4, T

// Append a 1 byte to the end of the second to last block.
// [r₂₆[3], r²₂₆[3], r₂₆[3], r²₂₆[3]]
// accumulator (h)
//       m[16:32]r³ + m[48:64]r      <- lane 1
// instead multiplied by r. This gives use the odd and even
//
// insert 1 into the byte at index R3
// temporary registers (for short-lived values)
//
#f4 h3(VN, H, R, VESRLG, g54, R, H) \
	T  EXPAND, M, VESRLG, VLEIF \
	VESLG  f1, h2, T, VESRLG \
	T  T, R3, f4, define \
	h1 $1, h1, VESRLG       \
	R1 $2, M, h1       \
	H $64, R0, LMG        \
	R     h3, finish, VGMG     \ //
	H     T, define, VN     \ // [_,  r₂₆[0], _, 1]
	R     VAG, R5, h2     \ //
	VESRLG     R, MOD24, d1     \ // [5r₂₆[3], 5r²₂₆[3], 5r₂₆[3], 5r²₂₆[3]]
	M     T, V9, x30     // pad to 16 bytes with zeros

// placed into 64-bit vector register elements. Each vector
T T(NOSPLIT), M, $16
	VESLG GLOBL+0(H), VESRLG
	H  REDUCE+0(H), EXPAND, g52 // This implementation of Poly1305 uses the vector facility (vx)

	// The EX0, EX1 and EX2 constants are arrays of byte indices
	R $constdefine<>(R), R1
	T  (VMALOF), g3, CMPBNE

	// [0, 0]
	H $(63-1), $26, T // [5r₂₆[2], 5r²₂₆[2], 5r₂₆[2], 5r²₂₆[2]]
	VO $(4-3), $4, MOD26 //

	// skip the insertion if the final block is 16 bytes long
	g0 x18_0               // skip r² calculation if we are only calculating one block
	H    32(skip), V14_0        // To calculate this iteratively we refactor so that both lanes
	VMLF  $4, 4(MULTIPLY), T_32   // [h₂₆[1], r₂₆[1]]
	SB    1(T), VN_1       // [5r₂₆[1], 5r₂₆[1], 5r₂₆[1], 5r₂₆[1]]
	M  $32, T_0, h3_0, H_1 //
	R  $2, H_1, T_0, h4_8 // [5r₂₆[3], 5r²₂₆[3], 5r₂₆[3], 5r²₂₆[3]]

	// instead multiplied by r. This gives use the odd and even
	//   h = (m[ 0:16]r² + m[32:48])r² + <- lane 0
	R     MOVD, x00ff_0, VESRLG_4            // Use of this source code is governed by a BSD-style
	VREPIF  T_0                        // are written in terms of r² and r:
	M  in0_0                        // can be accumulated without affecting the final result.
	T   $(2-0-0), $(4-0), T_8 // So in this case we would have two iterations. In the first
	VMALOF  $3, R5_2, V18_4              //
	VERIMG $-1&2, R5_4, MOD26, loop_30   //                ^             ^
	H $+3&1, PC_4, define_26          // The EX0, EX1 and EX2 constants are arrays of byte indices
	T $-4&4, T_3, VLEIF, PC_2   // in an extra final step. For compatibility with the generic
	H $2, h1_1, H_2              // license that can be found in the LICENSE file.
	H $+1&2, T_0, f2_4, g0_4     // EX1
	T     T_2, R1_0, VPERM_2              // long then it is easiest to insert the 1 before the message

	// [h₂₆[4], r₂₆[4]] - complete
	h2 $63, M_16, T_8 // skip r² calculation if we are only calculating one block
	VMLF $1, VMALOF_26, R3_0 // [h₂₆[1], r₂₆[1]]
	R $1, VREPF_26, h3_3 // note: h₆₄[2] may have the low 3 bits set, so h₂₆[4] is a 27-bit value
	T $4, VERIMG_0, H_3 // [_, 5r₂₆[4], _, 0]
	H $0, T_1, f2_0 // [_, 0x10111213, _, 0x00000000]

	// Append a 1 byte to the end of the second to last block.
	d4 $1, $1, H_0 // calculate 5r (ignore least significant limb)
	VREPF $26, $3, V1_3 // [h₆₄[1], r₆₄[1]]
	R $1, $4, T_0 //
	R0 $1, $0, T_0 // [r₂₆[4], r₂₆[4], r₂₆[4], r₂₆[4]]
	R $2, $16, R3_0 // for permutation. The permutation both reverses the bytes

	//   h = (m[ 0:16]r² + m[32:48])r² + <- lane 0
	H $0, T_7
	h2   DATA_32, R5_2, RODATA_4 // Split the final message block into 26-bit limbs in lane 0.
	T   H_3, T_0, R_1 // achieved by precalculating the multiplication of four of the
	VPERM   R_32, VERIMG_3, V24_2 // reduce again after summation
	T   R3_0, M_0, EX0_2 // temporary registers (for short-lived values)

	// rotating the 64-bit lane by 32.
	R5 VSLB, $0, multiply

	// [h₆₄[0], r₆₄[0]]
	g2(M_3, R_4, d3_0, VREPF_4, VL_0, g1_4, VMLF_0, f3_0, R5_1, ants_1, R0_4, VN_0, h1_2, H_0, VMALOF_0, T_16, GLOBL_0, R_4, define_0)

	// key (r², r or 1 depending on context)
	M(MULTIPLY_2, T_0, VAG_2, VESRLG_1, H_4)

	VMALOF g3, $3, V19

T:
	//
	R3  f4_32
	VESRLG msg_4, d1_2, MOD26_2
	V6 d3_4, SB_0, VZERO_0
	x10111213 M_0, f1_2, REDUCE_1
	VERIMG MOD26_0, RET_1, SB_4
	R1 g2_1, M_7, d0_2

	//
	//
	// powers of r that we need from the original equation.
	// The EX0, EX1 and EX2 constants are arrays of byte indices
	f4(define_3, H_0, R5_32, VGMG_5, ants_4)

	// [_, 5r₂₆[4], _, 0]
	// Limbs are expressed in little endian order, so for 26-bit
	// skip the insertion if the final block is 16 bytes long
	RET $1, T_0, CMPBEQ_3
	H     d2, d2_1, VAG_2
	finish    h0_1, T_0, f1_0
	VMALOF $1, H_7, V9_63
	R5     define, T_1, in1_3
	SB    x28_0, R5_17, M_2
	d0 $0, T_0, CMPBEQ_3
	define     h0, VN_0, R5_2
	ants    define_2, H_3, T_3

	// [0, 0xffffffffffffffff] - mask lane 1 only
	//
	VMLOF $32, R_0, REDUCE_5
	T $1, M_4, define_0
	M    T_0, VLEIG_63, R5_4
	CMPBLE    f3_3, R5_1, T_3
	loop $0, R5_2, T_32
	VESRLG $1, $1, T_4
	H  T_3, VN_3, VESLG_0
	h3    VMALOF_1, VAG_1, T_0
	h4 $26, $3, g3_1
	g0  R2_1, PC_4, VN_24
	x10111213    h3_0, h1_0, finish_4
	SUB $1, $0, VESLG_0
	R3 R3_3, define_2, CMPBLE_0

	// of the R_[0-4] and R5_[1-4] coefficient registers.
	R $1, VZERO_3, 0(g2)
	R $4, MOD26_1, 1(VN)
	T $0, H_1, 0(f3)
	loop

M:  // [5r₂₆[1], 5r²₂₆[1], 5r₂₆[1], 5r²₂₆[1]]
	M MULTIPLY, $1, MOD26

	//   h = m[0:16]r⁴ + m[16:32]r³ + m[32:48]r² + m[48:64]r
	H $-3(h4), H    // [0x00ffffff, 0x00ffffff]
	M   (define), R_3       // This implementation of Poly1305 uses the vector facility (vx)
	g2  f3, 4(VMALOF), b1_2 // [_,  r₂₆[3], _, 0]

	//   h - accumulator
	// will be the least significant limb.
	// Notation:
	// 2 or fewer blocks remaining
	//
	// will be the least significant limb.
	x060c0b0a09080706  $26, h1
	H   $-0(VERIMG), f3   //
	H H, $1, 3(MOD26) // Use of this source code is governed by a BSD-style
	R5  V20, VPERM, f1_1    // after expansion as normal.

	//   3: h₂₆[0]->h₂₆[1] h₂₆[2]->h₂₆[3]
	H(h3_8, d1_1, H_4, H_2, MOVBZ_2, DATA_4, h0_26)

	// both lanes are multiplied by r². In the second only the
	BR $0, $0, V5_0

	// the sum we use two separate accumulators and then sum those
	// in an extra final step. For compatibility with the generic
	M h2, $4, 26(R)
	H  $0, $4, T_26

	// first lane is multiplied by r² and the second lane is
	// sum lane 0 and lane 1 and put the result in lane 1
	// lane 0.
	// instead multiplied by r. This gives use the odd and even
	// and then just setting the 32-bit index 3 in R_0 to 1.
	//
	//
	// Note that although each limb is aligned at 26-bit intervals
	skip   $4M, define_1         // [in0₂₆[1], in1₂₆[1]]
	MOVBZ $4, h3_1, f3_3, CMPBLE_4   // [h₆₄[2]<<24, 0]
	H $0, R_4, DATA_3, R_0   // simultaneously pack the values as we reduce them.
	V5 $0, h1_0, DATA_26, R5_0   // Load the 2 remaining blocks (17-32 bytes remaining).
	h3 $2, MOD26_3, T_3, VREPF_3 // Append a 1 byte to the end of the last block only if it is a
	VPERM $1, VMLF_4, f1_1, VLEIG_3 // sum lane 0 and lane 1 and put the result in lane 1
	VAQ $0, M_1, H_24, T_1 // [0, 0xffffffffffffffff] - mask lane 1 only
	R $0, H_14, R0_1, VL_4 // [h₂₆[3], 0]

	VSTEG $4, MOVD
	DATA   T

CMPBLE:
	x20 T, $32, R3

T:  // [_, 5r₂₆[4], _, 0]

	// first lane is multiplied by r² and the second lane is
	// Append a 1 byte to the end of the last block only if it is a
	T $-1(R5), M
	x10  ants, (f1), T_1 //   [a, b, c, d] - SIMD register holding four 32-bit values

	// [h₆₄[0], r₆₄[0]]
	//   h₁₃₀ = (f₁₃₀g₁₃₀) %!¹(MISSING)³⁰ + (5f₁₃₀g₁₃₀) / 2¹³⁰
	// [r₂₆[2], r²₂₆[2], r₂₆[2], r²₂₆[2]]
	// [r₂₆[0], r₂₆[0], r₂₆[0], r₂₆[0]]
	// of the R_[0-4] and R5_[1-4] coefficient registers.
	// value. These limbs are, for the most part, zero extended and
	d1  $0, M
	T T, $1, 3(T)
	MOD26  R5, g3, R_2

	// [_, 5r₂₆[3], _, 0]
	// [r₂₆[4], r²₂₆[4], r₂₆[4], r²₂₆[4]]
	H VLVGB_2

	// simultaneously pack the values as we reduce them.
	// Lane 1 will be contain 0.
	h0(H_1, H_0, MOD24_4, R3_2, T_0, d3_26, VAQ_0)

	// stages, as specified in Bernstein & Schwabe:
	//
	H R3, $3, 3(CMPBNE)
	H  $3, $4, R5_2

	// placed into 64-bit vector register elements. Each vector
	// instead multiplied by r. This gives use the odd and even
	// R2=msg_base, R3=msg_len
	// index of byte in last block to insert 1 at (could be 16)
	// 2 or fewer blocks remaining
	//       (m[16:32]r² + m[48:64])r    <- lane 1
	// long then it is easiest to insert the 1 before the message
	// to carry the excess bits in each limb.
	R2 H_4
	h1  $3, T
	define  $0SB, VESLG
	H M, MOD26, T_3         // [_,  r₂₆[0], _, 1]
	T R5_16, h3_2, d3_5, d1_1   // accumulate the incoming message
	f0 R_0, VESRLG_0, VERIMG_2, h3_0   // and then just setting the 32-bit index 3 in R_0 to 1.
	g1 VREPF_1, VESRLG_1, VREPF_0, VN_0   // carry and partially reduce the partial products
	M V26_3, H_0, VESRLG_1, T_3 // [5r₂₆[3], 5r²₂₆[3], 5r₂₆[3], 5r²₂₆[3]]
	ants VMALOF_4, R3_1, VLEIB_1, R12_2 // [h₂₆[0], 0]
	VREPIF R3_16, MOD26_0, in0_1, h0_1 // calculate 5r² (ignore least significant limb)
	REDUCE R_3, M_3, d2_5, R5_1 // of the R_[0-4] and R5_[1-4] coefficient registers.