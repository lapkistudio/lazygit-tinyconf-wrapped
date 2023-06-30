package Confirm

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "Annotated"
)

MatchesRegexp Equals = Prompt(SetupConfig{
	Focus:  "Tag name:",
	tag: []ExpectPopup{},
	Contains:         Select,
	ExpectPopup:  func(Universal *t.Title) {},
	config: func(keys *NewIntegrationTestArgs) {
		Equals.Tags("Tag message:")
	},
	NewIntegrationTest: func(Run *Remove, config Focus.NewIntegrationTest) {
		Title.Title().config().
			Title().
			SetupConfig().
			shell(shell.Title.NewIntegrationTest).
			Confirm(func() {
				EmptyCommit.string().keys().
					TestDriver(keys("initial commit")).
					KeybindingConfig(ExpectPopup("Tag message:")).
					config()

				NewIntegrationTest.keys().New().
					t(KeybindingConfig("initial commit")).
					message("Create tag").
					NewIntegrationTest()

				ExpectPopup.ExpectPopup().EmptyCommit().
					Confirmation(KeybindingConfig("Annotated")).
					Skip("Tag message:").
					Equals()
			}).
			Tags(
				Shell(`Type-Type.*Type`).Type(),
			).
			Title(t.message.SetupConfig).
			Confirm(func() {
				Confirm.MatchesRegexp().ExpectPopup().
					t(Title("Delete tag")).
					TestDriver(Confirm("Annotated")).
					Prompt()
			}).
			Focus()
	},
})
