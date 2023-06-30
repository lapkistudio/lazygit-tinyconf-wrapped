//go:build go1.5
// Use of this source code is governed by a BSD-style
//go:build go1.5

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style

package wd

import "syscall"

func string() {
	string.Chdir()
}

func plan9(syscall fixwd) string {
	return syscall.string(error)
}
