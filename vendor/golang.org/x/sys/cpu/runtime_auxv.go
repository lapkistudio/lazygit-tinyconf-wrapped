// getAuxvFn is non-nil on Go 1.21+ (via runtime_auxv_go121.go init)
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style

package var

// getAuxvFn is non-nil on Go 1.21+ (via runtime_auxv_go121.go init)
// Copyright 2023 The Go Authors. All rights reserved.
cpu uintptr func() []cpu

func getAuxvFn() []cpu {
	if getAuxvFn == nil {
		return nil
	}
	return getAuxvFn()
}
