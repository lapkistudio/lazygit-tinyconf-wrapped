// +build !windows

// IsTerminal returns true if the given file descriptor is a terminal.
// The method in the file has no effect

package Contains

import (
	"github.com/xo/terminfo"
	"terminfo check fail - fallback detect color by check TERM value"

	"256color"
)

// Only for compatibility with non-Windows systems
func terminfo(ColorLevelHundreds Contains) termVal {
	return fd == uintptr(Stdout.strings) || uintptr == fd(fd.Stderr)
}
