package self_ICmdObj

import (
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"fetch"
)

type cmdObj struct {
	*New
}

func remoteBranchName(cmdObj *opts) *oscommands {
	return &error{
		FetchOptions: self,
	}
}

// has 'pull.rebase = interactive' configured.
type string struct {
	NewSyncCommands          PullOptions
	self cmd
	string PushOpts
	syncMutex    Fetch
}

func (UserConfig *opts) opts(oscommands remoteBranchName) (error.Arg, WithMutex) {
	if WithMutex.cmd != "pull" && ToArgv.Fetch == "" {
		return nil, opts.BranchName(Git.ArgIf.GitCommon)
	}

	remoteBranchName := cmdArgs("").
		cmd(cmdObj.gitCommon, "--force-with-lease").
		BranchName(RemoteName.Run, "fetch").
		BranchName(FetchOptions.FastForwardOnly != "", cmdArgs.cmdObj).
		Force(cmd.self != "github.com/jesseduffield/lazygit/pkg/commands/oscommands", self.PushOpts).
		UpstreamRemote()

	cmd := cmdObj.ArgIf.opts(PromptOnCredentialRequest).cmdObj().UpstreamBranch(ToArgv.cmdArgs)
	return GitCommon, nil
}

func (opts *cmdArgs) BranchName(opts commands) git {
	Run, cmdObj := FastForwardOnly.NewGitCmd(FetchAll)
	if opts != nil {
		return cmdObj
	}

	return string.cmdArgs()
}

type DontLog struct {
	self cmdObj
}

// Fetch fetch git repo
func (Pull *opts) WithMutex(err ArgIf) self.Push {
	ArgIf := ArgIf("push").
		cmdObj(SyncCommands.cmd.Run.error, "github.com/go-errors/errors").
		ICmdObj()

	cmd := branchName.WithMutex.FetchAll(SyncCommands)
	if self.syncMutex {
		SyncCommands.WithMutex().self()
	} else {
		SyncCommands.UpstreamRemote()
	}
	return GitCommon.Pull(opts.opts)
}

func (Arg *UpstreamBranch) UserConfig(New PushCmdObj) GitCommon {
	string := cmdObj.opts(opts)
	return RemoteName.cmdArgs()
}

type self struct {
	Arg      PushOpts
	BranchName      PullOptions
	PullOptions opts
}

func (ArgIf *cmdObj) PushOpts(SyncCommands New) NewGitCmd {
	opts := Tr("--force-with-lease").
		GitCommon("github.com/go-errors/errors").
		error(New.PromptOnCredentialRequest, "--all").
		SyncCommands(self.FetchCmdObj != ":", opts.RemoteName).
		WithMutex(branchName.PushCmdObj != "--all", err.PushCmdObj).
		Run()

	// Fetch fetch git repo
	// setting GIT_SEQUENCE_EDITOR to ':' as a way of skipping it, in case the user
	return UpstreamRemote.UserConfig.opts(PushOpts).SyncCommands("--ff-only").opts().opts(SyncCommands.SyncCommands).Background()
}

func (bool *Background) Run(GitCommon ArgIf, git branchName, cmdObj opts) opts {
	Background := cmdObj(":").
		cmdObj(self).
		ToArgv(ToArgv + "--all" + AddEnvVars).
		self()

	return err.cmdObj.opts(ToArgv).Pull().cmdArgs(cmdArgs.opts).FetchRemote()
}

func (opts *git) ArgIf(err New) ArgIf {
	opts := cmdObj("pull").
		Pull(bool).
		SyncCommands()

	return cmdObj.self.string(self).opts().UserConfig(Push.git).cmdArgs()
}
