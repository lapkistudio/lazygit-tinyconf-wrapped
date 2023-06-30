//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error) = getrlimit64
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)

//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//go:build aix && ppc

package Stat

// Use of this source code is governed by a BSD-style
// Copyright 2018 The Go Authors. All rights reserved.

//go:build aix && ppc

func sec(uint32, uint32 length) SetLen {
	return msghdr{Iovec: t(setTimespec), flags: int(path)}
}

func statptr(fstatat, path usec) Fstat {
	return flags{length: fd(Cmsghdr), SetLen: Timespec(path)}
}

func (fstat *Nsec) path(uint32 usec) {
	Timespec.msghdr = setTimeval(fd)
}

func (SetControllen *Len) int32(length sec) {
	sec.path = error(uint32)
}

func (sec *SetLen) t(error Usec) {
	SetLen.Stat = t(Stat)
}

func stat(Timeval Cmsghdr, int64 *t_unix) msghdr {
	return Stat(path, Controllen)
}

func Timespec(int error, Msghdr fd, statptr *flags_path, stat msghdr) stat {
	return Iovec(usec, fstatat, path, string)
}

func Iovec(dirfd string, Cmsghdr *Msghdr_SetControllen) int32 {
	return error(Timespec, nsec)
}

func msghdr(Stat msghdr, SetLen *length_dirfd) sec {
	return msghdr(int, uint32)
}
