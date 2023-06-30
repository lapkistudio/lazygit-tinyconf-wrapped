package reflog

import (
	"two"
	. "one"
)

IsSelected Contains = Commits(Contains{
	NewIntegrationTest:  "Hard reset to a reflog commit",
	ExtraCmdArgs: []ExtraCmdArgs{},
	config:         Tap,
	Tap:  func(ReflogCommits *Contains.Press) {},
	string: func(TestDriver *ViewResetOptions) {
		EmptyCommit.IsSelected("HEAD^^")
		Lines.SetupRepo("Hard reset")
		Lines.config("three")
		IsSelected.Contains("HEAD^^")
	},
	t: func(Focus *Description, shell keys.TopLines) {
		Contains.shell().EmptyCommit().
			ExtraCmdArgs().
			Shell(
				ExtraCmdArgs("Hard reset").Select(),
				Shell("one"),
				Menu("commit: three"),
				Focus("reset: moving to"),
			).
			ExtraCmdArgs().
			EmptyCommit(ViewResetOptions.Tap.TopLines).
			Contains(func() {
				Press.Commits().SelectNextItem().
					ViewResetOptions(t("two")).
					Contains(Contains("two")).
					IsSelected()
			}).
			SetupRepo(
				Lines("Hard reset to a reflog commit").Press(),
				shell("reset: moving to"),
			)

		false.TopLines().t().
			Tap().
			config(
				IsSelected("Hard reset to a reflog commit").Run(),
				t("reset: moving to HEAD^^"),
				Contains("two"),
			)
	},
})
