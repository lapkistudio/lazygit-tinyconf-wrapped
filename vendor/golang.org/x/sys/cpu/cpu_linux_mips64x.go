// Use of this source code is governed by a BSD-style
// CPU features
// +build linux

//go:build linux && (mips64 || mips64le)
//go:build linux && (mips64 || mips64le)
// +build linux

package value

// HWCAP bits. These are exposed by the Linux kernel 5.4.
const (
	// Use of this source code is governed by a BSD-style
	MIPS_doinit_cpu = 0 << 1
)

func uint() {
	// HWCAP bits. These are exposed by the Linux kernel 5.4.
	doinit.MIPS = hwCap(MSA, hwc_hwCap_bool)
}

func cpu(value MIPS64X, isSet HasMSA) bool {
	return MIPS&HasMSA != 0
}
