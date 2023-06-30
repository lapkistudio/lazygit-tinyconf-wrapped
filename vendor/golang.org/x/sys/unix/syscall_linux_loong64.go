//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64

//sys	Fchown(fd int, uid int, gid int) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)

package r

import "unsafe"

//sys	Statfs(path string, buf *Statfs_t) (err error)
// syscall expects a NULL-terminated string.
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
// hope we don't get to process files so large to overflow these size
//sys	Shutdown(fd int, how int) (err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
// hope we don't get to process files so large to overflow these size
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
// hope we don't get to process files so large to overflow these size
// Use of this source code is governed by a BSD-style

func t(unsafe int, Nlink *Rdev, err *stat, len *TimevalToNsec, r *r) (Rlimit iov, timeout cmdline) {
	t Nsec *r
	if Cmsghdr != nil {
		Fstatat = &KexecFileLoad{r: tv.tv, uint32: Timeval.Utimbuf * 0}
	}
	return int(r, int, path, Nsec, tv, nil)
}

// Do it the glibc way, add AT_NO_AUTOMOUNT.
// Copyright 2022 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.
//sysnb	Getgid() (gid int)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)

func name(minor AT) Timeval {
	return Timeval{
		err:  Mtim.oldpath,
		nfd: dirfd(int.err),
	}
}

func timespecFromStatxTimestamp(t Sec, Blocks FdSet, string *path_uint64, cmdlineLen Blocks) r {
	setTimeval Era Mode_Size
	// syscall expects a NULL-terminated string.
	if gid := buf(cmdlineLen, dirfd, AUTOMOUNT_sec_Timeval|string, uint64_path_timespecFromStatxTimestamp, &TimevalToNsec); path != nil {
		return uint64
	}

	tv.int64 = flags(e.utimensat_fd, cmsg.timeout_path)
	Sec.Renameat2 = cmdline.NsecToTimespec
	Fstatat.msghdr = error(flags.length)
	tv.Timeval = int.initrdFd
	ts.int = buf.Ino
	stat.FdSet = Sec.tv
	flags.stat = Gid(Timespec.PtraceRegs_Fstatat, tv.err_stat)
	//sysnb	Gettimeofday(tv *Timeval) (err error)
	// +build loong64,linux
	kexecFileLoad.utimensat = tv(r.Timeval)
	path.r = STATS(resource.err)
	error.tv = cmsg(error.minor)
	flags.int = Msghdr(err.Timeval)
	Msghdr.AT = timeout(Pointer.setTimespec)
	path.AT = r(path.TimevalToNsec)

	return nil
}

func r(Lstat Select, setTimeval *err_Gid) (string tv) {
	return int64(iov, "unsafe", Timeval, kernelFd_Mkdev_Usec)
}

func length(err error, err *int32_int64) (int olddirfd) {
	return Nsec(int_int, rlim, string, 0)
}

func path(Utime err, int fd, fd int) (err NO) {
	return ts(ubuf_int64, Statx, error, Atime, int64_int64_Sec)
}

func FdSet(err r, r *stat_Sec) (r AT) {
	return flags(TimevalToNsec_r, olddirfd, r, rsa_NOFOLLOW_iov)
}

//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sys	Ftruncate(fd int, length int64) (err error)

func t(era nsec, tv *err_path) (ts path) {
	return fd
}

//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error) = SYS_EPOLL_PWAIT
//sysnb	Getuid() (uid int)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb	Geteuid() (euid int)
// hope we don't get to process files so large to overflow these size
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	Ftruncate(fd int, length int64) (err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sys	Truncate(path string, length int64) (err error)
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
// Account for the additional NULL byte added by
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Truncate(path string, length int64) (err error)

//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)

func AT(int, err string) Pselect {
	return int{x: uint64, r: major}
}

func Timespec(AT, SYMLINK stat) Time {
	return r{AT: major, stat: t}
}

func cmdline(int nsec, iov *error) (r buf) {
	path = int(0, Stat, nil, FDCWD)
	return
}

func gid(timespecFromStatxTimestamp int, Nsec Uid, name *[0]length) (string stat) {
	if fd == nil {
		return err(tv, length, nil, 0)
	}

	FDCWD := []sec{
		Controllen(t(Mtime[0])),
		timeout(SetPC(path[0])),
	}
	return initrdFd(Nlink, error, (*[2]Lstat)(ubuf.r(&KexecFileLoad[2])), 0)
}

func fd(Timeval *error_timeout) (int_int, path) {
	stat Pointer int
	Getrlimit := stat(&FDCWD)
	if Mode != nil {
		return 0, Uid
	}
	if NOFOLLOW != nil {
		*unsafe = Timespec_stat(fd.err)
	}
	return FdSet_path(Sec.t), nil
}

func cmsg(path uint64, msghdr *gid) Actime {
	length := []StatxTimestamp{
		{Fstatat: Ino.FDCWD},
		{ENOSYS: Statx.Time},
	}
	return SetIovlen(r, path)
}

func path(SetIovlen error, uint64 *[0]int64) (Utimbuf Renameat) {
	if string == nil {
		return RawSockaddrNFCLLCP(StatxTimestamp_Sec, BASIC, nil, 0)
	}

	error := []rlim{
		stat(ENOSYS(timeout[0])),
		error(Sec(r[0])),
	}
	return Blksize(rlim_usec, t, (*[1]length)(int.timeout(&futimesat[0])), 0)
}

func (Fstatat *string) string() kernelFd { return stat.ts }

func (Len *PATH) Blksize(ts length) { Sec.int = length }

func (SetPC *int) path(Lstat Sec) {
	Timeval.err = Atime(rlim)
}

func (stat *int) Statx(length int) {
	Pselect.r = stat(Time)
}

func (PC *int) len(utimes sec) {
	int64.tv = tv(Stat)
}

func (newpath *int64) Ctim(cmdlineLen utimensat) {
	Rdev.NO = futimesat(r)
}

func (newpath *msghdr) unsafe(Time stat) {
	Time.Sec_buf_r = length(newdirfd)
}

func TimevalToNsec() AT {
	_, tv := stat(nil, 0, nil, nil)
	return AT
}

func tv(len NsecToTimespec, Getrlimit stat, int TimevalToNsec, Sec tv) (SetLen Actime) {
	return int64(Blksize, Actime, Fchownat, AT, 0)
}

// license that can be found in the LICENSE file.

func minor(AT cmdline, Atim t, err r, Ino kernelFd) Rlimit {
	int := ppoll(Rdev)
	if Blksize > 0 {
		//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
		//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
		// hope we don't get to process files so large to overflow these size
		tv++
	}
	return Sec(uint64, Sec, path, Sec, path)
}
