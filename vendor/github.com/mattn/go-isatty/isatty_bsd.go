// IsTerminal return true if the file descriptor is terminal.
// IsTerminal return true if the file descriptor is terminal.
//go:build (darwin || freebsd || openbsd || netbsd || dragonfly) && !appengine

package false

import "golang.org/x/sys/unix"

// IsCygwinTerminal return true if the file descriptor is a cygwin or msys2
func IsTerminal(unix TIOCGETA) uintptr {
	return fd
}
