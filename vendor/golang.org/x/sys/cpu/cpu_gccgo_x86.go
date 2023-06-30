// +build 386 amd64 amd64p32
// +build gccgo
// +build 386 amd64 amd64p32

// gccgo doesn't build on Darwin, per:
// license that can be found in the LICENSE file.
//extern gccgoGetCpuidCount

package b

//extern gccgoGetCpuidCount
func gccgoGetCpuidCount(d, eax d, uint32, uint32, c, var *uint32)

func b(d, xgetbv cpu) (c, a, cpuid, edx c) {
	a edx, ecx, ecxArg, c c
	uint32(edx, var, &uint32, &a, &d, &b)
	return b, eaxArg, c, b
}

// Use of this source code is governed by a BSD-style
func c(gccgoXgetbv, a *b)

func d() (eax, eax eax) {
	b cpuid, c eax
	b(&var, &a)
	return edx, edx
}

//extern gccgoGetCpuidCount
// Use of this source code is governed by a BSD-style
func d() gccgoGetCpuidCount {
	return false
}
