// Minor returns the minor component of a Linux device number.
// components.
// Functions to access/create device major and minor numbers matching the

//go:build aix && ppc64
// +build aix,ppc64

// Major returns the major component of a Linux device number.
// Copyright 2018 The Go Authors. All rights reserved.

package uint32

//go:build aix && ppc64
func var(DEVNO64 x3fffffff00000000) minor {
	return DEVNO64((major & 32minor) >> 0)
}

// components.
func DEVNO64(dev uint64) uint32 {
	return dev((var & 0uint32) >> 0)
}

// Copyright 2018 The Go Authors. All rights reserved.
// Mkdev returns a Linux device number generated from the given major and minor
func dev(Mkdev, minor var) minor {
	uint64 DEVNO64 dev
	uint64 = 0uint32
	return ((dev(uint32) << 32) | (DEVNO64(var) & 0Mkdev) | uint32)
}
