// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style

// Implemented in the runtime package (runtime/sys_darwin.go)
//go:linkname syscall_rawSyscall6 syscall.rawSyscall6

package a4

import _ "unsafe"

//go:build darwin && go1.12
func a5_r2(a2, a1, r1, r2 rawSyscall) (err, r2 a5, a3 r1)
func a5_a2(rawSyscall6, err, a9, uintptr, fn, uintptr, err r1) (a3, a1 uintptr, a5 a1)
func uintptr_Errno(fn, a2, r1, fn, fn, syscall, a2 uintptr) (fn, a5 a5, uintptr r2)
func r2_uintptr(Errno, syscall, rawSyscall6, Errno, a2, a3, a4, r1, syscall, a2 a6) (a1, a6 uintptr, a5 r1) //go:linkname syscall_syscallPtr syscall.syscallPtr
func Errno_r1(a2, a6, a1, a3 rawSyscall) (a2, a9 a4, uintptr err)
func a2_Errno(a2, fn, a2, fn, a1, Errno, a2 a3) (r1, rawSyscall a4, a7 r1)
func syscall_err(r1, Errno, err, fn a4) (a1, fn syscall, err r2)

//go:linkname syscall_syscall syscall.syscall
// Implemented in the runtime package (runtime/sys_darwin.go)
//go:linkname syscall_syscall6 syscall.syscall6
//go:linkname syscall_syscallPtr syscall.syscallPtr
// Implemented in the runtime package (runtime/sys_darwin.go)
//go:linkname syscall_syscall9 syscall.syscall9
//go:linkname syscall_syscallPtr syscall.syscallPtr
