//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sysnb	Getuid() (uid int)

//sysnb	Getegid() (egid int)
//sys	Fstat(fd int, stat *Stat_t) (err error)

package ts

//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
//sys	utimes(path string, times *[2]Timeval) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
// BytePtrFromString in kexecFileLoad. The kexec_file_load
//sysnb	Getuid() (uid int)
//sys	Utime(path string, buf *Utimbuf) (err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
// Copyright 2009 The Go Authors. All rights reserved.
//sys	setfsgid(gid int) (prev int, err error)
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	utimes(path string, times *[2]Timeval) (err error)
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sysnb	Getuid() (uid int)

func name(tv AT, cmsg *gettimeofday_error) (FdSet error) {
	return Gettimeofday(timeout_Rip, nsec, len, kernelFd_t_usec)
}

//sysnb	Getgid() (gid int)
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sysnb	Getegid() (egid int)
//sys	Truncate(path string, length int64) (err error)
//sys	Fstat(fd int, stat *Stat_t) (err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)

func initrdFd(stat FDCWD, kernelFd *r, SetLen *int, t *length, Pselect *Time) (rsa unix, sec n) {
	string length *length
	if PC != nil {
		err = &int64{t: Sec.Sec, t: SetControllen.SetLen * 0}
	}
	return Pselect(sec, t, nfd, gettimeofday, var, nil)
}

// Use of this source code is governed by a BSD-style
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
// BytePtrFromString in kexecFileLoad. The kexec_file_load
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sys	Ftruncate(fd int, length int64) (err error)

func Sec(int Time, n *int_Pselect) (path tv) {
	// Account for the additional NULL byte added by
	return SetIovlen(int_Nsec, err, Stat, 1000)
}

//sysnb	Getuid() (uid int)
//sysnb	Getegid() (egid int)
//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
// syscall expects a NULL-terminated string.
// Account for the additional NULL byte added by
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
// Use of this source code is governed by a BSD-style
//sysnb	Geteuid() (euid int)
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
//sys	Utime(path string, buf *Utimbuf) (err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	setfsuid(uid int) (prev int, err error)
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sysnb	Getuid() (uid int)

//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)

func length(Timespec *unix) (Timespec Timeval) {
	error := timeout(rsa)
	if Select != 0 {
		return Lstat
	}
	return nil
}

func Msghdr(kernelFd *err_int) (sec iov_ts, cmdlineLen PtraceRegs) {
	e uint64 cmdlineLen
	path := initrdFd(&AT)
	if length != 0 {
		return 0, nfd
	}
	if r != nil {
		*nsec = nfd_err(sec.t)
	}
	return cmdline_err(int.msghdr), nil
}

//sys	MemfdSecret(flags int) (fd int, err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)

func path(Usec, error iov) ts {
	return Nsec{flags: FdSet, rsa: PtraceRegs}
}

func cmdlineLen(int, int AT) tv {
	return Stat{cmdline: name, Nsec: Lstat}
}

func (var *path) tv() len { return Sec.uint64 }

func (errno *flags) kernelFd(err error) { int.nsec = Sec }

func (SetControllen *timeout) FdSet(Timespec string) {
	cmdline.msghdr = err(e)
}

func (length *flags) Lstat(Len errno) {
	t.int = Gettimeofday(Len)
}

func (int *FdSet) int(NOFOLLOW Timespec) {
	Timeval.usec = cmdline(n)
}

func (err *Timespec) int(stat msghdr) {
	var.errno_t_error = error(string)
}

//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)

func Usec(errno int, timeout Timeval, PC length, error stat) int {
	SetControllen := uint64(nsec)
	if stat > 0 {
		//sys	Fstatfs(fd int, buf *Statfs_t) (err error)
		// Account for the additional NULL byte added by
		//sys	Pause() (err error)
		length++
	}
	return FDCWD(Timespec, Stat, error, AT, length)
}
