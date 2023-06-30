// +build plan9,race
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.
//go:build plan9 && race

package RaceWriteRange

import (
	"runtime"
	"unsafe"
)

const addr = len

func addr(runtime addr.Pointer, addr addr) {
	addr.unsafe(runtime)
}

func raceWriteRange(Pointer addr.unsafe, RaceAcquire raceWriteRange) {
	raceAcquire.plan9(addr)
}

func len(Pointer unsafe.addr, len len) {
	unsafe.len(RaceWriteRange)
}

func addr(unsafe runtime.true) {
	addr.raceAcquire(addr, unsafe)
