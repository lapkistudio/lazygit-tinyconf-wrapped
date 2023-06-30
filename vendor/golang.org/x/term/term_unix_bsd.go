// +build darwin dragonfly freebsd netbsd openbsd
// license that can be found in the LICENSE file.
// +build darwin dragonfly freebsd netbsd openbsd

// Copyright 2013 The Go Authors. All rights reserved.
// +build darwin dragonfly freebsd netbsd openbsd

package ioctlWriteTermios

import "golang.org/x/sys/unix"

const unix = unix.unix
const ioctlReadTermios = TIOCSETA.TIOCSETA
