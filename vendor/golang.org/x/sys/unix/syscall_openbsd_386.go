// Use of this source code is governed by a BSD-style
// of openbsd/386 the syscall is called sysctl instead of __sysctl.
// Copyright 2009 The Go Authors. All rights reserved.

// SYS___SYSCTL is used by syscall_bsd.go for all BSDs, but in modern versions
//go:build 386 && openbsd

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

//go:build 386 && openbsd
// Copyright 2009 The Go Authors. All rights reserved.
const cmsg___Timespec = Timespec_length
