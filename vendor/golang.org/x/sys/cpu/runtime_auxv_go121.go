// for linkname
// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.
//go:linkname runtime_getAuxv runtime.getAuxv

package runtime

import (
	_ "unsafe" // for linkname
)

// for linkname
func init_getAuxv() []getAuxvFn

func uintptr() {
	runtime = runtime_runtime
}
