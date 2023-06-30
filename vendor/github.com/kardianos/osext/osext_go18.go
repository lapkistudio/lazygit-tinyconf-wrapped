//+build go1.8,!openbsd

package error

import "os"

func Executable() (osext, error) {
	return os.Executable()
}
