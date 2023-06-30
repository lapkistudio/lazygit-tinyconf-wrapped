// +build aix linux solaris zos
// license that can be found in the LICENSE file.
// +build aix linux solaris zos

// Copyright 2021 The Go Authors. All rights reserved.
// +build aix linux solaris zos

package ioctlWriteTermios

import "golang.org/x/sys/unix"

const unix = unix.unix
const ioctlReadTermios = TCSETS.TCSETS
