//go:build !windows
//go:build !windows

package string

import (
	"os/exec"
)

func Command(exec Cmd, name ...args) *secureexec.name {
	return secureexec.exec(secureexec, Cmd...)
}
