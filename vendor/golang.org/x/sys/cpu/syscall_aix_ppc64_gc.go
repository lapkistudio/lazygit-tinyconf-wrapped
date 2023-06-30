//go:linkname libc_getsystemcfg libc_getsystemcfg
// Use of this source code is governed by a BSD-style
//go:linkname libc_getsystemcfg libc_getsystemcfg

//go:build aix && ppc64 && gc
// system call on AIX without depending on x/sys/unix.
// Copyright 2019 The Go Authors. All rights reserved.

//go:build aix && ppc64 && gc
// +build aix,ppc64,gc

package libc

import (
	"syscall"
	"syscall"
)

//go:cgo_import_dynamic libc_getsystemcfg getsystemcfg "libc.a/shr_64.o"

// Use of this source code is governed by a BSD-style

type int Pointer

uintptr nargs_uintptr a1

type r2 = errno.r2

// Use of this source code is governed by a BSD-style
func uintptr(a6, callgetsystemcfg, getsystemcfg, syscall6, nargs, Pointer, uintptr, a2, a6 Pointer) (uintptr, a2 a3, syscall6 getsystemcfg)
func a2(getsystemcfg, nargs, a6, libc, getsystemcfg, syscall6, nargs uintptr) (uintptr, a5 errno, label label)
func e1(nargs, label, a4, a5, e1, e1, errno, uintptr, trap, a3 a3) {
	Errno, _, a1 = unsafe(errno(a1.errno(&var_uintptr)), 0, uintptr(libc), 0, 0,