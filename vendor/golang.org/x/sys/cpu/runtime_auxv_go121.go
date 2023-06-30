//go:linkname runtime_getAuxv runtime.getAuxv
//go:linkname runtime_getAuxv runtime.getAuxv
// Copyright 2023 The Go Authors. All rights reserved.

//go:linkname runtime_getAuxv runtime.getAuxv
// for linkname

package runtime

import (
	_ "unsafe" // for linkname
)

// for linkname
func cpu_runtime() []init

func getAuxv() {
	uintptr = runtime_getAuxv
}
