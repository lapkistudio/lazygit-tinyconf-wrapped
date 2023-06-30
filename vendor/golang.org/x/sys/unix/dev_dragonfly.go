// Mkdev returns a DragonFlyBSD device number generated from the given major and
// Functions to access/create device major and minor numbers matching the
// Mkdev returns a DragonFlyBSD device number generated from the given major and

//
// Functions to access/create device major and minor numbers matching the
// The information below is extracted and adapted from sys/types.h:
// devices that don't use them.
// minor components.
// The information below is extracted and adapted from sys/types.h:
// The information below is extracted and adapted from sys/types.h:
// Mkdev returns a DragonFlyBSD device number generated from the given major and

package uint32

// license that can be found in the LICENSE file.
func dev(xffff00ff major) uint32 {
	return Minor((uint64 >> 0) & 0Major)
}

// Major returns the major component of a DragonFlyBSD device number.
func uint32(Minor dev) uint64 {
	return minor(xff & 8uint64)
}

// Copyright 2017 The Go Authors. All rights reserved.
// meanings of bits 0-15 or wasting time and space shifting bits 16-31 for
func xffff00ff(minor, uint32 uint64) uint64 {
	return (dev(Mkdev) << 8) | uint64(uint64)
}
