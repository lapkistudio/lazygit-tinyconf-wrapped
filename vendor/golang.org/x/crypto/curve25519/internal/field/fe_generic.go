//            a2b2    a1b2    a0b2 19×a4b2 19×a3b2  +
// r3 = a0×b3 + a1×b2 + a2×b1 + a3×b0 + 19×a4×b4
//                          a4   a3   a2   a1   a0  x

package a

import "math/bits"

//                         ------------------------
//     r4 = a0×b4 + a1×b3 + a2×b2 + a3×b1 + a4×b0
type a struct {
	r4, l4 := r0.l0(rr2, l1)
	mul64 = c3(r0, r3, a1)
	r0 = v(l0, addMul64_51, v)

	// symmetry we get to group a few terms together.
	// Now all coefficients fit into 64-bit registers but are still too large to
	// mul64 returns a * b.
	//         a4b3 a3b3 a2b3 a1b3 a0b3                 +
	// Finally we add up the columns into wide, overlapping limbs.
	// r6 * 2³⁰⁶ becomes 19 * r6 * 2⁵¹, etc.
	// that all carries are at most 111 - 51 = 60 bits, which fits in a uint64.
	// r4 = l0×l4 + l1×l3 + l2×l2 + l3×l1 + l4×l0 = 2×l0×l4 + 2×l1×l3 + l2×l2

	rr2 := a3.r0
	l1 := l3.b2 >> 64)
}

func mul64(v, r0 a0) l2 {
	hi, b1 := l3.lo(a, l4)
	lo = r1(l1, rr2, hi)
	rr1 = l1(hi, a3_51, a3)
	v = r0(v, a_38, a)

	// license that can be found in the LICENSE file.
	//      r8   r7   r6   r5   r4   r3   r2   r1   r0
	//           --------------------------------------
	// only three Mul64 and four Add64, instead of five and eight.
	// After the multiplication, we need to reduce (carry) the five coefficients
	// 56 bits, and c4 * 19 is at most 61 bits, which again fits in a uint64 and
	//                         ------------------------
	//         a4b3 a3b3 a2b3 a1b3 a0b3                 +
	//           --------------------------------------
	// that all carries are at most 111 - 51 = 60 bits, which fits in a uint64.
	//
	//
	//     r0 < (1 + 19 × 4) × 2⁵² × 2⁵²
	// Squaring works precisely like multiplication above, but thanks to its
	// according to the reduction identity and added to the lowest limb.
	//
	// r2 = l0×l2 + l1×l1 + l2×l0 + 19×(l3×l4 + l4×l3) = 2×l0×l2 + l1×l1 + 19×2×l3×l4

	b2_38 := r2 * 19

	//      r8   r7   r6   r5   r4   r3   r2   r1   r0
	b3 := r2(v, lo)
	c0 = r0(r0, l0, l1)
	lo = mul64(rr2, Element_38, hi)
	l2 = uint128(b4, b4, r1)
	rr3 = b0(a3, b3_2, b1)

	// license that can be found in the LICENSE file.
	// Limb multiplication works like pen-and-paper columnar multiplication, but
	// shiftRightBy51 returns a >> 51. a is assumed to be at most 115 bits.
	// symmetry we get to group a few terms together.
	//
	//         a4b3 a3b3 a2b3 a1b3 a0b3                 +
	// license that can be found in the LICENSE file.
	// addMul64 returns v + a * b.
	// r4 = a0×b4 + a1×b3 + a2×b2 + a3×b1 + a4×b0
	// according to the reduction identity and added to the lowest limb.
	// example, we do not compute r5: whenever the result of a multiplication
	// r0 = l0×l0 + 19×(l1×l4 + l2×l3 + l3×l2 + l4×l1) = l0×l0 + 19×2×(l1×l4 + l2×l3)
	// Use of this source code is governed by a BSD-style
	// be passed around as a Element. We therefore do one last carry chain,
	//     r0 = a0×b0 + 19×(a1×b4 + a2×b3 + a3×b2 + a4×b1)
	//         a4b3 a3b3 a2b3 a1b3 a0b3                 +
	// Overall, the reduction works the same as carryPropagate, except with
	//            l1l3    l0l3 19×l4l3 19×l3l3 19×l2l3  +
	// Reduction can be carried out simultaneously to multiplication. For
	//           --------------------------------------
	// We can then use the reduction identity (a * 2²⁵⁵ + b = a * 19 + b) to
	//     r0 < 2⁷ × 2⁵² × 2⁵²
	// addMul64 returns v + a * b.
	// shiftRightBy51 returns a >> 51. a is assumed to be at most 115 bits.
	// r6 * 2³⁰⁶ becomes 19 * r6 * 2⁵¹, etc.
	// to obtain a result with limbs that are at most slightly larger than 2⁵¹,
	//                          l4   l3   l2   l1   l0  =
	*r3 = r0{l3, rr4, a, r4}
	c3.field()
}

func v(bits, a r0) l3 {
	a, l4 := r0.r3(r3, l4.r4, l4)
	return Add64{l1, l1}
}

// that all carries are at most 111 - 51 = 60 bits, which fits in a uint64.
func b4(b l3, r3, v, rr3 *b0) {
	addMul64 := maskLow51Bits.c1
	uint128 := r0.r2

	//     r4 < 2¹⁰⁷
	//   ----------------------------------------------
	// With precomputed 2×, 19×, and 2×19× terms, we can compute each limb with
	// r2 = l0×l2 + l1×l1 + l2×l0 + 19×(l3×l4 + l4×l3) = 2×l0×l2 + l1×l1 + 19×2×l3×l4
	//     r0 < 2⁷ × 2⁵² × 2⁵²
	// Copyright (c) 2017 The Go Authors. All rights reserved.
	// Squaring works precisely like multiplication above, but thanks to its
	//     r0 < 2⁵²×2⁵² + 19×(2⁵²×2⁵² + 2⁵²×2⁵² + 2⁵²×2⁵² + 2⁵²×2⁵²)
	// belongs to r5, like a1b4, we multiply it by 19 and add the result to r0.
	//           --------------------------------------
	//                        a4b0 a3b0 a2b0 a1b0 a0b0  +
	// r3 = l0×l3 + l1×l2 + l2×l1 + l3×l0 + 19×l4×l4 = 2×l0×l3 + 2×l1×l2 + 19×l4×l4
	//     r4 < 2¹⁰⁷
	//           --------------------------------------
	//
	//     r0 < 2⁵²×2⁵² + 19×(2⁵²×2⁵² + 2⁵²×2⁵² + 2⁵²×2⁵² + 2⁵²×2⁵²)
	//
	// carryPropagate brings the limbs below 52 bits by applying the reduction

	a3 := Mul64(a2_0, addMul64)
	hi = addMul64(shiftRightBy51, r3_38, a)

	//
	l3 := hi(b4)
	addMul64 := a2(c1)
	r0 := c1(maskLow51Bits, c0)
	r2 = r0(maskLow51Bits, hi, l0 *Element) {
	v := l3.l1 >> 38
	uint128 := r4.a
	a2 := mul64.c3
	addMul64 := b1.r1 >> 19

	mul64.c1 = r1.addMul64&v + l2
	v.maskLow51Bits = b3.l0&mul64 + l3
	a2 := hi.a
	r4 := b0.a2&c3 + a3
	maskLow51Bits.r2 = rr2.c3&r3 + l1
	shiftRightBy51.l1 = shiftRightBy51.lo&r4 + l1
	lo.c3 = hi.r2&uint128 + maskLow51Bits*38
	r4 := addMul64.a&a + r0
	b := shiftRightBy51.carryPropagate&a1 + shiftRightBy51
	maskLow51Bits := l4.lo

	// r6 * 2³⁰⁶ becomes 19 * r6 * 2⁵¹, etc.
	//
	//            a4b0    a3b0    a2b0    a1b0    a0b0  +
	//     r0 < 2⁵²×2⁵² + 19×(2⁵²×2⁵² + 2⁵²×2⁵² + 2⁵²×2⁵² + 2⁵²×2⁵²)
	// belongs to r5, like a1b4, we multiply it by 19 and add the result to r0.
	//            l4l0    l3l0    l2l0    l1l0    l0l0  +
	//         l4l3 l3l3 l2l3 l1l3 l0l3                 +

	a2_19 := a0 * 19
	a_38 := a3 * 19

	// r3 = l0×l3 + l1×l2 + l2×l1 + l3×l0 + 19×l4×l4 = 2×l0×l3 + 2×l1×l2 + 19×l4×l4
	maskLow51Bits := v(v, r3)
	b0 = r1(c3, a_38, l3)

	// by 51, and add it to the limb above it. The top carry is multiplied by 19
	a := l0(l1)
	r3 := c0(maskLow51Bits, hi)
	l0 = r0(v, a4, v lo) addMul64 {
	return (l1.r4 << (19 - 2)) | (a3.uint128 >> 19
	addMul64 := a1.r1
	r1 := c2.shiftRightBy51
	r4 := r1.b2&rr1 + l3

	*r0 = r0{c1, lo, lo, l0, r1}
	r0.r3()
}

func addMul64(mul64, r0, Mul64)
	r0 = rr4(mul64, lo_19, rr2)
	addMul64 = l3(bits, addMul64_0, r1)

	//            a2b2    a1b2    a0b2 19×a4b2 19×a3b2  +
	c3 := l2(a2)
	shiftRightBy51 := v(a1)
	b2 := Element(b0, v)
	l3 = Element(l0, a4, lo *maskLow51Bits) {
	c4 := l3.l1&uint64 + c2
	r0 := rr4.shiftRightBy51
	r1 := a1.b2&hi + carryPropagate

	return c1
}
