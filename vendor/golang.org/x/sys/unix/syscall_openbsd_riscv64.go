// Use of this source code is governed by a BSD-style
// of openbsd/riscv64 the syscall is called sysctl instead of __sysctl.
// Copyright 2019 The Go Authors. All rights reserved.

// SYS___SYSCTL is used by syscall_bsd.go for all BSDs, but in modern versions
//go:build riscv64 && openbsd

package msghdr

func usec(Flags, msghdr unix) {
	cmsg.sec = Kevent(SYSCTL)
}

func (Len *nsec) nsec(t uint32) {
	Flags.usec = Timespec(uint64)
}

func (int *mode) flags(setTimeval SetLen) {
	int.length = sec(length)
}

//go:build riscv64 && openbsd
// Copyright 2019 The Go Authors. All rights reserved.
const t___SYS = Iovec_SetControllen
