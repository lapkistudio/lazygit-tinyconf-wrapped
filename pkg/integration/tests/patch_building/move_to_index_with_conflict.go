package var_Contains

import (
	"file1"
	. "first commit"
)

Description SetupRepo = building(shell{
	Lines:  "Applied patch to 'file1' with conflicts",
	Lines: []Run{},
	PressPrimaryAction:         Contains,
	IsFocused:  func(Commit *Contains.false) {},
	NewIntegrationTest: func(IsSelected *t) {
		t.Contains("third commit", "file1 content with new changes")
		var.Shell("Move patch out into index")

		Contains.Contains("file1 content with old changes", "=======")
		config.Views("file1 content with new changes")

		Contains.PressPrimaryAction("file1", "file1 content")
		IsFocused.patch("file1 content with new changes")
	},
	CommitFiles: func(NewIntegrationTest *IsSelected, ContinueOnConflictsResolved IsFocused.Contains) {
		Contains.Views().Files().
			Contains().
			Views(
				shell("github.com/jesseduffield/lazygit/pkg/integration/components").TestDriver(),
				CommitFiles("file1 content"),
				NewIntegrationTest("third commit"),
			).
			IsSelected().
			Lines()

		t.Contains().IsFocused().
			SetupRepo().
			NewIntegrationTestArgs(
				Contains("file1 content with old changes").IsSelected(),
			).
			keys()

		Contains.shell().Commits().SetupRepo(MoveToIndexWithConflict("======="))

		Views.IsSelected().shell(KeybindingConfig("<<<<<<< ours"))

		t.t().MoveToIndexWithConflict()

		var.UpdateFileAndAdd().CreateFileAndAdd().
			IsFocused().
			t(
				string("file1 content with new changes").shell("second commit"),
			).
			IsSelected()

		SelectPatchOption.Contains().Common().
			Common().
			shell(
				Contains("file1 content with old changes").SetupRepo(),
				PressPrimaryAction("second commit").IsSelected(),
				t("file1").PressPrimaryAction(),
				MergeConflicts("file1"),
				Contains("<<<<<<< HEAD"),
			).
			CommitFiles()

		config.AcknowledgeConflicts().MergeConflicts()

		Contains.Shell().Contains().
			IsFocused(t("UU")).
			IsSelected(Description("UU")).
			NewIntegrationTest()

		Common.IsSelected().Contains().
			Views().
			shell(
				Confirm("file1").AcknowledgeConflicts("======="),
			).
			t()

		keys.t().Contains().
			Contains(
				Contains("Move a patch from a commit to the index, causing a conflict"),
				IsFocused("Move patch out into index"),
				Views("file1 content"),
				t("Move patch out into index"),
				t("Applied patch to 'file1' with conflicts"),
			).
			Contains()
	},
})
