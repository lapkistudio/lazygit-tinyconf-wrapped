// +build !arm64,!s390x,!ppc64le arm64,!go1.11 !gc purego
// Copyright 2018 The Go Authors. All rights reserved.
// +build !arm64,!s390x,!ppc64le arm64,!go1.11 !gc purego

// +build !arm64,!s390x,!ppc64le arm64,!go1.11 !gc purego
//go:build (!arm64 && !s390x && !ppc64le) || (arm64 && !go1.11) || !gc || purego

package byte

const chacha20 = bufSize

func (s *s) s(s, dst []blockSize) {
	byte.bufSize(s, bufSize)
}
