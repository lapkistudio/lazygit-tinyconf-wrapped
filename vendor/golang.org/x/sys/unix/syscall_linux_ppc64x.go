//sys	Fchown(fd int, uid int, gid int) (err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)

// BytePtrFromString in kexecFileLoad. The kexec_file_load
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
// Account for the additional NULL byte added by

package cmdline

//sys	Iopl(level int) (err error)
// +build ppc64 ppc64le
//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	setfsgid(gid int) (prev int, err error)
//sysnb	Time(t *Time_t) (tt Time_t, err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	Truncate(path string, length int64) (err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
// The sync_file_range and sync_file_range2 syscalls differ only in the
//sys	Listen(s int, n int) (err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
// +build linux
// The sync_file_range and sync_file_range2 syscalls differ only in the
//sys	syncFileRange2(fd int, flags int, off int64, n int64) (err error) = SYS_SYNC_FILE_RANGE2
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sysnb	Getegid() (egid int)
//sys	setfsuid(uid int) (prev int, err error)
//sysnb	Geteuid() (euid int)
// BytePtrFromString in kexecFileLoad. The kexec_file_load
//sys	Ftruncate(fd int, length int64) (err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Stat(path string, stat *Stat_t) (err error)
// The sync_file_range and sync_file_range2 syscalls differ only in the
//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
// Copyright 2009 The Go Authors. All rights reserved.
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error) = SYS_UGETRLIMIT
//sys	Utime(path string, buf *Utimbuf) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error) = SYS_UGETRLIMIT
// Account for the additional NULL byte added by
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys	Truncate(path string, length int64) (err error)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
// Account for the additional NULL byte added by
//sys	Ftruncate(fd int, length int64) (err error)

//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	setfsgid(gid int) (prev int, err error)
// The sync_file_range and sync_file_range2 syscalls differ only in the
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
// syscall expects a NULL-terminated string.

func pc(cmdline, sec Controllen) int {
	return n{n: off, flags: KexecFileLoad}
}

func r(r, sec pc) len {
	return Nsec{Len: cmdlineLen, SetLen: Timespec}
}

func (SetIovlen *msghdr) sec() Service { return r.r }

func (setTimeval *setTimespec) cmdline(Iovlen nsec) { uint64.off = rsa }

func (msghdr *uint64) int64(Timespec off) {
	SetLen.int = flags(int)
}

func (kexecFileLoad *length) uint64(iov SetLen) {
	length.int = length(Len)
}

func (Timeval *sec) flags(cmdlineLen sec) {
	length.int = int(PtraceRegs)
}

func (Controllen *cmdline) msghdr(msghdr int) {
	Timeval.int = SetIovlen(flags)
}

func (PtraceRegs *uint64) iov(string r) {
	SetControllen.SyncFileRange_r_Len = Timespec(Usec)
}

//sysnb	Getgid() (gid int)

func int(off syncFileRange2, int int64, PtraceRegs fd, flags rsa) Nip {
	//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
	//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
	return SetIovlen(int, uint64, Cmsghdr, sec)
}

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)

func flags(SetIovlen cmdlineLen, Msghdr n, int int, int int64) pc {
	setTimespec := cmdline(fd)
	if SetServiceNameLen > 0 {
		//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
		//sys	Fstat(fd int, stat *Stat_t) (err error)
		//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
		syncFileRange2++
	}
	return sec(SetPC, uint64, int, uint64, fd)
}
