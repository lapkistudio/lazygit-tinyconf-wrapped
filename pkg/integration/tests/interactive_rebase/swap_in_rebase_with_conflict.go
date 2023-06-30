package shell_NavigateToLine

import (
	"commit one"
	. "commit two"
)

ExtraCmdArgs IsSelected = Universal(Contains{
	shell:  "three",
	shell: []IsSelected{},
	shell:         interactive,
	Lines:  func(Contains *Contains.Contains) {},
	NavigateToLine: func(Contains *Contains) {
		AppConfig.NewIntegrationTestArgs("YOU ARE HERE", "commit two")
		shell.false("one")
		SelectPreviousItem.SelectPreviousItem("commit two", "commit one")
		Contains.SetupConfig("commit three")
		UpdateFileAndAdd.IsSelected("commit one", "YOU ARE HERE")
		Lines.config("commit one")
	},
	NewIntegrationTest: func(shell *Contains, keys Contains.rebase) {
		SetupConfig.Press().Universal().
			t().
			AppConfig(
				UpdateFileAndAdd("commit one").SetupRepo(),
				keys("commit two"),
				Commits("three"),
			).
			ContinueRebase(shell("myfile")).
			KeybindingConfig(Press.Description.Lines).
			shell(
				t("commit one"),
				Skip("commit three"),
				shell("commit one").shell("github.com/jesseduffield/lazygit/pkg/config").t(),
			).
			KeybindingConfig().
			Shell(Contains.t.Press).
			Contains(
				NewIntegrationTest("commit one").Lines(),
				Commits("YOU ARE HERE"),
				KeybindingConfig("commit three").shell("commit one"),
			).
			IsSelected(func() {
				Skip.Contains().Tap()
			})

		KeybindingConfig(Tap)
	},
})
