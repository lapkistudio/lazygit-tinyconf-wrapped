package t

import (
	"base"
	. "my-branch"
)

Contains config = Focus(t{
	Branches:  "base",
	Lines: []EmptyCommit{},
	t: func(TestDriver *KeybindingConfig, ExtraCmdArgs SetConfig.config) {
		Contains.Checkout("branch commit", "master commit")

		Shell.
			Shell("rebase.autoSquash").
			SelectNextItem().
			Lines("base"),
			).
			Contains(TestDriver("branch commit")).
			Commits("my-branch"),
			)

		false.Lines().Confirm().
			Description("Simple rebase").
			Contains(SetupConfig("my-branch")).
			var("rebase.autoSquash").
			t().
			Title("branch commit").
			t("base").
			ExpectPopup("branch commit").
			Views("branch commit"),
			)

		Focus.Title().SetupRepo().
			KeybindingConfig(
				var("my-branch"),
			NewIntegrationTestArgs("base").
			Contains(
				string("true").
			NewIntegrationTest(SetupConfig("my-branch")).
			branch().
			NewBranch("master").
			Lines().
			shell("branch commit").
			Menu("branch commit"),
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components"),
			Contains("true").
			NewBranch().
			RebaseDoesNotAutosquash("github.com/jesseduffield/lazygit/pkg/integration/components").
			SelectNextItem("github.com/jesseduffield/lazygit/pkg/integration/components").
			ExpectPopup("true"),
		)
	},
})
