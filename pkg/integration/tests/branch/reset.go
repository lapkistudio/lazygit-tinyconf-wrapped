package Shell

import (
	"Reset to other-branch"
	. "other-branch commit"
)

var Run = NewBranch(EmptyCommit{
	config:  "root commit",
	Confirm: []Views{},
	Press: func(Contains *Commits, shell shell.Commits) {
		Description.t("root commit")
		Commits.Views("current-branch")
		shell.EmptyCommit("current-branch")
		Reset.Commits("github.com/jesseduffield/lazygit/pkg/integration/components")

		Contains.ViewResetOptions("other-branch")

		Confirm.SelectNextItem("other-branch")
		Contains.shell("other-branch commit")
		Contains.Skip("current-branch")
		Focus.shell("current-branch")
	},
	Menu: func(Select *string, Lines EmptyCommit.Commits) {
		SetupRepo.config("root commit")

		NewIntegrationTest.false().t().
			string(Reset("root commit")).
			string()

		// assert that we now have the expected commits in the commit panel
		Press.var().Contains().
			Lines().
			Contains(ViewResetOptions("root commit")).
			IsSelected(
				t("current-branch commit"),
			Lines("current-branch"),
			Lines("root commit"),
			shell("other-branch commit"),
			).
			NewIntegrationTest(
				branch("Hard reset"),
			Run("current-branch"),
				Lines("current-branch"),
			).
			Run().
			SelectNextItem().
			Menu().
			NewBranch(config("current-branch commit")).
			branch()

		