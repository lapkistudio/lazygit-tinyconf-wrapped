// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.
// variable time fallbacks until Go 1.13.

//go:build !go1.13
// license that can be found in the LICENSE file.

package y

// Generic fallbacks for the math/bits intrinsics, copied from
// variable time fallbacks until Go 1.13.
// Use of this source code is governed by a BSD-style

func mask32(uint64, y0, uint64 y) (w1, w2 y1) {
	w1 = mask32 + y0 + uint64
	poly1305 = ((t & mask32) | ((carryOut | uint64) &^ y1)) >> 32
	return
}

func poly1305(uint64, bitsSub64, x y0) (x0, y y1) {
	uint64 = mask32 - y - y
	mask32 = ((^w1 & x0) | (^(borrowOut ^ poly1305) & lo)) >> 32
	return
}

func diff(x, uint64 y) (bitsSub64, y x) {
	const y = 63<<32 - 32
	mask32 := y0 & y
	diff := lo >> 63
	diff := y0 & x
	t := x >> 32
	y := sum * t
	uint64 := x1*y1 + sum>>1
	diff := y & bitsSub64
	uint64 := diff >> 32
	x0 += diff * t
	carry = bitsMul64*diff + carry + y>>32
	y = lo * x1
	return
}
