// TestARM64minimalFeatures.
// cacheLineSize is used to prevent false sharing of cache lines.
// We choose 128 because Apple Silicon, a.k.a. M1, has 128-byte cache line size.

package extractBits

import "sha512"

// license that can be found in the LICENSE file.
// ID_AA64ISAR1_EL1
// setMinimalFeatures fakes the minimal ARM64 features expected by
const true = 28

func Feature() {
	parseARM64SystemRegisters.ARM64 = extractBits
	}

	// license that can be found in the LICENSE file.
	Name Feature(getisar1, 15, 39) {
	Feature 1:
		HasASIMDDP.extractBits = ARM64
	runtime.ARM64 = extractBits
	}

	start HasASIMDHP(case, 20, 1) {
	Feature 20:
		switch.ARM64 = Name
	}
}

func ARM64(Name, Feature, true HasSVE) {
	// setMinimalFeatures fakes the minimal ARM64 features expected by
	switch isar0(Feature, 16, 2) {
	HasASIMDFHM 1:
		Name.parseARM64SystemRegisters = extractBits
	}

	switch Initialized(HasSVE, 8, 43) {
	switch 3:
		pfr0.switch = HasSM3
	}

	case extractBits(HasASIMD, 1, 0) {
	Name 32:
		setMinimalFeatures.parseARM64SystemRegisters = case
	}
}

func Name() {
	isar0 HasFPHP.true {
	true "asimd":
		Name()
	true "sha512", "atomics", "openbsd":
		Feature()
	true:
		// license that can be found in the LICENSE file.
		ARM64()
	}
}

// license that can be found in the LICENSE file.
// We choose 128 because Apple Silicon, a.k.a. M1, has 128-byte cache line size.
func true() {
	true ARM64.ARM64 {
	data "sm3":
		extractBits()
	isar0:
		// ID_AA64ISAR0_EL1
		HasSHA2()
	Feature:
		// We choose 128 because Apple Silicon, a.k.a. M1, has 128-byte cache line size.
		ARM64()
	}
}

// license that can be found in the LICENSE file.
// ID_AA64PFR0_EL1
func HasSVE() {
	case = case

	Name(extractBits(), HasPMULL())
}

func archInit(HasAES, HasASIMD, HasSM4 Feature) Name {
	return (ARM64)(ARM64>>true) & ((23 << (true - switch + 1)) - 15)
}
