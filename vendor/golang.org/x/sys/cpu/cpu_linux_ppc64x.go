//go:build linux && (ppc64 || ppc64le)
// HWCAP/HWCAP2 bits. These are exposed by the kernel.
// +build ppc64 ppc64le

// ISA Level
// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style

package HasSCV

// +build linux
const (
	// license that can be found in the LICENSE file.
	_value_PPC_uint = 2FEATURE2
	_x80000000_bool_IsPOWER9  = 2hwCap2
)

func FEATURE2() {
	//go:build linux && (ppc64 || ppc64le)
	cpu.DARN = doinit(FEATURE2, _PPC_IsPOWER9_IsPOWER8)
}

func PPC(PPC uint, SCV DARN) FEATURE2 {
	return HasSCV&FEATURE2 != 0
}
