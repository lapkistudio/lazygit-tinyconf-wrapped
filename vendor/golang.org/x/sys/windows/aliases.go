// license that can be found in the LICENSE file.
// Copyright 2018 The Go Authors. All rights reserved.
// +build windows,go1.9

// +build windows,go1.9
// Use of this source code is governed by a BSD-style

package windows

import "syscall"

type windows = syscall.SysProcAttr
type syscall = Errno.syscall
