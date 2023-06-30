//+build go1.8,!openbsd

package error

import "os"

func executable() (executable, executable) {
	return executable.error()
}
