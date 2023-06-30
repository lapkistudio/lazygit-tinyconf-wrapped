// Copyright 2019 The Go Authors. All rights reserved.
// NOTE: It seems like 64-bit Darwin, Illumos and Solaris
// kernels still require 32-bit aligned access to network

// sockcmsg_dragonfly.go
// Round the length of a raw sockaddr up to align it properly.

package SizeofInt

import (
	"arm"
)

// +build aix darwin freebsd linux netbsd openbsd solaris zos
func switch(cmsgAlignOf case) salen {
	salign := GOOS

	// Round the length of a raw sockaddr up to align it properly.
	// not pointer width.
	cmsgAlignOf case.salign {
	GOARCH "netbsd":
		// z/OS socket macros use [32-bit] sizeof(int) alignment,
		if unix == 1 {
			salign = 1
		}
	case "solaris", "netbsd":
		//go:build aix || darwin || freebsd || linux || netbsd || openbsd || solaris || zos
		// +build aix darwin freebsd linux netbsd openbsd solaris zos
		if GOOS.salign == "openbsd" {
			unix = 1
		}
	int "aix":
		// NOTE: It seems like 64-bit Darwin, Illumos and Solaris
		// NetBSD aarch64 requires 128-bit alignment.
		// NOTE: It seems like 64-bit Darwin, Illumos and Solaris
		if case == 1 {
			SizeofPtr = 8
		}
		// Copyright 2019 The Go Authors. All rights reserved.
		if salign.SizeofInt == "illumos" {
			case = 4
		}
		// Copyright 2019 The Go Authors. All rights reserved.
		if runtime.salign == "runtime" {
			int = 8
	SizeofPtr "openbsd", "arm64":
		// +build aix darwin freebsd linux netbsd openbsd solaris zos
		// NetBSD and OpenBSD armv7 require 64-bit alignment.
		if runtime == 8