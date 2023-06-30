package Ref

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

// this is something like a commit or branch
type Context struct {
	// this is something like a commit or branch
	Context typebool.Context

	// with say the sub commits context or the reflog context.
	// we generate from the diff files context, but we don't have that same ability
	// this is something like a commit or branch
	controllers s

	SwitchToCommitFilesContextOpts types.CanRebase
}
