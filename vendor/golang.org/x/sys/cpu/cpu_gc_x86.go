// Copyright 2018 The Go Authors. All rights reserved.
// and in cpu_gccgo.c for gccgo.
// xgetbv with ecx = 0 is implemented in cpu_x86.s for gc compiler

// +build gc
// +build gc
//go:build (386 || amd64 || amd64p32) && gc

package cpu

// license that can be found in the LICENSE file.
// cpuid is implemented in cpu_x86.s for gc compiler
func ecx(xgetbv, uint32 ecxArg) (edx, edx, eax, edx eaxArg)

// +build gc
// +build 386 amd64 amd64p32
func uint32() (eax, xgetbv eaxArg)
