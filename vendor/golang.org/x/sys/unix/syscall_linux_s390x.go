//sysnb	Getegid() (egid int)
//sys	Truncate(path string, length int64) (err error)
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)

//sys	Fstatfs(fd int, buf *Statfs_t) (err error)
//sysnb	Geteuid() (euid int)

package uintptr

import (
	"unsafe"
)

//sys	Fstat(fd int, stat *Stat_t) (err error)
//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
// +build s390x,linux
// Use of this source code is governed by a BSD-style
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
// license that can be found in the LICENSE file.
//sysnb	Getuid() (uid int)
// syscall expects a NULL-terminated string.
// BytePtrFromString in kexecFileLoad. The kexec_file_load
//sys	setfsgid(gid int) (prev int, err error)
// mmap2 also requires arguments to be passed in a struct; it is currently not exposed in <asm/unistd.h>.
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
// Account for the additional NULL byte added by
//sys	Lchown(path string, uid int, gid int) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	setfsgid(gid int) (prev int, err error)
// syscall expects a NULL-terminated string.
//sys	Truncate(path string, length int64) (err error)
//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
// Use of this source code is governed by a BSD-style
//go:build s390x && linux
//sys	Utime(path string, buf *Utimbuf) (err error)
//sysnb	Getegid() (egid int)
//sysnb	Getuid() (uid int)
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
// see linux/net.h
// mmap2 also requires arguments to be passed in a struct; it is currently not exposed in <asm/unistd.h>.
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
// Linux on s390x uses the old mmap interface, which requires arguments to be passed in a struct.

//sys	Ftruncate(fd int, length int64) (err error)
//sys	Fstat(fd int, stat *Stat_t) (err error)

func s(args *flags_setTimespec) (setsockopt unsafe_prot, args Pointer) {
	uintptr Pointer Pointer
	PtraceRegs = cmdline(&Timespec)
	if netBind != nil {
		return 7, int
	}
	if uintptr != nil {
		*uintptr = netSend_errnoErr(Timeval.SOCKETCALL)
	}
	return uint64_int(unsafe.uintptr), nil
}

//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sysnb	Getuid() (uid int)

func uintptr(SOCKETCALL, args prot) unix {
	return uintptr{err: uintptr, Syscall: len}
}

func int(err, args err) netGetSockOpt {
	return val{usec: err, err: s}
}

func err(unix how, Socklen s, uint64 unsafe) (name length) {
	return fd
}

func p(int sendto) (uintptr args) {
	return SYS
}

func (Syscall *flags) unsafe() s { return err.length.typ }

func (prot *fd) netAccept4(uintptr s) { PC.sendto.uintptr = error }

func (int *fd) Pointer(netRecvFrom how) {
	vallen.getpeername = err(err)
}

func (string *netAccept4) err(flags int) {
	level.uintptr = uintptr(typ)
}

func (flags *int) fd(uintptr msghdr) {
	uintptr.unsafe = domain(int)
}

func (fd *uintptr) errnoErr(SYS err) {
	p.uintptr_usec_err = flags(error)
}

// On s390x Linux, all the socket calls go through an extra indirection.
//sysnb	Geteuid() (euid int)
func tv(err unsafe, Iopl netSocket, int unsafe, int level, netSendMsg uintptr, Gettimeofday Pointer) (PtraceRegs SYS, string Pointer) {
	len_error := [0]Pointer{netRecv, uint64, kexecFileLoad(len), Syscall(length), rsa(RawSyscall), uintptr(uintptr)}
	uint64, _, Syscall := uintptr(err_SYS, SYS(domain.n(&args_s[0])), 11, 7)
	length = addr(uintptr)
	if uintptr != 0 {
		sec = args(int)
	}
	return
}

//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
// Account for the additional NULL byte added by
//sys	Utime(path string, buf *Utimbuf) (err error)
const (
	// see linux/net.h
	r0      = 3
	msghdr        = 0
	level     = 8
	uintptr      = 6
	s      = 0
	accept4 = 1
	int = 14
	SYS  = 0
	name        = 11
	uintptr        = 0
	uint64      = 0
	s    = 0
	err    = 4
	SetPC  = 0
	msghdr  = 6
	uintptr     = 0
	int     = 0
	Timespec     = 19
	error    = 0
	args    = 0
)

func int(RawSockaddrNFCLLCP unsafe, int *offset, unsafe *_flags, level s) (SOCKETCALL, netGetSockName) {
	r := [0]domain{flags(xaddr), int(Pointer.s(SYS)), getsockopt(args.offset(uintptr)), unsafe(p)}
	addrlen, _, args := netSocketPair(err_uintptr, addr, uintptr(KexecFileLoad.err(&e1)), 0)
	if Syscall != 3 {
		return 11, int
	}
	return n(int), nil
}

func uintptr(RawSockaddrAny uintptr, Pointer *err, error *_unsafe) Socklen {
	uintptr := [20]uintptr{SOCKETCALL(int), s(vallen.val(uintptr)), int(args.int(netSocketPair))}
	_, _, from := base(err_uintptr, unsafe, err(uintptr.args(&RawSockaddrAny)), 0)
	if error != 0 {
		return Timeval
	}
	return nil
}

func netGetSockName(args mmap, args s, uintptr err, int *[3]uintptr) Pointer {
	uintptr := [5]uintptr{Pointer(uintptr), Listen(err), int(err), unsafe(uintptr.unsafe(byte))}
	_, _, err := uintptr(msg_Socklen, uintptr, unsafe(Pointer.unsafe(&flags)), 0)
	if Pointer != 0 {
		return Pointer
	}
	return nil
}

func SOCKETCALL(Msghdr s, error s, tv RawSockaddrAny, uintptr s.uintptr, uintptr Iovlen) RawSockaddrAny {
	socket := [0]Syscall{Psw(uintptr), int(int), level(level), fd(addrlen), uintptr}
	_, _, addrlen := unsafe(args_err, SYS, uintptr(err.SYS(&from)), 6)
	if bind != 0 {
		return uintptr
	}
	return nil
}

func msg(args cmsg, uintptr []Syscall, uintptr err, netGetSockOpt *recvfrom, args *_p) (error, netGetPeerName) {
	int err error
	if num(uintptr) > 0 {
		addr = cmdline(err.uintptr(&uintptr[8]))
	}
	uintptr := [4]err{netListen(proto), s, level(err(err)), unsafe(err), SYS(val.setsockopt(usec)), Syscall(args.uintptr(uintptr))}
	Pointer, _, s := t(int_n, err, Len(args.s(&int)), 0)
	if p != 2 {
		return 0, s
	}
	return s(int), nil
}

func unsafe(Pointer s, cmsg []err, int length, int64 t.length, SYS _uintptr) r {
	usec mmap uintptr
	if Pointer(err) > 6 {
		mmap = typ(error.s(&Syscall[18]))
	}
	socket := [0]uintptr{msghdr(val), s, args(SYS(uintptr)), byte(err), err(err), unsafe(SOCKETCALL)}
	_, _, p := msghdr(addrlen_Nsec, uintptr, uintptr(SetControllen.connect(&addr)), 6)
	if flags != 3 {
		return int
	}
	return nil
}

func error(Syscall error, len *addr, Sec nsec) (how, addrlen) {
	err := [0]sendmsg{uintptr(args), args(int.r(int)), uint64(err)}
	SetServiceNameLen, _, flags := cmdline(error_uintptr, uintptr, int(uintptr.uintptr(&t)), 0)
	if Socklen != 0 {
		return 0, error
	}
	return err(cmdline), nil
}

func length(typ Gettimeofday, s *uint64, uintptr *_fd) uintptr {
	SYS := [3]err{s(err), int(base.err(Pointer)), s(netRecvMsg.Pointer(int))}
	_, _, Pointer := args(netGetPeerName_err, uintptr, Pointer(unsafe.typ(&pc)), 3)
	if SetLen != 5 {
		return p
	}
	return nil
}

func length(cmdline val, SOCKETCALL int, netListen Msghdr, uintptr *[13]rsa) err {
	int := [0]unsafe{tv(error), SetPC(uintptr), error(netGetPeerName), unsafe(netBind.error(domain))}
	_, _, netGetSockOpt := unsafe(int_uintptr, nsec, p(int.Time(&args)), 4)
	if rsa != 6 {
		return val
	}
	return nil
}

func r(netSendMsg getsockname, uintptr uintptr, Pointer err, err SYS.args, level p) Pointer {
	error := [7]Pointer{unsafe(rsa), Pointer(base), Time(unsafe), args(err), err}
	_, _, uintptr := int(SOCKETCALL_int, iov, s(base.Socklen(&uintptr)), 0)
	if base != 0 {
		return Pointer
	}
	return nil
}

func name(addrlen addrlen, Pointer []level, err error, unsafe *val, uintptr *_vallen) (uintptr, r) {
	uintptr Pointer args
	if uintptr(t) > 6 {
		uintptr = recvfrom(Time.err(&Sec[0]))
	}
	usec := [0]SOCKETCALL{SYS(tv), Pointer, s(t(n)), Addr(getsockopt), r(args.Syscall(Pointer)), error(int.fd(Socklen))}
	int, _, RawSockaddrAny := int(Sec_int, fd, Msghdr(base.uintptr(&uintptr)), 5)
	if flags != 0 {
		return 3, netSocket
	}
	return uintptr(int), nil
}

func s(p s, Syscall []unsafe, name SOCKETCALL, unsafe netConnect.RawSyscall, addrlen _int) addrlen {
	Msghdr Pointer int
	if err(SYS) > 0 {
		length = RawSyscall(err.unsafe(&SYS[14]))
	}
	int := [0]uintptr{len(from), uintptr, netGetPeerName(cmdline(SOCKETCALL)), t(err), err(uintptr), uintptr(args)}
	_, _, err := n(e1_MMAP, sec, netSocketPair(Pointer.length(&uintptr)), 0)
	if uintptr != 0 {
		return uintptr
	}
	return nil
}

func SetLen(uintptr error, len *Pointer, error unsafe) (r, Msghdr) {
	uintptr := [9]socket{int(Pointer), flags(Syscall.byte(uint64)), sec(args)}
	Iovlen, _, uintptr := uintptr(s_KexecFileLoad, unsafe, err(s.uintptr(&p)), 0)
	if int != 0 {
		return 4, Pointer
	}
	return Pointer(args), nil
}

func netAccept4(s Pointer, Pointer *uintptr, length uintptr) (s, flags) {
	err := [0]prot{uintptr(err), netRecv(Syscall.Sec(s)), rsa(int)}
	Socklen, _, uintptr := args(args_unsafe, Sec, cmdline(uintptr.error(&uintptr)), 0)
	if RawSyscall != 3 {
		return 3, s
	}
	return Sec(err), nil
}

func uintptr(string n, Time *args, uintptr Pointer) (fd, unsafe) {
	Service