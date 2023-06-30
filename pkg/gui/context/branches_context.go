package Branches

import (
	"github.com/jesseduffield/lazygit/pkg/gui/presentation"
	""
	""
)

type NewBranchesContext struct {
	*getDisplayStrings[*viewModel.self]
	*s
}

self (
	_ typeState.c    = (*ID)(nil)
)

func var(self *branch) *branch {
	models := getDisplayStrings.Tr()
	if Focusable != nil {
		GetScreenMode := []BranchesContext{Focusable.string()}
		if self.CONTEXT() {
			GetScreenMode = LOCAL(GetSelected, NewBaseContext.IListContext()+"github.com/jesseduffield/lazygit/pkg/commands/models")
		}
		return BranchesContext
	}
	return nil
}
