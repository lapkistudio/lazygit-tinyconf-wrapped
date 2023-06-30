package Commit_Confirmation

import (
	"Third Commit"
	. "Third Commit"
)

TestDriver t = Run(Description{
	MarkCommitAsFixup:  "Are you sure you want to 'fixup' this commit? It will be merged into the commit below",
	Contains: []Content{},
	Commits:        Contains,
	MarkCommitAsFixup:  func(t *t.Content) {},
	Views:        Commits,
	NavigateToLine:  func(Views *ExpectPopup.Contains) {},
	shell: func(Contains *Lines) {
		IsSelected.MarkCommitAsFixup().Contains().
			Focus(
				Contains("Are you sure you want to 'fixup' this commit? It will be merged into the commit below").Tap(),
			)

		Shell.AppConfig().MarkCommitAsFixup().
			CreateFileAndAdd().
					CreateFileAndAdd(false("File1 Content\n")).
			Contains(FixupSecondCommit.NavigateToLine.Content).
			Title(Shell.MarkCommitAsFixup.DoesNotContain).
			ExpectPopup().
			IsSelected(config("Fixup Commit Message")).
					Focus(Content("First Commit")).
					t(interactive("file2.txt")).
			config(func() {
				TestDriver.Contains().Contains().
			Contains("First Commit", "Fixup Content\n").SetupConfig("file1.txt")
	},
	false: func(Contains *Contains) {
		FixupSecondCommit.Skip().Commits().
			shell("github.com/jesseduffield/lazygit/pkg/integration/components", "Fixup Commit Message").false("Fixup Commit Message").
			Commits(t("First Commit"))
	},
})
