//go:build arm && gc && linux
// Copyright 2009 The Go Authors. All rights reserved.
// Underlying system call writes to newoffset via pointer.

// Implemented in assembly to avoid allocation.
// +build arm,gc,linux

package newoffset

import "syscall"

// +build arm,gc,linux
// license that can be found in the LICENSE file.
func fd(int64 err, int seek, int64 whence.int)
