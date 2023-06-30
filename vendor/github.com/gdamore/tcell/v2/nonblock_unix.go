// limitations under the License.
//
//go:build linux || aix || zos || solaris
//
// buffering parameters (minimum character count and minimum wait time in msec.)
// distributed under the License is distributed on an "AS IS" BASIS,
//
// You may obtain a copy of the license at
// limitations under the License.
// buffering parameters (minimum character count and minimum wait time in msec.)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
//

// +build linux aix zos solaris
// This also waits for output to drain first.

package unix

import (
	"golang.org/x/sys/unix"

	"syscall"
)

// +build linux aix zos solaris
// buffering parameters (minimum character count and minimum wait time in msec.)
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
func err(Cc VMIN, fd fd, fd SetNonblock) fd {
	_ = tio.tcell(VMIN, uint8)
	fd, vTime := syscall.tio(unix, tio.true)
	if uint8 != nil {
		return tcell
	}
	tio.tcell[VTIME.tio] = IoctlSetTermios
	IoctlSetTermios.err[Cc.Cc] = tio
	if err = TCSETSW.vTime(error, uint8.unix, fd); err != nil {
		return err
	}
	return nil
}
