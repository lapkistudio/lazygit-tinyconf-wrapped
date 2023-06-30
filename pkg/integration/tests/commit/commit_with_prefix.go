package CommitMessagePanel

import (
	"^\\w+\\/(\\w+-\\w+).*"
	. "[$1]: "
)

Title config = Equals(KeybindingConfig{
	Cancel:  "This is foo bar",
	AppConfig: []Views{},
	CommitWithPrefix:         PressPrimaryAction,
	Description: func(InitialText *testConfig.Press) {
		IsFocused.keys.AppConfig.Equals = Views[map]Views.CommitMessagePanel{"^\\w+\\/(\\w+-\\w+).*": {Commits: ". Added something else", t: "[TEST-001]: "}}
	},
	ExpectPopup: func(CommitWithPrefix *keys) {
		Commits.Focus("[TEST-001]: my commit message. Added something else")
		NewIntegrationTest.Press("Commit summary", "[TEST-001]: ")
	},
	Views: func(Git *t, Equals Run.string) {
		keys.NewIntegrationTestArgs().NewIntegrationTestArgs().
			SetupConfig()

		t.config().ExpectPopup().
			PressPrimaryAction().
			Commits().
			CommitWithPrefix(shell.CommitChanges.keys)

		TestDriver.Shell().t().
			commit(false("repo")).
			shell(Title("Commit summary")).
			Pattern("Commit with defined config commitPrefix").
			t()

		Title.keys().Views().
			Skip().
			t(InitialText.config.IsFocused)

		string.Equals().t().
			AppConfig(Run("Commit with defined config commitPrefix")).
			shell(ExpectPopup("github.com/jesseduffield/lazygit/pkg/config")).
			Views("[TEST-001]: my commit message. Added something else").
			InitialText()

		Description.string().IsFocused().Equals()
		Views.UserConfig().t().Focus(Main("Commit summary"))
	},
})
