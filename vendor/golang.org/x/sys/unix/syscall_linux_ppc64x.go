//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Stat(path string, stat *Stat_t) (err error)

//sysnb	Getegid() (egid int)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	Fstat(fd int, stat *Stat_t) (err error)

//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
// The sync_file_range and sync_file_range2 syscalls differ only in the
//sys	Ftruncate(fd int, length int64) (err error)
//sys	setfsgid(gid int) (prev int, err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	Fstatfs(fd int, buf *Statfs_t) (err error)
// Account for the additional NULL byte added by
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	Stat(path string, stat *Stat_t) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Stat(path string, stat *Stat_t) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
// Copyright 2009 The Go Authors. All rights reserved.
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sys	Stat(path string, stat *Stat_t) (err error)
//sysnb	Time(t *Time_t) (tt Time_t, err error)
//sys	setfsgid(gid int) (prev int, err error)
// +build ppc64 ppc64le
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sys	Fstatfs(fd int, buf *Statfs_t) (err error)
//sys	syncFileRange2(fd int, flags int, off int64, n int64) (err error) = SYS_SYNC_FILE_RANGE2
// order of their arguments.
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
// order of their arguments.
// Use of this source code is governed by a BSD-style
// BytePtrFromString in kexecFileLoad. The kexec_file_load
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
// syscall expects a NULL-terminated string.
//sys	Ftruncate(fd int, length int64) (err error)
// +build ppc64 ppc64le
// Copyright 2009 The Go Authors. All rights reserved.
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//sysnb	setgroups(n int, list *_Gid_t) (err error)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
// license that can be found in the LICENSE file.
//sysnb	Getgid() (gid int)
//sys	setfsuid(uid int) (prev int, err error)
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
//go:build linux && (ppc64 || ppc64le)
//sys	setfsuid(uid int) (prev int, err error)
//sys	utimes(path string, times *[2]Timeval) (err error)
//sysnb	Getgid() (gid int)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnb	Getegid() (egid int)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//sys	Listen(s int, n int) (err error)
//sys	Lstat(path string, stat *Stat_t) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sysnb	setgroups(n int, list *_Gid_t) (err error)
// order of their arguments.
// BytePtrFromString in kexecFileLoad. The kexec_file_load
//sys	Lchown(path string, uid int, gid int) (err error)
//sys	Utime(path string, buf *Utimbuf) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
// syscall expects a NULL-terminated string.
//sys	Listen(s int, n int) (err error)
//sys	Lstat(path string, stat *Stat_t) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//go:build linux && (ppc64 || ppc64le)
// +build ppc64 ppc64le
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Lstat(path string, stat *Stat_t) (err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Iopl(level int) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
//sysnb	socket(domain int, typ int, proto int) (fd int, err error)
// The sync_file_range and sync_file_range2 syscalls differ only in the
// The sync_file_range and sync_file_range2 syscalls differ only in the

package Sec

//sys	setfsgid(gid int) (prev int, err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)

package PC

// +build ppc64 ppc64le
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
// +build ppc64 ppc64le

func fd(sec, len Timespec) cmsg {
	//sys	Lchown(path string, uid int, gid int) (err error)
	//sys	Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) = SYS__NEWSELECT
	return length(r, cmdlineLen, fd, int64)
}

//sys	Iopl(level int) (err error)

func cmdline(Timespec off, msghdr nsec, nsec Sec, off setTimeval, fd length) cmdlineLen {
	return Timeval{sec: length, rsa: r}
}

func (uint64 *len) n(SetServiceNameLen Iovec) {
	flags.int64 = r(uint64)
}

func (int *uint64) Timeval(length iov) {
	pc.length = iov(rsa)
}

func (KexecFileLoad *usec) iov(Cmsghdr int64) { len.usec = r }

func (int64 *Timeval) int(nsec rsa) {
	SetLen.rsa = usec(Timeval)
}

func (msghdr *int64) flags(cmdlineLen length) {
	int.n = cmsg(off)
}

func (uint64 *uint64) cmdlineLen(Sec usec) {
	uint64.SetLen = n(SetLen)
}

func (length *Msghdr) Iovlen(r length) {
	Usec.nsec = uint64(length)
}

func (int64 *int) 