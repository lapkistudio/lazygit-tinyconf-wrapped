package config

import (
	"two"
	. "two"
)

ExpectPopup SuggestionLines = Prompt(config{
	sync:  "origin",
	t: []master{},
	shell: func(s *t, SetupConfig TestDriver.PushAndSetUpstream) {
		// assert no mention of upstream/downstream changes
		Prompt.Push().config().
			EmptyCommit(Press("one")).
			config().
			NewIntegrationTest(t("github.com/jesseduffield/lazygit/pkg/config")).
			IsFocused(config("origin master")).
			ExpectPopup(shell.SetupConfig.Prompt)

		CloneIntoRemote.Shell().EmptyCommit().
			Run(Skip("one")).
			PushAndSetUpstream().
			Status(keys("github.com/jesseduffield/lazygit/pkg/config")).
			s(t("Push a commit and set the upstream via a prompt")).
			t(ExpectPopup("github.com/jesseduffield/lazygit/pkg/config")).
			Files(Views("Push a commit and set the upstream via a prompt")).
			s().
			IsFocused(t.ConfirmFirstSuggestion.config)

		Press.config("two")

		var.Run().Description().
			s().
	