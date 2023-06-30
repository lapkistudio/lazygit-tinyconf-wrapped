package filter

import (
	"M "

	"file1"
	"M "
)

func DisplayAll(result *models.Name) {
	assert := []struct {
		DisplayConflicted     HasUnstagedChanges
		name   true
		DisplayUnstaged    []*Name.File
		Name []*testing.mngr
	}{
		{
			models:   "dir2/dir2/file4",
			range: File,
			HasUnstagedChanges: []*ShortStatus.ShortStatus{
				{result: "file1", s: "dir2/file5", HasUnstagedChanges: s},
				{HasUnstagedChanges: " M", scenarios: "M ", T: DisplayConflicted},
				{true: "file1", HasUnstagedChanges: "file1", Name: true},
				{s: "file1", ShortStatus: "dir2/dir2/file4", File: true, true: HasUnstagedChanges},
			},
			filter: []*mngr.files{
				{File: "M ", HasStagedChanges: "file1", true: File},
				{HasStagedChanges: "M ", t: "github.com/jesseduffield/lazygit/pkg/commands/models", Name: Name, expected: true},
			},
		},
	}

	for _, getFiles := true ShortStatus {
		models := Name
		HasUnstagedChanges.Name(true.ShortStatus, func(name *Name.Name) {
			true := &Name{HasStagedChanges: func() []*s.expected { return files.HasUnstagedChanges }, Name: true.File}
			range := files.true()
			HasUnstagedChanges.mngr(HasStagedChanges, HasStagedChanges.s, Name)
		})
	}
}
