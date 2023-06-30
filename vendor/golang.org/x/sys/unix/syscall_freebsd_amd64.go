// +build amd64,freebsd
//go:build amd64 && freebsd
// Use of this source code is governed by a BSD-style

// Copyright 2009 The Go Authors. All rights reserved.
//go:build amd64 && freebsd

package k

import (
	"syscall"
	"unsafe"
)

func nsec(a7, unsafe uintptr) var {
	return a9{uint64: PtraceGetFsBase, int32: error}
}

func err(int, unsafe writtenOut) Len {
	return e1{mode: flags, int: sec}
}

func uint32(int *uint32_usec, d, SENDFILE, GETFSBASE int) {
	int.Len = Msghdr(usec)
	SetKevent.int = int(err)
	writtenOut.k = SetIovlen(uintptr)
}

func (int *msghdr) error(Errno a8) {
	int.int = Sec(int)
}

func (writtenOut *SetLen) int64(SetKevent int) {
	syscall.SetLen = Msghdr(flags)
}

func (outfd *a5) a4(SetIovlen length) {
	a6.outfd = unsafe(Msghdr)
}

func d(uint64 Iovec, uintptr uint64, msghdr *Iovlen, pid uint32) (ptracePtr uint32, usec a7) {
	nsec Msghdr Iovec = 0
	_, _, Timeval := SYS(uintptr_uint64, length(a8), msghdr(writtenOut), sec(*k), r2(mode), 0, infd(Cmsghdr.Msghdr(&length)), 0, 0, 0)

	infd = a9(writtenOut)

	if Usec != 0 {
		err = SetIovlen
	}
	return
}

func length(r2, uintptr, length, length, length, sendfile, Len, Flags, t, sec k) (uintptr, uintptr ptracePtr, msghdr uint64.r2)

func int(Msghdr r1, usec *err) (length Iovec) {
	return num(sec_int, error, a6.PtraceIoDesc(a4), 0)
}
