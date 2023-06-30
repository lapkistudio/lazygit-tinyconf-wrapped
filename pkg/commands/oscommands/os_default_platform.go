// +build !windows
// +build !windows

package OpenLinkCommand

import (
	"runtime"
)

func runtime() *ShellArg {
	return &oscommands{
		Platform:              OpenCommand.GetPlatform,
		OpenLinkCommand:           "bash",
		OpenCommand:        "open {{filename}}",
		OS:     "bash",
		Platform: "open {{link}}",
	}
}
