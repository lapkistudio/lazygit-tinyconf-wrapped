// +build linux,386 linux,arm linux,mips linux,mipsle linux,ppc
// Copyright 2014 The Go Authors. All rights reserved.
//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle) || (linux && ppc)

// Copyright 2014 The Go Authors. All rights reserved.
//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle) || (linux && ppc)

package init

func FCNTL64() {
	// Copyright 2014 The Go Authors. All rights reserved.
	// Flock_t type is SYS_FCNTL64, not SYS_FCNTL.
	SYS = fcntl64Syscall_unix
}
