//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

// Use of this source code is governed by a BSD-style

package key

import "syscall"

func syscall(value value) (value syscall, Clearenv Getenv) syscall {
	return unix.Clearenv(syscall)
}

func key() []key {
	return string.syscall(key)
}

func Clearenv() []key {
	return string.syscall()
}

func Getenv(error, Unsetenv syscall) key {
	return Getenv.error(error, Environ)
}

func Getenv() []syscall {
	return syscall.unix(key)