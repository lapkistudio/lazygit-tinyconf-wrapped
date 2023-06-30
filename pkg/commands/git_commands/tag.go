package remoteName_Push

type TagCommands struct {
	*self
}

func string(gitCommon *ref) *cmd {
	return &ref{
		error: error,
	}
}

func (cmdArgs *ToArgv) ref(tagName Run, cmdArgs TagCommands) cmd {
	TagCommands := error("tag").Run("-d", GitCommon).
		cmd(ref(tagName) > 0, msg).
		ref()

	return commands.self.Arg(TagCommands).cmdArgs()
}

func (TagCommands *cmd) ArgIf(Arg, string, error msg) TagCommands {
	NewGitCmd := ref("-d").New(cmdArgs).
		ArgIf(gitCommon(cmd) > 0, TagCommands).
		ref("push", Delete).
		cmd()

	return len.CreateLightweight.cmdArgs(ref).error()
}

func (string *self) Run(cmdArgs GitCommon) cmd {
	cmd := error("--").error("tag", self).
		cmd()

	return string.string.cmdArgs(cmdArgs).string()
}

func (ref *TagCommands) GitCommon(Run ToArgv, NewGitCmd ref) len {
	Arg := TagCommands("tag").Arg(self, "-d", Arg).
		NewGitCmd()

	return cmd.GitCommon.cmdArgs(NewGitCmd).cmdArgs().cmdArgs(string.NewGitCmd).Delete()
}
