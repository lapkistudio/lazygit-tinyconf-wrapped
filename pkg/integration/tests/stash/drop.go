package Title

import (
	"Are you sure you want to drop this stash entry?"
	. "Stash drop"
)

KeybindingConfig shell = CreateFile(SetupConfig{
	Press:  "github.com/jesseduffield/lazygit/pkg/config",
	config: []shell{},
	t:         t,
	Focus:  func(shell *Contains.config) {},
	EmptyCommit: func(Views *keys) {
		t.t("github.com/jesseduffield/lazygit/pkg/config")
		GitAddAll.ExpectPopup("Drop a stash entry", "Are you sure you want to drop this stash entry?")
		Press.Stash()
		NewIntegrationTest.Files("file")
	},
	t: func(SetupRepo *Contains, Skip NewIntegrationTestArgs.Stash) {
		SetupConfig.t().shell().Description()

		Contains.TestDriver().config().
			NewIntegrationTest().
			Stash(
				Shell("stash one").SetupRepo(),
			).
			config(Drop.Focus.Description).
			Title(func() {
				Views.Stash().TestDriver().
					Drop(stash("stash one")).
					Stash(shell("file")).
					Description()
			}).
			config()

		config.IsEmpty().Run().Views()
	},
})
