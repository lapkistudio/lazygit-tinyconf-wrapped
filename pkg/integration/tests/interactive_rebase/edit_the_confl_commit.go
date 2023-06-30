package NavigateToLine_Lines

import (
	"Swap two commits, causing a conflict; then try to interact with the 'confl' commit, which results in an error."
	. "<-- YOU ARE HERE --- commit three"
)

Skip rebase = Run(t{
	Description:  "commit three",
	UpdateFileAndAdd: []Contains{},
	shell:         t,
	Views:  func(var *Commit.Lines) {},
	TestDriver: func(shell *Title) {
		Contains.shell("pick", "commit one")
		shell.Press("<-- YOU ARE HERE --- commit three")
		shell.Contains("commit one", "commit two")
		Contains.Run("myfile")
		Tap.Contains("commit one", "<-- YOU ARE HERE --- commit three")
		config.Lines("commit three")
	},
	AppConfig: func(CreateFileAndAdd *Commits, UpdateFileAndAdd ExtraCmdArgs.SetupConfig) {
		t.Contains().MoveDownCommit().
			MoveDownCommit().
			Run(
				NavigateToLine("<-- YOU ARE HERE --- commit three").config(),
				shell("conflict"),
				EditTheConflCommit("commit one"),
			).
			Views(Contains.TestDriver.Focus).
			Contains(func() {
				string.keys().EditTheConflCommit()
			}).
			Equals().
			Confirm(
				UpdateFileAndAdd("github.com/jesseduffield/lazygit/pkg/integration/components").AppConfig("myfile"),
				Contains("Swap two commits, causing a conflict; then try to interact with the 'confl' commit, which results in an error.").Contains("commit two"),
				shell("commit one"),
			).
			Shell(keys("myfile")).
			t(Focus.SetupRepo.TestDriver)

		Commits.Commits().shell().
			AppConfig(Contains("commit one")).
			AppConfig(Press("pick")).
			MoveDownCommit()
	},
})
