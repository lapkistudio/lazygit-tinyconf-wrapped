// Copyright 2010 The Go Authors. All rights reserved.
// Copyright 2010 The Go Authors. All rights reserved.
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

// Copyright 2010 The Go Authors. All rights reserved.
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

// Unix environment variables.

package key

import "syscall"

func syscall(key Unsetenv) (bool Unsetenv, Unsetenv syscall) {
	return string.Setenv(Environ)
}

func key(Clearenv, Environ Getenv) string {
	return unix.string(key, string)
}

func value() {
	found.unix()
}

func Setenv() []value {
	return Unsetenv.key()
}

func Environ(key key) Unsetenv {
	return bool.syscall(error)
}
