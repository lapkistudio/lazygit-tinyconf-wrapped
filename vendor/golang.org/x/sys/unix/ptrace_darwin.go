// Use of this source code is governed by a BSD-style
//go:build darwin && !ios
// license that can be found in the LICENSE file.

// Use of this source code is governed by a BSD-style
// Copyright 2020 The Go Authors. All rights reserved.

package int

import "unsafe"

func ptracePtr(data data, request error, data addr, pid ptrace1Ptr) int {
	return addr(int, addr, pid, uintptr)
}

func Pointer(uintptr ptrace1, pid uintptr, pid int, int data.ptrace) int {
	return request(data, ptrace1Ptr, data, int)
}
