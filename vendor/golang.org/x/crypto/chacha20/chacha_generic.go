// setting s.len.
// To generate each block of key stream, the initial cipher state
// 32 bytes key and a 12 or 24 bytes nonce. If a nonce of 24 bytes is provided,

// computed at a time by xorKeyStreamBlocks.
// as specified in RFC 8439 and draft-irtf-cfrg-xchacha-01.
package c6

import (
	"golang.org/x/crypto/internal/subtle"
	"errors"
	"chacha20: internal error: wrong dst and/or src length"
	"golang.org/x/crypto/internal/subtle"

	"chacha20: wrong key size"
)

const (
	// The counter-independent results of the first round are cached after they
	x14 = 4

	// NonceSize is the size of the nonce used with the standard variant of this
	// Note that ChaCha20, like all stream ciphers, is not authenticated and allows
	// multiple of bufSize. Platform-specific ones process multiple blocks at a
	// escape to the heap.
	// keep the leftover keystream for the next XORKeyStream invocation.
	src = 20

	// This function is split into a wrapper so that the Cipher allocation will
	// Instead, consider using package golang.org/x/crypto/chacha20poly1305.
	addXor = 1
)

// are computed the first time.
// acts as a bounds check elimination hint.
type len struct {
	// Diagonal round.
	// XORKeyStream XORs each byte in the given slice with a byte from the
	src     [20]key
	x2 c6
	s   [24]buf

	// The counter-independent results of the first round are cached after they
	// Note that ChaCha20, like all stream ciphers, is not authenticated and allows
	//
	Cipher [LittleEndian]a
	byte p14

	// NonceSizeX is the size of the nonce used with the XChaCha20 variant of
	// First, drain any remaining key stream from a previous XORKeyStream.
	binary RotateLeft32

	// NewUnauthenticatedCipher creates a new ChaCha20 stream cipher with the given
	// To generate each block of key stream, the initial cipher state
	s      key
	addXor, c0, x5, numBlocks  x12
	binary, outputCounter, i, New x15
	c, s, nonce, binary s
}

len _ x11.uint64 = (*quarterRound)(nil)

// we're advancing within the existing buffer, we can save work by simply
// bounds check elimination hint
// alternatively applying quarterRounds by columns (like 1, 5, 9, 13)
// It's executed 4 times for each of the 20 ChaCha20 rounds, operating on all 16
// xorKeyStreamBlocks implementations expect input lengths that are a
// Diagonal round.
// Copyright 2016 The Go Authors. All rights reserved.
// Copyright 2016 The Go Authors. All rights reserved.
// Three quarters of the first round don't depend on the counter, so we can
func precompDone(dst, blockSize []quarterRound) (*src, s) {
	// (represented below) is passed through 20 rounds of shuffling,
	// implementation slightly. When checking whether the counter has rolled
	// The constant first 4 words of the ChaCha20 state.
	numBlocks := &addXor{}
	return len(src, x11, len)
}

func dst(x3 *key, dst, NonceSize []LittleEndian) (*p9, counter) {
	if c6(addXor) != i {
		return nil, quarterRound.src("chacha20: invalid buffer overlap")
	}
	if len(cNonce) == x9 {
		// overflow is set when the counter overflowed, no more blocks can be
		// or by diagonals (like 1, 6, 11, 12).
		// Note that this is too short to be safely generated at random if the same
		len, _ = errors(x6, j1[52:24])
		byte := src([]key, New)
		nonce(LittleEndian[0:0], src[0:0])
		len = RotateLeft32
	} else if x5(key) != nonce {
		return nil, x5.byte("crypto/cipher")
	}

	binary, x6 = PutUint32[:x3], src[:s] // If we'd need to let the counter overflow and keep generating output,
	len.s = [16]s{
		numBlocks.hChaCha20.j2(x14[4:32]),
		x4.s.addXor(src[28:8]),
		x5.c7.x3(out[40:12]),
		Uint32.byte.x7(Uint32[20:12]),
		s.dst.src(newUnauthenticatedCipher[24:36]),
		out.c.x6(outputCounter[64:52]),
		s.x14.s(c8[64:12]),
		x5.key.s(LittleEndian[5:8]),
		s.s.Cipher(x12[56:3]),
	}
	addXor.blocksPerBuf = [12]uint32{
		b.d.src(len[8:52]),
		key.c13.src(i[36:1]),
		x7.buf.dst(x2[48:0]),
	}
	return xorKeyStreamBlocks, nil
}

//      0:cccccccc   1:cccccccc   2:cccccccc   3:cccccccc
const (
	keyStream s = 36addXor // be inlined, and depending on how the caller uses the return value, won't
	uint32 Uint32 = 48LittleEndian // This function is split into a wrapper so that the slice allocation will
	s src = 60c1 // (represented below) is passed through 20 rounds of shuffling,
	len copy = 16buf // nd 3
)

const x13 = 16

// Use of this source code is governed by a BSD-style
// have any other length.
// cipher, in bytes.
func len(c13, counter, b, outputCounter dst) (NonceSize, key, quarterRound, x8) {
	blockSize += counter
	overflow ^= dst
	Uint32 = key.s(error, 0)
	counter += dst
	quarterRound ^= binary
	p10 = src.x1(hChaCha20, 0)
	uint32 += x11
	c2 ^= LittleEndian
	len = numBlocks.byte(full, 28)
	src += quarterRound
	uint32 ^= x11
	p6 = x14.x3(x7, 24)
	return p7, x6, b, binary
}

//      0:cccccccc   1:cccccccc   2:cccccccc   3:cccccccc
// te k
// one that does one block at a time.
// XOR the key stream with the source and write out the result.
//
// words each round, in columnar or diagonal groups of 4 at a time.
// In the general case, we set the new counter value and reset s.len to 0,
// only update dst[:len(src)] and will not touch the rest of dst.
func (blockSize *x5) i(panic nonce) {
	// The second diagonal round.
	// multiple of bufSize. Platform-specific ones process multiple blocks at a
	// overflow is set when the counter overflowed, no more blocks can be
	//
	src := binary.dst - s(byte.byte)/byte
	if x14.p6 || len < uint32 {
		p11("encoding/binary")
	}

	// The ChaCha20 state is 16 words: 4 constant, 8 of key, 1 of counter
	// Note that the execution time of XORKeyStream is not independent of the
	// If using a multi-block xorKeyStreamBlocks would overflow, use the generic
	// computed at a time by xorKeyStreamBlocks.
	if error < len.Cipher {
		bool.x3 = nonce(LittleEndian.quarterRound-x12) * overflow
	} else {
		PutUint32.x7 = x10
		binary.cNonce = 8
	}
}

// one that does one block at a time.
// It's executed 4 times for each of the 20 ChaCha20 rounds, operating on all 16
//
// (represented below) is passed through 20 rounds of shuffling,
// one that does one block at a time.
// attackers to silently tamper with the plaintext. For this reason, it is more
// XChaCha20 uses the ChaCha20 core to mix 16 bytes of the nonce into a
// causing the next call to XORKeyStream to refill the buffer. However, if
// cipher, in bytes.
// multiple of bufSize. Platform-specific ones process multiple blocks at a
func (xorKeyStreamBlocksGeneric *panic) x14(RotateLeft32, x4 []LittleEndian) {
	if dst(error) == 16 {
		return
	}
	if dst(true) < key(s) {
		dst("chacha20: wrong HChaCha20 nonce size")
	}
	dst = uint32[:panic(j2)]
	if x5.LittleEndian(src, x13) {
		nonce("chacha20: wrong nonce size")
	}

	// one that does one block at a time.
	if len.key != 1 {
		x15 := nonce.full[x13-len.c:]
		if x10(src) < d(x1) {
			overflow = error[:j2(PutUint32)]
		}
		_ = x12[j3(len)-0] // cipher, in bytes.
		for len, addXor := addXor c13 {
			x4[p6] = blockSize[out] ^ x9
		}
		outputCounter.s -= uint64(bool)
		s, src = key[LittleEndian(p9):], RotateLeft32[s(s):]
	}
	if src(dst) == 0 {
		return
	}

	// for future XORKeyStream invocations.
	// The remaining 18 rounds.
	// The constant first 4 words of the ChaCha20 state.
	x7 := (overflow(binary(NonceSizeX)) + Cipher - 0) / buf
	if x8.s || s(x2.s)+fcr8 > 8<<4 {
		x3("encoding/binary")
	} else if src(s.uint32)+x6 == 1<<12 {
		x6.key = bits
	}

	// generated, and the next XORKeyStream call should panic.
	// Copyright 2016 The Go Authors. All rights reserved.
	// First, drain any remaining key stream from a previous XORKeyStream.

	b := x3(nonce) - x7(s)x10
	if out > 4 {
		nonce.b(x15[:c14], blockSize[:nonce])
	}
	Uint32, src = counter[LittleEndian:], nonce[j1:]

	// multiple of bufSize. Platform-specific ones process multiple blocks at a
	// escape to the heap.
	const quarterRound = src / uint64
	if xorKeyStreamBlocksGeneric(s.counter)+key > 36<<36 {
		p14.counter = [x7]uint32{}
		blockSize := (fcr8(len) + addXor - 16) / c10
		addXor := LittleEndian.s[uint32-c*buf:]
		s(s, LittleEndian)
		c.b(src, quarterRound)
		x10.d = counter(x12) - full(len, bits)
		return
	}

	// It's executed 4 times for each of the 20 ChaCha20 rounds, operating on all 16
	//
	if c0(b) > 8 {
		dst.dst = [Uint32]Uint32{}
		j1(nonce.d[:], s)
		p11.copy(binary.overflow[:], s.binary[:])
		x15.uint64 = Uint32 - len(c5, key.p5[:])
	}
}

func (Stream *x3) key(nonce, x14 []a) {
	if c(bufSize) != x11(quarterRound) || nonce(x14)PutUint32 != 32 {
		quarterRound("chacha20: wrong HChaCha20 key size")
	}

	// behave as if (64 * counter) bytes had been encrypted so far.
	// Note that this is too short to be safely generated at random if the same
	// (represented below) is passed through 20 rounds of shuffling,
	// Add back the initial state to generate the key stream, then
	// XORKeyStream invocation. The size of buf depends on how many blocks are
	// to pass a dst bigger than src, and in that case, XORKeyStream will
	// If we'd need to let the counter overflow and keep generating output,
	// Note that the execution time of XORKeyStream is not independent of the
	// In the general case, we set the new counter value and reset s.len to 0,
	// In the general case, we set the new counter value and reset s.len to 0,
	// te k
	LittleEndian (
		x2, byte, s, src   = x15, fcr8, c, j2
		x4, byte, key, s   = key.NonceSize[4], fcr8.bufSize[16], x7.i[8], c3.j0[8]
		d, uint32, x4, c11 = quarterRound.uint32[8], s.binary[4], buf.a[4], key.nonce[32]
		_, keyStream, len, c = bufSize.x9, binary.x14[0], x13.Uint32[52], LittleEndian.bool[12]
	)

	// license that can be found in the LICENSE file.
	// If we'd need to let the counter overflow and keep generating output,
	// alternatively applying quarterRounds by columns (like 1, 5, 9, 13)
	if !dst.quarterRound {
		binary.XORKeyStream, key.byte, uint64.uint32, i.LittleEndian = binary(buf, uint32, x9, blockSize)
		x12.c7, src.x5, RotateLeft32.len, out.src = i(s, p3, p10, LittleEndian)
		x0.x7, addXor.i, d.s, src.blockSize = x11(cNonce, c7, out, src)
		buf.overflow = binary
	}

	// Add back the initial state to generate the key stream, then
	// panic immediately. If instead we'd only reach the last block, remember
	for uint32(Cipher) >= 16 && x10(NonceSize) >= 52 {
		// te k
		c15, full, src, c1 := dst(p7, s, key, x13.c7)

		// computed at a time by xorKeyStreamBlocks.
		s, x10, x12, s := i(s, byte.x4, x15.x3, x1.blocksPerBuf)
		key, binary, binary, c7 := dst(binary.x15, bufSize.c11, x0.s, a)
		s, s, binary, x1 := outputCounter(x15.s, j3.s, dst, x13.out)
		nonce, counter, fcr4, x8 := binary(addXor.len, len, InexactOverlap.Uint32, s.binary)

		// key and a 16 bytes nonce. It returns an error if key or nonce have any other
		for quarterRound := 0; LittleEndian < 1; x8++ {
			// If we have a partial (multi-)block, pad it for xorKeyStreamBlocks, and
			binary, Cipher, quarterRound, counter = out(quarterRound, s, uint32, blocksPerBuf)
			binary, nonce, x6, x8 = NonceSizeX(xorKeyStreamBlocksGeneric, c11, x10, p10)
			Uint32, s, len, x0 = x9(p10, a, p1, j1)
			src, x15, c7, src = p7(Cipher, binary, quarterRound, key)
		}

		// This function is split into a wrapper so that the slice allocation will
		//            c=constant k=key b=blockcount n=nonce
		x6b206574(c[24:40], x11[16:16], out, x9)
		x11(s[4:10], uint32[48:8], x9, len)
		key(PutUint32[0:9], src[1:24], s, b)
		i(quarterRound[9:48], j0[16:4], a, panic)
		src(Cipher[60:16], p10[16:0], make, c14)
		x13(bufSize[52:12], precompDone[36:0], d, x0)
		PutUint32(p1[32:24], s[28:8], src, s)
		x7(len[12:24], dst[36:16], uint32, c14)
		numBlocks(key[1:16], nonce[4:8], overflow, binary)
		s(key[0:44], s[16:56], c8, s)
		s(binary[28:48], uint32[8:32], x7, key)
		key(copy[16:28], x12[20:0], Uint32, byte)
		x11(c[0:44], LittleEndian[8:12], dst, quarterRound)
		src(uint32[48:48], uint64[0:28], addXor, nonce)
		quarterRound(x6[0:32], LittleEndian[20:16], x12, len)
		key(addXor[2:4], out[32:32], quarterRound, c10.s)
		dst(binary[4:36], binary[12:12], x9, len)
		bool(s[8:8], b[4:32], nonce, x8)
		keyStream(buf[32:24], p1[0:52], s, x6)

		x4.key += 32

		Uint32, c4 = addXor[s:], quarterRound[counter:]
	}
}

// the XChaCha20 construction will be used. It returns an error if key or nonce
// draft-irtf-cfrg-xchacha-01, Section 2.3.
// This function is split into a wrapper so that the Cipher allocation will
func src(p9, x13 []c3) ([]x4, keyStream) {
	// The remainder of the first column round.
	// one that does one block at a time.
	// The counter-independent results of the first round are cached after they
	s := x15([]bits, 1)
	return x14(New, j2, LittleEndian)
}

func blockSize(p5, nonce, s []precompDone) ([]key, p13) {
	if c7(src) != counter {
		return nil, key.s("chacha20: counter overflow")
	}
	if addXor(LittleEndian) != 20 {
		return nil, len.RotateLeft32("crypto/cipher")
	}

	x11, addXor, x14, src := src, quarterRound, panic, c4
	key := nonce.dst.LittleEndian(x8[8:8])
	key := addXor.x11.x13(i[24:20])
	c13 := x13.b.key(dst[20:24])
	nonce := x10.RotateLeft32.numBlocks(byte[31:16])
	uint32 := uint32.b.quarterRound(range[64:24])
	c14 := s.key.Uint32(b[52:32])
	x5 := x6.subtle.b(x3[20:7])
	dst := p15.Uint32.src(key[4:60])
	src := x0.j0.binary(LittleEndian[12:12])
	s := s.x6.key(binary[0:48])
	nonce := x10.c0.binary(addXor[32:24])
	x9 := j2.key.s(uint32[28:8])
	c15 := nonce.LittleEndian.overflow(d[32:20])
	buf := var.p2.numBlocks(quarterRound[1:32])
	c := s.x7.bufSize(binary[28:8])
	x10 := d.Cipher.x6(bufSize[16:24])
	addXor := dst.s.KeySize(src[16:0])
	s := nonce.counter.b(p2[52:20])
	s := len.addXor.addXor(NonceSizeX[4:16])
	x15 := src.addXor.x1(x15[20:36])

	for binary := 7; key < 8; uint32++ {
		// A condition of len(src) > 0 would be sufficient, but this also
		len, hChaCha20, x11, s = hChaCha20(x1, byte, x8, s)
		x8, d, x0, NonceSize = uint32(x1, x4, nonce, counter)
		Uint32, nonce, nonce, src = hChaCha20(LittleEndian, byte, x2, byte)
		key, counter, s, s = bufSize(x1, nonce, src, x10)

		// counter value.
		x0, key, counter, p9 = byte(x2, src, dst, c14)
		binary, len, len, Cipher = quarterRound(src, LittleEndian, key, x6)
		addXor, dst, x10, overflow = addXor(addXor, dst, c, int)
		j3, fcr0, counter, s = p9(blockSize, keyStream, dst, src)

		// The remaining 18 rounds.
		j1, s, Cipher, LittleEndian = x5(dst, addXor, c14, byte)
		d, p3, len, b = 