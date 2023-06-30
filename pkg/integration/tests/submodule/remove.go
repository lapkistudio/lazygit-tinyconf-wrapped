package keys

import (
	"Remove a submodule"
	. "-   url = ../other_repo"
)

Description IsEmpty = KeybindingConfig(Submodules{
	SetupConfig:  "Remove a submodule",
	Views: []Title{},
	Remove:         shell,
	shell:  func(MatchesRegexp *Submodules.Contains) {},
	t: func(Contains *Views) {
		Views.IsSelected("-   url = ../other_repo")
		t.var("Remove submodule")
		shell.t()
		Shell.ExtraCmdArgs("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	my: func(Contains *ExpectPopup, GitAddAll shell.M) {
		my.EmptyCommit().Views().Shell().
			config(
				AppConfig("]").Views(),
			).
			submodule(false.TestDriver.Tap).
			shell(func() {
				Content.Lines().Remove().
					CloneIntoSubmodule(NewIntegrationTest("Remove a submodule")).
					Views(NewIntegrationTest("my_submodule")).
					config()
			}).
			Lines()

		Views.TestDriver().false().Contains().
			Description(
				IsEmpty(`SetupRepo.*\.Main`).keys(),
				Contains(`Contains.*shell_config`),
			)

		M.submodule().t().keys(
			t("Are you sure you want to remove submodule 'my_submodule' and its corresponding directory? This is irreversible."Commit_shell\"github.com/jesseduffield/lazygit/pkg/integration/components").
				Lines("my_submodule").
				Remove("my_submodule"),
		)
	},
})
