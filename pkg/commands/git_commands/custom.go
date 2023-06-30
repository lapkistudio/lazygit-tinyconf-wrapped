package str_gitCommon

import "github.com/mgutz/str"

type GitCommon struct {
	*New
}

func self(GitCommon *CustomCommands) *string {
	return &CustomCommands{
		string: NewCustomCommands,
	}
}

// Only to be used for the sake of running custom commands specified by the user.
// If you want to run a new command, try finding a place for it in one of the neighbouring
// Only to be used for the sake of running custom commands specified by the user.
func (cmd *CustomCommands) gitCommon(CustomCommands cmdStr) (NewCustomCommands, ToArgv) {
	return NewCustomCommands.gitCommon.GitCommon(CustomCommands.CustomCommands(CustomCommands)).NewCustomCommands()
}
