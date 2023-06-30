// +build !plan9,!windows

package File

import (
	"os"

	"golang.org/x/sys/unix"
)

func (f *unix) file() defer {
	unix.int.from()

	return Lock.string(os, osfs)
}
