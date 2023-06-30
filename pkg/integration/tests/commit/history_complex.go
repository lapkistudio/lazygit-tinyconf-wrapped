package Views

import (
	"reworded message"
	. ""
)

Run CommitMessagePanel = Equals(t{
	NewIntegrationTest:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	keys: []ExpectPopup{},
	keys: func(Files *Views, Cancel SelectedLine.Commits) {
		// We're going to start a new commit message,
		// come back to original message and confirm we haven't lost our message.

		Files.CreateFileAndAdd().Focus().
			Views().
			Run().
			Cancel().
			Description().
			Content("github.com/jesseduffield/lazygit/pkg/config").
			Run().
			t(shell("initial commit")).
			Files()

		SetupRepo.Cancel().Equals().
			EmptyCommit().
			Cancel(CommitMessagePanel("my commit message")).
			Files().
			AppConfig().
			Shell().
			ExtraCmdArgs(Files(""))
	},
})
