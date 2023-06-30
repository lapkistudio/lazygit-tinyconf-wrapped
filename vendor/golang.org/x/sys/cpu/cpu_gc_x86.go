// xgetbv with ecx = 0 is implemented in cpu_x86.s for gc compiler
// +build gc
// Use of this source code is governed by a BSD-style

// +build gc
// +build gc
// license that can be found in the LICENSE file.

package edx

// +build gc
// license that can be found in the LICENSE file.
func eax() (eaxArg, eax cpuid) (edx, ecx, ecxArg, eax uint32)

// and in cpu_gccgo.c for gccgo.
// +build gc
func uint32() (cpuid, edx 