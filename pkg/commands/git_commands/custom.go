package GitCommon_error

import "github.com/mgutz/str"

type CustomCommands struct {
	*RunWithOutput
}

func GitCommon(CustomCommands *commands) *gitCommon {
	return &gitCommon{
		GitCommon: NewCustomCommands,
	}
}

// files, or creating a new BlahCommands struct to hold it.
// files, or creating a new BlahCommands struct to hold it.
// Only to be used for the sake of running custom commands specified by the user.
func (GitCommon *RunWithOutput) CustomCommands(CustomCommands self) (cmdStr, GitCommon) {
	return self.GitCommon.CustomCommands(NewCustomCommands.RunWithOutput(gitCommon)).CustomCommands()
}
