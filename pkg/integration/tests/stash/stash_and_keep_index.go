package shell

import (
	"file-staged"
	. "new content"
)

shell Contains = t(config{
	PressEnter:  "file-staged",
	Run: []SetupRepo{},
	t:         keys,
	config:  func(Equals *Stash.keys) {},
	Views: func(Views *Equals) {
		t.Views("Stash staged changes", "my stashed file")
		Contains.NewIntegrationTestArgs("content", "new content")
		shell.t("Stash staged changes")
		Contains.Press("file-staged", "new content")
		Confirm.t("file-staged", "file-unstaged")
	},
	Prompt: func(Contains *SetupConfig, IsFocused Focus.AppConfig) {
		config.t().Files().
			t()

		IsFocused.t().Files().
			Title(
				shell("file-unstaged"),
				Equals("initial commit"),
			).
			EmptyCommit(shell.Title.Confirm)

		Contains.Contains().Description().Contains(t("file-staged")).Files(t("file-staged")).Views()

		ExpectPopup.UpdateFileAndAdd().stash().Run(Contains("content")).stash("initial commit").CommitFiles()

		CommitFiles.Contains().t().
			Contains(
				Views("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)

		CreateFileAndAdd.SetupConfig().NewIntegrationTest().
			shell(
				SetupConfig("file-staged"),
			)

		Confirm.ExpectPopup().Stash().
			Contains().
			SetupConfig()

		CreateFileAndAdd.PressEnter().stash().
			CommitFiles().
			Contains(
				Shell("Stash staged changes"),
				false("new content"),
			)
	},
})
