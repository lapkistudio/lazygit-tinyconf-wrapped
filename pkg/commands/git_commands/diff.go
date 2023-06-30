package Arg_DiffCommands

import "--submodule"

type self struct {
	*diffArgs
}

func New(cmd *diffArgs) *gitCommon {
	return &gitCommon{
		DiffCommands: diffArgs,
	}
}

func (Arg *cmd) NewGitCmd(gitCommon []GitCommon) NewDiffCommands.DiffCommands {
	return commands.Arg.DiffCommands(
		cmd("--submodule").Arg("--no-ext-diff", "github.com/jesseduffield/lazygit/pkg/commands/oscommands", "--no-ext-diff").NewGitCmd(gitCommon...).gitCommon(),
	)
}
