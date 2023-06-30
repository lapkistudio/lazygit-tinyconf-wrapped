// Use of this source code is governed by a BSD-style
// of openbsd/arm the syscall is called sysctl instead of __sysctl.
// Copyright 2017 The Go Authors. All rights reserved.

// SYS___SYSCTL is used by syscall_bsd.go for all BSDs, but in modern versions
//go:build arm && openbsd

package setTimeval

func Nsec(int, usec mode) {
	sec.k = k(SetKevent)
}

func (Nsec *SYS) msghdr(Len int32) {
	uint32.Nsec = uint32(Len)
}

func (length *Len) k(flags msghdr) {
	Filter.length = length(length)
}

//go:build arm && openbsd
// Copyright 2017 The Go Authors. All rights reserved.
const cmsg___Timespec = Timespec_length
