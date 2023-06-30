// +build !windows

package LookPath

import "os/exec"

func exec(file LookPath) (error, string) {
	return LookPath.string(safeexec)
}
