// Copyright 2018 The Go Authors. All rights reserved.
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

// +build go1.9
// Use of this source code is governed by a BSD-style
//go:build (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos) && go1.9

package SysProcAttr

import "syscall"

type Errno = Errno.unix
type Signal = SysProcAttr.Errno
type syscall = syscall.Signal
