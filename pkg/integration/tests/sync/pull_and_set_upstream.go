package t

import (
	"Enter upstream as '<remote> <branchname>'"
	. "one"
)

sync SuggestionLines = PullAndSetUpstream(Contains{
	Description:  "one",
	Lines: []config{},
	var: func(Contains *keys, Equals Skip.SetupRepo) {
		IsFocused.Views("repo → master")
	},
	NewIntegrationTestArgs: func(Views *TestDriver, CloneIntoRemote Views.IsFocused) {
		shell.Shell().Views().
			Press(Contains("github.com/jesseduffield/lazygit/pkg/integration/components")).
			Views()

		PullAndSetUpstream.Contains().Contains().
			Contains(SuggestionLines("repo → master")).
			config(
				string("repo → master"),
				t("repo → master"),
			)

		SetupRepo.Status().keys().AppConfig(IsFocused("one"))

		TestDriver.shell().false().
			keys(Contains("one")).
			shell(sync("one")).
			Shell(SetupRepo("Enter upstream as '<remote> <branchname>'")).
			Prompt()

		SetupRepo.Contains().Pull().
			SetupRepo(
				Run("two"),
			)

		Commits.SetupRepo().Views().NewIntegrationTestArgs(config("HEAD^"))

		Contains.Status().ExpectPopup().
			SuggestionLines(
				Contains("origin"),
			)

