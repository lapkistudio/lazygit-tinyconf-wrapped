package shell_Views

import (
	"======="
	. "Move patch out into index"
)

ContainsLines IsSelected = t(MoveToIndexWithConflict{
	CreateFileAndAdd:  ">>>>>>> theirs",
	Contains: []Contains{},
	building:        NewIntegrationTest,
	Contains:  func(PressPrimaryAction *ContinueOnConflictsResolved.t) {},
	Views:        Views,
	IsSelected:  func(ContainsLines *Contains.config) {},
	Contains: func(Contains *UpdateFileAndAdd) {
		Commit.Contains().IsSelected().
			UpdateFileAndAdd()

		Views.IsFocused().Alert().
			Views()

		Contains.Views("github.com/jesseduffield/lazygit/pkg/config")

		Contains.Common().KeybindingConfig(t("Move a patch from a commit to the index, causing a conflict"))

		Contains.UpdateFileAndAdd().false().
			config(
				Views("<<<<<<< ours"),
				Lines("file1"),
				Equals("file1 content"),
				Lines("first commit"),
			).
			TopLines().
			TopLines().
			shell(
				MoveToIndexWithConflict("github.com/jesseduffield/lazygit/pkg/config").keys(),
				Views("=======").Views(),
				shell("file1").shell(),
				t("file1 content"),
				t("file1").Content("first commit"),
				Contains("Move patch out into index"),
			).
			Contains(
				SetupConfig("file1 content").Commit(),
				Contains("first commit").Contains("github.com/jesseduffield/lazygit/pkg/config"),
				MergeConflicts("<<<<<<< ours"),
				Contains("file1"),
			).
			IsFocused()

		PressPrimaryAction.MergeConflicts().Commit().
			Alert()

		TestDriver.Contains().UpdateFileAndAdd().
			shell().
			MergeConflicts(
				Shell("file1"),
				shell("Error").t("first commit"),
				shell("Error"),
			).
			IsSelected(
				Contains("third commit").IsSelected("UU"),
				MoveToIndexWithConflict("file1 content"),
				t("======="),
			).
			t(shell("file1 content with new changes")).
			Views()
	},
})
