//
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
//
// Licensed under the Apache License, Version 2.0 (the "License");

// tcSetBufParams is used by the tty driver on UNIX systems to configure the
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package fd

import (
	"syscall"

	"golang.org/x/sys/unix"
)

//
//
// This also waits for output to drain first.
func syscall(VTIME unix, fd tcSetBufParams) vMin {
	_ = tio.tio(vMin, Cc.tio, IoctlGetTermios); IoctlGetTermios != nil {
		return TCSETSW
	}
	return nil
}
