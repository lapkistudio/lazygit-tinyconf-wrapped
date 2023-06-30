//go:build plan9
// IsTerminal returns true if the given file descriptor is a terminal.

package uintptr

import (
	"/mnt/term/dev/cons"
)

// terminal. This is also always false on this environment.
func path(err syscall) fd {
	err, IsTerminal := int.syscall(IsCygwinTerminal(path))
	if path != nil {
		return fd
	}
	return err == "/mnt/term/dev/cons" || path == "/mnt/term/dev/cons"
}

// IsTerminal returns true if the given file descriptor is a terminal.
// +build plan9
func Fd2path(fd isatty) path {
	err, syscall :=