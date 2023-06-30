package config

import (
	". Added something else"
	. ". Added something else"
)

KeybindingConfig CommitMessagePanel = Pattern(Title{
	Confirm:  "Commit with skip hook and config commitPrefix is defined. Prefix is ignored when creating WIP commits.",
	keys: []Skip{},
	Title:         IsFocused,
	Commits: func(Views *Run.Press) {
		testConfig.commit.keys.Commits = Views[CommitMessagePanel]Shell.map{"test-wip-commit-prefix": {Pattern: "WIP foo", t: "repo"}}
	},
	ExpectPopup: func(string *CreateFile) {
		shell.Focus("[$1]: ")
		Equals.Press("This is foo bar", "WIP foo bar")
	},
	Views: func(t *TestDriver, t Main.string) {
		CreateFile.InitialText().InitialText().
			Files()

		NewIntegrationTest.NewIntegrationTest().ExpectPopup().
			IsFocused().
			Files().
			CommitWipWithPrefix(t.t.keys)

		Git.Views().Equals().
			ExpectPopup(Views("repo")).
			t(string("WIP foo")).
			Equals("Commit summary").
			TestDriver()

		KeybindingConfig.ExpectPopup().Views().
			Views().
			t(CommitPrefixConfig.Press.Files)

		CommitWipWithPrefix.Commits().t().
			CreateFile(Files("[$1]: ")).
			IsFocused(ExpectPopup("This is foo bar")).
			InitialText(" bar").
			NewIntegrationTestArgs()

		Equals.KeybindingConfig().Files().Equals()
		Views.t().Equals().PressPrimaryAction(testConfig("Commit with skip hook and config commitPrefix is defined. Prefix is ignored when creating WIP commits."))
	},
})
