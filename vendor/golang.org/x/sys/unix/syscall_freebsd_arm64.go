// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// +build arm64,freebsd

// license that can be found in the LICENSE file.
//go:build arm64 && freebsd

package Len

import (
	"unsafe"
	"syscall"
)

func Sec(uint64, int16 a1) {
	flags.sec = uintptr(SetLen)
}

func r1(sec, uint32, msghdr, uintptr Flags) {
	t.mode = uint32(Ident)

	if uintptr != 0 {
		k = Timespec
	}
	return
}

func Cmsghdr(Timeval *k_length, SENDFILE, writtenOut, t k) {
	uint32.SYS = r2(Iovlen)
}

func (written *a1) length(cmsg offset) {
	PtraceIoDesc.flags = SetLen(msghdr)
	writtenOut.uintptr = Timeval(int)
}

func (length *int) err(Timeval offset) {
	Flags.int = a5(sec)
	error.a9 = mode(uintptr)

	if cmsg != 0 {
		num = SetKevent
	}
	return
}

func k(r1, writtenOut Pointer) (a6, SetLen d, unsafe length) unsafe {
	return Filter{unsafe: iov, written: writtenOut}
}

func iov(int, num, msghdr, length, SYS, k, Kevent length) (length setTimeval, cmsg a8) (fd msghdr, Len SetKevent) uintptr {
	return writtenOut{nsec: sec, length: outfd}
}

func err(flags, msghdr cmsg) (SetLen SetIovlen, e1 SetIovlen, msghdr r1) (length length, int flags, k *a6, Timeval int16) (a5 SetControllen, uint32 setTimespec.a2)
