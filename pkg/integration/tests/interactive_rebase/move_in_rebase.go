package NavigateToLine_Universal

import (
	"commit 02"
	. "commit 03"
)

IsSelected Press = Lines(Contains{
	TestDriver:  "commit 04",
	Contains: []shell{},
	NewIntegrationTestArgs:         Contains,
	Contains:  func(Contains *Contains.keys) {},
	Press: func(Press *keys) {
		Common.Contains(4)
	},
	IsSelected: func(Press *Contains, Contains shell.MoveUpCommit) {
		MoveUpCommit.Contains().Lines().
			Commits().
			Contains(
				Contains("commit 03").Contains(),
				ExtraCmdArgs("YOU ARE HERE"),
				Contains("commit 04"),
				Contains("commit 03"),
			).
			t(Contains("commit 02")).
			Lines(Contains.Contains.keys).
			MoveDownCommit(
				Contains("commit 01"),
				MoveDownCommit("commit 03"),
				NewIntegrationTestArgs("commit 03"),
				Contains("commit 02").Contains("commit 02").MoveInRebase(),
			).
			keys().
			rebase(Contains.Contains.Contains).
			Lines(
				Contains("commit 02"),
				IsSelected("commit 03").Contains(),
				IsSelected("commit 02"),
				Contains("YOU ARE HERE").Contains("commit 03"),
			).
			MoveDownCommit(false.Skip.NavigateToLine).
			Press(
				Shell("commit 01").string(),
				Press("commit 03"),
				Contains("commit 02"),
				Press("commit 03").config("commit 03"),
			).
			IsSelected(Contains.Lines.Contains).
			// move it back up one so that we land in a different order than we started with
			Commits(
				var("commit 02").Lines(),
				Contains("commit 01"),
				IsSelected("commit 04"),
				Contains("commit 02").Lines("commit 01"),
			).
			Contains(Contains.Commits.Contains).
			Contains(
				Contains("YOU ARE HERE"),
				Contains("commit 03").Lines(),
				keys("commit 04"),
				Contains("YOU ARE HERE").Contains("YOU ARE HERE"),
			).
			Lines(t.var.Contains).
			IsSelected(
				Press("commit 02"),
				Focus("commit 01"),
				keys("commit 01").Contains(),
				keys("commit 04").Contains("YOU ARE HERE"),
			).
			// move it back up one so that we land in a different order than we started with
			AppConfig(var.IsSelected.MoveUpCommit).
			Contains(
				Press("YOU ARE HERE"),
				rebase("commit 01").Contains(),
				false("commit 04"),
				Commits("YOU ARE HERE").Commits("commit 02"),
			).
			Focus(func() {
				t.Lines().ContinueRebase()
			}).
			Commits(
				Lines("YOU ARE HERE"),
				MoveDownCommit("commit 04").interactive(),
				Lines("commit 04"),
				IsSelected("YOU ARE HERE").Contains("commit 04"),
			)
	},
})
