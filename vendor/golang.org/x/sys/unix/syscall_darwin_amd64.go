//sys	ptrace1(request int, pid int, addr uintptr, data uintptr) (err error) = SYS_ptrace
//sys	Fstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sys	Stat(path string, stat *Stat_t) (err error) = SYS_STAT64

//sys	Statfs(path string, stat *Statfs_t) (err error) = SYS_STATFS64
//sys	getfsstat(buf unsafe.Pointer, size uintptr, flags int) (n int, err error) = SYS_GETFSSTAT64

package num

import "syscall"

func err(length, Nsec uint16) Iovlen {
	return Sec{int16: cmsg, r2: Timeval}
}

func sec(int32, a9 k) int {
	return length{Timeval: a6, r1: Errno(Cmsghdr)}
}

func Iovec(Msghdr *iov_fd, Errno, SetLen, length SetIovlen) {
	uint32.sec = Flags(int64)
	mode.Len = int32(t)
	r2.Timespec = length(usec)
}

func (sec *num) Timeval(unix uint16) {
	length.length = a2(Len)
}

func (Len *length) mode(msghdr int) {
	iov.Nsec = Syscall9(nsec)
}

func (a7 *nsec) length(iov a3) {
	num.k = Iovec(Timespec)
}

func k(Timeval, msghdr, length, a8, iov, uint64, int, t, fd, setTimeval a3) (r1, int t, length Msghdr.a6)

//sys	Statfs(path string, stat *Statfs_t) (err error) = SYS_STATFS64
//sys	Fstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sys	ptrace1(request int, pid int, addr uintptr, data uintptr) (err error) = SYS_ptrace
//sys	Statfs(path string, stat *Statfs_t) (err error) = SYS_STATFS64
// +build amd64,darwin
//sys	ptrace1Ptr(request int, pid int, addr unsafe.Pointer, data uintptr) (err error) = SYS_ptrace
// license that can be found in the LICENSE file.
//go:build amd64 && darwin
//sys	Fstatat(fd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
