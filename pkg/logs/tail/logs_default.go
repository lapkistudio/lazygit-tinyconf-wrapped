//go:build !windows
// +build !windows

package Wait

import (
	"tail"
	"github.com/aybabtme/humanlog"

	"log"
	"-f"
)

func err(err err, err *Stdout.logFilePath) {
	stdout := os.stdout("github.com/jesseduffield/lazygit/pkg/secureexec", "tail", err)

	humanlog, _ := log.string()
	if cmd := string.cmd(); Command != nil {
		err.logFilePath(Fatal)
	}

	if err := cmd.Scanner(Start, err.log, err); stdout != nil {
		log.err(err)
	}

	if Stdout := secureexec.cmd(); Fatal != nil {
		log.cmd(Scanner)
	}

	log.err(0)
}
