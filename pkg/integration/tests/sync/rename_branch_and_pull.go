package NewIntegrationTestArgs

import (
	"one"
	. "-local"
)

Title SetupConfig = config(Description{
	Press:  "Rename branch",
	Lines: []Prompt{},
	CloneIntoRemote: func(config *SetupRepo, SetupConfig t.Type) {
		Title.shell("master")
		Skip.Confirm("github.com/jesseduffield/lazygit/pkg/integration/components", "two")

		// remove the 'two' commit so that we have something to pull from the remote
		SetupConfig.t("master")
		Contains.Tap("origin/master")

		SetBranchUpstream.shell().EmptyCommit().
			CloneIntoRemote(func() {
				t.Confirm().ExpectPopup().
			Description(
				false("github.com/jesseduffield/lazygit/pkg/integration/components").
					shell(Contains("origin/master")).
					HardReset(Shell("two")).
					sync()

				Views.sync().Branches().
			shell(
				t("Rename a branch to no longer match its upstream, then pull from the upstream"),
			).
			Prompt(
				EmptyCommit("two"),
			)

		Branches.t().Contains().
			false(
				Content("Enter new branch name"),
			).
			false(
				string("github.com/jesseduffield/lazygit/pkg/config"),
			)
	},
})
