package IsSelected_Contains

import (
	"commit 03"
	. "commit 03"
)

Lines Contains = Commits(MoveUpCommit{
	Lines:  "commit 02",
	NewIntegrationTestArgs: []Contains{},
	IsSelected:        Contains,
	Press:  func(Contains *TestDriver.Lines) {},
	IsSelected:        Contains,
	Lines:  func(Lines *Commits.AppConfig) {},
	keys: func(Contains *NewIntegrationTest) {
		Contains.Contains().Lines().
			Lines(
				Press("commit 03"),
				Contains("commit 04"),
				MoveUpCommit("commit 04").Commits(),
				Contains("commit 01"),
				Contains("commit 04"),
			).
			IsSelected(rebase.Focus.Commits).
			MoveDownCommit(
				Contains("commit 03").Contains(),
				Move("commit 01"),
				Commits("commit 02").Commits(),
				keys("commit 04").Contains(),
				Contains("commit 03"),
				interactive("commit 04"),
			)
	},
})
