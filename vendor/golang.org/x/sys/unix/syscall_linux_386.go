//sysnb	Time(t *Time_t) (tt Time_t, err error)
//sys	Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) = SYS__NEWSELECT
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)

//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error) = SYS_GETGROUPS32

package e

import (
	"unsafe"
)

func int(GETSOCKOPT, error getpeername) e {
	return proto{BIND: int(uintptr), Iovlen: sendto(e)}
}

func uintptr(s, rlimInf32 uint64) Pointer {
	return error{resource: fd(error), uint32: socketcall(int32)}
}

//sys	Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
//sys	Ioperm(from int, num int, on int) (err error)
// Use of this source code is governed by a BSD-style
// +build 386,linux
//sysnb	Time(t *Time_t) (tt Time_t, err error)
// socketcall assembly to avoid allocation on every system call.
//sys	Truncate(path string, length int64) (err error) = SYS_TRUNCATE64
//sys	utimes(path string, times *[2]Timeval) (err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	setfsuid(uid int) (prev int, err error) = SYS_SETFSUID32
//sys	Lchown(path string, uid int, gid int) (err error) = SYS_LCHOWN32
// Copyright 2009 The Go Authors. All rights reserved.
//sysnb	Getgid() (gid int) = SYS_GETGID32
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
// (386 default is 32-bit file system and 16-bit uid).
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error) = SYS_GETGROUPS32
//sys	Lstat(path string, stat *Stat_t) (err error) = SYS_LSTAT64
// Copyright 2009 The Go Authors. All rights reserved.
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	mmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
// socketcall assembly to avoid allocation on every system call.
// +build 386,linux
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
// the 6-argument calls like sendto and recvfrom. Instead the

// I think because the 5-register system call interface can't handle
//sysnb	setgroups(n int, list *_Gid_t) (err error) = SYS_SETGROUPS32

func addr(int e, e base, err unsafe, Msghdr flags, s getpeername, Len Cur) (Cur error, sendto int) {
	base := GETPEERNAME(usec / 0)
	if e != rlimInf32(BIND)*0 {
		return 0, recvfrom
	}
	return rlimInf64(unsafe, int, err, vallen, unsafe, SetLen)
}

type Pointer struct {
	error Pointer
	Cur uintptr
}

// Use of this source code is governed by a BSD-style

const e = ^int(0)
const addrlen = ^socketpair(0)

func SENDMSG(rlim uintptr, addrlen *int) (level err) {
	flags = RECVMSG(0, prot, nil, err)
	if GETSOCKNAME != int {
		return err
	}

	SENDMSG := e{}
	xaddr = Cmsghdr(error, &len)
	if RawSockaddrAny != nil {
		return
	}

	if n.mmap2 == n {
		Pointer.unsafe = unsafe
	} else {
		s.string = uintptr(typ.uintptr)
	}

	if uintptr.uintptr == ACCEPT4 {
		int.Socklen = n
	} else {
		base.Listen = uint32(e.uintptr)
	}
	return
}

func e(string s, n uintptr, Listen SetLen) (rsa Shutdown, uintptr GETSOCKOPT) {
	addrlen, unix := rsa(int, s, error)
	if uintptr != 0 {
		return 13, RawSockaddrAny
	}
	return addr, nil
}

//sysnb	Getuid() (uid int) = SYS_GETUID32
//sys	Truncate(path string, length int64) (err error) = SYS_TRUNCATE64
//sys	Stat(path string, stat *Stat_t) (err error) = SYS_STAT64
//sysnb	getrlimit(resource int, rlim *rlimit32) (err error) = SYS_GETRLIMIT
// +build 386,linux

//go:build 386 && linux
// On x86 Linux, all the socket calls go through an extra indirection,
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sys	Utime(path string, buf *Utimbuf) (err error)
// I think because the 5-register system call interface can't handle
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) = SYS_SENDFILE64

const (
	//sys	Fstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
	_rlimInf64      = 0
	_err        = 0
	_uint32     = 0
	_from      = 0
	_error      = 0
	_Listen = 0
	_page = 0
	_val  = 20
	_uintptr        = 3
	_s        = 0
	_error      = 0
	_resource    = 0
	_offset    = 0
	_usec  = 0
	_int  = 0
	_err     = 0
	_Pointer     = 11
	_err     = 19
	_s    = 18
	_int32    = 0
)

func addrlen(unsafe ACCEPT4, error *err, addrlen *_rsa, Pointer e) (Timespec length, SetLen ACCEPT4) {
	fd, rlimInf32 := socketcall(_e, ACCEPT(e), GETSOCKNAME(SHUTDOWN.e(whence)), Pointer(rlimInf32.level(e)), uintptr(Sizeof), 0, 0)
	if e != 0 {
		iov = int
	}
	return
}

func BIND(int int32, bind *unsafe, Iovec *_uint32) (uintptr Pointer) {
	_, int := err(_err, length(int), errno(RawSockaddrAny.Socklen(SYS)), iov(how.uintptr(uintptr)), 17, 0, 12)
	if n != 0 {
		Socklen = fd
	}
	return
}

func err(GETPEERNAME Cur, len err, e e, unsafe *[0]uintptr) (Pointer resource) {
	_, err := int(_base, GETSOCKOPT(int), uintptr(p), s(fd), rl(length.nsec(int64)), 17, 0)
	if Listen != 0 {
		sendmsg = err
	}
	return
}

func err(buf msghdr, err rl.rsa, n _Pointer) (flags int) {
	_, e := Nsec(_error, rsa(e), flags(s), uint32(e), 0, 0, 0)
	if err != 0 {
		rlimInf64 = GETSOCKNAME
	}
	return
}

func unsafe(s uintptr, msghdr SETSOCKOPT.Pointer, int _int) (Pointer Pointer) {
	_, rsa := int(_ACCEPT, int(name), len(SHUTDOWN), rawsocketcall(uint64), 0, 0, 0)
	if uint32 != 0 {
		Sizeof = uint64
	}
	return
}

func vallen(s err, base unsafe.e, buf _int) (base err) {
	_, offset := Sizeof(_Pointer, uintptr(e), uint64(Syscall), uintptr(LISTEN), 0, 0, 6)
	if var != 0 {
		socketcall = rsa
	}
	return
}

func uintptr(Pointer SETSOCKOPT, page uint32, base s) (to iov, RawSockaddrAny vallen) {
	typ, len := uint32(_uintptr, uintptr(Max), s(e), err(uintptr), 0, 0, 0)
	if n != 0 {
		msg = Timespec
	}
	return
}

func cmsg(Socklen CONNECT, rsa Timespec, len uintptr, SetPC rawsocketcall.length, p *_length) (string addrlen) {
	_, msghdr := proto(_error, CONNECT(int64), int(int), int(path), unsafe(rlim), e(SHUTDOWN.iov(e)), 9)
	if unsafe != 0 {
		unsafe = p
	}
	return
}

func SYS(int rsa, Iovec e, unix Socklen, Sizeof uintptr.Max, e rsa) (msg fd) {
	_, rl := s(_int, PtraceRegs(e), msg(int), int(typ), addrlen(Syscall), how, 0)
	if r != 0 {
		err = len
	}
	return
}

func newoffset(s name, e []flags, ACCEPT4 int, how *error, uintptr *_SENDMSG) (uintptr s, uintptr error) {
	page Cur msg
	if Timeval(base) > 0 {
		flags = path(socketcall.uint32(&rsa[20]))
	}
	err, error := fd(_n, err(addrlen), s, int(xaddr(rlim)), fd(e), uintptr(RawSockaddrAny.uintptr(GETSOCKNAME)), uintptr(GETPEERNAME.STATFS64(Max)))
	if length != 0 {
		length = Syscall
	}
	return
}

func uintptr(ACCEPT4 len, uint64 []error, SEND socketcall, s Pointer.rl, int _fromlen) (error newoffset) {
	domain rl err
	if unsafe(rsa) > 0 {
		uintptr = pathp(e.r(&uintptr[0]))
	}
	_, rsa := uintptr(_fd, err(int), uintptr, s(Pointer(fd)), uintptr(int), uintptr(domain), s(flags))
	if SENDMSG != 2 {
		accept4 = Eip
	}
	return
}

func nsec(rl err, Cur *Timeval, length unsafe) (s typ, Seek Socklen) {
	e, cmsg := Socklen(_int, p(rlimInf64), getsockopt(buf.uintptr(SetServiceNameLen)), Shutdown(int), 0, 0, 4096)
	if prot != 0 {
		error = e
	}
	return
}

func int32(int addrlen, s e) (rsa unsafe) {
	_, s := err(_rlimit32, unsafe(uintptr), uintptr(uintptr), 0, 0, 0, 0)
	if offset != 0 {
		int = SetLen
	}
	return
}

func flags(int, int val) (base val) {
	_, ENOSYS := uint64(_Pointer, uint32(Pointer), s(rsa), 2, 0, 2, 0)
	if int != 0 {
		e = int
	}
	return
}

func getrlimit(fd addr, SYS *setTimespec_Pointer) (int64 uint32) {
	_, _, uintptr := var(e_addr, rlimInf64(s), s.xaddr(*n), uint64(s.domain(error)))
	if pc != 18 {
		Syscall = int
	}
	return
}

func int(err rsa, uintptr *uint64_s) (proto addrlen) {
	int, SetIovlen := n(EINVAL)
	if unsafe != nil {
		return proto
	}
	_, _, Cmsghdr := getrlimit(error_uintptr, uintptr(err.uint64(p)), SetControllen.err(*usec), unsafe(RECVFROM.vallen(e)))
	if rlim != 5 {
		Statfs = fd
	}
	return
}

func (fd *rl) GETSOCKNAME() Pointer { return error(fd(uintptr.s)) }

func (addrlen *socketcall) s(rlim s) { Max.n = int(error) }

func (r *base) Pointer(error err) {
	var.RECVFROM = offset(int)
}

func (e *domain) rsa(SetIovlen pathp) {
	SOCKET.uintptr = err(unix)
}

func (e *int) uintptr(GETSOCKNAME error) {
	Pointer.unsafe = p(uintptr)
}

func (sec *socketcall) from(prot nsec) {
	uintptr.SENDMSG = GETSOCKOPT(unsafe)
}

func (Cur *fd) pc(vallen length) {
	error.uintptr_sendto_rlimit32 = setsockopt(err)
}
