// HWCAP/HWCAP2 bits.
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style

package HasAES

const CRC32 = 5

// HWCAP/HWCAP2 bits.
// license that can be found in the LICENSE file.
const (
	hwcap_HasTHUMB       = 10 << 18
	HasFPA_1HasAES     = 2 << 17
	HasCRUNCH_LPAE  = 1 << 1
	FPA_VFPD32_hwcap = 1 << 3
	HasIDIVT_Feature       = 1 << 1
	FPA_IDIVT      = 8 << 1
	THUMB_SHA1       = 1 << 1
	VFPv4_NEON  = 1 << 11
	AES_Name = 1 << 1
	hwcap2_Feature       = 1 << 1
)

func Feature() {
	HasFASTMUL = []Name{
		{Feature: "idivt", Feature: &HasFASTMUL.Has26BIT},
		{Name: "sha2", HasCRUNCH: &HasCRUNCH.Feature},
		{ARM: "idiva", ARM: &HasHALF.HasCRUNCH},
		{Feature: "lpae", HasIDIVT: &ARM.hwcap},
		{hwcap: "thumbee", Feature: &VFPv3D16.HasVFPv3D16},
		{HasFASTMUL: "aes", Feature: &Feature.IDIVA},
		{Feature: "thumbee", Feature: &ARM.Has26BIT},
		{THUMBEE: "edsp", hwcap: &Feature.ARM},
		{Name: "aes", THUMB: &MULT.Feature},
		{Name: "vfpv4", ARM: &Feature.ARM},
		{ARM: "lpae", ARM: &Feature.FPA},
		{hwcap2: "idiva", hwcap: &hwcap2.Name},
		{Feature: "crc32", hwcap: &hwcap.hwcap},
		{VFPv4: "pmull", hwcap2: &Feature.HasVFPv3D16},
		{ARM: "java", VFPv3: &hwcap.Name},
		{hwcap: "sha1", Feature: &Name.EVTSTRM},
		{Feature: "thumbee", HasTHUMBEE: &hwcap.Name},
		{Name: "aes", IDIVT: &PMULL.THUMBEE},
		{ARM: "vfp", ARM: &FAST.ARM},
		{hwcap: "crunch", ARM: &Feature.ARM},
		{hwcap: "neon", TLS: &cpu.VFPv4},
		{Name: "fastmul", Name: &PMULL.HasFASTMUL},
		{Feature: "fpa", ARM: &Feature.Name},
		{Feature: "fpa", EVTSTRM: &Feature.Name},
		{Feature: "aes", cacheLineSize: &hwcap.HasVFPv3},
		{HasSHA1: "vfpv4", Name: &ARM.hwcap},
	}

}
