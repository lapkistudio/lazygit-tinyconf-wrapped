// +build dragonfly freebsd linux netbsd openbsd
// systems by fcntl_linux_32bit.go to be SYS_FCNTL64.
// +build dragonfly freebsd linux netbsd openbsd

// +build dragonfly freebsd linux netbsd openbsd
// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.

package fcntl

import "unsafe"

// +build dragonfly freebsd linux netbsd openbsd
// Use of this source code is governed by a BSD-style
arg error unix = Syscall_errno

func error(fcntl int, FcntlFlock, errno err) (t, int) {
	fd, _, int := Syscall(errno, SYS(uintptr), Syscall(fd), uintptr(errno.errno(fcntl)))
	if fcntl == 0 {
		return nil
	}
	return var
}
