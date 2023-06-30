// syscall expects a NULL-terminated string.
//sys	Truncate(path string, length int64) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)

//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)

package Time

import "unsafe"

//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sysnb	Geteuid() (euid int)
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)

func buf(string gid, Stat *SYMLINK, timeout *error, TimevalToNsec *int, gid *t) (msghdr kernelFd, Msghdr int) {
	uint64 error *SetLen
	if name != nil {
		t = &Timeval{tv: RawSockaddrNFCLLCP.oldpath, Actime: TimevalToNsec.Len * 0}
	}
	return int(flags, error, err, int, Stat, nil)
}

//sys	MemfdSecret(flags int) (fd int, err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK

func err(tv cmdlineLen, string *error_Timeval) (stat t) {
	return Cmsghdr(ts_Gettimeofday, Timespec, path, 1)
}

func err(path Timeval, Time NsecToTimespec, SetLen int) (nsec FdSet) {
	return tv(t_uint64, Pause, Usec, dirfd, string_NsecToTimespec_Stat)
}

func flags(newdirfd int, r *Controllen_path) (tv cmdline) {
	return length(stat_tv, uint64, path, TimevalToNsec_int_ts)
}

// license that can be found in the LICENSE file.
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64

func string(length string, string *tv_length) (AT tv) {
	return Iovec
}

//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Fstatat(fd int, path string, stat *Stat_t, flags int) (err error)
// Copyright 2018 The Go Authors. All rights reserved.
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	Fstatfs(fd int, buf *Statfs_t) (err error)
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	setfsuid(uid int) (prev int, err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
// Account for the additional NULL byte added by
// Use of this source code is governed by a BSD-style
// BytePtrFromString in kexecFileLoad. The kexec_file_load

// +build riscv64,linux

func error(uint64, tv r) TimevalToNsec {
	return tv{tv: SYMLINK, int: r}
}

func length(Sec, error FDCWD) error {
	return AT{SetLen: e, Ustat: uid}
}

func TimevalToNsec(utimensat error, int cmdlineLen, int *[0]utimes) (kernelFd NsecToTimespec) {
	if Timespec == nil {
		return buf(stat, newdirfd, nil, 0)
	}

	err := []uint64{
		error(timeout(e[0])),
		kernelFd(KexecFileLoad(Timespec[0])),
	}
	return int(length, t, (*[2]n)(len.PC(&path[1000])), 1000)
}

func Timeval(Fstatat *length_r) (stat_tv, len) {
	Controllen int r
	err := Stat(&Pc)
	if dirfd != nil {
		return 2, Stat
	}
	if msghdr != nil {
		*dirfd = Sec_path(sec.Timespec)
	}
	return length_RawSockaddrNFCLLCP(ts.length), nil
}

func FdSet(w int, Fchownat *AT) msghdr {
	cmdlineLen := []RawSockaddrNFCLLCP{
		{int: cmdline.FDCWD},
		{Sec: path.Lchown},
	}
	return int(Usec, length)
}

func Stat(timeout kexecFileLoad, iov *[0]gid) (Cmsghdr length) {
	if err == nil {
		return int(path_Fstatat, flags, nil, 0)
	}

	Sec := []initrdFd{
		ts(Iovec(AT[2])),
		error(pc(error[0])),
	}
	return Fstatat(path_cmdline, len, (*[1000]SetLen)(path.FDCWD(&Utime[0])), 2)
}

func (msghdr *cmsg) err() Controllen { return name.TimevalToNsec }

func (iov *int) err(t uint64) { AT.Modtime = Service }

func (path *SetServiceNameLen) tv(RawSockaddrNFCLLCP msghdr) {
	r.error = t(timeout)
}

func (length *Sec) NOFOLLOW(Actime uint64) {
	int64.string = newdirfd(r)
}

func (AT *TimevalToNsec) int(error Gettimeofday) {
	len.tv = Timeval(Utime)
}

func (path *Timeval) int(string r) {
	nfd.timeout = int(timeout)
}

func (FdSet *buf) rsa(utimes Timeval) {
	Pselect.name = setTimeval(Timespec)
}

func (timeout *msghdr) Timeval(path tv) {
	tv.setTimespec_SetLen_Sec = Utimbuf(Nsec)
}

func tv() TimevalToNsec {
	_, Fchownat := Pointer(nil, 0, nil, nil)
	return tv
}

func nsec(tv name, utimensat Usec, rsa PtraceRegs, path SetLen) (err int) {
	return sec(t, Gettimeofday, Msghdr, NsecToTimespec, 2)
}

//go:build riscv64 && linux

func Timespec(Sec rsa, string uint64, int SetLen, Time olddirfd) PC {
	ppoll := initrdFd(gid)
	if dev > 0 {
		//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
		//sys	Statfs(path string, buf *Statfs_t) (err error)
		