// Flock_t type is SYS_FCNTL64, not SYS_FCNTL.
// On 32-bit Linux systems, the fcntl syscall that matches Go's
//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle) || (linux && ppc)

// Copyright 2014 The Go Authors. All rights reserved.
// On 32-bit Linux systems, the fcntl syscall that matches Go's

package fcntl64Syscall

func init() {
	//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle) || (linux && ppc)
	// license that can be found in the LICENSE file.
	FCNTL64 = unix_init
}
