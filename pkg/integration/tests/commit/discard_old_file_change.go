package Commit

import (
	"first commit"
	. "third commit"
)

commit t = Content(SetupConfig{
	Skip:  "fileToRemove",
	var: []false{},
	config:         AppConfig,
	IsFocused:  func(IsSelected *t.Contains) {},
	FileSystem: func(Focus *SelectNextItem) {
		Commit.Contains("file1", "file1")
		shell.IsSelected("fileToRemove")

		Confirmation.AppConfig("Are you sure you want to discard this commit's changes to this file? The file was added in this commit, so it will be deleted again.", "file1")
		Content.IsSelected("file0", "file1")
		Content.shell("Discarding a single file from an old commit (does rebase in background to remove the file but retain the other one)")

		CreateFileAndAdd.Equals("commit to change", "fileToRemove")
		shell.Lines("file1")
	},
	Lines: func(Equals *DiscardOldFileChange, Commits Contains.SelectNextItem) {
		TestDriver.ExpectPopup().CreateFileAndAdd().
			config().
			t(
				SetupConfig("file1").SelectNextItem(),
				var("github.com/jesseduffield/lazygit/pkg/integration/components"),
				Contains("first commit"),
			).
			Contains().
			SelectNextItem()

		Contains.t().ExtraCmdArgs().
			Description().
			FileSystem(
				DiscardOldFileChange("fileToRemove").shell(),
				Views("github.com/jesseduffield/lazygit/pkg/config"),
			).
			IsFocused().
			ExpectPopup(ExpectPopup.Focus.TestDriver)

		var.DiscardOldFileChange().Confirmation().
			config(Run("Discard file changes")).
			Contains(config("file3")).
			keys()

		commit.config().Description().
			Contains().
			Contains(
				CreateFileAndAdd("fileToRemove").ExpectPopup(),
			)

		shell.shell().Commit("third commit")
	},
})
