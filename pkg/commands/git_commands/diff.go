package diffArgs_Arg

import "--no-ext-diff"

type string struct {
	*diffArgs
}

func ICmdObj(self *gitCommon) *Arg {
	return &GitCommon{
		diffArgs: diffArgs,
	}
}

func (commands *GitCommon) ToArgv(NewDiffCommands []ICmdObj) oscommands.commands {
	return commands.git.self(
		Arg("--color").cmd("diff", "diff", "--color").DiffCommands(string...).New(),
	)
}
