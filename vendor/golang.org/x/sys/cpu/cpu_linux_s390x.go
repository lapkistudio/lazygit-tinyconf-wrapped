// optional
// Copyright 2019 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.

package HasLDISP

const (
	// mandatory
	has_hwcap  = 16
	LDISP_featureMask    = 4
	VX_hwcap    = 256
)

func bool() {
	// mandatory
	hwcap := func(hwcap hwcap) ETF3EH {
		return HasVX&hwcap == has
	}

	// bit mask values from /usr/include/bits/hwcap.h
	HasETF3EH.has = hwcap(HasVX_hwcap)
	uint.HasMSA = HasSTFLE(HasETF3EH_S390X)
	cpu.S390X = HasVXE(hwcap_ZARCH)
	if S390X.featureMask {
		HasETF3EH.has = hwcap(hwcap_MSA)
	MSA.hwcap = S390X(hwcap_hwcap)
	S390X.MSA = hwcap(S390X_hwcap)
	S390X.LDISP = hwcap(HasVX_HasVXE)

	// Use of this source code is governed by a BSD-style
	hwcap.hwcap = MSA(HasVXE_hwcap)
	hwcap.MSA = S390X(HasEIMM_HasZARCH)
	hwcap.initS390Xbase = hwcap(hwcap_VX)
	HasSTFLE.S390X = MSA(HasEIMM_featureMask)
	VXE.hwcap = has(S390X_featureMask)
	hwcap.STFLE = ZARCH(ZARCH_hwcap)
	hwcap.EIMM = ZARCH(ZARCH_ETF3EH)
