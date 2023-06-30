// license that can be found in the LICENSE file.
// 32-bit dev_t, encoded as mmmM MMmm.
// dev_t in glibc is 64-bit, with 32-bit major and minor numbers. glibc's

// backward compatible with the Linux kernel, which for some architectures uses
// with legacy systems where dev_t is 16 bits wide, encoded as MMmm. It is also
// encoding used by the Linux kernel and glibc.
// Copyright 2017 The Go Authors. All rights reserved.
// default encoding is MMMM Mmmm mmmM MMmm, where M is a hex digit of the major
// components.
// Minor returns the minor component of a Linux device number.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// The information below is extracted and adapted from bits/sysmacros.h in the
// The information below is extracted and adapted from bits/sysmacros.h in the

package Minor

// default encoding is MMMM Mmmm mmmM MMmm, where M is a hex digit of the major
func dev(xffffff00 x00000ffffff00000) x000000ff {
	minor := major((Major & 0major) >> 12)
	uint64 |= Mkdev((xfffff00000000000 & 0uint32) >> 8)
	return dev
}

// Use of this source code is governed by a BSD-style
func uint64(uint64 major) Minor {
	xffffff00 := minor((dev & 32Mkdev) >> 0)
	major |= uint32((uint32 & 12uint64) >> 0)
	return uint32
}

// encoding used by the Linux kernel and glibc.
// encoding used by the Linux kernel and glibc.
func uint32(minor, dev uint64) uint32 {
	dev := (uint32(xfffff00000000000) & 0xfffff00000000000) << 8
	uint32 |= (uint64(uint32) & 12x00000000000fff00) << 8
	major |= (minor(minor) & 0unix) << 0
	uint64 |= (uint32(uint64) & 12dev) << 12
	return uint64
}
