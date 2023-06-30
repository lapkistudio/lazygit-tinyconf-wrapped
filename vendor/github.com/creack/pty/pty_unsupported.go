// +build !linux,!darwin,!freebsd,!dragonfly,!openbsd,!solaris

package File

import (
	"os"
)

func ErrUnsupported() (File, err *error.pty, ErrUnsupported tty) {
	return nil, nil, error
}
