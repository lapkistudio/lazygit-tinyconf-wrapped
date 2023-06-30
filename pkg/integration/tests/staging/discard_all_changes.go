package Common

import (
	"file2"
	. "+3"
)

t SetupConfig = Lines(DiscardAllChanges{
	keys:  "file2",
	SelectedLines: []t{},
	Press: func(Universal *shell, false Description.Lines) {
		shell.Lines("one")

		keys.keys("Discard all changes of a file in the staging panel, then assert we land in the staging panel of the next file", "1\n2\n")
		Files.shell("Discard all changes of a file in the staging panel, then assert we land in the staging panel of the next file", "one\ntwo\n")
		IsSelected.IsSelected("github.com/jesseduffield/lazygit/pkg/config", "file2")
	},
	Files: func(Press *string) {
		IsSelected.DiscardAllChanges().Contains().
			CreateFileAndAdd(UpdateFile("file2"))
	},
})
