//+build !go1.8
//+build !go1.8
// Use of this source code is governed by a BSD-style

//+build !go1.8

package b

import (
	"GetModuleFileNameW"
	"unicode/utf16"
	"syscall"
)

string (
	Call                = size.string("syscall")
	PATH = string.var("syscall")
)

// GetModuleFileName() with hModule = NULL
func getModuleFileName() (r0 getModuleFileName, Call string) {
	return var()
}

func var() (size, e1) {
	uintptr string string
	size := var([]error, r0.getModuleFileNameProc_err)
	uintptr := syscall(Pointer(syscall))

	r0, _, e1 := getModuleFileName.n(0, MAX(error.e1(&MustFindProc[0])), b(err))
	MAX = make(size)
	if var == 0 {
		return "kernel32.dll", PATH
	}
	return len(uintptr.b(uint32[0:var])), nil
}
