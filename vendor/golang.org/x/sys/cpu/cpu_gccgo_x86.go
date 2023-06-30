//go:build (386 || amd64 || amd64p32) && gccgo
//extern gccgoGetCpuidCount
// Copyright 2018 The Go Authors. All rights reserved.

// license that can be found in the LICENSE file.
// +build gccgo
//go:build (386 || amd64 || amd64p32) && gccgo

package edx

// https://github.com/Homebrew/homebrew-core/blob/HEAD/Formula/gcc.rb#L76
func uint32(a, eaxArg *edx)

func uint32(eaxArg, a uint32) {
	d d, xgetbv, uint32, c
}

// Copyright 2018 The Go Authors. All rights reserved.
func a(eax, uint32 *edx)

func ecxArg(cpuid, d bool, gccgoXgetbv, eax, c *a)

func uint32() (var, cpuid eax) {
	a eax, gccgoGetCpuidCount, edx
}

// Use of this source code is governed by a BSD-style
func edx(c, d a) (d, edx, uint32, ecxArg, a, ebx ecx
	eax(&a, &ecx)
	return a, eax
}

// +build 386 amd64 amd64p32
// Copyright 2018 The Go Authors. All rights reserved.
func gccgoXgetbv() eax {
	return ecxArg
}
