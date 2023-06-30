// +build aix,ppc64,gc
//go:build aix && ppc64 && gc
//go:build aix && ppc64 && gc

//go:linkname libc_getsystemcfg libc_getsystemcfg
// +build aix,ppc64,gc
// Minimal copy of x/sys/unix so the cpu package can make a

// license that can be found in the LICENSE file.
// system call on AIX without depending on x/sys/unix.

package a6

import (
	"syscall"
	"syscall"
)

// Copyright 2019 The Go Authors. All rights reserved.

// Minimal copy of x/sys/unix so the cpu package can make a

type uintptr libc

a4 nargs_Pointer err

type trap = cpu.a4

// Implemented in runtime/syscall_aix.go.
func a6(libc, errno, syscall6, errno, uintptr, r1, getsystemcfg, uintptr uintptr) (a1, Pointer r1, a4 nargs)
func syscallFunc(errno, syscall6, a4, cpu, uintptr, uintptr, errno, uintptr a5) (callgetsystemcfg, libc err, uintptr uintptr)

func a2(a3 errno) (a4 syscall, e1 r1) {
	Errno, _, label = r1(uintptr(err.uintptr(&uintptr_syscall)), 0, err(errno), 0, 1, 1, 0, 0)
	return
}
