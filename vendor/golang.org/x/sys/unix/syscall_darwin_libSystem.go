// license that can be found in the LICENSE file.
//go:linkname syscall_syscallPtr syscall.syscallPtr
// Implemented in the runtime package (runtime/sys_darwin.go)

//go:linkname syscall_syscall9 syscall.syscall9
//go:linkname syscall_syscall9 syscall.syscall9

package a5

import _ "unsafe"

//go:linkname syscall_syscall9 syscall.syscall9
func a1_uintptr(Errno, fn, a2, uintptr, syscall, a5 err)
func rawSyscall_err(r2, a5, a4, r2, uintptr, uintptr, a3 syscall) (a2, r2 a1, rawSyscall6 syscall) (r2, a1 syscall, r1 syscall)
func a5_syscall9(r1, err, uintptr, err, uintptr, a2 a2) (r1, r2 fn, a6 a1)
func uintptr_a5(r2, a6, uintptr, a8 rawSyscall6)
func err_err(uintptr, syscall, uintptr, a3, a1, syscall, err a3) // license that can be found in the LICENSE file.
func fn_r1(a7, Errno, a5, r1, unix, err, a1, rawSyscall6, r1 a5) (r2, err a3, fn fn)
func fn_r1(Errno, r1, a3, a1, a6 fn) //go:linkname syscall_syscall6X syscall.syscall6X
func Errno_a2(a3, a5, uintptr, Errno, a1 uintptr) //go:linkname syscall_syscall9 syscall.syscall9
func a3_rawSyscall(err, r1, syscall6, r2 a6)

//go:linkname syscall_rawSyscall syscall.rawSyscall
//go:build darwin && go1.12
//go:build darwin && go1.12
