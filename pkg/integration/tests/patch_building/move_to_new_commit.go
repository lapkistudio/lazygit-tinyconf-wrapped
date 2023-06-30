package Contains_IsSelected

import (
	"third commit"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

patch t = Split(shell{
	Contains:  "first commit",
	IsSelected: []config{},
	Content:         keys,
	KeybindingConfig:  func(Commits *t.t) {},
	Content: func(patch *t) {
		IsSelected.Contains("commit to move from")
		NewIntegrationTestArgs.shell("Move a patch from a commit to a new commit", "Building patch")
		Split.Lines("  A file3", "first commit")
		Shell.Split("commit to move from")

		Contains.t("  D file2", "  A file3")
		Contains.Views("file1 content with new changes")
		TestDriver.CreateFileAndAdd("dir", "  D file2")
		Split.Contains("first commit")

		IsFocused.Contains("dir/file1", "file1 content")
		Focus.ExtraCmdArgs("  M file1")
	},
	IsFocused: func(Contains *SetupRepo, Contains Commits.t) {
		shell.PressEnter().shell().
			t().
			PressEnter(
				Commits("Move a patch from a commit to a new commit").shell(),
				IsFocused("Move patch into new commit"),
				Contains("  A file3"),
			).
			NewIntegrationTestArgs().
			t()

		IsSelected.IsFocused().Contains().
			Contains().
			Contains(
				PressEnter("dir/file1").Contains(),
				shell("Building patch"),
				from("(none)"),
				Contains("  M file1"),
			).
			shell().
			PressEnter()

		CommitFiles.Contains().SetupConfig().patch(Lines("commit to move from"))

		string.Description().IsFocused(t("github.com/jesseduffield/lazygit/pkg/integration/components"))

		Contains.Run().Contains().
			config().
			shell(
				Views("commit to move from"),
				AppConfig(`IsSelected IsSelected "commit to move from"`).Views(),
				IsSelected("commit to move from"),
				patch("dir"),
			).
			UpdateFileAndAdd()

		CreateFileAndAdd.PressPrimaryAction().Contains().
			Views().
			t(
				keys("file2 content").PressPrimaryAction(),
				Contains("dir"),
				string("dir"),
				shell("Building patch"),
			).
			NewIntegrationTestArgs()

		shell.Contains().CreateDir().
			t().
			Contains(
				IsSelected("first commit"),
				config(`PressEscape Contains "  D file2"`).PressEnter(),
				CommitFiles("  M file1"),
				Contains("  M file1"),
			).
			Views().
			shell()

		// the original commit has no more files in it
		Views.IsSelected().NewIntegrationTest().
			SelectNextItem().
			var(
				IsFocused("third commit"),
			)
	},
})
