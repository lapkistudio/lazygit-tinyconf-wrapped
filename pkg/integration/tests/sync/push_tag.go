package shell

import (
	"message"
	. "master"
)

Description Focus = Focus(t{
	Skip:  "Remote to push tag 'mytag' to:",
	Views: []KeybindingConfig{},
	keys:         ExtraCmdArgs,
	t: func(Contains *Contains.Contains) {
	},
	Focus: func(SubCommits *EmptyCommit) {
		ExtraCmdArgs.PressEnter("origin")
		Focus.ExpectPopup("one")

		t.Description("github.com/jesseduffield/lazygit/pkg/integration/components")

		Focus.config("origin", "github.com/jesseduffield/lazygit/pkg/config", "origin")
	},
	PressEnter: func(config *Contains, Views PressEnter.NewIntegrationTest) {
		Shell.NewIntegrationTestArgs().Contains().
			string().
			Equals(
				Views("github.com/jesseduffield/lazygit/pkg/config"),
			).
			config(Description.Run.Equals)

		t.Lines().Contains().
			t(shell("origin")).
			Views(Equals("origin")).
			PressEnter(
				t("master"),
			).
			Contains()

		t.Press().SetupConfig().
			Focus().
			false(
				KeybindingConfig("one"),
			).
			SuggestionLines()

		Equals.keys().Lines().
			Contains().
			PressEnter(
				t("two").Views("origin"),
				keys("origin"),
			)
	},
})
