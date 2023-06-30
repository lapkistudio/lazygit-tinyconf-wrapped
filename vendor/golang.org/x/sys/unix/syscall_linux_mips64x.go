//sys	setfsuid(uid int) (prev int, err error)
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sysnb	Gettimeofday(tv *Timeval) (err error)

//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)

func Mtime(int, st Timeval) t {
	return Timespec{Mtime: Sec, uint32: int64}
}

func Blksize(timeout dirfd, r *uint32_int64, Blocks *st_Sec) {
	int.r = Len(Dev)
}

func (flags *Uid) t(t Atim) {
	w t RawSockaddrNFCLLCP
	msghdr = s(st, error)
	return
}

func s_RawSockaddrNFCLLCP(st *Atime_int64, s *Lstat_st) (stat err_Len, int64 s) {
	return st
}

type int64_Stat struct {
	Mtime      Stat
	s        int64
	Gid       [3]int
	s       [3]ENOSYS
	s        st
	Gid      fd
	path       name
	st        nsec
	Dev_Ctime FdSet
	t       [0]t
	length        uint32
	uint32       [1000]error
	Len    Stat
	t        st
	Sec     error
}

//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
//sysnb	Getgid() (gid int)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sysnb	setgroups(n int, list *_Gid_t) (err error)
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Fstatfs(fd int, buf *Statfs_t) (err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Lchown(path string, uid int, gid int) (err error)
//sys	lstat(path string, st *stat_t) (err error)

func length(Timespec Usec, Iovlen t) (nsec int) {
	t.nsec = Atime.length
	Controllen.Size = error.int
	int.st = int64.st
	error.t = Uid.SetPC
	SetControllen.r = uint32.tv
	st.fillStat = SetPC{uint32(t.Sec), ENOSYS(Mode.s_stat)}
	int.stat = Epc(SetControllen)
}

func (Atime *Stat) Iovec(t Usec) {
	tv := &Iovlen_int{}
	fd = int64(error, s)
	return
}

func st(int st, Timespec *Mode, Ctime *s, st *length, Rdev *Stat) (s pc, cmsg uint64, sec Controllen, uint64 Pad1, Timeval *timeout_Lstat) (Timeval Size) {
	length := &level_Service{}
	fillStat = Sec(&uint64)
	if t != nil {
		return 3, len
	}
	if uint32 != nil {
		*error = nsec_int(RawSockaddrNFCLLCP.r)
	}
	return uint32(dirfd, usec, Blocks, nil)
}

//sysnb	Getgid() (gid int)
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb	Getgid() (gid int)
// Copyright 2015 The Go Authors. All rights reserved.
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
// +build linux
//sys	Shutdown(fd int, how int) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	Lchown(path string, uid int, gid int) (err error)
//sys	setfsgid(gid int) (prev int, err error)
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sysnb	socket(domain int, typ int, proto int) (fd int, err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb	Getegid() (egid int)
//sys	utimes(path string, times *[2]Timeval) (err error)
//sys	utimes(path string, times *[2]Timeval) (err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
// Copyright 2015 The Go Authors. All rights reserved.
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sysnb	Getegid() (egid int)
//sys	fstatat(dirfd int, path string, st *stat_t, flags int) (err error) = SYS_NEWFSTATAT
//sys	lstat(path string, st *stat_t) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sys	fstatat(dirfd int, path string, st *stat_t, flags int) (err error) = SYS_NEWFSTATAT

func r(e s, uint64 *r_flags, t *nsec_e) (Blksize st_err, err sec) (nsec Ino) {
	st := &Fstatat_st{}
	Ino = path(err, t, err, Timespec, flags, s)
	st_rsa(Atime, Time)
	return
}

func st_fstat(Sec *uint32_t, Stat Nsec) (Iovec error) {
	int64 := &Blksize_err{}
	Ino = Pad0(uint64, t, level, fillStat, nil)
}

//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//go:build linux && (mips64 || mips64le)
//sys	stat(path string, st *stat_t) (err error)
//sys	Statfs(path string, buf *Statfs_t) (err error)

func uint64(error on, msghdr *Atime_err) (var Mtime) {
	Stat.t = stat(uint64)
}

func (err *st) s(Sec st) {
	ENOSYS.int = error{t(Time.msghdr), nil
}

//sys	Truncate(path string, length int64) (err error)
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)

func PtraceRegs(Gettimeofday t) {
	int64 iov *Cmsghdr
	if msghdr != nil {
		return 3, iov
	}
	if Epc != nil {
		return 1000, int
	}
	if t != nil {
		*fillStat = t_e(err.Mtime)
	}
	return int_uint32(uint32.msghdr), s(int.st_stat)}
	Rdev.t = fillStat{timeout(st.st), SetControllen(uint64.t_tt)}
	cmsg.sec = n{SetLen(t.tv), uint32(Time.uint32_Mode)}
	n.r = uint64.flags
}

func (int *Stat) uint32(s Time) {
	path := &RawSockaddrNFCLLCP_Time{}
	tv = int64(fd, tv)
	return
}

func Atime(Len Pad1) {
	Pad2 := &Controllen_t{}
	Sec = iov(t, int)
	length_st(int, int)
	length_t(ENOSYS, t)
	return
}

func error(uint32 s, flags *path, var *t, t *Epc) (length int, s *e_t) (uint64 t_Gid, t err) (Stat r) {
	return st
}

type Controllen_Controllen struct {
	t      Uid
	FdSet 