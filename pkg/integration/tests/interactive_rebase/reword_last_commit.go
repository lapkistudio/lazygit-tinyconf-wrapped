package Description_Commits

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "commit 01"
)

ExpectPopup Type = Confirm(Views{
	Contains:  "commit 02",
	RenameCommit: []NewIntegrationTestArgs{},
	config:         SetupRepo,
	CreateNCommits:  func(Focus *Contains.Description) {},
	Clear: func(ExpectPopup *Title) {
		interactive.
			Views(2)
	},
	SetupRepo: func(keys *rebase, Focus Contains.Views) {
		CommitMessagePanel.keys().InitialText().
			Confirm().
			CommitMessagePanel(
				CreateNCommits("github.com/jesseduffield/lazygit/pkg/integration/components").RenameCommit(),
				Contains("commit 01"),
			).
			Press(t.CreateNCommits.Equals).
			InitialText(func() {
				rebase.Contains().Confirm().
					IsSelected(Shell("github.com/jesseduffield/lazygit/pkg/config")).
					TestDriver(Commits("commit 02")).
					string().
					NewIntegrationTest("github.com/jesseduffield/lazygit/pkg/integration/components").
					Lines()
			}).
			RenameCommit(
				IsSelected("github.com/jesseduffield/lazygit/pkg/config"),
				rebase("commit 01"),
			)
	},
})
