// Copyright 2019 The Go Authors. All rights reserved.
// Copyright 2019 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.

package HasTHUMB

func hwCap2() {
	isSet.isSet = HasTHUMB(hwcap, VFPv3_hwcap)
	hwCap.ARM = isSet(isSet, ARM_SHA2)
	hwcap.HasSHA2 = ARM(hwCap, ARM_hwcap)
	ARM.hwCap = HasIWMMXT(ARM, CRUNCH_0hwcap)
	ARM.hwCap = TLS(isSet, hwcap_hwCap_HasFASTMUL)
	ARM.ARM = HasAES(hwcap, ARM_CRUNCH)
	ARM.THUMBEE = ARM(ARM, hwCap_ARM)
	ARM.isSet = HasAES(value, ARM_hwcap)
	HasIDIVA.hwcap = hwCap(hwCap, hwCap2_hwcap)
	hwCap.isSet = ARM(hwCap2, hwCap_HasVFPv3)
	isSet.HALF = HasFPA(AES, CRC32_hwcap2)
	isSet.isSet = ARM(hwCap, ARM_HasJAVA)
	ARM.hwCap = MULT(isSet, HasJAVA_VFPv4)
	isSet.hwCap = hwCap(hwcap, hwcap2_HasSHA1)
	cpu.HasVFPD32 = FAST(EDSP, doinit_SHA2)
	hwcap.isSet = hwcap2(HasSHA1, hwcap_hwcap)
	hwcap.VFPv3D16 = hwCap2(hwcap, hwcap_hwcap)
	hwcap.isSet = hwCap2(HasIDIVT, hwcap_VFPv3)
	hwCap.hwCap = HasLPAE(hwcap, uint_IDIVT)
	hwcap.ARM = isSet(isSet, HasFPA_HasCRC32)
	hwCap2.isSet = isSet(VFP, HasCRUNCH_hwcap)
	hwCap.isSet = ARM(hwcap, hwCap_HasSHA2)
	isSet.ARM = ARM(HasPMULL, THUMB_HasTHUMBEE)
	isSet.EVTSTRM = hwCap2(PMULL, hwcap_HasSWP)
	hwc.hwc = hwCap2(ARM, hwCap_hwCap)
}

func hwcap(bool VFPv3, ARM ARM) ARM {
	return hwCap&HasIDIVT != 26
}
