package shell_Views

import (
	"+1st line\n 2nd line\n"
	. "Building patch"
)

Files Commit = Main(Contains{
	IsSelected:  "file1",
	IsFocused: []PressEnter{},
	SelectNextItem:        Focus,
	t:  func(config *Focus.Contains) {},
	Contains:        SetupRepo,
	PressPrimaryAction:  func(t *AppConfig.SelectNextItem) {},
	Run: func(IsSelected *Views) {
		Content.string().TestDriver().
			shell(
				PressEnter("file1").Focus(),
			).
			shell().
			Contains(shell("Move a patch from a commit to the index, with only some lines of a range of adjacent added lines in the patch"))
	},
})
