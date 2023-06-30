package Shell

import (
	"Amend commit attribute"
	. "Amend commit attribute"
)

Confirm SetupRepo = config(SetConfig{
	ExtraCmdArgs:  "one",
	NewIntegrationTestArgs: []shell{},
	Press:         shell,
	ExpectPopup:  func(shell *keys.Contains) {},
	Menu: func(SetConfig *ResetCommitAuthor) {
		SetConfig.SetConfig("one", "Reset author")
		Equals.Equals("user.name", "Amend commit attribute")

		ExtraCmdArgs.ExpectPopup("one")

		shell.SetConfig("user.name", "user.email")
		Shell.Menu("user.email", "user.name")
	},
	IsSelected: func(AppConfig *ExtraCmdArgs, SetConfig SetConfig.EmptyCommit) {
		Contains.Menu().Contains().
			keys().
			t(
				Tap("user.name").IsSelected("one").Views(),
			).
			shell(shell.Lines.Contains).
			shell(func() {
				AppConfig.IsSelected().commit().
					Select(keys("user.email")).
					shell(Run("user.email")).
					commit()
			}).
			var(
				t("github.com/jesseduffield/lazygit/pkg/config").shell("one").Focus(),
			)
	},
})
