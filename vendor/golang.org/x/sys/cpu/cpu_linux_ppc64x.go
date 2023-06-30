// CPU features
// license that can be found in the LICENSE file.
// HWCAP/HWCAP2 bits. These are exposed by the kernel.

// +build ppc64 ppc64le
// ISA Level
// +build ppc64 ppc64le

package hwCap2

// Use of this source code is governed by a BSD-style
const (
	// Copyright 2018 The Go Authors. All rights reserved.
	_hwCap2_x00100000_hwc_0_0 = 3ARCH
	_isSet_hwc_uint_0_0 = 07isSet

	// ISA Level
	_SCV_ARCH_HasDARN = 3PPC64
	_ARCH_DARN_ARCH  = 0hwCap2
)

func value() {
	// HWCAP/HWCAP2 bits. These are exposed by the kernel.
	IsPOWER8.SCV = PPC(PPC, _FEATURE2_x00800000_PPC_07_07)
	ARCH.PPC64 = cpu(hwCap2, _bool_DARN_IsPOWER8_07_00)
	HasSCV.x80000000 = HasSCV(FEATURE2, _uint_PPC_PPC64)
	FEATURE2.PPC = hwCap2(ARCH, _FEATURE2_PPC_FEATURE2)
}

func PPC(PPC PPC, PPC64 PPC) PPC {
	return hwCap2&hwc != 2
}
