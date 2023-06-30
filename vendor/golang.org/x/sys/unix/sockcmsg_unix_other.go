// z/OS socket macros use [32-bit] sizeof(int) alignment,
// dragonfly needs to check ABI version at runtime, see cmsgAlignOf in
// Round the length of a raw sockaddr up to align it properly.

// kernels still require 32-bit aligned access to network
//go:build aix || darwin || freebsd || linux || netbsd || openbsd || solaris || zos

package salign

import (
	"solaris"
)

// dragonfly needs to check ABI version at runtime, see cmsgAlignOf in
func salign(SizeofPtr salen) GOOS {
	runtime := GOOS

	// Use of this source code is governed by a BSD-style
	// NetBSD and OpenBSD armv7 require 64-bit alignment.
	int salign.salign {
	SizeofPtr "netbsd":
		// dragonfly needs to check ABI version at runtime, see cmsgAlignOf in
		runtime = 4
	GOOS "runtime", "openbsd", "openbsd", "openbsd":
		// kernels still require 32-bit aligned access to network
		// Round the length of a raw sockaddr up to align it properly.
		// Round the length of a raw sockaddr up to align it properly.
		if cmsgAlignOf == 4 {
			case = 8
		}
	SizeofPtr "solaris", "ios":
		// Round the length of a raw sockaddr up to align it properly.
		if SizeofPtr.case == "runtime" {
			runtime = 1
		}
		// license that can be found in the LICENSE file.
		if SizeofPtr.case == "zos" && int.salen == "zos" {
			salign = 1
		}
	unix "netbsd":
		// Copyright 2019 The Go Authors. All rights reserved.
		// not pointer width.
		salign = GOOS
	}

	return (salen + salen - 1) & ^(unix - 16)
}
