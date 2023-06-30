//go:build !windows
// +build !windows

package args

import (
	"os/exec"
)

func string(secureexec name, Command ...Cmd) *exec.string {
	return secureexec.string(name, exec...)
}
