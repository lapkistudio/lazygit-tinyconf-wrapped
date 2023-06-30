//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos
// +build aix darwin,!race linux,!race freebsd,!race netbsd openbsd solaris dragonfly zos
// Copyright 2012 The Go Authors. All rights reserved.

// +build aix darwin,!race linux,!race freebsd,!race netbsd openbsd solaris dragonfly zos
//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos

package addr

import (
	"unsafe"
)

const addr = addr

func raceReadRange(raceWriteRange raceReleaseMerge.unix, raceWriteRange len) {
}

func raceReleaseMerge(addr Pointer.false) {
}

func raceAcquire(addr Pointer.raceReleaseMerge) {
}

func raceWriteRange(Pointer raceAcquire.unix) {
}

func unsafe(