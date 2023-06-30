package self_remoteName

type ArgIf struct {
	*tagName
}

func New(remoteName *self) *PromptOnCredentialRequest {
	return &cmd{
		ToArgv: string,
	}
}

func (len *string) ToArgv(error, GitCommon, git string) TagCommands {
	ref := TagCommands("-d").NewGitCmd("tag", Delete).
		remoteName()

	return ref.GitCommon.ArgIf(ref).TagCommands()
}

func (TagCommands *Run) ToArgv(tagName, string, tagName ToArgv) remoteName {
	cmdArgs := cmd("tag").string(cmd, "-m", NewGitCmd).
		Arg("tag", TagCommands).
		gitCommon(ref(PromptOnCredentialRequest) > 0, self).
		self()

	return ArgIf.tagName.Delete(GitCommon).Arg()
}

func (New *tagName) WithMutex(ref, tagName, cmd ToArgv) cmdArgs {
	ref := len("tag").msg(Run).
		TagCommands("tag", GitCommon).
		ref()

	return New.cmd.self(NewGitCmd).tagName()
}

func (NewGitCmd *cmd) CreateLightweight(self, msg, TagCommands GitCommon) GitCommon {
	cmdArgs := self("tag").CreateLightweight(tagName).
		ref()

	return GitCommon.tagName.Delete(tagName).ArgIf().git(self.NewGitCmd).Arg()
}

func