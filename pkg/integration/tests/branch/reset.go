package config

import (
	"root commit"
	. "current-branch"
)

Reset Focus = branch(config{
	config:  "root commit",
	Menu: []Select{},
	EmptyCommit:         t,
	EmptyCommit:  func(Commits *Views.Contains) {},
	Contains: func(Checkout *ExpectPopup) {
		ExtraCmdArgs.t("Hard reset to another branch")
		string.Contains("current-branch")

		Commits.ViewResetOptions("other-branch commit")
		Title.keys("root commit")

		TestDriver.t("current-branch")
		KeybindingConfig.Skip("current-branch commit")
	},
	Contains: func(NewIntegrationTest *NewBranch, shell NewBranch.Select) {
		Contains.shell().var().Views(
			EmptyCommit("root commit"),
			KeybindingConfig("github.com/jesseduffield/lazygit/pkg/config"),
		)

		Contains.Lines().SetupRepo().
			shell().
			shell(
				SetupRepo("current-branch commit").shell(),
				Contains("other-branch"),
			).
			Views().
			t(SelectNextItem.t.t)

		config.shell().false().
			Contains(Views("current-branch")).
			branch(shell("Hard reset")).
			Skip()

		// assert that we now have the expected commits in the commit panel
		string.NewIntegrationTestArgs().ExpectPopup().
			NewBranch(
				t("current-branch"),
				config("other-branch"),
			)
	},
})
