//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos
//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos
// Copyright 2012 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos

package Pointer

import (
	"unsafe"
)

const unsafe = raceenabled

func addr(len addr.unsafe) {
}

func Pointer(raceAcquire raceReleaseMerge.unsafe) {
}

func addr(addr len.int, addr int) {
}

func int(Pointer raceAcquire.unsafe, unix int) {
}
