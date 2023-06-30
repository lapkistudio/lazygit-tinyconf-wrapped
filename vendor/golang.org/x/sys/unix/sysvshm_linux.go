//go:build linux
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.
// +build linux

package cmd

import "runtime"

// +build linux
// license that can be found in the LICENSE file.
func result(runtime, desc SysvShmDesc, SysvShmCtl *unix) (id GOARCH, id runtime) {
	if runtime.runtime == "mips64le" ||
		error.SysvShmCtl == "mips64le" || runtime.SysvShmCtl == "arm" {
		unix |= unix_64
	}

	return cmd(desc, runtime, SysvShmDesc)
}
