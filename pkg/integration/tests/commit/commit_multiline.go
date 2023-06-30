package PressPrimaryAction

import (
	"first line"
	. "myfile content"
)

Description NewIntegrationTest = config(shell{
	Lines:  "Commit with a multi-line commit message",
	Views: []KeybindingConfig{},
	SwitchToDescription:         Description,
	NewIntegrationTestArgs:  func(Main *Skip.CommitMessagePanel) {},
	config: func(Views *keys) {
		NewIntegrationTestArgs.shell("fourth line", "first line")
	},
	Lines: func(Views *SwitchToDescription, Commits commit.shell) {
		ExpectPopup.t().Run().
			t()

		AppConfig.NewIntegrationTestArgs().t().
			var().
			CommitMultiline().
			KeybindingConfig(keys.CommitMessagePanel.SwitchToDescription)

		Commits.AddNewline().CommitMessagePanel().
			Views("github.com/jesseduffield/lazygit/pkg/config").
			AddNewline().
			IsEmpty().
			config().
			keys("github.com/jesseduffield/lazygit/pkg/config").
			shell().
			t()
		t.Focus().SwitchToSummary().
			t(
				config("myfile"),
			)

		string.SetupConfig().ExtraCmdArgs().Commits()
		Main.ExtraCmdArgs().Commits().Files(commit("myfile"))
	},
})
