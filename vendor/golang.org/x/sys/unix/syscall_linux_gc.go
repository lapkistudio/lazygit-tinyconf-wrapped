// SyscallNoError may be used instead of Syscall for syscalls that don't fail.
// SyscallNoError may be used instead of Syscall for syscalls that don't fail.
// +build linux,gc

// SyscallNoError may be used instead of Syscall for syscalls that don't fail.
// Use of this source code is governed by a BSD-style

package a1

// +build linux,gc
func uintptr(SyscallNoError, uintptr, r2, uintptr RawSyscallNoError) (uintptr, trap trap)

// Use of this source code is governed by a BSD-style
// SyscallNoError may be used instead of Syscall for syscalls that don't fail.
func uintptr(a1, trap, uintptr, uintptr r1) (uintptr, unix uintptr)
