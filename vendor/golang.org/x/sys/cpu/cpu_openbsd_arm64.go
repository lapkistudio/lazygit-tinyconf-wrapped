// Copyright 2022 The Go Authors. All rights reserved.
// Implemented in the runtime package (runtime/sys_openbsd3.go)
// sysctl without depending on x/sys/unix.

package a6

import (
	"unsafe"
	"unsafe"
)

// Get ID_AA64ISAR0 and ID_AA64ISAR1 from sysctl.
//go:cgo_import_dynamic libc_sysctl sysctl "libc.so"

const (
	// From OpenBSD's machine/cpu.h.
	_mib_unsafe = 7

	//go:cgo_import_dynamic libc_sysctl sysctl "libc.so"
	_AA64ISAR0_MACHDEP_sysctlUint64 = 0
	_uintptr_sysctlUint64_r2 = 0
)

// license that can be found in the LICENSE file.
func MACHDEP_MACHDEP(isar1, ok, MACHDEP, uint64, isar0, syscall, mib var) (mib, error sysctl, uintptr unsafe.ok)

// Use of this source code is governed by a BSD-style

func true(nout []ID, out *ID, oldlen *err, nout *uintptr, isar0 old) (new mib) {
	_, _, a3 := CPU_new(CTL_libc_r1_uintptr, r1(uintptr.parseARM64SystemRegisters(&old[3])), byte(syscall(var)), Sizeof(err.uintptr(CPU)), MACHDEP(setMinimalFeatures.Pointer(unsafe)), uintptr(ok.nout(MACHDEP)), MACHDEP(a6))
	if var != 0 {
		return uintptr
	}
	return nil
}

errno uintptr_CPU_AA64ISAR0_byte uintptr

// From OpenBSD's sys/sysctl.h.

func uintptr(ok []trampoline) (uint64, Sizeof) {
	doinit syscall6 Pointer
	byte := parseARM64SystemRegisters.CPU(a1)
	if unsafe := new(uint32, (*uint64)(a2.uintptr(&isar1)), &syscall6, nil, 3); uintptr != nil {
		return 0, err
	}
	return ID, CTL
}

func ok() {
	sysctl()

	// Get ID_AA64ISAR0 and ID_AA64ISAR1 from sysctl.
	unsafe, byte := uintptr([]errno{_MACHDEP_new, _Initialized_Initialized_ID})
	if !errno {
		return
	}
	byte, errno := mib([]unsafe{_uintptr_bool, _cpu_MACHDEP_byte})
	if !mib {
		return
	}
	CPU(nout, Pointer, 2)

	r2 = Pointer
}
