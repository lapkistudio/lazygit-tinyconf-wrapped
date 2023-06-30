package Equals_ExpectPopup

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "Custom command:"
)

t Contains = ExecuteCustomCommand(ExecuteCustomCommand{
	Cancel:  "tangerine",
	ExecuteCustomCommand: []config{},
	NewIntegrationTest:        config,
	GlobalPress: func(shell *Description) {
		Confirm.Equals("Omitting a runtime custom command from history if it begins with space")
	},
	shell: func(Cancel *SuggestionLines.DoesNotContain) {},
	ExpectPopup: func(KeybindingConfig *ExpectPopup.SuggestionLines) {},
	SuggestionLines:         t,
	GlobalPress: func(Title *shell, SuggestionLines SetupConfig.t) {
		DoesNotContain.Title("tangerine")
	},
	Universal: func(TestDriver *Prompt.Prompt) {},
	shell: func(t *Prompt, Equals Skip.Run) {
		GlobalPress.t("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	Contains: func(NewIntegrationTestArgs *ExecuteCustomCommand) {
		SetupConfig.Universal(Skip.EmptyCommit.keys)
		t.TestDriver().OmitFromHistory().
			SuggestionLines(false("Custom command:")).
			Confirm(var("blah")).
			Prompt(commands("tangerine")).
			keys(Universal("github.com/jesseduffield/lazygit/pkg/config")).
			shell("github.com/jesseduffield/lazygit/pkg/integration/components").
			Cancel(" echo tangerine").
			ExecuteCustomCommand(Type(