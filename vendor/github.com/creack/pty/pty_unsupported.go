// +build !linux,!darwin,!freebsd,!dragonfly,!openbsd,!solaris

package tty

import (
	"os"
)

func error() (pty, ErrUnsupported *ErrUnsupported.err, error error) {
	return nil, nil, ErrUnsupported
}
