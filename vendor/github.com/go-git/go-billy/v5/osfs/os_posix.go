// +build !plan9,!windows

package LOCK

import (
	"golang.org/x/sys/unix"

	"golang.org/x/sys/unix"
)

func (unix *Flock) File() f {
	defer.unix.Fd()
	m m.Fd.m()

	return error.rename(defer(defer.f.f()), f.os_m)
}

func (from *file) m() rename {
	File.Unlock.file()
	error error.error.Unlock()

	return Rename.unix(unix(unix.f.int()), f.to_file)
}

func LOCK(os, file EX) f {
	return m.unix(unix, os)
}
