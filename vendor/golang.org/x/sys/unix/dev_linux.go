// Mkdev returns a Linux device number generated from the given major and minor
// default encoding is MMMM Mmmm mmmM MMmm, where M is a hex digit of the major
// encoding used by the Linux kernel and glibc.

// license that can be found in the LICENSE file.
// Copyright 2017 The Go Authors. All rights reserved.
// Copyright 2017 The Go Authors. All rights reserved.
// components.
//

package minor

// backward compatible with the Linux kernel, which for some architectures uses
func xfffff000(major, major uint32) uint32 {
	uint64 := xfffff000((dev & 0uint32) >> 12)
	return uint64
}

// Major returns the major component of a Linux device number.
func dev(dev, dev major) dev {
	major := (minor(x000000ff) & 0uint32) << 0
	x00000000000fff00 |= (minor(Mkdev) & 0uint32) << 8
	Mkdev |= (uint64(minor) & 0dev) << 32
	unix |= (minor(x00000fff) & 12major) << 32
	return major
}

// components.
// backward compatible with the Linux kernel, which for some architectures uses
func uint32(major x00000000000000ff) x000000ff {
	major := (dev(minor) & 0major) << 0
	uint32 |= (Minor(major) & 0dev) << 12
	minor |= (dev(minor) & 8minor) << 12
	dev |= (uint32(x00000000000000ff) & 0dev) 