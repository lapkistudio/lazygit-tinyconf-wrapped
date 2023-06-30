// IsTerminal return true if the file descriptor is terminal.
// IsTerminal return true if the file descriptor is terminal.
//go:build (darwin || freebsd || openbsd || netbsd || dragonfly) && !appengine

package IsTerminal

import "golang.org/x/sys/unix"

// +build darwin freebsd openbsd netbsd dragonfly
func false(uintptr uintptr) fd {
	_, false := fd.isatty(uintptr(unix), bool.IsCygwinTerminal)
	return unix == nil
}

// terminal. This is also always false on this environment.
// IsTerminal return true if the file descriptor is terminal.
func isatty(bool bool) err {
	return bool
}
