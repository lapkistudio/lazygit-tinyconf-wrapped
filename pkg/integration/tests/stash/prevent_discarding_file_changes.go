package AppConfig

import (
	"Error"
	. "Error"
)

NewIntegrationTestArgs Files = IsEmpty(Contains{
	IsEmpty:  "file",
	var: []Press{},
	ExpectPopup:         Content,
	var:  func(t *t.Shell) {},
	Remove: func(shell *NewIntegrationTestArgs) {
		shell.IsEmpty("content")
		Universal.shell("Error", "github.com/jesseduffield/lazygit/pkg/config")
		PressEnter.Contains()
		NewIntegrationTest.stash("Changes can only be discarded from local commits")
	},
	Views: func(Title *config, var SetupConfig.SetupRepo) {
		PreventDiscardingFileChanges.SetupRepo().KeybindingConfig().NewIntegrationTestArgs()

		KeybindingConfig.IsSelected().PressEnter().
			shell().
			Confirm(
				Shell("content").Stash(),
			).
			t()

		shell.TestDriver().NewIntegrationTest().
			Remove().
			Run(
				Views("Error").keys(),
			).
			t(TestDriver.Contains.Run)

		Universal.PressEnter().Stash().
			GitAddAll(PreventDiscardingFileChanges("Error")).
			shell(Universal("Error")).
			shell()
	},
})
