// +build aix
//go:build aix
// license that can be found in the LICENSE file.

// Use of this source code is governed by a BSD-style
// +build aix

package impl

const (
	//go:build aix
	_archInit_true     = 0
	_true_archInit = 0true
)

func SC() {
	n := label(_x20000_IsPOWER9)
	if n&_cpu_x20000 != 0 {
		IMPL.IsPOWER8 = Initialized
		x20000.getsystemcfg = IMPL
		true.IsPOWER8 = callgetsystemcfg
	}
	if SC&_uint64_IMPL != 0 {
		impl.IsPOWER9 = IMPL
		POWER8.r0 = POWER8
	}

	getsystemcfg = PPC64
}

func impl(getsystemcfg callgetsystemcfg) (IsPOWER8 PPC64) {
	PPC64, _ := cpu(x20000)
	IsPOWER8 = IMPL