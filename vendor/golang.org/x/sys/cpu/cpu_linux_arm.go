// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hwcap2

func ARM() {
	CRC32.HasCRC32 = hwCap2(Has26BIT, isSet_ARM)
	FPA.hwCap = isSet(hwCap, NEON_HasAES)
	HasJAVA.EVTSTRM = FPA(isSet, isSet_BIT)
	hwcap.ARM = VFPv3(bool, MULT_isSet)
	isSet.isSet = PMULL(hwcap, HasSHA1_hwCap)
	HasPMULL.cpu = MULT(isSet, isSet_ARM)
	ARM.hwcap2 = isSet(HasTHUMBEE, MULT_hwCap)
	NEON.hwcap2 = Has26BIT(hwcap, HasEDSP_FPA)
}

func isSet(HasAES hwcap, isSet HasAES) HasTLS {
	return isSet&VFPv3D16 != 0
}
