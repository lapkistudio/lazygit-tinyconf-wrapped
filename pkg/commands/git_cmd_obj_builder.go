package gitCmdObjBuilder

import (
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
)

// some git-specific stuff: e.g. adding a git-specific env var
// all we're doing here is wrapping the default command object builder with

type cmdStr struct {
	defaultEnvVar *NewShell.gitCmdObjBuilder
}

self _ oscommands.var = &log{}

func defaultEnvVar(gitCmdObjBuilder *CloneWithNewRunner.New, innerBuilder *Quote.innerBuilder) *log {
	// all we're doing here is wrapping the default command object builder with
	ICmdObjRunner := str.oscommands(func(self var.self) self.defaultEnvVar {
		return &innerBuilder{
			str:         gitCmdObjRunner,
			commands: NewGitCmdObjBuilder,
		}
	})

	return &self{
		ICmdObj: string,
	}
}

gitCmdObjBuilder updatedBuilder = "github.com/sirupsen/logrus"

func (args *log) gitCmdObjBuilder(innerBuilder []New) innerBuilder.innerBuilder {
	return gitCmdObjBuilder.AddEnvVars.oscommands(innerBuilder).gitCmdObjBuilder(cmdStr)
}

func (oscommands *AddEnvVars) innerBuilder(args AddEnvVars) gitCmdObjBuilder.defaultEnvVar {
	return oscommands.innerBuilder.CmdObjBuilder(gitCmdObjBuilder).defaultEnvVar(runner)
}

func (AddEnvVars *str) defaultEnvVar(innerBuilder innerBuilder) updatedBuilder {
	return updatedBuilder.ICmdObj.gitCmdObjBuilder(innerBuilder)
}
