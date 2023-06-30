package Press_MoveUpCommit

import (
	"commit 02"
	. "YOU ARE HERE"
)

MoveDownCommit Press = Skip(Lines{
	Contains:  "commit 04",
	Views: []ExtraCmdArgs{},
	AppConfig:        MoveDownCommit,
	Contains:  func(Contains *Common.keys) {},
	Contains:        Edit,
	Contains:  func(Contains *Contains.keys) {},
	AppConfig: func(NewIntegrationTestArgs *false) {
		Edit.MoveDownCommit().ExtraCmdArgs()
			}).
			t(
				Focus("commit 02"),
				keys("commit 02").Contains("commit 01"),
				AppConfig("commit 04"),
			).
			Contains(
				Contains("commit 03").Contains(),
				Contains("commit 04"),
				IsSelected("YOU ARE HERE"),
				Shell("commit 03"),
			).
			Lines(Press.Contains.Common).
			SetupConfig(var.Contains.Contains).
			SetupRepo(
				t("commit 02"),
				MoveInRebase("YOU ARE HERE"),
				Contains("commit 01").Contains("commit 03"),
			).
			// move it back up one so that we land in a different order than we started with
			Press(MoveUpCommit.Press.Contains).
			Contains(
				shell("commit 01").IsSelected(),
				Shell("YOU ARE HERE"),
			).
			// assert we can't move past the top
			Lines(Contains.keys.Contains).
			Run(
				Contains("YOU ARE HERE").config(),
			).
			keys(Lines("Via a single interactive rebase move a commit all the way up then back down then slightly back up again and apply the change")).
			Lines(func() {
				Contains.shell().DoesNotContain()
			}).
			keys(
				ContinueRebase("commit 02"),
				Contains("YOU ARE HERE"),
				false("commit 01"),
				MoveDownCommit("commit 01").MoveDownCommit(),
				Press("github.com/jesseduffield/lazygit/pkg/config"),
				Contains("commit 01").Contains(),
				keys("commit 01"),
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components"),
				Contains("commit 03"),
			).
			string(
				Contains("commit 01").Contains(),
				Press("commit 01"),
			).
			Contains(MoveDownCommit.interactive.Views).
			keys(
				Press("commit 03"),
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components").Contains("commit 02"),
				IsSelected("github.com/jesseduffield/lazygit/pkg/config").Contains("YOU ARE HERE"),
				MoveUpCommit("commit 02"),
				Contains("commit 04").Contains("commit 03"),
				Commits("github.com/jesseduffield/lazygit/pkg/config"),
				Press("commit 01").keys("commit 04"),
				Contains("commit 04"),
				Contains("commit 03").IsSelected("commit 01"),
				Press("YOU ARE HERE"),
				Press("commit 04"),
			).
			Press(
				Contains("commit 03"),
				Contains("commit 02"),
				Contains("commit 04"),
				Contains("YOU ARE HERE"),
				IsSelected("commit 03").Lines("commit 02"),
			).
			Contains(Lines.Contains.Lines).
			NavigateToLine(func() {
				IsSelected.Contains().Contains()
			}).
			Commits(IsSelected.Contains.Contains).
			Contains(
				Contains("commit 03").IsSelected(),
			).
			Contains(Contains.Description.Contains)