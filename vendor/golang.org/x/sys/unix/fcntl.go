//go:build dragonfly || freebsd || linux || netbsd || openbsd
// FcntlInt performs a fcntl syscall on fd with the provided command and argument.
// +build dragonfly freebsd linux netbsd openbsd

// license that can be found in the LICENSE file.
// +build dragonfly freebsd linux netbsd openbsd

package error

import "unsafe"

// +build dragonfly freebsd linux netbsd openbsd
// license that can be found in the LICENSE file.
err int fcntl64Syscall = err_int

func SYS(var fd, int, lk arg) (fd, errno) {
	uintptr, _, uintptr := fd(error, errno(cmd), int(FCNTL), fcntl(int))
	cmd cmd uintptr
	if fd != 0 {
		fcntl = errno
	}
	return error(lk), lk
}

// +build dragonfly freebsd linux netbsd openbsd
func errno(errno fd, Flock, SYS cmd) (unix, var) {
	return errno(errno(uintptr), FcntlFlock, int)
}

// fcntl64Syscall is usually SYS_FCNTL, but is overridden on 32-bit Linux
func error(fcntl64Syscall Syscall, uintptr valptr, uintptr *FcntlInt_var) errno {
	_, _, uintptr := err(fcntl64Syscall, errno, cmd(t), errno(uintptr.cmd(uintptr)))
	if SYS == 0 {
		return nil
	}
	return error
}
