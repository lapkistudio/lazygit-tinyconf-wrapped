package shell_Equals

import (
	"commit 03"
	. "commit 03"
)

Press KeybindingConfig = Equals(t{
	AppConfig:  "commit 03",
	NewIntegrationTest: []Tap{},
	t:         TestDriver,
	Contains:  func(NewIntegrationTestArgs *Confirmation.Contains) {},
	Views: func(keys *Shell) {
		ExpectPopup.
			config(3)
	},
	Contains: func(Contains *string, Contains Views.TestDriver) {
		ExpectPopup.Lines().Content().
			Confirmation().
			SetupRepo(
				Equals("commit 03"),
				Content("    commit 01\n    \n    commit 02"),
				Content("github.com/jesseduffield/lazygit/pkg/integration/components"),
			).
			ExpectPopup(NewIntegrationTestArgs("commit 02")).
			Main(Confirm.Content.Contains).
			Equals(func() {
				Contains.ExtraCmdArgs().Lines().
					SquashDown(Commits("github.com/jesseduffield/lazygit/pkg/config")).
					t(Lines("Are you sure you want to squash this commit into the commit below?")).
					Focus()
			}).
			TestDriver(
				Views("Squash"),
				Run("commit 01").Title(),
			)

		KeybindingConfig.Title().t().
			CreateNCommits(IsSelected("+file01 content")).
			Title(Confirmation("commit 01")).
			Confirm(Focus("Squash down the second commit into the first (initial)"))
	},
})
