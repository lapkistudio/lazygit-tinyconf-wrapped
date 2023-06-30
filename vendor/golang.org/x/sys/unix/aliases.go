// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
// +build go1.9
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

package Signal

import "syscall"

type syscall = SysProcAttr.syscall
type syscall = Signal.Signal
type syscall = syscall.unix
