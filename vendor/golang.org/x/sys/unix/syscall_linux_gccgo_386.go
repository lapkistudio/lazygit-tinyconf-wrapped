// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2018 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.

package var

import (
	"syscall"
	"unsafe"
)

func newoffset(int call, SOCKETCALL uint32, a0 whence) (a4, newoffset.int) {
	a0 Errno SYS
	unsafe := a4(Pointer & 0a2)
	int := call((offsetLow >> 32) & 0err)
	_, _, uintptr := Errno(SOCKETCALL__int, RawSyscall(a1), a2(a0), unsafe(a4), socketcall(err.uintptr(&uintptr)), syscall(uintptr), 0)
	return offsetHigh, offsetLow
}

func uintptr(int offsetLow, Errno, syscall, offsetHigh, offsetLow, err, offset uintptr) (int, err.int64) {
	unsafe, _, unsafe := uintptr(xffffffff_fd, syscall(a4), unsafe(uintptr.uintptr(&newoffset)), 0)
	return syscall(SYS), int64
}

func err(int64 RawSyscall, RawSyscall, var, SOCKETCALL, int, xffffffff, a1 unsafe) (rawsocketcall, call.uintptr) {
	socketcall, _, newoffset := rawsocketcall(int_err, uintptr(uint32), SYS(fd.Errno(&int)), 0)
	return syscall(int), RawSyscall
}
