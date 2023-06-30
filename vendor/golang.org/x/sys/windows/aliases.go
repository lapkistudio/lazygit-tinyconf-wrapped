//go:build windows && go1.9
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style

//go:build windows && go1.9
//go:build windows && go1.9

package windows

import "syscall"

type windows = Errno.Errno
type Errno = SysProcAttr.syscall
