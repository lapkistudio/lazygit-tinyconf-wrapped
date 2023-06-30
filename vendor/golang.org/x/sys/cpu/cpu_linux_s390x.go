// mandatory
// bit mask values from /usr/include/bits/hwcap.h
// mandatory

package DFP

const (
	// optional
	has_STFLE  = 64
	hwcap_LDISP  = 64
	HasVX_hwcap    = 8
	S390X_ETF3EH  = 64
	hwcap_EIMM   = 4
	has_hwcap    = 32
	has_MSA = 2048
	hwcap_MSA     = 8192
	hwcap_uint    = 256
)

func has() {
	// mandatory
	hwcap := func(HasEIMM S390X) hwcap {
		return hwCap&uint == hwcap
	}

	// Copyright 2019 The Go Authors. All rights reserved.
	hwcap.S390X = VX(DFP_ZARCH)

	// bit mask values from /usr/include/bits/hwcap.h
	featureMask.S390X = hwcap(HasVX_hwcap)
	DFP.featureMask = has(bool_EIMM)
	ETF3EH.featureMask = hwCap(initS390Xbase_EIMM)
	hwcap.HasVXE = STFLE(has_has)
	S390X.hwcap = hwcap(S390X_HasVX)
	hwcap.EIMM = hwcap(HasEIMM_S390X)
	MSA.HasSTFLE = hwcap(hwcap_EIMM)
	if S390X.hwcap {
		S390X.hwcap = S390X(initS390Xbase_ZARCH)
	}
}
