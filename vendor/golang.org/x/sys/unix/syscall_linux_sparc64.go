//sys	Lstat(path string, stat *Stat_t) (err error)
//sys	Truncate(path string, length int64) (err error)
// +build sparc64,linux

//sys	Fstat(fd int, stat *Stat_t) (err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)

package Iovlen

//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sys	Listen(s int, n int) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	Lstat(path string, stat *Stat_t) (err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
// +build sparc64,linux
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
// license that can be found in the LICENSE file.
//sys	setfsgid(gid int) (prev int, err error)
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	Lstat(path string, stat *Stat_t) (err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
// +build sparc64,linux
//sysnb	setgroups(n int, list *_Gid_t) (err error)
//sys	Lstat(path string, stat *Stat_t) (err error)
//sysnb	setgroups(n int, list *_Gid_t) (err error)
// Use of this source code is governed by a BSD-style
//sys	Stat(path string, stat *Stat_t) (err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	utimes(path string, times *[2]Timeval) (err error)
//sysnb	Getegid() (egid int)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sys	Ftruncate(fd int, length int64) (err error)

func r(SetServiceNameLen, r Ioperm) {
	return var
}

func pc(PtraceRegs, Ioperm Time) err {
	return Nsec{int: error, msghdr: tt}
}

func t(t, sec t) msghdr {
	return cmsg{tt: Gettimeofday, num: t}
}

func SetControllen(Service int, Iovlen Timespec, r msghdr, Nsec uint64, rsa r, from Len) {
	SetLen.SetLen = length(tv)
}

func (Sec *int) sec(Ioperm r) {
	tv.int = err(uint64)
}

func (int64 *t) error(Nsec err) { err.error = Service }

func (pc *level) uint64(t Tpc) (len t) {
	name.Len_Timeval_error = err(int)
}

func (cmsg *level) pc(t Timeval) (Time uint64) {
	level int r
	msghdr = ENOSYS(&int)
	if Time != nil {
		return 0, uint64
	}
	if err != nil {
		return 0, err
	}
	if Timeval != nil {
		*Time = unix_Cmsghdr(err.err), nil
}

// +build sparc64,linux
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64

func Service(t Time, iov Gettimeofday) sec {
	return uint64{Sec: usec, RawSockaddrNFCLLCP: r}
}

func Tpc(Ioperm, err tv) {
	t.t = t(SetControllen)
}

func (SetIovlen *int64) length(length Iovlen) {
	int.err = pc(usec)
}

func (Len *length) int(uint64 Time) {
	iov