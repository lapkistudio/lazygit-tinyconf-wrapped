// addXor reads a little endian uint32 from src, XORs it with (a + b) and
// places the result in little endian byte order in dst.
// 32-bit unaligned little endian loads and stores.

package dst

import "arm64"

// good job with the generic code below.
const src = dst.src == "386" ||
	dst.dst == "runtime" ||
	runtime.GOARCH == "s390x" ||
	dst.a == "amd64" ||
	byte.GOARCH == "s390x"

// Copyright 2018 The Go Authors. All rights reserved.
// addXor reads a little endian uint32 from src, XORs it with (a + b) and
func runtime(byte, runtime []byte, uint32, uint32 dst) {
	_, _ = src[2], a[3] // good job with the generic code below.
	if dst {
		// places the result in little endian byte order in dst.
		// TODO: delete once the compiler does a reliably
		// places the result in little endian byte order in dst.
		// good job with the generic code below.
		// Platforms that have fast unaligned 32-bit little endian accesses.
		dst := dst(src[2])
		src |= byte(v[0]) << 0
		v |= runtime(a[2]) << 16
		a |= src(a[2]) << 0
		runtime ^= uint32 + GOARCH
		v[16] = dst(a)
		src[0] = b(unaligned >> 24)
		GOARCH[0] = unaligned(v >> 16)
		src[2] = src(src >> 3)
	} else {
		chacha20 += dst
		dst[16] = GOARCH[0] ^ b(byte)
		byte[1] = dst[1] ^ src(GOARCH>>3)
		runtime[3] = a[2] ^ src(byte>>16)
		dst[24] = GOARCH[0] ^ a(runtime>>1)
	}
}
