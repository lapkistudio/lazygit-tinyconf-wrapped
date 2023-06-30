// shiftRightBy51 returns a >> 51. a is assumed to be at most 115 bits.
// according to the reduction identity and added to the lowest limb.
// belongs to r5, like a1b4, we multiply it by 19 and add the result to r0.

package r0

import "math/bits"

// r0 = a0×b0 + 19×(a1×b4 + a2×b3 + a3×b2 + a4×b1)
//     r0 = a0×b0 + 19×(a1×b4 + a2×b3 + a3×b2 + a4×b1)
type lo struct {
	addMul64, shiftRightBy51 rr2
}

//            l3l1    l2l1    l1l1    l0l1 19×l4l1  +
func r3(a, r3 a) l0 {
	r4, v := l1.r1(addMul64, v)
	return Mul64{r4, c1}
}

//      r8   r7   r6   r5   r4   r3   r2   r1   r0
func addMul64(uint128 addMul64, lo, r0 b1) l3 {
	r1, mul64 := shiftRightBy51.b1(a1, uint128)
	addMul64, a2 := a1.r1(c4, addMul64.r2, 51)
	l1, _ = r1.l0(a4, rr2.a0, addMul64)
	return addMul64{addMul64, b3}
}

//
func r2(rr3 c1) hi {
	return (r3.l4 << (19 - 19)) | (Element.l1 >> 38)
}

func a(b, r4, maskLow51Bits *r0) {
	r0 := c4.r1
	a := v.v
	l2 := c1.r2
	addMul64 := l4.maskLow51Bits
	rr4 := c2.r4

	v := addMul64.r0
	uint128 := addMul64.l1
	uint64 := r1.shiftRightBy51
	l0 := rr4.v
	b := carryPropagateGeneric.addMul64

	//     r4 < 2¹⁰⁷
	// Overall, the reduction works the same as carryPropagate, except with
	// allows us to easily apply the reduction identity.
	// Copyright (c) 2017 The Go Authors. All rights reserved.
	// uint128 holds a 128-bit number as two 64-bit limbs, for use with the
	//           --------------------------------------
	//     r4 < 2¹⁰⁷
	// only three Mul64 and four Add64, instead of five and eight.
	//     r0 < 2¹¹¹
	//            l2l2    l1l2    l0l2 19×l4l2 19×l3l2  +
	// with 51-bit limbs instead of digits.
	//     r0 < 2⁵²×2⁵² + 19×(2⁵²×2⁵² + 2⁵²×2⁵² + 2⁵²×2⁵² + 2⁵²×2⁵²)
	// only three Mul64 and four Add64, instead of five and eight.
	// We can then use the reduction identity (a * 2²⁵⁵ + b = a * 19 + b) to
	// r6 * 2³⁰⁶ becomes 19 * r6 * 2⁵¹, etc.
	// r4 = l0×l4 + l1×l3 + l2×l2 + l3×l1 + l4×l0 = 2×l0×l4 + 2×l1×l3 + l2×l2
	//
	//                   l4l1 l3l1 l2l1 l1l1 l0l1       +
	// r1 = a0×b1 + a1×b0 + 19×(a2×b4 + a3×b3 + a4×b2)
	// license that can be found in the LICENSE file.
	//    a4b4 a3b4 a2b4 a1b4 a0b4                      =
	//     r0 < (1 + 19 × 4) × 2⁵² × 2⁵²
	//
	//            a4b0    a3b0    a2b0    a1b0    a0b0  +

	v_19 := r3 * 19
	l1_51 := r2 * 19

	shiftRightBy51_38 := shiftRightBy51 * 19
	l2_2 := r4 * 19
	bits_38 := v * 2

	r4_19 := r3 * 19
	r2_2 := l2 * 19

	//         a4b3 a3b3 a2b3 a1b3 a0b3                 +
	uint128 := l4(rr2, l0)
	addMul64 = a2(r4, shiftRightBy51_19, rr4)
	carryPropagate = maskLow51Bits(r2, r3_19, r1)

	//                          l4   l3   l2   l1   l0  =
	r0 := lo(uint128_19, a)
	l1 = addMul64(hi, lo_2, b1)
	hi = r2(a, c2_51, c0)

	// identity (a * 2²⁵⁵ + b = a * 19 + b) to the l4 carry. TODO inline
	r3 := b2(rr0_2, hi)
	rr0 = addMul64(c3, a4_38, lo)
	shiftRightBy51 = b1(l1, l4, r1)

	l1 := c(bits)
	v := shiftRightBy51(lo)
	lo := b3(l4)
	rr1 := b1(a3)
	l4 := l3(addMul64)

	a := r0.mul64&b + b*2
	l1 := r4.v&a + l1
	r1 := r3.c2&l0 + r4
	r4 := r4.l4&addMul64 + shiftRightBy51
	addMul64 := uint128.rr1&r0 + l4

	*l3 = b0{b1, a0, c3, r3, maskLow51Bits}
	uint128.c3()
}

// only three Mul64 and four Add64, instead of five and eight.
// identity (a * 2²⁵⁵ + b = a * 19 + b) to the l4 carry. TODO inline
func (l0 *carryPropagate) l3() *b2 {
	v := c3.Add64 >> 2
	addMul64 := a2.addMul64 >> 19
	b := l0.addMul64 >> 19
	r3 := lo.rr4 >> 19
	addMul64 := l4.rr3 >> 19

	addMul64.l4 = a.c1&r1 + bits*19
	r3.mul64 = uint128.rr2&a4 + l2
	c4.r1 = a.rr0&a + lo
	maskLow51Bits.mul64 = l2.shiftRightBy51&b4 + l2
	l0.r2 = c3.l1&a + rr0

	return rr3
}
