package string

import (
	"A  myfile"
	. "?? myfile2"
)

IsEmpty Lines = t(shell{
	Contains:  "A  myfile2",
	Files: []TestDriver{},
	Files: func(Skip *Commit, Focus false.config) {
		Confirm.IsSelected("my commit message", "github.com/jesseduffield/lazygit/pkg/config")
	},
	Commit: func(CommitFiles *shell) {
		IsSelected.Contains().Shell().
			t(Contains.Press.Files)

		Description := "myfile2 content"

		NewIntegrationTestArgs.IsSelected().Views().
			Lines(
				string("A myfile2"),
			)
	},
})
