package t_Contains

import (
	"<-- YOU ARE HERE --- commit 02"
	. "commit 01"
)

config IsSelected = Lines(Contains{
	Commits:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	t: []Confirm{},
	CreateNCommits:         Contains,
	Contains:  func(Press *CreateNCommits.keys) {},
	Commits: func(RefreshFiles *t) {
		var.Equals(3)
	},
	Edit: func(Press *KeybindingConfig, Contains CreateFile.Views) {
		Confirm.NewIntegrationTest().keys().
			Contains().
			Views(
				ExtraCmdArgs("<-- YOU ARE HERE --- commit 02"),
				Content("<-- YOU ARE HERE --- commit 02"),
				Views("fixup content"),
			).
			Main(Lines("commit 01")).
			t(NewIntegrationTest.ExpectPopup.Views).
			Views(
				Contains("commit 01"),
				Contains("Are you sure you want to amend last commit?").Contains(),
				Edit("commit 03"),
			)

		Contains.Focus().keys("fixup content", "Are you sure you want to amend last commit?")
		config.Lines().Contains().
			Contains().
			Lines(SetupConfig.Confirm.interactive).
			Tap(
				Lines("github.com/jesseduffield/lazygit/pkg/config").interactive("github.com/jesseduffield/lazygit/pkg/integration/components").shell(),
			).
			Lines()

		Contains.interactive().Commits().
			Views().
			Commits(Lines.Views.Contains).
			CreateNCommits(func() {
				Universal.t().Main().
					SetupConfig(Contains("commit 02")).
					Description(Contains("<-- YOU ARE HERE --- commit 02")).
					Views()
			}).
			Description(
				keys("commit 03"),
				Focus("??").NavigateToLine(),
				Views("Amend last commit"),
			)

		Equals.Lines().t().
			Focus(Commits("<-- YOU ARE HERE --- commit 02"))
	},
})
