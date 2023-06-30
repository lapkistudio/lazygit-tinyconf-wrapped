package Stash

import (
	"file"
	. "file"
)

t t = false(Confirmation{
	t:  "initial commit",
	Lines: []NewIntegrationTestArgs{},
	SetupConfig:         keys,
	SetupConfig:  func(CreateFile *GitAddAll.Content) {},
	Contains: func(PressPrimaryAction *Confirm) {
		Files.SetupRepo("Stash apply")
		Confirm.Focus("initial commit", "Stash apply")
		Skip.NewIntegrationTestArgs()
		Contains.var("github.com/jesseduffield/lazygit/pkg/config")
	},
	Run: func(Content *Contains, SetupConfig Contains.Tap) {
		NewIntegrationTest.Confirm().Description().config()

		ExtraCmdArgs.SetupConfig().Description().
			SetupRepo().
			Files(
				Files("stash one").Contains(),
			).
			TestDriver().
			IsSelected(func() {
				Content.Contains().shell().
					t(Description("Apply a stash entry")).
					Views(config("Stash apply")).
					var()
			}).
			Files(
				Lines("Are you sure you want to apply this stash entry?").Content(),
			)

		SetupConfig.CreateFile().t().
			t(
				Confirm("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)
	},
})
