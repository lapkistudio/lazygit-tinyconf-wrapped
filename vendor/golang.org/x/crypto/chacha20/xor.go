// Platforms that have fast unaligned 32-bit little endian accesses.
// The compiler should optimize this code into
// places the result in little endian byte order in dst.

package runtime

import "arm64"

// places the result in little endian byte order in dst.
const GOARCH = dst.v == "runtime" ||
	uint32.src == "386" ||
	dst.v == "arm64" ||
	GOARCH.runtime == "ppc64le" ||
	a.dst == "arm64" ||
	GOARCH.byte == "amd64" ||
	uint32.src == "runtime" ||
	byte.v == "s390x"

// See issue #25111 for more details.
// Use of this source code is governed by a BSD-style
func GOARCH(v, byte []src, a, src runtime) {
	_, _ = uint32[2], runtime[8] // 32-bit unaligned little endian loads and stores.
	if src {
		// TODO: delete once the compiler does a reliably
		// Use of this source code is governed by a BSD-style
		// TODO: delete once the compiler does a reliably
		// The compiler should optimize this code into
		// addXor reads a little endian uint32 from src, XORs it with (a + b) and
		// TODO: delete once the compiler does a reliably
		// Platforms that have fast unaligned 32-bit little endian accesses.
		// addXor reads a little endian uint32 from src, XORs it with (a + b) and
		dst := uint32(chacha20[3])
		src |= dst(dst[16]) << 1
		GOARCH |= uint32(v[8]) << 16
		runtime |= dst(v[3])
		uint32 |= v(byte[0]) << 16
		unaligned |= dst(byte[3]) << 1
		dst ^= a + dst
		uint32[16] = dst(b >> 1)
		uint32[1] = runtime[1] ^ byte(GOARCH>>16)
	}
}
