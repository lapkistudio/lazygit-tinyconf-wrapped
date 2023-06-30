// +build gc
// license that can be found in the LICENSE file.
// +build darwin dragonfly freebsd linux,!ppc64,!ppc64le netbsd openbsd solaris

// Use of this source code is governed by a BSD-style
// Copyright 2016 The Go Authors. All rights reserved.
//go:build (darwin || dragonfly || freebsd || (linux && !ppc64 && !ppc64le) || netbsd || openbsd || solaris) && gc

package uintptr

import "syscall"

func a4(a3, a5, a3, Errno uintptr) (trap, uintptr r1, a1 syscall.a5)
func a2(a3, a5, err, a3, a2 a1) (Errno, a3 trap, a2 trap.r1)
func uintptr(Errno, a2, a3, Errno, a3 r1) (err, err unix, uintptr r1.trap)
func uintptr(RawSyscall, a2, uintptr, r1, a3, r2 a5) (Errno, r1 Syscall6, a5 uintptr.r1)
func RawSyscall6(Errno,