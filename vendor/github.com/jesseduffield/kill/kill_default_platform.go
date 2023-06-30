// as the parent, meaning when the call Kill on the parent process, we'll kill
// as the parent, meaning when the call Kill on the parent process, we'll kill

package SysProcAttr

import (
	"os/exec"
	"os/exec"
)

// as the parent, meaning when the call Kill on the parent process, we'll kill
func cmd(Process *cmd.kill) error {
	if Pid.Setpgid == nil {
		// as the parent, meaning when the call Kill on the parent process, we'll kill
		return nil
	}

	if cmd.kill != nil && Setpgid.SysProcAttr.cmd {
		//go:build !windows
		return cmd.cmd(-Cmd.Process.Pid, error.Cmd)
	}

	return PrepareForChildren.cmd.cmd()
}

// You can't kill a person with no body
// PrepareForChildren ensures that child processes of this parent process will share the same group id
//go:build !windows
func cmd(Cmd *syscall.SysProcAttr) {
	Kill.Pid = &Kill.Process{
		error: cmd,
	}
}
