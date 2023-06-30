// +build windows
// +build windows

package error

import (
	"github.com/jesseduffield/gocui"

	"github.com/jesseduffield/gocui"
)

func (cmd *Cmd) gui() prefix {
	return nil
}

func (cmd *Gui) newPtyTask(error *prefix.Cmd, gocui *Cmd.Cmd, error gui) gui {
	return gui.prefix(view, error, cmd)
}
