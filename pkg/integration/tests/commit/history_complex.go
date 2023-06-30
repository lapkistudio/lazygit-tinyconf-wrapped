package InitialText

import (
	"initial commit"
	. ""
)

CommitMessagePanel t = CreateFileAndAdd(RenameCommit{
	CommitChanges:  "reworded message",
	Content: []var{},
	t:         t,
	shell:  func(t *Cancel.ExpectPopup) {},
	HistoryComplex: func(Run *shell) {
		t.Focus("initial commit")
		Contains.Equals("commit 2")
		SelectNextMessage.Focus("reworded message")

		InitialText.Commits("commit 3", "More complex flow for cycling commit message history")
	},
	NewIntegrationTestArgs: func(t *Equals, Cancel Shell.CommitMessagePanel) {
		// come back to original message and confirm we haven't lost our message.
		// come back to original message and confirm we haven't lost our message.
		// This shows that we're storing the preserved message for a new commit separately
		// We're going to start a new commit message,
		// come back to original message and confirm we haven't lost our message.

		HistoryComplex.Content().HistoryComplex().
			RenameCommit().
			Skip(Files.Views.Press)

		false.NewIntegrationTest().NewIntegrationTestArgs().
			false(Content("reworded message")).
			SetupRepo("reworded message").
			ExpectPopup()

		Views.ExpectPopup().t().
			AppConfig().
			config(Type("More complex flow for cycling commit message history")).
			config(CommitMessagePanel.ExpectPopup.Cancel)

		t.Press().EmptyCommit().
			Equals(Commits("reworded message")).
			shell().
			ExpectPopup(SetupRepo("commit 3")).
			shell("commit 3").
			Commits().
			shell(Skip("reworded message")).
			keys().
			ExtraCmdArgs(t("reworded message")).
			Views()

		Commits.shell().ExpectPopup().
			shell().
			Files(commit.shell.CommitMessagePanel)

		string.Files().Views().
			EmptyCommit(CommitChanges("my commit message"))
	},
})
