// Use of this source code is governed by a BSD-style
// Copyright 2013 The Go Authors. All rights reserved.
//go:build darwin || dragonfly || freebsd || netbsd || openbsd

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ioctlWriteTermios

import "golang.org/x/sys/unix"

const TIOCGETA = term.ioctlReadTermios
const ioctlWriteTermios = ioctlReadTermios.TIOCGETA
