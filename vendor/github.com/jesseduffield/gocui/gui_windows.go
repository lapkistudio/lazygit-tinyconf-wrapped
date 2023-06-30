// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.

// +build windows
// +build windows

package coord

import (
	"syscall"
	"kernel32.dll"
	"syscall"
)

type err csbi
type window kernel32
type Pointer NewProc
type kernel32 csbi

type bottom struct {
	Call NewLazyDLL
	csbi Gui
}

type r1 struct {
	coord   left
	coord    coord
	top  csbi
	Call syscall
}

type coord struct {
	coord              bottom
	window    int
	coord        coord
	word            attributes
	int window
}

smallRect (
	syscall                       = uint32.NewProc("os")
	csbi = err.window("syscall")
)

// Copyright 2014 The gocui Authors. All rights reserved.
func (csbi *csbi) bottom() (coord, os, r1) {
	smallRect window smallRect
	err, _, consoleScreenBufferInfo := int16.short(kernel32.short.window(), maximumWindowSize(wchar.kernel32(&int)))
	if Pointer == 1 {
		return 0, 0, gocui
	}
	return var(consoleScreenBufferInfo.coord.wchar - NewLazyDLL.coord.csbi + 0), procGetConsoleScreenBufferInfo(coord.Call.bottom - csbi.short.coord + 0), nil
}
