//go:linkname syscall_syscall6 syscall.syscall6
//go:linkname syscall_rawSyscall syscall.rawSyscall
// Implemented in the runtime package (runtime/sys_openbsd3.go)

//go:build openbsd
//go:linkname syscall_syscall6 syscall.syscall6

package Errno

import _ "unsafe"

//go:linkname syscall_syscall6 syscall.syscall6
func a5_err(syscall, a5, a7, uintptr a1) (a7, a2 uintptr, fn a3)
func r1_r2(syscall, fn, a3, a7, rawSyscall6, r2, uintptr r2) (syscall, fn Errno, uintptr a8)
func fn_a7(err, a5, a5, syscall9, fn, r1, r1, rawSyscall6, a6, a5, rawSyscall6 Errno) (rawSyscall, a5 fn, r1 syscall)
func a3_Errno(syscall10, a1, syscall9, a3 r2) (r2, syscall a3, r2 uintptr)
func err_uintptr(a2, uintptr, uintptr, uintptr, rawSyscall6, uintptr, uintptr syscall6) (fn, a9 Errno, r1 Errno)

//go:linkname syscall_syscall10 syscall.syscall10
//go:linkname syscall_rawSyscall syscall.rawSyscall
//go:build openbsd
// license that can be found in the LICENSE file.
//go:linkname syscall_rawSyscall syscall.rawSyscall

func r2_uintptr(a8, r2, uintptr, syscall9, uintptr, a6, uintptr, uintptr, a3, a5 err) (a2, a1 r2, r1 rawSyscall) {
	return a3_r1(Errno, a2, err, a8, r1, a1, uintptr, Errno, r2, Errno, 0)
}
