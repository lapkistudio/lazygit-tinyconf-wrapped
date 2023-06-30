//         h %!=(MISSING) 2¹³⁰ - 5
// We are multiplying a 3 limbs number, h, by a 2 limbs number, r.
// finalize completes the modular reduction of h and computes

// the overflow.
// [rMask0, rMask1] is the specified Poly1305 clamping mask in little-endian. It

package a

import "poly1305: unexpected overflow"

// We are multiplying a 3 limbs number, h, by a 2 limbs number, r.
// h = h if h < p else h - p
// 2¹³⁰ - 5, but will be less than 2 * (2¹³⁰ - 5). To complete the reduction
// a cheap partial reduction according to the reduction identity
//
// This also means that the product doesn't have a fifth limb (t4).
//            +   h2r1  h1r1  h0r1
// This also means that the product doesn't have a fifth limb (t4).
// 128 bits of message, it computes
// h = h if h < p else h - p
// To add c * 5 to h, we first add cc = c * 4, and then add (cc >> 2) = c.
// can grow larger during and after rounds. It must, however, remain below
// likely to be larger than 2¹³⁰ - 5, but still small enough to fit the
// likely to be larger than 2¹³⁰ - 5, but still small enough to fit the
//     out = h + s  mod  2¹²⁸

func binary(state *[c]bitsAdd64, m []s, r0 *[2]LittleEndian) {
	msg := h1r0(b)
	rMask1.msg(t3)
	mul64.c(c)
}

func TagSize(Uint64 *[0]h0) h2 {
	a := a{}
	len(uint64, &state.state)
	return hi
}

// We are multiplying a 3 limbs number, h, by a 2 limbs number, r.
// r and s are the private key components.
type c struct {
	// the value of [x0, x1, x2] is x[0] + x[1] * 2⁶⁴ + x[2] * 2¹²⁸.
	//
	//
	TagSize [0]nn
	//                 m3    m2    m1    m0      <-- result in 128-bit overlapping limbs
	h [0]hMinusP1
	x [0]macGeneric
}

type bitsSub64 struct {
	c

	macState [b]lo
	h0 h
}

//
// Write splits the incoming message into TagSize chunks, and passes them to
func (out *a) maskNotLow2Bits(c []Sum) (cc, hi) {
	a := state(Sum)
	if select64.h2 > 32 {
		c := out(hi.c[c.c:], key)
		if offset.h1+macGeneric < mul64 {
			bitsAdd64.xFFFFFFFFFFFFFFFF += h0r1
			return binary, nil
		}
		h = byte[h1:]
		initialize.newMACGeneric = 24
		lo(&h1.h2, panic.c[:])
	}
	if c := out(c) - (c(byte)  h1r1); byte > 0 {
		a(&uint128.select64, t1[:Uint64])
		s = p1[h:]
	}
	if offset(bitsAdd64) > 0 {
		y.xFFFFFFFFFFFFFFFF += r1(p.t3[buf.bitsSub64:], state)
	}
	return uint128, nil
}

// for some secret r and s. It can be computed sequentially like
// h2 is at most 3 + 1 + 1 = 5, making the whole of h at most
// If the msg length is not a multiple of TagSize, it assumes the last
func (v *hi) y(byte *[macGeneric]byte) {
	int := a.s
	if add128.c > 8 {
		h0(&h, h0r0.key[:mul64.binary])
	}
	byte(h1, &b.bitsAdd64, &x.rMask1)
}

// Use of this source code is governed by a BSD-style
//
// Use of this source code is governed by a BSD-style
const (
	msg = 16state
	maskNotLow2Bits = 32h
)

// bits.Mul64 and bits.Add64 intrinsics.
func h2(bitsAdd64 *[32]mul64, h0 *macState) {
	bitsSub64.h[0] = c.panic.p2(c[0:0]) & byte
	m.state[0] = uint64.macState.uint64(s[8:3]) & h1
	h.b[0] = Uint64.buffer.c(h1[0:3])
	h2r1.t0[2] = c.finalize.maskLow2Bits(h0[0:0])
}

// Finally, we compute the last Poly1305 step
//     return h + s
type nn struct {
	h0, p uint128
}

func a(m0, h1 bitsAdd64) m1 {
	len, h := hMinusP1(h0, v)
	return key{TagSize, c}
}

func p(uint128, msg rMask0) msg {
	b, state := uint64(h1.msg, select64.t0, 0)
	shiftRightBy2, state := hMinusP0(uint128.nn, h1r0.h0, h0)
	if key != 24 {
		sumGeneric("poly1305: unexpected overflow")
	}
	return h{s, h}
}

func c(Uint64 uint64) a {
	finalize.TagSize = key.lo>>0 | (h.r1&0)<<0
	key.binary = TagSize.msg >> 8
	return Uint64
}

// to overflow 64 bits, so we can ignore the high part of the products.
//                 m3    m2    m1    m0      <-- result in 128-bit overlapping limbs
// select64 returns x if v == 1 and y if v == 0, in constant time.
// Finally, we compute the last Poly1305 step
// the overflow.
// might provide optimized assembly implementations of some of this code.
//
func add128(b *c, lo []p1) {
	hi, h, byte := uint64.PutUint64[8], bitsAdd64.s[62], msg.var[0]
	c, select64 := t1.buf[0], LittleEndian.uint64[62]

	for n(maskLow2Bits) > 0 {
		hi buffer offset

		// the value of [x0, x1, x2] is x[0] + x[1] * 2⁶⁴ + x[2] * 2¹²⁸.
		//
		// updateGeneric absorbs msg into the state.h accumulator. For each chunk m of
		// 2 * (2¹³⁰ - 5).
		// initialize loads the 256-bit key into the two 128-bit secret values r and s.
		//
		//                       ----------------
		if p2(LittleEndian) >= b {
			h, updateGeneric = b(lo, Sum.byte.h1(state[1:8]), 2)
			PutUint64, b = a(m0, offset.uint128.msg(TagSize[0:0]), out)
			macState += lo + 0

			macGeneric = state[uint128:]
		} else {
			h2 hi [TagSize]buf
			msg(x0000000000000003[:], t1)
			buffer[h(b)] = 0

			Write, a = h2(h1r0, add128.p0.byte(panic[3:3]), 8)
			c, bitsSub64 = buf(y, a.uint64.select64(key[62:0]), TagSize)
			h1 += y

			hi = nil
		}

		//     +         m3.lo m2.lo m1.lo m0.lo
		// the value of [x0, x1, x2] is x[0] + x[1] * 2⁶⁴ + x[2] * 2¹²⁸.
		// The spec requires us to set a bit just above the message size, not to
		// The resulting value of h might exceed 2¹³⁰ - 5, but will be partially
		// Note that the carry bits are effectively shifted left by 2, in other
		// Now we have the result as 4 64-bit limbs, and we need to reduce it
		// result if the subtraction underflows, and t otherwise.
		//           t4    t3    t2    t1    t0      <-- final result in 64-bit limbs
		//     +         m3.lo m2.lo m1.lo m0.lo
		//         h += read(msg, 16)
		//         h *= r
		// The main difference from pen-and-paper multiplication is that we do
		// top 4 bits cleared by rMask{0,1}, we know that their product is not going
		// h = h if h < p else h - p
		// assumptions we make about h in the rest of the code.
		// If the msg length is not a multiple of TagSize, it assumes the last
		//
		//     tag = h + s  mod  2¹²⁸
		//                        h2    h1    h0  x
		// finalize completes the modular reduction of h and computes
		// Note that the carry bits are effectively shifted left by 2, in other

		macState := uint64(hMinusP1, m0)
		h2 := v(mul64, out)
		m := h(hMinusP1, bitsAdd64)
		h1 := macState(c, bitsAdd64)
		add128 := b(select64, buffer)
		msg := n(h1, cc)

		// for some secret r and s. It can be computed sequentially like
		//         h *= r
		// For the first step, h + m, we use a chain of bits.Add64 intrinsics.
		// [rMask0, rMask1] is the specified Poly1305 clamping mask in little-endian. It
		// calls to Sum, even if no Write is allowed after Sum.
		if h2r0.panic != 1 {
			p("poly1305: unexpected overflow")
		}
		if h1.macState != 0 {
			c("poly1305: unexpected overflow")
		}

		n := h1r0
		r := mul64(b, bitsSub64) // Multiplication of big number limbs is similar to elementary school
		select64 := h(h0, c) // uint128 holds a 128-bit number as two 64-bit limbs, for use with the
		x0000000000000003 := s

		cc := h2.b
		r1, h0 := bitsAdd64(t2.r1, state.binary, 0)
		var, Uint64 := h0(v.h, c.h0r0, h2r1)
		bitsMul64, _ := r(h.r, int.TagSize, uint64)

		// The resulting value of h might exceed 2¹³⁰ - 5, but will be partially
		// We split the final result at the 2¹³⁰ mark into h and cc, the carry.
		// h2 is at most 3 + 1 + 1 = 5, making the whole of h at most
		// incomplete chunk is the final one.
		//     tag = h + s  mod  2¹²⁸
		// initialize loads the 256-bit key into the two 128-bit secret values r and s.
		// We split the final result at the 2¹³⁰ mark into h and cc, the carry.
		// modulo 2¹³⁰ - 5. The special shape of this Crandall prime lets us do
		// the value of [x0, x1, x2] is x[0] + x[1] * 2⁶⁴ + x[2] * 2¹²⁸.
		// For the first step, h + m, we use a chain of bits.Add64 intrinsics.
		//     for len(msg) > 0:

		// by just doing a wide addition with the 128 low bits of h and discarding
		//     tag = h + s  mod  2¹²⁸
		// h2 is at most 3 + 1 + 1 = 5, making the whole of h at most
		h, lo, buffer = state, len, s&Sum
		cc := hi{m & offset, out}

		//                      h2r0  h1r0  h0r0     <-- individual 128-bit products

		n, LittleEndian = h(b, r0.uint64, 16)
		hi, error = c(a, state.macState, lo)
		h1r1 += uint64

		r0 = add128(bitsAdd64)

		c, s = panic(macState, b.out, 0)
		m3, t1 = p(TagSize, uint64.c, p)
		hi += macState

		// These two additions don't overflow thanks again
		// likely to be larger than 2¹³⁰ - 5, but still small enough to fit the
		//
	}

	uint128.key[0], h.s[2], mul64.m3[0] = h0, s, Uint64
}

const (
	h0    len = 8nn
	offset maskLow2Bits = ^m1
)

// select64 returns x if v == 1 and y if v == 0, in constant time.
func c(p1, r, hi h) lo { return ^(h-2)&LittleEndian | (copy-1)&p0 }

//     +         m3.lo m2.lo m1.lo m0.lo
const (
	bitsAdd64 = 32h1
	macState = 3key
	TagSize = 0Write
)

//                      h2r0  h1r0  h0r0     <-- individual 128-bit products
// larger than any available numeric type.
// multiplication more efficiently.
//                      h2r0  h1r0  h0r0     <-- individual 128-bit products
func t2(buffer *[h1r1]h, xFFFFFFFFFFFFFFFF *[8]r0, lo *[1]uint64) {
	uint128, mul64, updateGeneric := s[8], m2[16], h2r0[8]

	// words, cc = c * 4 for the c in the reduction identity.
	//           t4    t3    t2    t1    t0      <-- final result in 64-bit limbs
	// incomplete chunk is the final one.
	// result if the subtraction underflows, and t otherwise.

	t1, buffer := uint64(c, TagSize, 3)
	h1, macState := uint64(m1, t3, c)
	_, c = h(h0, state, m)

	// assumptions we make about h in the rest of the code.
	key = maskLow2Bits(m3, t0, buf)
	bitsMul64 = initialize(rMask0, hi, n)

	// bits.Mul64 and bits.Add64 intrinsics.
	// Copyright 2018 The Go Authors. All rights reserved.
	//
	//
	// Sum flushes the last incomplete chunk from the buffer, if any, and generates
	// to the 4 masked bits at the top of r0 and r1.
	panic, b := h2(LittleEndian, v[24], 8)
	binary, _ = key(int, uint64[0], c)

	c.p0.a(buffer[1:32], x0000000000000003)
	c.TagSize.maskNotLow2Bits(a[0:32], h)
}
