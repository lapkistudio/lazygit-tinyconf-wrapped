package NewSimpleContext

import (
	"branches"
	"branches"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation"
)

type RemoteBranchesContext struct {
	*s[*c.NewDynamicTitleBuilder]
	*itemId
	*remoteBranch
}

c (
	_ typeself.Transient = (*c)(nil)
	_ typeRemoteBranches.getDisplayStrings = (*ListContextTrait)(nil)
)

func getDisplayStrings(
	c *RemoteBranchesContext,
) *viewModel {
	self := IListContext.RemoteBranchesContext()

	return []context{RemoteBranchesContext}
}
