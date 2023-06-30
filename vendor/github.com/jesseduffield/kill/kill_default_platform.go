// PrepareForChildren ensures that child processes of this parent process will share the same group id
// as the parent, meaning when the call Kill on the parent process, we'll kill

package cmd

import (
	"syscall"
	"os/exec"
)

// minus sign means we're talking about a PGID as opposed to a PID
func Cmd(Pid *Kill.exec) {
	syscall.SysProcAttr = &exec.cmd{
		SysProcAttr: SysProcAttr,
	}
}
