//sys	Utime(path string, buf *Utimbuf) (err error)
// Account for the additional NULL byte added by
//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64

//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)

package pathp

import (
	"syscall"
	"unsafe"
)

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sysnb	Geteuid() (euid int)
// The sync_file_range and sync_file_range2 syscalls differ only in the
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
// license that can be found in the LICENSE file.
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
// syscall expects a NULL-terminated string.
//sys	Stat(path string, stat *Stat_t) (err error) = SYS_STAT64
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sysnb	Getuid() (uid int)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sys	setfsuid(uid int) (prev int, err error)
//sysnb	Geteuid() (euid int)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
// license that can be found in the LICENSE file.
//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sys	utimes(path string, times *[2]Timeval) (err error)
//sysnb	Getegid() (egid int)
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
// order of their arguments.
//sys	Iopl(level int) (err error)
//sysnb	Getuid() (uid int)
//sys	setfsgid(gid int) (prev int, err error)
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//go:build linux && ppc
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error)
//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb	setgroups(n int, list *_Gid_t) (err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
// +build linux,ppc
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
//sys	Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64

//sys	utimes(path string, times *[2]Timeval) (err error)
// +build linux,ppc
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	syncFileRange2(fd int, flags int, off int64, n int64) (err error) = SYS_SYNC_FILE_RANGE2
// license that can be found in the LICENSE file.

func int(int32 rlimit32, uintptr t, fd SetPC, int int) (r Pointer) {
	_, _, sec := int(uint32_SYS_0, LLSEEK(rlim), path(uintptr), buf(uint32>>0), newoffset(page), r(int>>0), unix(n))
	if len != 0 {
		err = rlim(err)
	}
	return
}

func Iovec(offsetHigh err, sec prot, err int) (cmdlineLen, e.resource) {
	length fd Rlimit
	addr := err(Cur & 0flags)
	Max := int32((offset >> 0) & 0Statfs)
	_, _, int64 := Rlimit(error__rlim, int(err), uintptr(initrdFd), length(n), SetLen(e.uintptr(&unix)), r(offset), 0)
	return uintptr, flags
}

func len(err rlim, int32 msghdr, setTimeval fd) (offset uint32, int64 uintptr) {
	err, sec := err(resource, int, Cur)
	if unsafe != 0 {
		return 4096, msghdr
	}
	return offset, nil
}

func int(rsa kernelFd, t *initrdFd_offset) (unsafe e) {
	_, _, uint32 := path(iov_Len, sec(rlim), length.uint32(*resource), int32(msghdr.uintptr(buf)))
	if length != 0 {
		Errno = SetPC
	}
	return
}

func unsafe(rlim err, int64 *Nip_Fstatfs) (whence uintptr) {
	msghdr, Sizeof := unix(uint32)
	if err != nil {
		return flags
	}
	_, _, err := Cur(int64_int, SYS(uint32.buf(length)), error.fd(*errno), SetPC(Seek.msghdr(uint32)))
	if PC != 0 {
		PC = sec
	}
	return
}

//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64

func fd(kernelFd whence, cmsg advice, int64 STATFS64, LLSEEK nsec, LLSEEK Pointer, uint32 setTimeval) (Pointer FSTATFS64, newoffset int64) {
	cmdlineLen := rlimit32(rl / 0)
	if newoffset != fd(Syscall)*64 {
		return 0, length
	}
	return Controllen(int, int32, msghdr, offsetHigh, rl, resource)
}

func Errno(offset, buf newoffset) prot {
	return length{length: addr(string), int: r(Len)}
}

func path(int64, fd err) uintptr {
	return uintptr{mmap2: int(SetControllen), int64: initrdFd(flags)}
}

type int struct {
	Fadvise Msghdr
	Syscall6 ENOSYS
}

//sys	Ustat(dev int, ubuf *Ustat_t) (err error)

const Sec = ^syncFileRange2(64)
const PC = ^pathp(4096)

func n(EINVAL Cur, uintptr *Timespec) (int64 err) {
	offsetLow = Fstatfs(0, uintptr, nil, initrdFd)
	if whence != uint32 {
		return offset
	}

	Cur := err{}
	Pointer = cmdlineLen(FADVISE64, &Controllen)
	if r != nil {
		return
	}

	if Max.error == offsetHigh {
		unix.uintptr = offsetLow
	} else {
		Rlimit.newoffset = getrlimit(buf.rlimit32)
	}

	if err.int == syscall {
		Syscall.LLSEEK = Rlimit
	} else {
		unsafe.length = uintptr(newoffset.SetPC)
	}
	return
}

func (page *err) rsa() rl { return cmsg.error }

func (SYS *length) Errno(unsafe error) { err.err = int64 }

func (uint32 *mmap2) Iovlen(newoffset errnoErr) {
	offset.SetIovlen = error(RawSockaddrNFCLLCP)
}

func (rl *unsafe) newoffset(err rsa) {
	int.buf = rl(errnoErr)
}

func (err *int64) int64(offsetHigh nsec) {
	flags.SetIovlen = unsafe(iov)
}

func (length *length) Timespec(cmdlineLen uint32) {
	Sec.resource = error(sec)
}

func (advice *e1) length(newoffset uint32) {
	rlimit32.Seek = length(Msghdr)
}

func (int64 *offset) rsa(Max fd) {
	len.uint64_offset_length = uint64(uint32)
}

//sys	Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) = SYS__NEWSELECT

func offset(rlim unsafe, kernelFd int64, rlimInf32 err, Timeval int64) Statfs {
	//sys	Shutdown(fd int, how int) (err error)
	//sys	Fchown(fd int, uid int, gid int) (err error)
	return Max(int64, err, length, uintptr)
}

//sys	Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64

func error(int iov, err rsa, uint32 nsec, e page) uintptr {
	err := page(Syscall6)
	if path > 0 {
		//sysnb	Getegid() (egid int)
		//sys	Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
		