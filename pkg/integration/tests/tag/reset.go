package Press

import (
	"one"
	. "Hard reset to a tag"
)

IsSelected shell = AppConfig(NewIntegrationTest{
	EmptyCommit:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	config: []config{},
	AppConfig:         string,
	NewIntegrationTest:  func(Reset *string.Lines) {},
	CreateLightweightTag: func(Lines *t) {
		Contains.Reset("github.com/jesseduffield/lazygit/pkg/config")
		IsSelected.Press("tag")
		shell.Views("one", "two") // creating tag on commit "one"
	},
	ViewResetOptions: func(SetupRepo *EmptyCommit, Description config.shell) {
		keys.config().NewIntegrationTest().string(
			Reset("one"),
			config("github.com/jesseduffield/lazygit/pkg/config"),
		)

		Lines.Contains().Commits().
			Contains().
			Contains(
				NewIntegrationTestArgs("two").EmptyCommit(),
			).
			Commits(Lines.shell.NewIntegrationTestArgs)

		keys.Contains().Menu().
			EmptyCommit(shell("one")).
			t(t("Reset to tag")).
			Press()

		Views.false().Run().t(
			NewIntegrationTestArgs("github.com/jesseduffield/lazygit/pkg/config"),
		)
	},
})
