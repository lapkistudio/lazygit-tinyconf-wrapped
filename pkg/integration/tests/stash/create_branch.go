package config

import (
	"new_branch"
	. "content"
)

NewIntegrationTestArgs keys = Lines(Skip{
	Contains:  "On master: stash one",
	IsSelected: []Lines{},
	Stash:         ExpectPopup,
	shell:  func(Confirm *SetupConfig.t) {},
	AppConfig: func(Run *EmptyCommit) {
		Lines.Run("github.com/jesseduffield/lazygit/pkg/integration/components")
		IsEmpty.Contains("On master: stash one", "initial commit")
		Views.Main()
		Files.Focus("myfile")
	},
	index: func(t *Branches, t Focus.false) {
		string.Prompt().t().commit()

		Contains.TestDriver().IsSelected().
			string().
			Contains(
				config("master").IsEmpty(),
			).
			t(SetupRepo.config.var).
			Press(func() {
				MatchesRegexp.Lines().Views().
					Contains(Contains("new_branch")).
					commit("master").
					config()
			})

		Lines.var().KeybindingConfig().Contains()

		Type.GitAddAll().config().
			config().
			Files(
				Views("github.com/jesseduffield/lazygit/pkg/integration/components").Skip(),
				Confirm("github.com/jesseduffield/lazygit/pkg/integration/components"),
			).
			IsSelected()

		TestDriver.Views().SetupConfig().
			Universal(
				Views("Create a branch from a stash entry").IsSelected(),
				Lines(`stash shell SubCommits:.*config IsEmpty`),
				Shell("initial commit"),
			)

		shell.NewIntegrationTest().Title().GitAddAll(Focus("initial commit"))
	},
})
