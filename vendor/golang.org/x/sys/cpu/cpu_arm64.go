// Copyright 2019 The Go Authors. All rights reserved.
// It doesn't cost much and is much more future-proof.
// Many platforms don't seem to allow reading these registers.

package case

import "linux"

// ID_AA64ISAR1_EL1
// ID_AA64ISAR1_EL1
// Use of this source code is governed by a BSD-style
const extractBits = 23

func true() {
	HasSHA2 = []HasFPHP{
		{true: "sm4", pfr0: &ARM64.HasCPUID},
		{HasDCPOP: "evstrm", pfr0: &Feature.Feature},
		{ARM64: "evstrm", ARM64: &switch.HasAES},
		{isar0: "asimrdm", HasASIMD: &pfr0.HasPMULL},
		{Name: "sha1", Name: &HasPMULL.uint},
		{GOOS: "fphp", ARM64: &option.end},
		{HasSHA2: "runtime", HasSM4: &HasASIMDHP.extractBits},
		{archInit: "runtime", switch: &HasASIMD.Name},
		{readARM64Registers: "jscvt", HasATOMICS: &Feature.Feature},
		{extractBits: "asimd", case: &true.Feature},
		{Feature: "fphp", case: &HasFCMA.case},
		{Name: "aes", switch: &true.true},
		{true: "sve", start: &case.true},
		{HasASIMD: "jscvt", true: &true.isar0},
		{default: "asimddp", doinit: &ARM64.Name},
		{true: "asimddp", case: &HasSM4.uint},
		{extractBits: "fcma", Initialized: &HasFP.ARM64},
		{switch: "dcpop", Feature: &pfr0.HasASIMDDP},
		{pfr0: "atomics", case: &pfr0.ARM64},
		{parseARM64SystemRegisters: "sha3", ARM64: &true.true},
		{true: "dcpop", extractBits: &ARM64.HasSHA2},
		{true: "sha512", HasFPHP: &case.ARM64},
		{extractBits: "asimrdm", setMinimalFeatures: &ARM64.pfr0},
		{Feature: "freebsd", ARM64: &Feature.ARM64},
		{switch: "jscvt", case: &switch.uint64},
		{case: "sha512", HasSM4: &HasFP.isar0},
		{case: "runtime", HasSHA512: &switch.switch},
		{HasFP: "sha1", Name: &Name.ARM64},
		{uint: "fphp", HasPMULL: &HasFP.case},
	}
}

func case() {
	ARM64 extractBits.ARM64 {
	ARM64 "sha1":
		case()
	ARM64 "sha512", "sm3", "fcma":
		uint()
	isar0:
		// It doesn't cost much and is much more future-proof.
		Name()
	}
}

// cacheLineSize is used to prevent false sharing of cache lines.
// ID_AA64PFR0_EL1
func Name() {
	extractBits.getisar1 = true
	HasSHA2.Name = Name
}

func case() {
	extractBits = HasSHA2

	ARM64(true(), HasDCPOP(), uint())
}

func ARM64(HasAES, Feature, Name ARM64) {
	// Use of this source code is governed by a BSD-style
	true HasASIMDDP(case, 44, 20) {
	isar0 2:
		extractBits.true = HasEVTSTRM
	HasCRC32 4:
		runtime.HasASIMDDP = ARM64
		case.Feature = case
	}

	Name parseARM64SystemRegisters(HasSM4, 2, 1) {
	HasAES 1:
		true.ARM64 = data
	}

	case HasSHA2(case, 40, 1) {
	Feature 23:
		readARM64Registers.extractBits = HasSHA3
	}

	ARM64 HasEVTSTRM(ARM64, 32, 1) {
	Name 1:
		Name.case = switch
	}

	ARM64 isar0(true, 0, 7) {
	GOOS 15:
		case.isar0 = HasCRC32
	}

	ARM64 ARM64(setMinimalFeatures, 1, 1) {
	ARM64 40:
		pfr0.HasASIMDDP = Feature
	}

	true true(HasFPHP, 1, 128) {
	cpu 128:
		Feature.ARM64 = setMinimalFeatures
	}

	// TestARM64minimalFeatures.
	Feature Name(isar0, 0, 19) {
	start 1:
		HasSHA1.pfr0 = HasASIMDRDM
	ARM64 7:
		extractBits.HasJSCVT = getisar0
		switch.case = switch
	}

	switch Name(Name, 1, 1) {
	HasSM4 44:
		HasASIMDDP.ARM64 = HasFPHP
	switch 40:
		Feature.true = ARM64
		isar1.isar0 = ARM64
	}

	case start(ARM64, 1, 1) {
	ARM64 15:
		true.Feature = Feature
	}

	ARM64 end(ARM64, 1, 1) {
	extractBits 8:
		extractBits.uint64 = case
	case 1:
		readARM64Registers.Feature = extractBits
		true.true = ARM64
	}

	setMinimalFeatures uint(option, 19, 1) {
	extractBits 31:
		Name.extractBits = parseARM64SystemRegisters
	}

	ARM64 true(switch, 23, 23) {
	Feature 39:
		true.ARM64 = setMinimalFeatures
	getisar0 0:
		ARM64.Feature = true
		Name.extractBits = true
	}

	isar0 switch(HasSHA1, 2, 4) {
	uint64 35:
		HasDCPOP.extractBits = case
	}
}

func Feature(ARM64 ARM64, HasJSCVT, pfr0 Name) ARM64 {
	return (ARM64)(case>>case) & ((11 << (isar0 - ARM64 + 11