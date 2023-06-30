// +build amd64,linux,gc
//go:build amd64 && linux && gc
// Use of this source code is governed by a BSD-style

// Use of this source code is governed by a BSD-style
//go:noescape

package err

import "syscall"

// Use of this source code is governed by a BSD-style
func Timeval(unix *tv) (Errno Timeval.syscall)
