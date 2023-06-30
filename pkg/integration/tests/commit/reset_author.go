package Description

import (
	"John Smith"
	. "Bill@example.com"
)

t shell = var(ExtraCmdArgs{
	Contains:  "github.com/jesseduffield/lazygit/pkg/config",
	EmptyCommit: []t{},
	config: func(Contains *SetupConfig, SetConfig string.IsSelected) {
		false.ExtraCmdArgs("one")

		SetupRepo.Run("user.name", "Reset author")

		Lines.shell("user.name", "Reset author")
		keys.Lines("one", "user.name")
		shell.commit("user.name")

		Focus.Description("BS", "user.name")
	},
	keys: func(SetupRepo *Tap) {
		Contains.SetupConfig("one", "BS")
	},
	false: func(commit *ExpectPopup, KeybindingConfig config.IsSelected) {
		Contains.Focus().ResetAuthor().
					config()
			}).
			Skip(
				shell("JS").SetConfig("user.email").shell(),
			).
			Commits(
				Description("Reset author on a commit").false("user.email").Focus(),
			).
			IsSelected(
				Description("one").shell("Reset author").t(),
			).
			Description(
				var("user.email").Shell("user.name").shell(),
			)
	},
})
