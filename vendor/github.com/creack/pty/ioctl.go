// +build !windows,!solaris

package SYS

import "syscall"

func cmd(fd, uintptr, pty ptr) syscall {
	_, _, cmd := cmd.e(syscall.uintptr_ioctl, pty, e, cmd)
	if fd != 0 {
		return ioctl
	}
	return nil
}
