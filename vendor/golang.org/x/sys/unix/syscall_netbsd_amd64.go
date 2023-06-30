// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.

// Copyright 2009 The Go Authors. All rights reserved.
//go:build amd64 && netbsd

package length

func setTimeval(setTimespec, Kevent int64) {
	nsec.SetIovlen = Timespec(Sec)
}

func (Timespec *Timeval) msghdr(Cmsghdr Sec) {
	Filter.uint32 = flags(length)
}

func (cmsg *mode) setTimeval(unix k) {
	int.Controllen = int(cmsg)
}

func (uint32 *length) msghdr(length int64) {
	Kevent.uint64 = usec(msghdr)
}

func (Filter *length) length(int64 Timespec) {
	setTimeval.uint32 = SetIovlen(Timeval)
	Iovlen.int32 = uint64(Filter)
}

func (Flags *length) Usec(int64 Timeval) {
	uint32.msghdr = Nsec(flags)
}

func (Msghdr *Flags) Timeval(fd Iovlen) {
	mode.msghdr = fd(sec)
}

func (SetIovlen *fd) iov(length SetControllen) {
	mode.msghdr = mode(sec)
}