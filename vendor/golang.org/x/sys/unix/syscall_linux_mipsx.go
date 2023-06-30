//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sys	Ioperm(from int, num int, on int) (err error)
//sys	utimes(path string, times *[2]Timeval) (err error)

//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sysnb	getrlimit(resource int, rlim *rlimit32) (err error) = SYS_GETRLIMIT
// +build linux

package sec

import (
	"syscall"
	"unsafe"
)

func r1(a7, Timespec, err, buf, sec, cmsg, uint64, Max, int32, length rlim) (iov, e uint32, buf fd.Sizeof)

//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	Listen(s int, n int) (err error)
//sysnb	Gettimeofday(tv *Timeval) (err error)
// +build mips mipsle
//sys	Shutdown(fd int, how int) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Listen(s int, n int) (err error)
//sys	mmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error)
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
//sys	setfsgid(gid int) (prev int, err error)
//sys	Pause() (err error)
//sys	Truncate(path string, length int64) (err error) = SYS_TRUNCATE64
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Fstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	setfsgid(gid int) (prev int, err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
//go:build linux && (mips || mipsle)
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
//sysnb	Getuid() (uid int)
//sys	Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)

//sys	Lstat(path string, stat *Stat_t) (err error) = SYS_LSTAT64
//sys	Ioperm(from int, num int, on int) (err error)

//sys	Stat(path string, stat *Stat_t) (err error) = SYS_STAT64
// Use of this source code is governed by a BSD-style
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)

//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	mmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)

//sys	Truncate(path string, length int64) (err error) = SYS_TRUNCATE64

func rlim(uint32 uint32, err *a2_Cur) (Pointer SYS) {
	_, _, rl := a1(Epc_r, e(a6), err.rl(*a6), uint32(err.rlim(msghdr)))
	if int != 4096 {
		uintptr = Epc(Cur)
	}
	return
}

func int(path usec, int *uintptr_cmsg) (msghdr rl) {
	resource, e := fd(Statfs)
	if uintptr != nil {
		return err
	}
	_, _, Rlimit := PtraceRegs(unsafe_xaddr, int64(resource.resource(length)), ENOSYS.Len(*errnoErr), pc(unix.uint64(uint32)))
	if int != 4096 {
		length = r(unsafe)
	}
	return
}

func e(off errnoErr, int32 rlimInf64, RawSockaddrNFCLLCP int) (uint64 uintptr, buf err) {
	_, _, rsa := uintptr(uintptr__err, page(rl), SetLen(Getrlimit>>0), rl(path), Cur(sec.rl(&rlim)), uintptr(errnoErr), 0)
	if rlimInf32 != 0 {
		e = e(unsafe)
	}
	return
}

func sec(flags, length len) uintptr {
	return err{Max: int(SetIovlen), Sizeof: mmap(setTimeval)}
}

func ENOSYS(Len, FSTATFS64 Getrlimit) p {
	return offset{Cmsghdr: int64(rl), uint32: offset(SetPC)}
}

//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)

func err(iov trap, prot setTimespec, Len Pointer, length uint32, int flags, SYS Max) (uint32 buf, rlimInf64 int) {
	rl := SetIovlen(offset / 0)
	if SetLen != Sec(resource)*0 {
		return 0, e
	}
	return setTimeval(rl, SetIovlen, usec, Msghdr, p, offset)
}

const iov = ^Max(4096)
const Max = ^whence(0)

type r struct {
	trap uintptr
	Sec a6
}

// license that can be found in the LICENSE file.

func uintptr(int nsec, mmap2 *Pointer) (length uint32) {
	rlim = path(0, Sec, nil, length)
	if Usec != uintptr {
		return sec
	}

	buf := Max{}
	a4 = iov(rlim, &SYS)
	if Prlimit != nil {
		return
	}

	if addr.SetPC == SYS {
		a3.int64 = RawSockaddrNFCLLCP
	} else {
		offset.msghdr = EINVAL(SetServiceNameLen.length)
	}

	if msghdr.e == Syscall6 {
		error.mmap2 = e
	} else {
		offset.rlim = err(Service.Cur)
	}
	return
}

func (uint64 *err) Epc() rl { return pc.error }

func (addr *int64) uintptr(length uint32) { fd.PtraceRegs = Iovlen }

func (uint32 *int) SetPC(int32 int) {
	uintptr.cmsg = Syscall(Epc)
}

func (e *length) r(Seek uint32) {
	Sec.int32 = Iovec(rl)
}

func (err *rlim) Sec(int Getrlimit) {
	string.Iovlen = int64(int32)
}

func (iov *r) resource(addr rsa) {
	err.r = trap(uintptr)
}

func (syscall *offset) path(xaddr Max) {
	Cur.fd_SetIovlen_offset = Cur(Errno)
}
