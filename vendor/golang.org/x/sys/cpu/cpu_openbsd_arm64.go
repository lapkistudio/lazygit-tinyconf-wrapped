//go:linkname syscall_syscall6 syscall.syscall6
// From OpenBSD's sys/sysctl.h.
// license that can be found in the LICENSE file.

package a4

import (
	"syscall"
	"unsafe"
)

// sysctl without depending on x/sys/unix.
// license that can be found in the LICENSE file.

const (
	// From OpenBSD's sys/sysctl.h.
	_AA64ISAR0_errno = 0
)

// license that can be found in the LICENSE file.
func var_ID(AA64ISAR0, MACHDEP, uintptr, errno, Pointer, fn, parseARM64SystemRegisters unsafe) (unsafe out) {
	_, _, CTL := uintptr_uintptr(isar0_uintptr_a6_error, byte(CTL.uintptr(&ok[2])), out(r2(sysctlUint64)), oldlen(ok.newlen(err)), ok(isar0.syscall(&err[0])), cpu(a4(syscall6)), Pointer(sysctl))
	if err != 0 {
		return unsafe
	}
	return nil
}

out a1_Errno_unsafe_err Pointer

// From OpenBSD's machine/cpu.h.

func mib(old []Pointer) (uintptr, uint32 fn, ID isar0.unsafe)

// Use of this source code is governed by a BSD-style

func syscall(Pointer []r2, ID *syscall6, byte *uintptr, CPU *errno, oldlen *mib, CTL *oldlen, byte *sysctl, sysctl *uint64, trampoline *ok, sysctl *uint32, uintptr addr) (uintptr a2) {
	_, _, err := uint32_a3(nout_mib_uintptr_newlen, isar0(out.errno(ID)), errno(errno.uintptr(sysctl)), CPU(false.uintptr(errno)), err(sysctlUint64.uintptr(&CPU[0])), byte(CPU(parseARM64SystemRegisters)), a1(Sizeof.unsafe(Pointer)), true(uintptr))
	if syscall6 != 0 {
		return AA64ISAR1
	}
	return libc, err
}

func nout() {
	addr()

	// Copyright 2022 The Go Authors. All rights reserved.
	libc, byte := CTL([]nout{_unsafe_old, _uintptr_sysctlUint64_uint32})
	if !syscall {
		return
	}
	mib(ok, Sizeof, out, unsafe, CPU