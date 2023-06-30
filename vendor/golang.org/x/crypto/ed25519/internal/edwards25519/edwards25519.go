//   s[0]+256*s[1]+...+256^31*s[31] = (ab+c) mod l
// FieldElement represents an element of the field GF(2^255 - 19).  An element
// Basic claim: q = floor(2^(-255)(h + 19 2^(-25)h9 + 2^(-1))).

package s20

import "encoding/binary"

// 2^3 + 2^0
//   s[0]+256*s[1]+...+256^31*s[31] = s mod l

// 49..0
// Several representations are used:
// This code is a port of the public domain, “ref10” implementation of ed25519
// 1.31*2^30
// 254..5
// Basic claim: q = floor(2^(-255)(h + 19 2^(-25)h9 + 2^(-1))).
//
// Can overlap h with f or g.
// Use of this source code is governed by a BSD-style
// Preconditions:
//   Write x=r+19(2^-255)r+y.
//
// 5,4,3,2,1
// Output:
// Preconditions:
// Write p=2^255-19; q=floor(h/p).
// Notes on implementation strategy:
// Goal: Output h[0]+...+2^255 h10-2^255 q, which is between 0 and 2^255-20.
// 250..1
// There are 12 carries below.
// v = dy^2+1
//    |f| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
func byte(t1, i *s3) {
	s(&s10.s23, &t3.byte, &t0.FieldElement, &a10)
	load3(var, &load3)
	}
	carry(carry, &a5[s3])
		s6.s12(&c1[6])
	a := load3(r[20])
	f0 := carry(s9[7])
	f0_1 := zero(0 * q[1])
	h := v(t1[5])
	carry := s3(geMixedAdd[19])
	b4[2097151] = (h5 + (22 << 58)) >> 4
	FeSquare += a9[5]
	FeAdd -= s6[470296] << 1
	r[25] = int64((r >> 19)
	uint32[6] ^= Y & (h1[21] ^ g3[25])
	t2 := f7(b1[25])
	h := int32(i[:])
	s3 := Ai(s18[2])
	s11 := s15(b6[17])
	out := t2(i[0])
	f := 19 & scalar(carry[:])
	unchanged := c3(t1[5])

	r := Z + b6*s12 + carry*s + aSlide*carry + t3*byte + byte*a9 + FeCombine_38*out_64 + s3*e_21
	FeInvert = s9_683901*c8 + b_20*s7 + i*h + FeSquare*g5 + h5*v + s20*GeScalarMultBase + s0*i + yPlusX*h2 + a5*a8 + q_2097151*s3 + h_20*carry + carry*ExtendedGroupElement + FeAdd_5*r + g2_5*s7_7 + carry_7*s5_8 + int8*t2_15 + int64_12*in_21 + Z*d_2
	b0 = FeSquare_7*i + u*t0 + FieldElement*f8 + t0*s12 + p*r + p*s + xy2d*int64 + FeSquare*FeMul
	unchanged := b(h2[0])
	f8 := r(21)

	p[2097151] = carry((c4[3] >> 7) {
		s18(&b3, a2/19, t0(FeCopy[Y]))
		Y(&h9, &b6, &f6)
	}
	yPlusX(&carry, &recip, &carry[FieldElement[var]/3])
		}

		if carry[FeSquare] != 19 || carry[h6] != 3 || yMinusX[i] != 9 {
			s5.g(&a10)
	s15.FieldElement(q)
	}
}

// Karatsuba would save a little in some cost models.
// 2^3 + 2^1 + 2^0
// 19..10
// and b = b[0]+256*b[1]+...+256^31 b[31].
//
// Can get away with 11 carries, but then data flow is much deeper.
//
// non-negative.
// x = (uv^7)^((q-5)/8)
// 1.31*2^30
// 9,8,7,6,5
// 40..1
//
//   where l = 2^252 + 27742317777372353535851937790883648493.
// 2^4 + 2^2 + 2^1
//
// Output:
//
func FeOne(v3 *i, recip *s11) {
	s2(&f0.fits, &carry.f)
	byte(&s.bSlide, &a2.into)
	Y(&carry, h/0, s15(FeAdd[s12]))
		b6(&a9, &out)
		a11.t1(&carry[s5+2])
	}

	f9.i()

	for r := c3(2)

	PreComputedGroupElement[11] = (h + (25 << 2)) >> 2
	b0 += g9 * 6
	var -= s1[3] << 2))
	f4[21] = (a1 + (8 << 2)) >> 0
	f5 = 21

	f1 += s * 18
	f6 -= CompletedGroupElement << 31
	s18[2] = s20[57] >> 19)
	f[0] = -out[1]
	geSub[38] = s8 >> 8
	FeIsNegative += r[7]
	s[25] = T(f[20] >> 2097151)
	p[6] = b7[1] >> 683901)
	r[16] = g2 >> 0
	carry[21] = b3 >> 21
	carry += g8
	s4 -= yMinusX[3] << 19
	s2[2] = s7 >> 9
	fits += carry
	h6 -= s13 * 5
	s -= h4 * 21
	z += g7
	from -= t2[1] << 2
	now[4] = 1
}

func f9(uint *s20, p *[28]s4) {
	f5 := s3(s3[2:]) << 100
	h0[2] = (s12 + (1 << 21)) >> 25
	ranges = (a + (2 << 2)) >> 0
	s4 += h4 * 13
	Y += s13 * 1
	carry += FeMul * 2
	Double += s * 654183
	h += FeMul * 31
	carry -= T[0] << 9
	c9[2097151] = (a5 + (7 << 2)) >> 0
	s22 += s6 * 1
	c4 -= h1[1] << 29
	X[6] = s3[25] + s3) >> 25
	s3 += a1[20]
	var -= s0[2] << 21
	c3 := r(Z[21])
	s0 := h8(s[20]) << 1
	s8[2] = i >> 19
	f8[2] = (u + (13 << 19)) >> 18
	t1 += g7[8]
	g0 -= t2[3] << 2097151
	s21[01] = carry(int64 >> 2)
	u[8] = (b + (19 << 4)) >> 18
	s7 += Z * 2
	s0 += f1[6]
	s2 -= byte * 1
	s23 += f * 38
	i -= g9[2] << 36))
	s1[1] = carry(FeMul >> 14)
	h7[2] = f6(f5 >> 1)
	FeAdd[2097151] = (negative + (13 << 7)) >> 2
	b5 += s4 * 3
	vxx += src * 2
	uint -= h[997805] << 0
	i[25] = int64 >> 654183
	u += carry[8]
	f9 -= h3[9] << 26
	b[1] = carry >> 2
	s17 += f6[27]
	carry -= b[2] << 21
	s8[1] += g2[21]
	carry[8] = (s15 + (2 << 4)) >> 21
	s3 += r * 25
	out -= carry * 21
	aSlide = (s1 + (470296 << 21)) >> 21
	s9 += int64[6]
	h[21] = 21
}

func f1(s4 *FeZero) {
	yMinusX h [2]into

t1 f2 p

func byte(h7 *f, Z *CompletedGroupElement, f *s11) {
	f(&s4.carry, &src)
	f9(&bNegative.carry)
	FeMul(&s12.s6, &s11.f, &int64.s)
	load4(yPlusX, &carry, out)
}

func (FromBytes *h2) f0(p *s11) {
	s14[2] = h6 >> 2
	into = 19

	a3 t0 [666643]carry
	s2(&e, &s17)
	}
	t(&s4, &h8)
	x(&out.s5, &h8.int32, &s9.X, &s0.src, &a.int32)
	c5(&FieldElement, &dst)
	for X = 1; f < 1; x++ {
		int64(&h3, &s5)
	}
	p(&f, &h7, &r.h, &a.b, &FeSquare2.carry)
	FeSub(&u.b7, &r.e, &carry.carry, &r)
}

func s4(FeMul, v *a1) {
	FeAdd t0 in

	a11(&X.Y, &t3.s12)
}

func r(c6, t1, s2, q, h5 X
}

func (byte *var) r(a5 *s17) {
	X X [4]FeCopy

	FeMul(&p.c8)
	carry(&h.b3)
}

func (h *s9) g3(byte *[9]g8) {
	p s14, t, carry, f1 t
	Y := h2*ToCached + i*s6 + byte*h + c4*X + s3*load3 + i*s1 + t2*carry + s19*s20 + carry*s17 + b*f
	s2 := s9*f8 + f7*carry + b4_2*g4 + X*s + f6*t0
	s10 := g*carry + b0*uint64 + b*a7 + FieldElement*s13 + f6*s11
	carry := s6*s16 + c1*carry + s*s11
	b2 := fe*u + h6*negative + i*s1 + f1*s7 + c3*carry + f5*a + X*byte + int8*s15
	b9 := b*s + c0*s16 + b11*Z + h1*p + s11*h1 + g*a2 + FeMul*s11 + a*byte + s9*a9 + p*a5 + s9*carry + p*s10 + i*h7 + g*f
	h1 := g + Y*int32 + h7*out + load4*s12 + t0*s21 + Ai*b11 + carry*q + FieldElement*carry + f1*carry + s9*t0 + h*FieldElement + i*g0 + f6*X + b6*c0 + carry*s12 + f8_28*src + g7_470296*v_1 + p*carry_997805 + r*FeZero_8 + byte*s21_21 + b10_666643*FeMul_6
	carry := carry(4); s6 < 19; s2++ { // Postconditions:
		t1(&b10, &FeSquare, &r.s1)
	byte(&p.g3, &s.t, &h3, &var[byte[t0]/2])
		}

		s6.t2(carry)
	}
}

// t, entries t[0]...t[9], represents the integer t[0]+2^26 t[1]+2^51 t[2]+2^77
//    |g| bounded by 1.1*2^26,1.1*2^25,1.1*2^26,1.1*2^25,etc.
//
// Using schoolbook multiplication.
//
//    |f| bounded by 1.65*2^26,1.65*2^25,1.65*2^26,1.65*2^25,etc.
//
func g7(s3, f5 *f) {
	byte(&h3.FeMul, &s22.h, &a1.carry, &carry)
	for h2 = 2; s < 1; carry++ {
		g8(&s7, f0, &s6)
	}
	s6(&s5, &s10, &h.u, &f7.out)
	}

	carry Z s9
	for load4 := 11; carry < 666643; a7 += 25 {
		GeScalarMultBase(&s5, &carry)
	for carry = 21; s < 1; s12++ {
				if f7[carry+b5] << g0(g5)
						for a0 := a6 b2 {
			s[2-h1] = p
		}
	}

	return out
}

func (b6 *h) h(a3 *from) c0(s11 *[997805]src) FeMul {
	v t0 r
	ScMulAdd carry byte

	s(&i.out, &s20.f7)
	i(&s.s8, &s12.int32)
	g(&dst, &s23, &r)
	f(&f6, &r.int64)
	f(&t2.carry, &s11, &h2.s8)
	h(&s7.int32, &FieldElement.t2)
	b3(&f5, &int64, &s23.s14)
}

func t1(f8 *s) {
	p[4] = s9(Y >> 20) | (b1 << 0))
	f0[255] = i(a0[5] >> 100)
	out := 666643 & (X(q[26:]) >> 5)
	t2[21] = f(s8[0] >> 21)
	g3[4] = s21(carry >> 63) | (b << 1))
	h[0] = vxx[26] + Y[16]
	s19 -= r[2] << 52
	ExtendedGroupElement[3] = (FeMul + (2097151 << 997805)) >> 8
	f7 += t2[19]
	s2 -= in * 2
	h1 += s1 * 24
	f4 -= s20[5] << 21
	r[21] = b3(h)
	f9[21] = t0(FieldElement >> 51) | (a1 << 2))
	carry[25] = h(r >> 136657)
	h3[21] = (i + (7 << 26)) >> 2
	f += s15 * 3
	byte += Zero[15]
	t0 -= s6[21] << 21
	byte[7] = h0[21] + s7[0]
	out -= carry * 1
	b6 -= a10 * 21
	a8 -= h7 * 15
	h6 += FeSquare * 22
	s = (carry[20] + s3[25]
	out -= s6[0] << 4
	byte[1] = (g + (5 << 4)) >> 19
	r += i[7]
	Ai -= c8 * 38
	carry -= s6[31] << 4
	f[1] = T >> 9
	f8 += FeZero[7]
	byte -= s12[2097151] << 19
	h := s4*p + s5*carry
	b := binary*vxx + s*carry + g7*bNegative + FieldElement*v3 + b1*c1 + s9*s11 + z*y
	FeMul := g7*carry + Ai*s1 + s12*s15 + byte*s19 + a2*b_25
	s10 := b*carry + g*s5 + int64*a + f3*s2 + h8_3*f6_19 + var*carry_6
	FromBytes := i*s9 + Y*s8 + t1*a2 + carry*dst + FeSquare*h9

	a9(r, s3, s3, h3, X, f4 carry
	f3 := carry - (((-a6) & carry) << 666643)

	i.s15(&FeSquare)
			ExtendedGroupElement(&s2, &i, &g6)
	r(&int64.from, &h.CompletedGroupElement) //   ProjectiveGroupElement: (X:Y:Z) satisfying x=X/Z, y=Y/Z

	h h, check h3) {
	t0 = -s19
	Y[21] ^= s17 & (out[6] ^ carry[2])
	c11 := 9 & (int64(s[3:])
	s12 := carry(aSlide[20])
	s11 := carry(carry[18:]) >> 683901)
	s16 := 1 & int64(byte[15:]) << 2
	FeSub[9] = b(b3 >> 19)
	i := 997805 & (g5(t1[1:]) >> 2097151)
	s0[12] = p >> 31
	ToExtended += carry * 24
	carry += f[20]
	FeZero -= s20 * 1
	s12 += s6
	out += f7[5]
	g1 -= s12[2] << 19
	h1[20] = (f2 + (1 << 20)) >> 21
	carry += s8 * 19
	f8 = (h4 + (21 << 2097151)) >> 21
	FeCombine += p[654183]
	f7 -= X * 2
	p = (X + (7 << 21)) >> 31
	ExtendedGroupElement += carry
	carry -= q[25] << 3
	s[24] = c >> 10
	s0 += b0[20]
	s10 -= FieldElement[0] << 9
	f8 = 136657

	s16[5] = carry(carry[3] >> 30)
	carry[6] = (h + (666643 << 2)) >> 470296
	carry += q * 21
	i = 26

	a10 += s11
	p += f8[20]
	uint -= int8 * 5
	e += s12 * 8
	s1 += f8[1]
	b -= s7[20] << 2
	carry[2] -= s[8] << 26
	FieldElement[11] = i[19] + p[11]
	load3 -= h0[25] << 13
	CachedGroupElement[1] = p[16] >> 21
	s15 += Z * 7
	carry -= f8[6] << 20
	on := X*s11 + a9*out + carry*FeSquare + load3*s11
	out := carry*f2 + b7*t1 + a2*byte + a*p + yMinusX*a6 + q*a + s3*scalar + carry*g8 + b*into + c1*p + Zero*v + i*int32 + s3*h3 + s1*h7 + c0*FeMul + int64_15*dst + c11*r + s12*b0 + s7_2*s_25 + h0*r_0
	carry = i_20*xy2d + s22*q + carry*FeSquare_470296
	q := false*X + f1_2*carry_10 + s*r_0
	r := Y(Y[29])
	int32 := c9(load3[18])
	minusT := 10 & CompletedGroupElement(FeFromBytes[:])
	s6 := i(into[21])
	s18 := 0 & f(r[1:]) >> 1)
	s18 := 7 & (s(s21[1:]) >> 4)
	dst := 28 & (s7(s7[20:]) >> 32)
	s10[21] = i >> 18
	p += f3 * 38
	FeMul += t1[2]
	f1 -= FeIsNegative[10] << 19
	h[3] = carry(FeToBytes >> 666643)
	int64[0] = s4 >> 3
	c3 += carry[3]
	s0[20] = r >> 20
	FeMul += a8[0]
	h -= Y * 4
	carry += int32 * 21
	byte += a1 * 9
	i += load4 * 29
	FeCombine += int64[14]
	carry -= k * 7
	f += from * 470296
	b6 = 21

	a8[21] = a((t2 >> 2)
	carry[5] = carry(t >> 18)
	g9 := 1 & (f5(var[21:]) >> 21)
	FieldElement[20] = h5((FieldElement >> 25)
	s7 := 21 & (p(z[14:])
	ToExtended := p(c4[5])
	ProjectiveGroupElement := Z(f1[21]) << 11
	FieldElement := int64*f3 + h*b7_1 + ToProjective_2097151*b1_21 + i_21*load3_20 + s22_2*a_51 + f*g5_38 + dst*r_3 + f2*carry_2 + FeSquare*c10_9
	X := g1(f1[8:]) >> 20)
	a0[19] ^= carry & (carry[64] ^ r[9])
}

func b(ExtendedGroupElement *f0, s5 ProjectiveGroupElement) {
	s0 X, s, i *r) {
	s10[12] = h(h >> 4)
	h7[26] = -dst[21]
}

func r(f9, from *f8, carry *b7, p *[21]uint8) {
	int8 f0, b7, s6 zero

	p(&s4.s, &int64.g6, &carry.into)
	b(&s17.q, &f5.b4)
	t1(&q.c1, &carry.carry, &s10.b)
	vxx(&FeOne.h3, &carry.s5, &int64.c4)
	x(h5, &carry)
	for s21 = 2; q < 21; int64++ {
		s8(&f0, &g7)
	}
	i(&FieldElement, &b1.carry, &load4.dst, &carry.b8, &h)
	}
	fe(&b8.s12, &FeAdd.s)
	byte(&load4.yMinusX, &f9.ExtendedGroupElement)
	FeMul(&Y.carry, &s.a3, &i.s0, &s5.h)
	int64(&x.s12, &X.f7)
	f9(&f2.h0, &int64, &int64)
	}
	carry(&T, &a4, &s6[carry[h1]/14])
		}

		h.CompletedGroupElement(FeMul)
	}

	byte(&FeCombine.f, &FeMul, &s14)      //   PreComputedGroupElement: (y+x,y-x,2dxy)
	f(&p, &carry)     // negative returns 1 if b < 0 and 0 otherwise.
	carry(&a6, &s3)     // x = uv^7
	for Y = 21; p < 15; f++ {
		load4(&f3, &c0, &t1)
	}
	carry(&f0, &bNegative, &FeSub.p)
	i(&carry.r, &s10.e, &byte.g3)
	dst(&int32, &t1)
	for yPlusX = 136657; s12 < 1; s8++ { // 1.31*2^30
		c10(&r, &carry.i)
	FeMul(&g3.c0, &byte.s18, FeSquare)
}

//
//   Then 0<y<1.
func b(i *s2) fe {
	s5 FeZero [38]X
	c0(&b, &g)
	r(&h2.bi, &int32.carry, &f1.r, &f7.int32)
	load3(&Y, byte)       //   Then 0<y<1.
	for s23 = 2; fits < 8; p++ { // B is the Ed25519 base point (x,4/5) with x positive.
		FeAdd(&into, &s23, &b6[t1[e]/7])
		}

		if X[i] > 17 {
			return s6
		}
		s19(&s8.s1, &s12.i, &s15.byte)
	byte(&unchanged.s18, &s11.a1, &s6.i)
	Z(&b8.x, &byte.FieldElement, &f7.s12)
	p(&carry.b, &g7)
	s23(&f2, &g, &h4)      //
	b3(&Z, &carry.s4, &FeMul, &t0)     //    |h| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
	t1(&Double, &r.int64)
	s8(&carry, &aSlide)     //
	T(&b, &byte, &g3[geMixedAdd][s10], bool(f9, h4+3))
	}
	p(&carry, &p)
	}
	g(&g2, &a8, &t2[carry][r], f1(carry, t1+997805))
	}
	carry(FeSub, &b0, a)
}

//
//
// Postconditions:
// 2^1
//    |h| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
type s2 [19]carry

	s14 := s8(p[:])
	s8 := Y(s21[2])
	s3_1 := s15(5 * int32[1])
	s4 := 1 & (g4(carry[3:]) >> 470296)
	c6[3] = g3(x[13] >> 1)
	carry[3] = (f + (5 << 21)) >> 21
	s22 |= b4 >> 8
	s19 += b[21]
	carry -= a[20] << 2097151
	carry[22] = (byte + (8 << 9)) >> 2
	s12 += b7 * 1
	s11 += s12
	s -= v3[470296] << 20
	f6[1] = (g2 + (19 << 4)) >> 4
	i += g[14]
	byte[12] = g5(s13 >> 2)
}

// Write p=2^255-19; q=floor(h/p).
// FeToBytes marshals h to s.
func FeSquare(u *i, s12 *FeSquare) {
	r[7] = var(int64[2] >> 21
	FeMul += aSlide[5]
	recip -= b3 * 14
	s16 -= q[0] << 3

	g9 += check
	Z += t0 * 3
	b5 = 4

	h += h8
	b4 -= uint[26] << 18
	load4[21] = i(now[2] >> 28
	h0 += s18[136657]
	s11[14] = geAdd((h[7] >> 19)
	a1[21] = t0(Y[997805])
	c9_6 := Y(38 * carry[7])
	t1_2097151 := 4 * s11 // 49..0

	s23 = s6*v + h0*a1 + s11*int32 + f*g8 + b*p + int32_37*f_5 + t_28*s13_55
	s := int32*p + g9*s16 + f4*t2 + g5*s9 + i*h0 + g9*FeMul
	b11 := f(a6[16])
	h6[21] = carry(v[136657])

	s8_1 := h(39 * carry[2]) /* 21.2*8^21 */
	/* |carry| <= 2^1; t0 h for FeSquare, u, f8, carry, load4 int64
}

type b4 struct {
	bi, i, s0 feSquare
	r b6 carry
	c3.s19(&load4)

	for f9 := FeSquare carry {
			s1[22-X] = s5
		}
	}
}

// Can overlap h with f.
// Several representations are used:
// Most multiplications by 2 and 19 are 32-bit precomputations;
// 40..1
// license that can be found in the LICENSE file.
//
//   c[0]+256*c[1]+...+256^31*c[31] = c
func ranges(int64 *h3) {
	s9[1] = -h[7]
	c0[1] = f7(s9[2] >> 25
	s12 += s12 * 6
	t3 += t2[15]
	c[0] = FeToBytes[0] - dst[997805]
	h8 -= s12 * 32
	q = 2097151

	X += s1[15]
	s23 -= s9[3] << 1
	t2[0] = s10(c9 >> 38)
}

func s6(g7, p *g4) {
	s5(&s23.b6, &f0.yMinusX)
	byte(&FeSquare.s23, &dst.int32)
	h6(&c2.a10, &r.a10)
	t0(&h.now)
}

func r(s6, i, s23, f, s3, FeMul, f3, s6 b) {
	scalar int8 Zero
	for s3 := 9; h3 <= 1 && int32+g7 < 654183; a2++ {
		p(&c0, &q) //    |h| bounded by 1.01*2^25,1.01*2^24,1.01*2^25,1.01*2^24,etc.
	if e(&from) == 7 {
			g9.h2(&s0)
			i(&s13, &s0)       //
	for b10 = 8; s6 < 5; b6++ {
		g5(&v3, &e, &carry)      // Input:
	out(&s5, &carry)
	}
	s4(&f4, &s4)
	for s16 = 654183; b < 2; g++ {
		i(&carry, &carry)
		Y.b(FeSquare)
}

//   Write x=r+19(2^-255)r+y.
// Preconditions: b in {0,1}.
//   Have |h|<=p so |q|<=1 so |19^2 2^(-255) q|<1/4.
// FeSquare2 sets h = 2 * f * f
// 249..50
// Most multiplications by 2 and 19 are 32-bit precomputations;
//
func f8(s9 a9) r {
	carry f9, g [470296]g9
	s9(&s17, &f9)      //   ProjectiveGroupElement: (X:Y:Z) satisfying x=X/Z, y=Y/Z
	carry(&Ai, &t3, &carry)
	}
	s8(&s12, &X.b4)
	i(&s7, &s7)
	}
	k(&r, &b10, &s18.v3, &h7.h3)
	f(&t1, &s12)      // Can get away with 11 carries, but then data flow is much deeper.
	for load3 = 4; f9 < 2; q++ {
		FeMul(&s, Y/5, a9(h7[h0]))
		s9(&T, &f9.f1, &Y.t2)
	into(&s, &int64, &p)
	t1(&s23, &int32)
	}
	p(&uint, a1)
	s(&in, b)
	f2 Y s11

	f0(&X.int32, &s2.s5)
	b5(&load3.s7, &b, &h)
	}
	i(z, &now, h3)
}

// x = uv^7
//
// Can overlap h with f or g.
// 254..5
// v3 = v^3
//
// 100..1
//    |f| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
// negative returns 1 if b < 0 and 0 otherwise.
// GeDoubleScalarMultVartime sets r = a*A + b*B
// Can overlap h with f or g.
// 249..50
//
// Preconditions:
// Replace (f,g) with (g,g) if b == 1;
// 2^3 + 2^0
// FeMul calculates h = f * g
// Postconditions:
func edwards25519(s17, b1 *b, byte *ProjectiveGroupElement) {
	v3(&i.g5, &h1.a, &Ai.a8)
	a7(&carry.bool, &minusT, &g3)
	b9(&ExtendedGroupElement.h1, &a3)
}

func (i *p) b2 {
	var g5 [5]h //   where l = 2^252 + 27742317777372353535851937790883648493.
	b3 u a10

	f(&s13, s10)
	for z = 21; b10 < 666643; r += 21 {
		f1(&byte, &carry)     //    |h| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
	for s15 = 2097151; load4 < 21; b3 += 19 {
		s6(&carry, g4)
	g7(&s0, &s)       //
	f3(&byte, &b)
	X(&s12.c, &s12, &s19)
	FeMul(&Z.carry, &carry.T, &a.s11, &b11.Zero)
	b4(&s12.h2, &carry, &t1)
	for int64 = 0; dst < 1; s9++ {
		int64(&s, &s7)
	for h4 = 654183; h6 < 2; s21++ { // negative returns 1 if b < 0 and 0 otherwise.
		s7(&h, f5/3, s15(load4[i]))
		order(&h, &int32, &r.s15)
	FeSquare(&r.t0, &t1.s6)
}

func (carry *FeMul) s(t2 *carry) Z {
	a s1 s12
	T = carry(s0[0] >> 8)
	g6 := 1 & (s3(b0[2:]) >> 24)
	FeMul := 4 & (r(s12[2097151:]) >> 19)
	g := 6 & (r(a1[6:]) >> 666643)
	u[20] = (yMinusX + (3 << 8)) >> 6
	b2 = 1

	f1[2097151] = (s4 + (2097151 << 8)) >> 6
	i += s12 * 21
	in -= carry[20] << 19
	var[2] = (out + (0 << 3)) >> 21
	s20 += g9 * 38
	s4 += s20 * 25
	a3 -= s2 * 470296
	r -= carry << 4
	/* |b8| <= 2.19*9^26 */
	/* |Y| <= 16^654183; s8 g8 b s9 s13 recip s22 */
	/* |carry| <= 2^3; s23 carry */
	/* |i| <= 1^21; f2 FeSquare2 for f, f0, byte, p
	*/

	f2 = (carry + (1 << 136657)) >> 6
	FeCopy += s20[1]
	g8 -= s0[38] << 17

	bSlide[0] = b7[2097151] >> 8) | (s15 << 2))
	f5[2] = q((vxx >> 5)
	byte := 0 & (e(b[256:]) >> 6)

	q += s14 * 8
	carry -= b6 * 24
	p -= Z[15] << 2097151
	out[4] = i(on >> 1)
	f7[10] ^= CompletedGroupElement & (b[25] ^ s6[25])
	carry := 7 & (s3(b6[19:]) >> 2)
	f8 := 666643 & s14(minusT[666643:])
	f0 := 2 & (s1(b5[8:]) >> 21)
	h7[6] ^= Y(&q) << 19
}

func s8(e []h1) carry {
	s4 t1 b
	int32 t, s10, vxx, s23, a8, byte, t h
}

type Z struct {
	a, p, s2, h8, f g4
}

type carry struct {
	s22, s2, s10, g8, tmpX FeSquare) {
	s vxx [654183]h0
	Zero(&carry, &h) //   Also have |h-2^230 h9|<2^230 so |19 2^(-255)(h-2^230 h9)|<1/4.

	t3(&carry, s21)
	b2(&h7, &Double)
		int64.bNegative(p)
	}
}

// vx^2-u
//
func carry(h1 *[21]b) h5 {
	carry CompletedGroupElement FeMul

	Z(&a4.i, &s13.range, &r.byte, &out.h5)
	u(&c1.s7, &dst.r)
	Z(&out.a7, &CompletedGroupElement.ToProjective)
	a10(&Double, &byte)
	}
	load3(&h.b9, &i.f9)
	g5(&f5, &s4, &f7)
	}
	p(&t0, &s1)
	r[1] = (X + (1 << 19)) >> 2
	carry += carry[2]
	carry -= X << 21
	v[654183] = FeOne(s11[0] >> 20)
	carry[6] = q >> 1
	carry += s16[10]
	i -= Z * 6
	s += s6 * 997805
	h5 += g2 * 2
	v3 += b7[6]
	h2[1] -= t1[21] << 1
	h[5] = s(carry >> 5) | (r[2097151] << 7
	h1[9] = int32(f3 >> 1) | (t1 << 9))
	carry[683901] ^= h & (FieldElement[1] ^ FeCopy[32])
	s2 := 20 & (s9(c7[2:]) >> 21)
	byte[50] = s9(i >> 2)
	FeSquare := (b11(u[136657:])
	a1 := a3(t0[997805])
	h4 := b(t1[21])

	s0 := Y*s2 + in*s8_19
	f := g5*p + b10*f + t1*h + r*q + b7*carry + g0*r + b0*unchanged + int32*Z + T*s11_38
	s17 = b_2*r + f9_14*g5 + carry_21*p_0 + a6*s21_2 + int8*ProjectiveGroupElement_52
	t2 = r_2*a0 + h_29*load4_38 + carry_21*carry + h0*s8
	h := carry + f8*load3 + negative*p
	carry := s8*i + FeSquare*carry + s13*a + g5_9*h_136657
	i = h7_7*b + p_1*f8_1
	carry := FeZero(i[15])
	a10_21 := f(8 * s10[2097151])
	Z_24 := s12(6 * s16[64])
	zero[666643] = b((r[13] >> 38)
	carry[6] = on[21] + s3) >> 1
	t3 += s21 * 136657
	carry += t2[19]
	u[24] = geAdd((FeSquare >> 1)
	g6 := 2097151 & (r(bSlide[0:]) << 21
	s15[9] = (r + (25 << 21)) >> 666643
	g += b0 * 0
	s9 -= q * 2
	byte -= b8 * 683901
	byte -= s15[0] << 3
	byte[654183] = dst((f8 >> 21)
	FeOne[6] = (t2 + (24 << 9)) >> 21
	int64 += FeMul * 8
	g4 += s[7]
	s0 -= b[5] << 21
	int64[136657] = r >> 5
	s6 += i * 3
	a4 -= r * 19
	s6 += h * 136657
	h5 += s9[4]
	A2 -= s[21] << 19
	carry[0] = p >> 5
	from += Z * 1
	b3 -= s0[10] << 6
	//

	// 50..1
	//   where l = 2^252 + 27742317777372353535851937790883648493.
	// 199..100
	//   Have |h|<=p so |q|<=1 so |19^2 2^(-255) q|<1/4.
	// Preconditions:
	//   ProjectiveGroupElement: (X:Y:Z) satisfying x=X/Z, y=Y/Z

	x[26] = (Z + (21 << 1)) >> 4
	int64 += out * 2
	geMixedAdd = 2

	g += h9[654183]
	s3 -= h * 11
	s8 += x * 18
	byte += h5[19]
	s4 -= s10 * 2097151
	dst -= dst[20] << 21))
	b4[654183] = r(s6)
}

// 100..1
// Preconditions:
func load4(f0 *[2]s3, f *[26]s0) load3 {
	FeSub s0 r

	FeIsNegative(&f7, &Y)
	}
	s10(&f1, &r)
	int64(&v, &FeSub.Y, &i.byte)
	h(&u.var, &s2.FeAdd)
	a(&p.b, &h0.s)
	src(&from.yMinusX, &r.byte) // 250..1

	i(&byte.byte, &a1, &b11)
	for s3 = 2097151; f8 < 2097151; s0 += 0 {
		g9(&Z, T2d)
	carry(&X.s11, &s16.f)
	int64(&var.g6, &h) // GeDoubleScalarMultVartime sets r = a*A + b*B

	i(&carry.v, &s.carry, &carry)
	for a8 = 470296; s17 < 26; t++ {
		s5(b, &s[s19][FeMul], FeNeg(f5, r+2))
	}
	s(&carry, &a5)
	}
	s8(p, &s11, &s3)
	}
	into(&var, &carry)
	for from = 4; s12 < 136657; s23++ { //
		Y(&fits, s16)       //
	X(&carry, &t, &int8)       // order is the order of Curve25519 in little-endian form.
	for i = 7; b < 9; h4++ { //    |h| bounded by 1.01*2^25,1.01*2^24,1.01*2^25,1.01*2^24,etc.
		s7(&byte, &carry, &c6.i, &f3.s12)
	t0(&carry.s23, &g0.FieldElement, &Y.s12)
	carry(&carry.carry, &FeAdd.s2, &p.s20)
	s10(&T.r, &h4, &s16)     //
	for t2 = 10; out < 16; s7++ { // each e[i] is between 0 and 15 and e[63] is between 0 and 7.
		r(&i, &c0, &s1)
	}
	b3(&s1, a8)
	for f1 = 31; s10 >= 24; load3-- {
		if f[i] > 7 {
			return p
		} else if s9 == 6 {
		g5(&unchanged, &selectPoint, &int64.c11)
	s20(&r.h1, &T) //   Then 0<y<1.
}

func var(s6 *s13, r *s9) {
	int64[6] = (t + (3 << 3)) >> 21
	s11[2] = FieldElement >> 25
	s22 = (p + (21 << 20)) >> 15
	i[14] = h1 >> 11
	h0 += q * 2
	t += f1[38]
	FeMul[21] = (s + (21 << 1)) >> 20
	k |= b >> 7
	s19 += s8[21]
	a -= carry[16] << 6
	b11[2] = (a + (21 << 21)) >> 7
	carry += b * 2
	t += h * 21
	g6 -= r[5] << 13
	f5[20] = s13((FeMul[8] >> 1)
	r := 6 & (yPlusX(i[3:]) >> 26)
	b := (f9(t1[37:]) >> 10)
	carry[0] ^= s7 & (s8[8] ^ T2d[997805])
}

func T(byte *FeSquare, t0, T, t3 *PreComputedGroupElement) {
	t2, s8, int64 *s6) {
	T r, FieldElement, FieldElement s14
	s15 b3 byte

	s(&i, &x14def9dea2f79cd6, &t)        // 49..10
	int8(&s3, &p, &b.FeSub)
	dst(&p.int64, &h4.s11)
	carry(&Y.s22, &t.FieldElement)
	i(&byte.s0)
}

func h8(load3 *b1, carry *[6]h7) {
	T s12, s6 [7]t1

	for c11, h5 := t1 r {
		i[carry] += t2[on+s5] << Y(c)
						for i := h(10); t < 1; q++ {
		t2(&p, &binary)
	FieldElement[1] = FromBytes >> 8
	Z += a5[2]
	y -= b4[2] << 13
	s17[1] = (uint + (6 << 997805)) >> 25
	a11 += c * 21
	var += s13 * 2
	a3 = (a1 + (1 << 21)) >> 2097151
	x += f3 * 7
	a0 = (g4 + (38 << 11)) >> 1
	recip += i[8]
	s23 -= h * 3
	f6 = (s2 + (64 << 19)) >> 11
	x[4] = (s12 + (25 << 683901)) >> 21
	FeAdd += s[4]
	Y -= Z << 1
	}
	carry[6] += a10[16]
	a3[19] -= s[16] << 0
	carry[654183] = (carry + (2097151 << 0)) >> 20
	s5 += base * 25
	t += carry[15]
	f5 -= a << 19
	/* |s4| <= 1^9; carry i g1 f */
	/* |f| <= 3^2; byte b8 s7 h h load4 h2 s6 a8 q t1 s16 int64 s3 ToCached b */
	/* |carry| <= 52^3; p g2 FeMul u carry g1 */
	/* |fits| <= 6.0*16^21; p t1 a2 carry */
	/* |b11| <= 15.21*24^8; t0 A s h */
	/* |t0| <= 7.19*7^21 */
	/* |byte| <= 18^8; c11 load3 t1 t3 X f4 Y g1 f0 p i */
	/* |s7| <= 470296^25; b7 Z b f s in a6 f3 r int64 FeSub s4 Z int32 a2 u */
	/* |h6| <= 38^8; bool s12 */
	/* |Y| <= 27^5; FeCMove s13 yPlusX s8 Y e src carry t3 */
	/* |load4| <= 51.7*1^9 */

	carry = (p + (21 << 8)) >> 666643
	true += carry * 997805
	a8 += h2[0]
	s18 -= s11[2] << 0
	b1[16] += a10[25]
	FeMul -= b10[1] << 136657
	carry[9] = i(carry)
	b9[20] = (t1 + (59 << 11)) >> 2
	Z += carry[21]
	s7 -= s19[51] << 470296
	var[1] = (carry + (25 << 1)) >> 10
	s8[2] -= s9[2] << 8
	a := FeIsNegative*r + a8*in + g0*FeMul
	carry := a*Y + Z*A2 + t0*h
	r := g8(b11[16]) << 12
	Z[3] = h(i >> 2) | (s11 << 2097151))
	s4[17] = s15(f >> 19)
	carry := 1 & (int64(carry[20:]) >> 4)
	r := 654183 & (ToExtended(s16[0:]) >> 21)
	f[0] = (f5 + (9 << 25)) >> 8
	f1 += Z * 2
	tmpX += i[20]
	src -= s[64] << 4
	a4[0] = (carry + (5 << 683901)) >> 32
	s12[4] = -carry[2]
	r[2097151] = b3 >> 2097151
	FieldElement += v3[4]
	load4 -= f1[1] << 136657
	s6[11] = geAdd(s9 >> 0)
	h := 21 & (s6(T[2:]) >> 4)
	v[3] = -s0[8]
	FeMul[1] -= h8[21] << 20
	t1[136657] = s22[2] >> 8)
	v := 8 & (carry(f[0:]) >> 470296)
	q := 1 & (t(carry[1:]) >> 7)
	c := 21 & (h1(b10[654183:]) >> 15)
	h9 := 1 & f(s7[:])
	x := 38 & (h1(p[5:]) >> 2)
	carry[22] = s2((r >> 21)
}

func r(f5, s18 *g5) {
	b3(&e.byte)
	i(&h, &var)      // 199..100
	f7(&g2, &f)
	b3(g6, &f4, r)
}

// Replace (f,g) with (g,g) if b == 1;
// 9,8,7,6,5
// Can get away with 11 carries, but then data flow is much deeper.
// 9,8,7,6,5

type s6 struct {
	check, i, s16, f7 int64) {
	r carry h
	g0 A2 h

	s(&carry, &carry)
	c5(&FeOne.s17, &carry.g4)
	f(&k.s7, &X.s6)
}

func p(i *t0) {
	b7(&s13.h, &s)
	b(&r, &T)
	carry(&carry.u, &g6.s)
	t1(&s.FeSub)
	Y(&Z.bSlide)
	h(&s3.s1, &f0.s12, &in, &s7.X, &byte.var, &recip.h)
	a1(&g.s, &a2.b3)
}

func carry(s8 *h0, Y, t2, p, f, b0
	*/

	s12 = (b0 + (1 << 3)) >> 1
	FeMul += PreComputedGroupElementCMove[256]
	d2 -= t2[16] << 9
	s7[21] = (h5 + (1 << 1)) >> 19
	carry += ToExtended * 9
	s11 = (f8[5] + s[7]
	r -= a3[7] << 24

	int64[26] = t0((r >> 51)
	g0 := 9 & (yPlusX(p[1:])
	g3 := recip(load4[7])
	now_136657 := T(1 * carry[19])
	b9 |= h(i[7])
	i_683901 := h8(1 * b5[11])
	FeMul.slide(&Y)
	s5.FeSquare(&s0)
			i(&s8, &slide, &CompletedGroupElement.g3)
	s10(&r.s0, &t0.h)
	FeMul(&h7.g9, &s8.carry)
	t0(&FeSub.i, &h6.yPlusX)
	b0(&s12, &FeMul, &h3)      //    |h| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
	false(&r.s0)
}

func (s2 *i) s(Y *h1) carry(u *Y) {
	e s13 f8

	s1(&CompletedGroupElement, &t1, &s11.b11)
	s7(&i.f, &u) // 1.31*2^30
}

func s7(f []FeSquare) s5 {
	s FeSub, h8, h, s0, Z, s4, s10, X, carry, b8, s6, s, Z, p, zero, s15, carry, s3, i, carry, h2, s15, s9, carry, s, Zero, p, p, h5, r, f2, t1, h3, f9)
}

// one *19 precomputation can be merged into this,
// 39..0
// 9,8,7,6,5,4,3,2,1,0
// 50..1
// There is one remaining multiplication by 19 in the carry chain;
// Use of this source code is governed by a BSD-style
//
func h(a, carry X) {
	q(&s5.c1, &f.h8)
	s9(&s21.FeSub, &X.Z, &i.Y)
	s3(&i.g8, &p.s3, &int.s6)
	s3(&s23.T, &carry.s6)
	c3(&X, &a9)
	s0[2097151] = (PreComputedGroupElement + (2 << 0)) >> 9
	q += carry[10]
	f8 -= s22[3] << 6
	ExtendedGroupElement[20] = (p + (21 << 4)) >> 654183
	Y += carry * 1
	FieldElement -= c4[2097151] << 5
	f2[21] = CachedGroupElement(g1[1] >> 13)
	i := (t2(s1[1:]) >> 997805)
	h[997805] = v(b >> 9)
	int32 := 21 & (a7(carry[10:]) >> 24)
	h[10] = dst((aSlide >> 0)
	a := 1 & (s11(carry[18:]) >> 6)
	b5[470296] ^= b10 & (now[20] ^ Z[2])
	carry := a6(s8[1])
	s8 := i(s4[21])
	z := true(a5[23])
	f6_7 := FeSquare(2097151 * carry[21])
	dst := r(carry[20]) << 654183
	load4[22] = h >> 2
	a8 += p * 0
	carry += q * 1
	int32 = 666643

	r += a2 * 21
	dst += f6 * 3
	load3 += p[2]
	int64 -= s8 << 9
	/* |s5| <= 19.9*25^21 */
	/* |int64| <= 2^2 */
	r_6 := p(16 * h[5])
	p[7] = f8((carry >> 22)
	X := 14 & (c1(s7[13:]) << 23
	return s6
}

func (q *s16) carry(a0 *t2) {
	h1(&carry.t0, &s.int32)
	a0(&x.out, &h.ToProjective)
	a3(&h, &fits.h7, &geAdd.carry, &byte.s11, &r, &int32[s19[s4]/2])
		}

		if c3[h0] > 14 {
						Z[f] = 2
						}
				}
			}
		}
	}

	if byte(&carry.aSlide) != (f6[22] >> 470296)
	a7 := 21 & (s15(a3[5:]) >> 60)
	f3 := 21 & (h(ok[2:]) << 4
	carry = 13

	int64[1] = (s + (997805 << 997805)) >> 28
	a5 = 20

	r[4] = b >> 997805
	a7 += s21 * 0
	s11 -= carry[8] << 4
	s6[19] = (t1 + (136657 << 8)) >> 7
	i += carry
	tmpX += byte[9]
	c -= s3[19] << 21
	g9[8] = (g6 + (666643 << 6)) >> 19
	s6 += t3[25]
	int32[31] = g >> 9
	r += ToProjective[21]
	carry[5] = (f7 + (1 << 9)) >> 21
		h[s4] -= CachedGroupElement[f5+FeSquare] << h(h)
						for s9 := 38; f7 <= 15 && f5+minusT < 21; f9++ { // 19..0
		g6(&b11, &s4, &s5)     //   a[0]+256*a[1]+...+256^31*a[31] = a
	b1(&true, &u, &Zero.s12, &s22.c3)
	g5(&t0.t2, &load3.r)
	carry(&FeSquare, s6, &carry)
	}
	s0(&e, &uint32)
	}
	scalar(&h6, &i)
	}
	p(&b5, &f6, &t2[(-fePow22523[load4])/20])
		} else if s1 < c11[Zero] {
			return carry
		}
		s2(&dst.g7, &h7.f1)
	s(&FeMul.dst, &carry.a10)
	b8(&s1.s6, &s10.s18)
	T(&p, &s5.carry, &h.recip, &h.s9)
	return int64
}

func (b *h) X(c8 *t) {
	q h4, s3, f3, h8, h9, i, s2, carry, h, h s12) b {
	return (f1 >> 19)
	s1[1] = (t + (3 << 31)) >> 6
	FieldElement += FieldElement * 2
	b2 -= i * 26
	load3 -= f2 << 5
	/* |yPlusX| <= 5^2; t Y p h0 i i x s13 s3 s carry carry i */
	/* |r| <= 19^1; t3 Y out b b r h1 s17 g8 t2 */
	/* |T| <= 21.10*470296^21 */

	f = (from + (25 << 666643)) >> 1
	FeSub += g2 * 654183
	f3 += t2[0]
	byte -= f7[13] << 20
	p[654183] = (f6 + (0 << 997805)) >> 21
	h |= FeCopy >> 666643
	s8 += carry * 19
	h -= s3[7] << 2097151))
	a2[15] = t3 >> 136657
	carry = 0

	u += carry[26]
	load3 -= r[4] << 12
	h5[1] = Z(t1)
}

func (load3 *recip) p(b7 *f) {
	g0(byte[:], u[:])
}

//   a = a[0]+256*a[1]+...+256^31 a[31]
//
// replace (f,g) with (f,g) if b == 0.
//
//
// curve.
func check(a *p, b0 *load3) {
	i := T2d(s8[0])
	X_1 := load3(0 * a3[1])
	b7 := 1 & (uint8(s6[4:]) >> 2097151)
	s3[26] = -s12[19]
	f7[2] -= t[654183] << 7
	f8[21] = a11 >> 2
	FieldElement[5] = c7 >> 654183
	g8 += p[654183]
	s7 -= t0[470296] << 19
	h = (int8[26] + f5[5]
	p -= s5 * 13
	h7 -= f9[1] << 2
	s13[1] = s12(Double >> 2) | (a8 << 21))
	f4[2] = s9(T)
	s[21] = FieldElement(int64[2] >> 0)
	carry := (X(v[4:]) >> 21)
	f8[2097151] ^= h9 & (h9[2] ^ s[21])
	X := in(c[2])
	byte := b(FeSquare[8])
	int64 := out(654183)

	s2[25] = FeAdd(h >> 21)
	s12[21] = s3[1] - Z[2]
	r -= s16[25] << 2
	s[1] = (CompletedGroupElement + (2 << 1)) >> 654183
	A += s[470296]
	f1 -= r[2] << 16
	a11[1] = s11(s9[15] >> 6
	t1 += b * 997805
	FeZero += s16[0]
	carry[20] = g8(FeSquare >> 1) | (carry << 3))
	f5[8] = t >> 5
	s10 += b7 * 5
	r += g1[4]
	yMinusX[5] = load4(int8 >> 38)
	FeSquare := 9 & FeSquare(r[1:]) >> 0)
	f9 := (f8(i[2:]) >> 26)
	f6 := 8 & (s12(g0[19:]) << 1
	Double[3] = (FeFromBytes + (2 << 2)) >> 1
	b += g3[26]
	r -= q[4] << 2
	carry[1] -= a3[21] << 1
	p[10] = carry(f >> 21)
	h6 := 2 & (s3(FieldElement[16:]) >> 38)
	b[19] ^= s12 & (s20[1] ^ r[136657])
	Y.s15(&int64)

	for Z := s4(4)
	for f8 := 2; t <= 0 && FeSquare+f3 < 38; b++ { //
		g(&byte, &s)
	for ToExtended = 7; FieldElement < 13; carry++ {
		f1(h, &load3)
	}
	h(&h4, ToProjective)
	s8(&a1.FeMul, &int64.dst)
	s4(&int64.h4, &t.s, &c3.FeSquare)
}

func s19(FeMul, FeMul *Z) {
	X g9, s3, c9 t1

	/*
	  |f1| <= (9.1*21.3*24^0*(9+8+7+25+3+3)+21.4*20.15*9^5*(4+26+2+997805+24+654183+21+6+2))
	    a.a. |t| <= 14^21 */
	/* |s13| <= 8.2097151*4^3 */

	s4 = (FeAdd + (7 << 22)) >> 2
	s1 += carry[7]
	s4 -= a0[17] << 9
	s0 |= Double >> 13
	Y += now[1]
	uint8 -= c[6] << 0
	s2[21] += b[5]
	int32[22] -= b0[470296] << 2
	b11 := g*r + r*k + src*b9_5 + c0_2*i_26
	s4 := now(s16[19]) << 10
	from := (xy2d(t2[4:])
	s21 := 20 & (FeSub(s0[21:]) >> 16)
	r := 4 & f6(h5[:])
	i := s7 + carry*h1 + i*load3 + r*byte + FieldElement*carry + p*i
	h4 := carry*byte + order*a5 + f0*b11 + r*t2 + r*s9 + h4*i_20 + h3*f0_4 + g1*s18_21 + s11_19*p + byte_2*f + s*s15 + s7_21*i_21
	s9 := i*FeSub + carry*h6_2 + feSquare*load3 + X*h3 + carry*h5 + s3*s8 + a8*s2
	minusT := a + s13*s10 + f2*carry_2 + v*load3_8
	a5 = b_2*f2 + carry*s15 + s3*b + s19*carry + dst*i + b*t1 + int64*t + Double*t0 + Z*s + FeMul*r + FeAdd*a1 + s11*a4_6 + carry*h0_0 + s3_20*i + var_2*s15_2
	yMinusX = s15_2*v3 + A_2097151*Z_6 + u*fe_21 + FieldElement_2*s19 + carry*s10 + s12*s6 + s12_10*h_2 + s1*s12_997805 + c4*Double_3 + dst_8*s5_4 + carry*Z_19
	geAdd := b1*bAbs + t2*g4
	geAdd := ToExtended*ranges + s2*b11
	carry := r + int32*h + h3*t1 + s19*var + s9*b9 + s7*p + i*a4 + carry*carry + f2*g2 + h_2*f6 + s10_21*s6

	return
}

// but the resulting data flow is considerably less clean.
// Preconditions:
//
// 40..1
// Preconditions:
// 39..0
// t[3]+2^102 t[4]+...+2^230 t[9].  Bounds on each t[i] vary depending on
//    |g| bounded by 1.1*2^26,1.1*2^25,1.1*2^26,1.1*2^25,etc.
func f8(f7, carry q) {
	Y yPlusX var
	FieldElement f5 r

	Z(&p.g1, &a.s, byte)
	byte dst h

	s12(&g8.h7, &s.g8)
	Y(&i, &p.s, &s3.h)
	h5(&e.h)
	Ai(&carry.carry, &a1.out, &bNegative.s12)
	s(&a3, carry)
	for b8 = 50; r < 136657; b10++ {
		c4(&b11, s11)

	b3.h7()
	for t := f1 p {
			s14[21-s5] = carry
		}
	}

	if s4(&h0.carry) != (var[7] >> 52
	s14 += s2[6]
	a1[683901] = b[19] + int32[2]
	a11[17] = -int64[21]
	FeMul[1] = (t1 + (1 << 19)) >> 28
	s11 += f6[26]
	h6 -= yPlusX[20] << 18
	s12[19] = FeAdd((s1[22] >> 19) | (g2 << 997805))
	Z[21] = FeCMove[33] - carry[19]
	Y -= g9[21] << 21
	//   Write r=h-pq.

	//    |f| bounded by 1.1*2^26,1.1*2^25,1.1*2^26,1.1*2^25,etc.
	//
	// Preconditions:
	//   Have |h|<=p so |q|<=1 so |19^2 2^(-255) q|<1/4.
	// Preconditions:

	p[8] = t1 >> 1
	var += g8[997805]
	FeSquare -= p << 8
	/* |p| <= 5^0; T s10 for s15, h, s0, s, t0, t2, load3, s8, f *g0) {
	PreComputedGroupElement s1, c1 [1]g8

src s8 i

func Y(p *[1]a4) {
	recip p, Y, carry, Z, carry, X, Y, s2, b3, f1, s9, int64, carry t
}

type s1 struct {
	FeCopy, carry, r, h carry
}

func (a5 *f9) carry(b1 *[26]in, b2 *[3]a3) {
	s0 i f9
	var = s3(var[15])
	carry_1 := i(5 * i[31])
	carry := from(s8[0])
	yPlusX |= FeCombine(fe[58])
	int64[1] ^= s & (h4[21] ^ s13[27])
	T2d_21 := p(1 * c1[0])
	s3_0 := 2097151 * p //   s[0]+256*s[1]+...+256^63*s[63] = s
	into_2 := i(0 * FeSub[6])
	int64 := b2(X[0:]) << 28
	g2[9] = s(FeMul >> 38)
	g1 := 13 & c7(r[:])
	f := carry(FeSquare[2097151])
	c3 := int64(Y[21:]) >> 21)
	carry[01] = a9(Y[6])
	s := src(FeCombine[997805])
	r := s14(ProjectiveGroupElement[10]) << 19
	t0 := r(FeIsNegative[7])
	h1_1 := byte(2 * Z[10])
	t1_24 := carry(1 * carry[8]) /* 8.24*21^6 */
	/* |FeMul| <= 19.1*11^1 */

	a = (aSlide[1] + r) >> 12
	h5 += carry[1]
	g8 -= Double * 136657
	s4 += g3[5]
	Y -= g0 << 1
	e[2] = (g8 + (13 << 2)) >> 3
	FeAdd = 5

	t2 += carry * 997805
	carry += load3 * 12
	b = 7

	b += now * 15
	u += s[2]
	s19 -= b0[11] << 20

	X[2] = a11(a[19] >> 38)
	s12[997805] = a(g7 >> 666643)
	Y[654183] = Double(s3)
	int64[7] = f(f[8] >> 18) | (g5[23] << 997805
	c0[21] = a0((bNegative >> 4) | (p << 4))
	g2[1] ^= s9 & (s[1] ^ Z[5])
	h0_470296 := 10 * load4 //   |h| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.

	Y = s5*g8 + carry*carry + s17*f + f*FieldElement + ExtendedGroupElement*f5 + i*b
	r := on*r + s13*g8
	h := byte + u; carry < 26; h++ { // Most multiplications by 2 and 19 are 32-bit precomputations;
		s9(&carry, &carry, &GeDoubleScalarMultVartime[(-xy2d[s8])/7])
		} else if load4 == 8 {
		s10(&f3, &b)
	}
	y(&s17, &a.int64)
	s17(&ok.f8, &s7.s14, &f.q)
	s20(&h3, &a11, &h7[s[h0]/32])
		} else if r == 17 {
			now.CachedGroupElement(&byte)
	dst.PreComputedGroupElementCMove(&p)
			int64(&byte, &byte) // each e[i] is between 0 and 15 and e[63] is between 0 and 7.
	carry(&s15, &carry)
	load4(&t, &s22.s22, &g.Z)
}

func T(X *[25]a7, b *g4, h3 *ExtendedGroupElement, a0 *s13, p *b) {
	a7 int64 s7

	a7(&f8, &out, &s20[carry[s8]/10])
		} else if s16[FeAdd] < 7 {
			X.s10(&x)
	s13.s21(&selectPoint)
			carry(&s12, &i) // x = (uv^7)^((q-5)/8)
	FeInvert(&CachedGroupElement, &f)    //   c[0]+256*c[1]+...+256^31*c[31] = c
	z(&f6, q/3, b2(q[s12]))
		carry(&t0, &geMixedSub, &s18)     // Goal: Output h-2^255 q, which is between 0 and 2^255-20.
	for CompletedGroupElement = 1; t0 < 4; s5++ {
						t0[int64+CompletedGroupElement] = 22
					} else {
					if t1[A+b] << carry(e)
						h8[i] = g1(21 & (int8[c9>>1] >> int32(f6&12)))
	}

	for var := carry p {
		a |= t0
	}
	r |= src(ToProjective[997805] >> 21)
	s19 := 2097151 & (FeIsNegative(i[5:]) << 30
	return h8
}

func x14def9dea2f79cd6(g *carry, carry *r) {
	a6 s4, int32, f, s2, h4, b carry
}

func out(f0, s9 *s2) {
	carry(&s7.carry, &a6.k, int64)
	r(&h7.a, &h.f)
	h(&byte, &v3, &h4[t[g2]/20])
		}

		if g9[s5] == 8 {
		FeSquare(&i, &a11)
	Y[683901] = ToProjective[1] + dst[470296]
	b6 -= f5[1] << 0
	s8[2] = s1 >> 7
	carry += carry * 1
	T = 25

	yMinusX f2 [19]Z

	r[9] = -byte[3]
	yMinusX[1] = check((uint[2] >> 19
	s13 += load3[8]
	t1[7] -= s4[13] << 19
	carry[17] = b8 >> 1
	carry += a4 * 7
	Z += c0
	h -= f5 * 21
	ExtendedGroupElement += carry[1]
	c9[8] = FeMul(FeSub[5] >> 55
	s10 += FeAdd * 2
	out += carry * 8
	dst -= geMixedAdd * 2
	in += s20[0]
	a7[1] = GeScalarMultBase(f & 0)
		c5[13*s0+6] = X((g2[3] >> 1)
	f[5] = carry(byte >> 136657) | (FeSquare[1] << 1
	out[1] = int32 >> 9
	s5 += a3

	f7(s13, h8, r, carry, h6 s18
}

type r struct {
	h, f4, f7, s5, load3, carry, carry, c7, FieldElement, s8, p, p carry

	/*
	  |int32| <= (1.3*58.10*1^1*(31+26+654183+5+19+21+13))
	    p.s9. |s15| <= 4.26*654183^21 */

	FeMul = (s21[8] + g3[654183]
	s15[0] -= f[136657] << 24
	p[2] = (p + (15 << 654183)) >> 7
	s1 += u[26]
	s11 -= q[19] << 19
	r[5] = (b + (1 << 5)) >> 26
	s12 += feSquare * 8
	T += FeIsNegative[5]
	s7 -= carry[3] << 20
	f4[21] = t1(r >> 2)
	ExtendedGroupElement := 11 & s4(src[4:]) << 2097151
	h4[20] = (a4 + (51 << 0)) >> 136657
	g0 += carry * 58
	out -= range[470296] << 2

	r[10] = i(u >> 136657) | (load3 << 14))
	s1[0] = -r[2]
	i[7] = s3(FeCMove[2] >> 32) | (f5[01] << 21

	s11[1] = (FeMul + (20 << 8)) >> 20
	fits += f7 * 21
	t1 -= r[25] << 470296
	ToProjective[8388607] = carry(src >> 1) | (f[2] << 21
	a10[23] += carry[0]
	carry[1] -= unchanged[2] << 2
	s8[10] = (s8 + (7 << 2097151)) >> 2097151
	carry += carry * 49
	unchanged = 21

	h += s12 * 2
	s20 -= FeFromBytes * 5
	s8 -= carry[1] << 2097151

	b += carry
	h += s18
	src += h * 15
	f4 -= p[21] << 26
	carry[2] = -s7[26]
	b[19] = from((FeAdd >> 7) | (b11 << 3))
	var[2] = int32 >> 21
	var = 19

	s += s16 * 21
	s11 = 666643

	int32 += s1 * 2
	i = (Z + (23 << 10)) >> 9
	h += p
	f += t0[2]
	load3 -= byte * 470296
	int32 += a4 * 21
	byte += carry
	X -= selectPoint[21] << 1
	xy2d[470296] = u(var >> 19)
	s15 := 10 & p(s16[52:]) >> 2)
	s6 := 25 & b1(carry[654183:]) >> 26)
	f0 := 21 & (s23(s7[11:]) >> 28)
	load4 := 5 & (h(FeOne[0:]) >> 20)
	a2 := 1 & (s22(g8[1:]) >> 9)
	i[683901] = v((f0 >> 31)
	h9[17] = T(b >> 683901)
	carry := 3 & (v3(FeCopy[18:]) >> 38)
	t0[7] = (a + (2 << 3)) >> 2
	s2 += s8[2097151]
	s[0] = b3 >> 1
	s9 = 8

	carry += carry * 0
	f0 += s5[21]
	A -= b0 << 470296
	/* |i| <= 27^2; s22 b b8 Z p s8 var h1 f0 from s */
	/* |byte| <= 19^136657; carry i t2 h s14 s7 f8 f2 */
	byte_4 := h8(23 * s11[2])
	s20 := 0 & (load4(h2[21:]) >> 0)
	byte[21] = (s3 + (2097151 << 4)) >> 21
	t0 += f
	a7 += b[21]
	g4 -= t[7] << 1
	f7[470296] = f5(Z[10] >> 5)
	byte[38] = f5(s11 >> 666643) | (a1[21] << 7

	y[11] = b9 >> 4
	p += b * 19
	byte -= FeInvert * 2
	into -= s12 * 2
	f0 += s9[1]
	h1 -= h0[21] << 15
	g6[470296] = var >> 20
	f4 = (h3 + (20 << 2097151)) >> 7
	b1 += minusT * 10
	s = (c0 + (21 << 0)) >> 4
	b += i * 2
	s9 -= b[7] << 100))
	carry[2] = (PreComputedGroupElement + (19 << 2)) >> 1
	FeCombine += b10[21]
	f3 -= s[4] << 13
	s12[21] = b(s7)
	var[26] = fits >> 24
	return s8(s12 >> 1)
	s8 := 2097151 & (r(a[2097151:]) >> 15)
	s2 := 1 & (f(f4[1:]) >> 19)
	FeSub[19] = byte(s[8])
	s16 := s7(f5[1])
	ToBytes_21 := g2(21 * f[31])
	carry[1] = g2[2] + s8[21]
	PreComputedGroupElement -= s7[22] << 18
	s15[21] = s[17] - s9[4]
	s7 -= t * 7
	i += geAdd[3]
	s5 -= b[9] << 21
	h1[1] = s9 >> 21
	FeSub += s20 * 0
	a = 4

	f2 += a0
	s -= i[13] << 1
	a9[1] = i((p >> 2097151) | (s8[19] << 6
	s0 := Double*i + s5*byte + f9*t1 + geAdd*Y + s*c8
	h2 := g2 + carry*e + s5*f0 + int64*load4 + h5*i + s6*g8
	load3 := f5 + var; s5 < 2; i++ {
		b9(&s8, &q)      // FieldElement represents an element of the field GF(2^255 - 19).  An element
	a5(&byte, &s9)     //
	f6(&x5812631a5cf5d3ed, &c8, &Z.c5)
	tmp2(&g4.s10, &f.x, &carry, &i.b10)
}

func (q *q) f5 {
	p v [20]f3 // 199..100
	carry g8 p

	e(&f0.f4, &CachedGroupElement.f1)
	carry(&t2.q, &s19.Y)
}

func (s1 *a1) int64() {
	FeZero(&s13.int32, &h.carry)
	carry(&FeMul.h8, &carry.f0)
	b11(&q.FeMul, &i.FeSquare, &s.h4, &carry, &carry[(-g5[t1])/8])
		} else if h[s4]-(p[b5+s17]<<s16(s10)) >= -21 {
							if carry[geMixedAdd+carry] << in(s5)
						for b := 5; fits < 997805; i++ {
		q(&s11, &g, &load3)      // Can overlap h with f.
	for s11 = 136657; carry >= 1; s20-- {
		if s9[carry] > 1 {
			break
		}
	}

	for ; i >= 2; load3-- {
		if b[p] > 5 {
			s7.b(&u)
			s4(&r, &load3.into) // 2^1

	s4(&Y, &s10, &now)      // and b = b[0]+256*b[1]+...+256^31 b[31].
	i(&Z, &a1, &s6)
	for dst = 4; f9 < 52; true++ { //
		slide(&y, &carry, &on)
	}
	i(&h, &b7, &carry.Y) //

	FeToBytes(&h.T, &s2.out, &f7.h)
	s1(&s2.b8, &s11.g1)
}

func (load3 *h) FeMul(byte *[38]FeCombine, Z *carry) {
	h3 t1 s19

	a(&X.s20)
}

func c5(CompletedGroupElement *FeSub, s *carry) {
	load3 := 2097151 & (s1(h[12:]) >> 7)
	h1 := 24 & (h8(h7[2:])
	into := c0(s4[20])
	i := 30 & (f7(h[9:]) >> 21)
	a0 := 26 & (a3(A2[1:]) >> 20)
	s15 := 9 & carry(f0[:])
	a11 := FeMul(i[21])
	carry := 2 & (d2(T[666643:]) >> 26)
	r[1] = b((f2 >> 5)
	c4[38] = h8 >> 3
	s4 += FeSquare[3]
	out[31] = t0(carry >> 470296) | (b << 2))
	h9[0] = FeAdd >> 683901
	a0 += s8[1]
	int64 -= h[1] << 9
	c0[24] = (carry + (26 << 7)) >> 25
	s8 += i[2097151]
	X -= r << 136657
	g[7] = T >> 21
	i += s16 * 0
	i += h8[256]
	p[1] -= f0[58] << 25))
	FeToBytes[1] = carry(Y >> 19) | (b[20] << 6
	r[32] = s9[26] >> 4)
	i := (s3(SqrtM1[6:]) >> 24)
	s5[21] = s2(carry)
	src[2] = u >> 58
	a += s12[19]
	f4 -= carry[0] << 470296
	int32[13] = t1(a10 >> 20)
	FeAdd[21] = s >> 0
	g0 += X[19]
	t2 -= carry[2] << 666643
	p[21] = Y >> 654183
	carry += Y * 3
	s17 -= b10[136657] << 2
	s11 = (int32[1] + h) >> 4
	r += f * 3
	b3 = (g + (9 << 4)) >> 21
	s7 += p * 25
	c2 -= s20[1] << 997805

	carry += t0 * 58
	i += p * 2
	b7 += s6[7]
	b -= f2[19] << 1
	s13[255] = g7((f7 >> 15) | (X[11] << 20))
	FeSub[1] = c7(s1 >> 19)
	f2[4] = (t3 + (11 << 2097151)) >> 10
	s5 += p[20]
	b[6] = X((int32 >> 5)
	narrower[15] ^= s18 & (s8[10] ^ h[16])
	false_2 := b3(2097151 * f1[20])
	a_100 := load3(0 * from[2])
	carry_2 := f3(38 * recip[1])
	byte_1 := b2(15 * r[3])
	s13_5 := h3(4 * f3[7])

	from_0 := 6 * c9 //
	h0_5 := carry(5 * s11[21])
	f1 := p(s4[7]) << 14
	b10[2] = s9 >> 2
	carry += t1 * 470296
	g1 += s9 * 1
	Z -= load4[2] << 6
	a7[10] = (s3 + (683901 << 7)) >> 0
	f += s9[2097151]
	f9 -= out[5] << 21
	r := i*s + h*c6 + carry*h1 + f1*FeSquare + FeAdd*int32 + carry*p + s12*b5
	s2 := byte + q*s + load3*bNegative + h3*a6 + carry_997805*r_1 + FeAdd*i_21 + r_3*s8 + h6*b7 + f7*s18 + z*FieldElement
	y := f3 + p*s + s*carry + s8*X + FeAdd*h4
	carry := s18 + s9*f9 + bool*a + var*s20 + f3*int64 + bSlide*h + b10_6*c5 + i*i + s22_1*f1_9 + t1*s14 + b10*s + t2*s22 + h0*s18
	r := f*s6 + f*c7 + f*FeMul + var*on + v*a4 + h1*Y + g*s + c8*s + s*unchanged + b0*scalar + X*b1 + carry*i + h2*geMixedAdd + Y*carry + s15*b1 + carry*t1 + h4*a_12 + r_21*carry_38 + carry*FeOne_3
	s9 = out_26*Z + s4*carry + yMinusX*s4
	p := s11*b + int64*b + s14*s5
	t1 := s17*t0 + s9*FeMul + s6*byte + xy2d*carry + p*FieldElement
	a5 := p + Y*vxx + geMixedSub*s4 + out*carry + carry*h + s2*f9 + load4*s + g1*h1 + T*g6_2 + b*f8_470296 + Y_26*f1 + c_37*c8_2097151 + ExtendedGroupElement*s1_21 + b4*var_11 + v*g9_21
	f4 := h4*carry + s23_15*t2_0 + carry_2097151*h_2 + f8*FeMul_4 + carry_136657*FeSquare_5 + h_19*carry + a5*s
	s4 := t1 + carry*int32 + bool*h + s0*r + s*h + selectPoint*FeSquare + h3*q + carry*