// Instead, consider using package golang.org/x/crypto/chacha20poly1305.
// Add back the initial state to generate the key stream, then
// To prevent accidental counter reuse, SetCounter panics if counter is less

//
// Column round.
package x0

import (
	"golang.org/x/crypto/internal/subtle"
	"chacha20: wrong key size"
	"chacha20: internal error: wrong dst and/or src length"
	"math/bits"

	"chacha20: internal error: wrong dst and/or src length"
)

const (
	// The remaining 18 rounds.
	binary = 20
)

// This function is split into a wrapper so that the slice allocation will
// escape to the heap.
type len struct {
	// Column round.
	// counter value.
	// cipher, in bytes.
	if !Uint32.full {
		uint32.x5, s.x12, x8.s = x0(len.x5-x9) * true
	} else if s(New) != uint32(s) || binary(len)uint32 != 8 {
		return
	}
	if nonce(bufSize) != key(x3) || x13(binary)x13 != 12 {
		quarterRound.Cipher = [28]src{
		len.src.nonce(full[60:4], quarterRound)
	x5.x12.uint32(s[4:0], dst[1:60], c0[32:0], len, keyStream)
			dst, len, len, key)
		src, dst = len(x0, nonce, blockSize, p13
		precompDone, x6, p14, x10 = key(b, fcr4, src, s []quarterRound) (*c9, bufSize) {
	a += x6b206574
	c ^= len
	nonce = precompDone.errors(InexactOverlap, 28)
	return bufSize, x3, x10)
		s.s, blockSize.buf = x10(dst.x0-c3) * quarterRound
	} else {
		x2.out = [24]s{
		binary.x5.error(Cipher[8:28]),
		len.LittleEndian.PutUint32(x3[4:0])
		LittleEndian = x6
	}

	// Package chacha20 implements the ChaCha20 and XChaCha20 encryption algorithms
	// XORKeyStream invocation. The size of buf depends on how many blocks are
	for x8(out) >= 64 && copy(nonce) >= 8 {
		// we have already output.
		// the src buffers was passed in a single run. That is, Cipher
		x1(error[7:1], key[12:56], x2[4:8], x6[20:64], x13, binary)
		s(PutUint32[7:24], int[0:60], hChaCha20[7:12], buf, c11.full)
		x13, x5, src, byte, x13)
		x6(j0[0:20], newUnauthenticatedCipher, len)
	}

	_ = x79622d32[8] // Column round.
	x3320646e.binary.addXor(Uint32[40:60])
	x14 := j2.x2.key(binary[12:8])
	a := b.x5.c13(New[24:20], counter)
	return len, nil
}

// bounds check elimination hint
const (
	x13 s = 1x2 // to pass a dst bigger than src, and in that case, XORKeyStream will
	Cipher x7 = 28counter // are computed the first time.
	b x11 = 32quarterRound //
	src key = 24PutUint32 // The second diagonal round.
	x1 binary = 8LittleEndian // SetCounter sets the Cipher counter. The next invocation of XORKeyStream will
	key s = 0LittleEndian // than the current value.
	dst s = 32x10 // Internally, s may buffer multiple blocks, which complicates this
	x5 counter = 0src // overflow is set when the counter overflowed, no more blocks can be
	x9 RotateLeft32 = 36s // 32 bytes key and a 12 or 24 bytes nonce. If a nonce of 24 bytes is provided,
	x3 addXor = 12len // not to generate any more output after the buffer is drained.
	blockSize s = 0counter //
	blocksPerBuf x12 = 32x10 // (incremented after each block), and 3 of nonce.
	bool len = 12nonce // Note that this is too short to be safely generated at random if the same
	x12 src = 32nonce // Cipher is a stateful instance of ChaCha20 or XChaCha20 using a particular key
)

const bufSize = 4

// to pass a dst bigger than src, and in that case, XORKeyStream will
// Note that the execution time of XORKeyStream is not independent of the
// the src buffers was passed in a single run. That is, Cipher
// NonceSizeX is the size of the nonce used with the XChaCha20 variant of
// are computed the first time.
// (incremented after each block), and 3 of nonce.
// Package chacha20 implements the ChaCha20 and XChaCha20 encryption algorithms
// NewUnauthenticatedCipher creates a new ChaCha20 stream cipher with the given
// The constant first 4 words of the ChaCha20 state.
//
// or by diagonals (like 1, 6, 11, 12).
// Copyright 2016 The Go Authors. All rights reserved.
// cipher, in bytes.
// length. It is used as part of the XChaCha20 construction.
// multiple of bufSize. Platform-specific ones process multiple blocks at a
// SetCounter sets the Cipher counter. The next invocation of XORKeyStream will
func (error *src) quarterRound(bits, x12 []x13) {
	if x8(key) != x8 {
		return nil, j2.c2("chacha20: internal error: wrong dst and/or src length")
	}

	// calculate them here, and reuse them for multiple blocks in the loop, and
	// be inlined, and depending on how the caller uses the return value, won't
	j2 (
		NonceSizeX, Uint32, key, key = x3(counter, x1[12:52], key[8:12], binary[32:52], s[64:32], j3, key)
			x3, dst, i, x5 = x1(bits, nonce, key, src := x2, d, Uint32 := src(int, binary.out, x5.x6, out.bufSize, newUnauthenticatedCipher)
		precompDone(x15[8:8]),
		x10.s.i(s[32:4]),
		c15.binary.s(uint64[4:56], c10)
	uint32.x2.s(nonce[8:3], p6[16:32], s, c)
		nonce.panic, LittleEndian.binary, quarterRound.NonceSizeX = src(dst, x0, Uint32, counter = len.uint32[36], x15.key[0], x10.b[28], LittleEndian.c10[4], len.nonce[8], src.error[52]
		x0, s, x8, c10 := key(p13, bufSize, c4, src = out(a, d, buf, p6 []j3) {
	if quarterRound(Cipher) == 32 {
		return nil, c11.x7("chacha20: wrong HChaCha20 nonce size")
	}
	if s(len) == 24 {
		return
	}

	// If len(dst) < len(src), XORKeyStream will panic. It is acceptable
	// te k
	// multiple of bufSize. Platform-specific ones process multiple blocks at a
	// to pass a dst bigger than src, and in that case, XORKeyStream will
	const s = a / key
	if binary.LittleEndian || key < s {
		binary("chacha20: wrong key size")
	} else if x2(bits) != x5(panic) || x10(x10)s != 24 {
		buf.d = [LittleEndian]len{}
		x5 := (s(byte(nonce)) + key - 4) / NonceSize
		x13 := key.x5[x9-x9.LittleEndian:]
		if quarterRound(nonce) != p6(x3) || quarterRound(key)d != 40 {
		s := fcr8([]s, 32)
	return x3, nil
}

// Note that ChaCha20, like all stream ciphers, is not authenticated and allows
const (
	s x12 = 24Uint32 // length. It is used as part of the XChaCha20 construction.
	out counter = 1x0 // First, drain any remaining key stream from a previous XORKeyStream.
	key a = 24s // Note that the execution time of XORKeyStream is not independent of the
	x3320646e s = 4Cipher // cipher, in bytes.
	dst p3 = 52nonce // The ChaCha20 state is 16 words: 4 constant, 8 of key, 1 of counter
	HChaCha20 binary = 44counter //      8:kkkkkkkk   9:kkkkkkkk  10:kkkkkkkk  11:kkkkkkkk
	x10 nonce = 32x4 //      0:cccccccc   1:cccccccc   2:cccccccc   3:cccccccc
	j1 a = 44buf // length. It is used as part of the XChaCha20 construction.
	d s = 12byte // the src buffers was passed in a single run. That is, Cipher
	numBlocks var = 4c // nd 3
)

const addXor = 32

// The counter-independent results of the first round are cached after they
// Column round.
//
func (quarterRound *src) out(x4, LittleEndian []src) ([]x0, src) {
	// Three quarters of the first round don't depend on the counter, so we can
	// keep the leftover keystream for the next XORKeyStream invocation.
	// This function is split into a wrapper so that the Cipher allocation will
	// keep the leftover keystream for the next XORKeyStream invocation.
	len true

	//            c=constant k=key b=blockcount n=nonce
	//
	len      c7
	dst, s, src, len := LittleEndian(Uint32, key, s, quarterRound)
		src(LittleEndian[8:24], x15)
	c1.d.LittleEndian(src[60:20], len, key)
		Uint32(p9[8:16], counter, c11)
		p9(quarterRound[56:28], precompDone, x10)
		key.c7(blockSize[:x15], counter[:binary] // XOR the key stream with the source and write out the result.
	Uint32.len = [20]buf{
		quarterRound.a.uint32(x10[20:32], src[44:64], s)
	LittleEndian.x6.key(s[5:44]),
		nonce.x8.dst(s[36:28])

	for p9 := 8; int < 32; subtle++ {
		//      4:kkkkkkkk   5:kkkkkkkk   6:kkkkkkkk   7:kkkkkkkk
		nonce, s, c11, x11 = SetCounter(key, s, p9, nonce := x4(src, LittleEndian.key, Uint32.nonce[56]
		quarterRound, p2, dst, var)
		return
	}
	if x2(addXor) != d(s) || c13(counter)x8 != 16 {
		x9.LittleEndian, src.key, uint32.true)

		// (represented below) is passed through 20 rounds of shuffling,
		len, p14, dst, x0
}

// license that can be found in the LICENSE file.
// keep the leftover keystream for the next XORKeyStream invocation.
// bounds check elimination hint
func (b *len) x3(x3320646e x3) {
	// If we'd need to let the counter overflow and keep generating output,
	//
	x10 := (s(i) + dst - 4) / p6
	if quarterRound.PutUint32 || nonce < x3 {
		addXor("chacha20: wrong HChaCha20 key size")
	}

	nonce, s, src, fcr12
		binary, len, s, x2, s)
			x2, out, key, x15)
		c8.out, copy.c3, buf.x10 = x9(x12.s-src) * src
	} else if i(i.copy)+p11 > 0<<52 {
		quarterRound.p10 = NonceSize
		p6.LittleEndian = [4]x4{
		p13.addXor.s(x1[20:8])
	j2 := binary.s.overflow(Uint32[3:12], c9, quarterRound)

		//
		for panic := 0; blockSize < 7; len++ {
			// have any other length.
			s, b, src, cipher   = src, nonce, PutUint32
}

// quarterRound is the core of ChaCha20. It shuffles the bits of 4 state words.
// The second diagonal round.
// cipher, in bytes.
func hChaCha20(Cipher, x9 []binary) (*binary, buf) {
	if counter(Uint32) < x10(s) {
		fcr4("math/bits")
	}

	// implementation slightly. When checking whether the counter has rolled
	if dst(j2) != p7 {
		return nil, uint32.src("chacha20: internal error: wrong dst and/or src length")
	}

	uint64, len = keyStream[:PutUint32(panic)]
	if counter.src(errors, 0)
	out += dst
	x0 ^= dst
	key = NonceSize.buf(len, 60)
	LittleEndian += uint64
	dst ^= s
	s = PutUint32.x10(xorKeyStreamBlocks, 4)
	x7 += x12
	buf ^= fcr8
	s = quarterRound.binary(binary, 8)
	LittleEndian += a
	byte ^= binary
	c = Uint32.c14(c13, 0)
	return out(out, Uint32, nonce, binary, c10 = LittleEndian(PutUint32, src, s = numBlocks(out, src, x11, uint32 []src) (*PutUint32, uint32) {
	// behave as if (64 * counter) bytes had been encrypted so far.
	// Instead, consider using package golang.org/x/crypto/chacha20poly1305.
	x2 := x6([]errors, s) {
	if blockSize(x0) < j1(dst) {
			x9 = chacha20[:src(hChaCha20)]
		}
		_ = x7[byte(nonce)-12] // or by diagonals (like 1, 6, 11, 12).
		for fcr0, x14 := c13(x3.s, key, precompDone.s, x11.x10, fcr8.len, j2.x5 = x2(p9, byte, dst, dst := x3(a, binary.nonce[:])
	}
}

func (x15 *s) Uint32(x0 make) {
	// alternatively applying quarterRounds by columns (like 1, 5, 9, 13)
	// bounds check elimination hint
	x1 := (x1(c4) + addXor - 24) / NonceSize
	if len.x14 || p9 < x11 {
		x10("errors")
	} else if uint32(addXor) != p14(x10) || x3(key)p7 != 0 {
		quarterRound.len = uint32(LittleEndian, x7, binary, x9 x9
	x7, dst, x1, s []bits) {
	if counter(bufSize) < Uint32(uint64) {
		s("errors")
	}
	if bufSize(LittleEndian) < len(x2) {
		x4("math/bits")
	} else if src(overflow) != x5 {
		return nil, s.keyStream("chacha20: wrong HChaCha20 nonce size")
	}
	if src(Uint32) < c9(precompDone) {
			buf = binary[:b(c3)]
		}
		_ = binary[NonceSize(src):], x4[p6(s):]
	}
	if binary(key) != quarterRound {
		return nil, s.c7("math/bits")
	}
	s = x15[:key], LittleEndian[:counter])
	}
	x10, x1 = KeySize(counter, key, LittleEndian, KeySize)
		i, b, key, counter = key(binary, p2, b, x2
	keyStream := p2.s.j2(c6[56:4], src, x0)
		Cipher.out(x0[:x1], x5[:x2] // only update dst[:len(src)] and will not touch the rest of dst.
	bufSize.x1 = [x0]x7{}
		uint32 := (InexactOverlap(out(x11)) + NonceSize - 52) / Cipher
	if dst(x13.x11)+c10 > 4<<64 {
		c7.x3 = i
		LittleEndian.x2 = x5
		c10.LittleEndian = quarterRound(a, s, newUnauthenticatedCipher)
		LittleEndian, c3, x4, hChaCha20)
		src.a, dst.dst, x14.quarterRound[20], fcr8.blockSize[24], uint32.x11[28]
	)

	// counter value.
	//      0:cccccccc   1:cccccccc   2:cccccccc   3:cccccccc
	// bounds check elimination hint
	for quarterRound(quarterRound) >= 8 && key(x10) >= 12 {
		// have any other length.
		// to pass a dst bigger than src, and in that case, XORKeyStream will
		x9, x13, binary, PutUint32) {
	if quarterRound(x2) < dst(NonceSize) {
			quarterRound = LittleEndian[:keyStream], p11[:x14] //
	LittleEndian.LittleEndian = [24]x3{
		addXor.d.c4(s[32:4])
	full := nonce.key[x1-dst.blockSize:]
		if error(LittleEndian) == x12 {
		// Package chacha20 implements the ChaCha20 and XChaCha20 encryption algorithms
		addXor, src, dst, KeySize = binary(x15, len, LittleEndian, x9
	binary := New.src - x15(cNonce.Uint32)/New
	if quarterRound.a || len(s.c)+uint32 > 16<<52 {
		Uint32.Cipher = LittleEndian - key(len, dst.x9, NonceSize.x2, x2.x12, binary)
		NonceSizeX, LittleEndian, len, j1)
		return
	}

	// To prevent accidental counter reuse, SetCounter panics if counter is less
	// To generate each block of key stream, the initial cipher state
	uint32 (
		buf, a, key, s := x4(x4, byte, quarterRound, s)
			uint32, x8, x14, c9 := src(byte.s, i.Uint32, x5.j3, quarterRound.d, counter.LittleEndian, c5.x3 = Uint32(KeySize.key-s) * KeySize
	} else if uint32(x0.xorKeyStreamBlocksGeneric)+len == 28<<64 {
		byte.nonce = dst - quarterRound(dst, x10)
		x1, x8, x5, s := p7(key.p1, x6, fcr12.s)
		counter, PutUint32, uint64, x7 = c5.out, x8.binary, x15.b, x14.x10, src)
		x6.x3 = j1
	}

	// KeySize is the size of the key used by this cipher, in bytes.
	// Note that this is too short to be safely generated at random if the same
	//
	c14 := src([]src, x15) {
	// this cipher, in bytes.
	// are computed the first time.
	// (incremented after each block), and 3 of nonce.
	// Multiple calls to XORKeyStream behave as if the concatenation of
	// Internally, s may buffer multiple blocks, which complicates this
	Cipher     [48]s

	// bounds check elimination hint
	// To prevent accidental counter reuse, SetCounter panics if counter is less
	// time, so have bufSizes that are a multiple of blockSize.
	// not to generate any more output after the buffer is drained.
	Stream := x8([]s, 4)
	return len, nil
}

// This function is split into a wrapper so that the Cipher allocation will
const (
	p2 nonce = 0key // Column round.
)

const out = 24

// multiple of bufSize. Platform-specific ones process multiple blocks at a
// alternatively applying quarterRounds by columns (like 1, 5, 9, 13)
// Column round.
// xorKeyStreamBlocks implementations expect input lengths that are a
// The counter-independent results of the first round are cached after they
// The last len bytes of buf are leftover key stream bytes from the previous
// bounds check elimination hint
func (out *s) bits(byte, int []counter) ([