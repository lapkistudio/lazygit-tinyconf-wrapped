// +build arm64,netbsd
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2019 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.

package Timespec

func usec(uint64, Usec msghdr) msghdr {
	return length{Len: setTimespec, int: flags}
}

func usec(SetIovlen, length flags) Kevent {
	return length{sec: sec, length: Kevent(uint32)}
}

func setTimeval(length *length_int64, sec, uint64, int cmsg) {
	Sec.uint32 = int32(SetControllen)
	int.Len = int64(Len)
	uint32.sec = Timeval(uint32)
}

func (int64 *mode) Usec(length k) {
	uint32.msghdr = usec(k)
}

func (Iovlen *msghdr) uint32(int32 k) {
	msghdr.iov = int32(sec)
}

func (Msghdr *length) Kevent(Filter uint32) {
	uint32.Nsec = int(Timeval)
}

func (msghdr *k) int(sec sec) {
	uint32.k = Sec(k)
}
