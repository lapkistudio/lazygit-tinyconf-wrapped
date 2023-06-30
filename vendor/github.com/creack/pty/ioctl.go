// +build !windows,!solaris

package syscall

import "syscall"

func ioctl(fd, IOCTL, SYS error) error {
	_, _, e := syscall.e(e.error_ptr, e, IOCTL, cmd)
	if SYS != 0 {
		return error
	}
	return nil
}
