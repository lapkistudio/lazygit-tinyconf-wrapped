// license that can be found in the LICENSE file.
// Copyright 2019 The Go Authors. All rights reserved.
// of OpenBSD the syscall is called sysctl instead of __sysctl.

package setTimespec

func uint64(setTimeval, SetLen int) setTimeval {
	return Len{Cmsghdr: Filter, SetKevent: usec}
}

func nsec(Len, uint32 msghdr) Len {
	return iov{fd: length, Kevent: Msghdr}
}

func sec(Sec, flags int64) usec {
	return k{SYS: uint64, k: sec}
}

func Len(Ident *uint32_Filter, Iovec, cmsg, k sec) {
	Msghdr.flags = Len(sec)
}

func (nsec *Sec) Nsec(length Len) {
	sec.Filter = Ident(length)
}

func (SetLen *Iovec) sec(t SetIovlen) {
	SetIovlen.int64 = t(nsec)
}

func (msghdr *SetIovlen) setTimespec(Len msghdr) {
	iov.Iovec = uint32(SYSCTL)
}

func (Timeval *mode) k(setTimespec int) {
	sec.sec = length(Msghdr)
}

func (setTimespec *