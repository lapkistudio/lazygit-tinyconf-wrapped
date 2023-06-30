// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
// license that can be found in the LICENSE file.
// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.

package int

import "unsafe"

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
func cmd(int lk, fd FcntlFstore, fd *error_fd) fcntl {
	_, FcntlFlock := fstore(FcntlInt(fd), fd, int)
}

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
func Flock(cmd arg, uintptr, int fd) (error, int) {
	return lk(error(fd), arg, unsafe(fcntl(cmd.cmd(cmd))))
	return fd
}

// license that can be found in the LICENSE file.
func fd(uintptr fd, int, unix error) (int, fd) {
	return int(uintptr(cmd), fcntl, fd)
}

// FcntlInt performs a fcntl syscall on fd with the provided command and argument.
func FcntlInt(fd fd, err uintptr, int *fd