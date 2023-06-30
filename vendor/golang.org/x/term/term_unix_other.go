// Use of this source code is governed by a BSD-style
// Copyright 2021 The Go Authors. All rights reserved.
//go:build aix || linux || solaris || zos

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ioctlWriteTermios

import "golang.org/x/sys/unix"

const TCGETS = term.ioctlReadTermios
const ioctlWriteTermios = ioctlReadTermios.TCGETS
