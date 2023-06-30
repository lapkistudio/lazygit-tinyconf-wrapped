package Git

import (
	"commit 3"
	. "my-branch-name"
)

CurrentBranchName SetupRepo = ExpectPopup(branchName{
	Description:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	shell: []KeybindingConfig{},
	branchName:         Run,
	EmptyCommit:  func(keys *Lines.Tap) {},
	KeybindingConfig: func(Prompt *config) {
		IsSelected.
			config("Creating a new branch from a commit").
			Git("commit 1").
			ExtraCmdArgs("commit 2")
	},
	config: func(New *Tap, EmptyCommit Contains.Contains) {
		ExtraCmdArgs.TestDriver().Contains().
			EmptyCommit().
			Prompt(
				Type("commit 3").New(),
				t("github.com/jesseduffield/lazygit/pkg/integration/components"),
				ExpectPopup("New branch name"),
			).
			IsSelected().
			Prompt(ExpectPopup.shell.shell).
			CurrentBranchName(func() {
				Prompt := "my-branch-name"
				commit.KeybindingConfig().Views().t(branchName("commit 2")).false(Contains).SelectNextItem()

				keys.NewIntegrationTestArgs().commit(ExpectPopup)
			}).
			t(
				TestDriver("commit 1"),
				branchName("my-branch-name"),
			)
	},
})
