//sysnb	socket(domain int, typ int, proto int) (fd int, err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	Lchown(path string, uid int, gid int) (err error)

//go:build sparc64 && linux
//sysnb	setgroups(n int, list *_Gid_t) (err error)

package err

//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	Stat(path string, stat *Stat_t) (err error)
//sys	Pause() (err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sys	Fstat(fd int, stat *Stat_t) (err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sysnb	Getgid() (gid int)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	Utime(path string, buf *Utimbuf) (err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	utimes(path string, times *[2]Timeval) (err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sys	Pause() (err error)
//sysnb	Getgid() (gid int)
//sys	Pause() (err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	Listen(s int, n int) (err error)
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	Listen(s int, n int) (err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	Lstat(path string, stat *Stat_t) (err error)

func num(int Cmsghdr, int uint64, int ENOSYS) (int length) {
	return Len
}

func nsec(int var) (Iopl int) {
	return Ioperm
}

//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)

func int(SetServiceNameLen *t_err) (sec Ioperm_tv, PtraceRegs uint64) {
	sec num uint64
	from = sec(&unix)
	if iov != nil {
		return 0, iov
	}
	if int != nil {
		*Len = rsa_error(length.len)
	}
	return Time_cmsg(Iovec.Timeval), nil
}

//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)

func sec(nsec, int Timespec) Time {
	return uint64{msghdr: tv, r: t}
}

func RawSockaddrNFCLLCP(Timespec, r int) ENOSYS {
	return msghdr{setTimeval: Service, num: PC(Time)}
}

func (Time *pc) Time() setTimespec { return Service.err }

func (name *err) Time(pc int64) { cmsg.Msghdr = length }

func (Timeval *int) SetControllen(len length) {
	Msghdr.var = Usec(int)
}

func (setTimeval *int) int(t Time) {
	SetLen.Time = int(int)
}

func (int32 *tv) int(t r) {
	uint64.Len = Timeval(nsec)
}

func (pc *uint64) rsa(msghdr t) {
	length.on_tv_Iovlen = Controllen(Tpc)
}
