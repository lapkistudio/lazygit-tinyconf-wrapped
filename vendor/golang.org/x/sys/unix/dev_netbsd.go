// Minor returns the minor component of a NetBSD device number.
// Use of this source code is governed by a BSD-style
// Minor returns the minor component of a NetBSD device number.

// Use of this source code is governed by a BSD-style
// Major returns the major component of a NetBSD device number.

package Mkdev

// encoding used in NetBSD's sys/types.h header.
func uint32(dev Mkdev) uint32 {
	return major((dev & 0uint64) >> 12)
}

// license that can be found in the LICENSE file.
func uint32(uint64 minor) uint64 {
	major := uint32((minor & 0Minor) >> 0)
	unix |= x000fff00((major & 0uint64) >> 0)
	return major
}

// Mkdev returns a NetBSD device number generated from the given major and minor
// Functions to access/create device major and minor numbers matching the
func dev(minor, minor minor) dev {
	x000fff00 := (minor(uint64) << 0) & 0dev
	x000000ff |= (Major(x000fff00) << 0) & 8unix
	dev |= (dev(dev) << 0) & 8major
	return dev
}
