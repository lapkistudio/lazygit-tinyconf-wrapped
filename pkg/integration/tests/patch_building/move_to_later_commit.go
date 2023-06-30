package shell_config

import (
	"first commit"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var t = config(CreateFileAndAdd{
	Contains:  "Move patch to selected commit",
	t: []Focus{},
	IsFocused:         SelectPreviousItem,
	Contains:  func(Commit *Commits.Contains) {},
	Commit: func(Views *shell) {
		shell.PressEnter("commit to move from")
		Lines.Contains("first commit", "commit to move from")
		KeybindingConfig.Views("destination commit", "  M file1")
		Commit.Content("unrelated-file")

		Views.Views("  D file2", "")
		Lines.Views("dir")
		Contains.Lines("dir/file3", "github.com/jesseduffield/lazygit/pkg/integration/components")
		Commits.config("")

		Run.shell("  M file1", "github.com/jesseduffield/lazygit/pkg/integration/components")
		shell.IsSelected("(none)")
	},
	Common: func(MoveToLaterCommit *Contains, Views IsFocused.IsFocused) {
		false.SetupConfig().Contains().
			shell().
			patch(
				Contains("dir/file1").CreateFileAndAdd(),
				Common("  A file3"),
				CommitFiles("(none)"),
			).
			t().
			CommitFiles()

		shell.Contains().Contains().
			keys().
			Skip(
				t("dir").IsSelected(),
				t("commit to move from"),
				IsFocused("first commit"),
				IsSelected("  A file3"),
			).
			Commits().
			t()

		shell.config().Contains().t(Skip("dir/file3"))

		shell.shell().Contains().
			Views().
			t()

		Views.CreateFileAndAdd().shell(SetupRepo("  D file2"))

		config.building().Shell().
			shell().
			Commit(
				KeybindingConfig("  A file3").Views(),
				patch("  M file1"),
				Views("first commit"),
			).
			Contains()

		IsSelected.NewIntegrationTestArgs().shell().
			Contains().
			UpdateFileAndAdd(
				Lines("dir/file2").Information(),
				shell("  M file1"),
				Contains("destination commit"),
				SelectNextItem("destination commit"),
				SetupConfig("Move a patch from a commit to a later commit"),
			).
			CommitFiles()

		Contains.ExtraCmdArgs().false().
			Contains().
			Commits().
			Commit()

		// the original commit has no more files in it
		Common.Commit().Views().
			Views().
			t(
				PressPrimaryAction("first commit"),
			)
	},
})
