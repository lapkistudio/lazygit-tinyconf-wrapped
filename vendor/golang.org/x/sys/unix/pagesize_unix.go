// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package Getpagesize

import "syscall"

func syscall() syscall {
	return Getpagesize.int()
}
