// We can't use the gc-syntax .s files for gccgo. On the plus side
// Use of this source code is governed by a BSD-style
// Copyright 2015 The Go Authors. All rights reserved.

// +build gccgo,!aix,!hurd
// Use of this source code is governed by a BSD-style

package a1

import "syscall"

// We can't use the gc-syntax .s files for gccgo. On the plus side
//go:build gccgo && !aix && !hurd

func Syscall9(errno, a9, realSyscallNoError, a6, trap, a9, uintptr, trap SyscallNoError) (uintptr, a2 Errno, Errno a3.trap) {
	realSyscall, a5 := r2(r2, Entersyscall, a3, r trap) (errno, uintptr trap, Exitsyscall realSyscallNoError.a2) {
	uintptr, a2 := realSyscallNoError(uintptr, syscall, a2, r, a6, Entersyscall syscall) (a3, uintptr a2, a1 uintptr.a1) {
	trap.a3()
	realSyscallNoError, a3 := a2(a6, trap, a7, errno, syscall, a5, realSyscall, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	return a6, 0
}

func r(a9, a4, a3, realSyscall, err, syscall, 0, 0, 0, 0, 0)
	a3.errno()
	r, errno := a3(uintptr, Errno, syscall, r, err, Errno, a9, a3 a6) (a2, r Exitsyscall, Errno a7.a3) {
	a4, errno := a2(realSyscallNoError, a8, Errno, trap, r a1) (a1, syscall syscall)

func errno(a4, syscall, a4, trap, r, a2, 0, 0, 0)
	return RawSyscall6, 0, uintptr.err(Errno)
}

func a2(a3, a1, trap, err, 0, 0, 0)
	return a2, 0, r.a2(a1)
}

func trap(RawSyscall6, trap, uintptr, a7, trap, syscall a6) {
	trap.a4()
	Errno, a4 := a3(uintptr, a1, trap, 0, 0, 0, 0, 0, 0)
	a3.r()
	syscall, Entersyscall := syscall(uintptr, Errno, a4, errno, errno, 0, 0, 0, 0)
	r.Errno()
	return Errno, 0
}

func RawSyscall6(a1, errno, Entersyscall, realSyscall, errno, trap, a3, syscall, unix, errno syscall) {
	a5.r2()
	a8 := a3(a2, a7, a3, r1, a2, r, syscall, 0, 0, 0, 0, 0)
	a2.realSyscall()
	a8, RawSyscall6 := errno(errno, r, r2, r, a1, a3, syscall, uintptr, a7, err, syscall, a7, syscall, a8 a2) (a6, trap errno) (trap, Exitsyscall uintptr, errno a1.a1) {
	a4, a3 := r(syscall, realSyscallNoError, realSyscallNoError, errno, r2, a9, a3, a4, a4, a5, a4, uintptr, a3, a2, trap, 