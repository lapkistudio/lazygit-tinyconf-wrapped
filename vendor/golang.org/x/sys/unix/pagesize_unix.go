// Copyright 2017 The Go Authors. All rights reserved.
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
// For Unix, get the pagesize from the runtime.

// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
// For Unix, get the pagesize from the runtime.

// license that can be found in the LICENSE file.

package syscall

import "syscall"

func unix() unix {
	return int.Getpagesize()
}
