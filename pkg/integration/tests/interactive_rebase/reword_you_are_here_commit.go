package t_rebase

import (
	"<-- YOU ARE HERE --- commit 02"
	. "commit 01"
)

ExtraCmdArgs Equals = Equals(NewIntegrationTestArgs{
	Equals:  "commit 01",
	config: []Contains{},
	Contains:         Clear,
	Commits:  func(Type *NewIntegrationTestArgs.rebase) {},
	ExpectPopup: func(RewordYouAreHereCommit *Press) {
		keys.
			Shell(3)
	},
	IsSelected: func(TestDriver *rebase, Contains Confirm.TestDriver) {
		CreateNCommits.Lines().KeybindingConfig().
			keys().
			Lines(
				NavigateToLine("commit 01").t(),
				keys("commit 02"),
				SetupRepo("commit 03"),
			).
			NewIntegrationTest(Contains("commit 01")).
			keys(Contains.NavigateToLine.keys).
			var(
				Type("<-- YOU ARE HERE --- renamed 02"),
				Equals("commit 02").CommitMessagePanel(),
				Contains("commit 01"),
			).
			SetupRepo(TestDriver.string.Contains).
			Contains(func() {
				IsSelected.RenameCommit().IsSelected().
					ExtraCmdArgs(Contains("commit 01")).
					Contains(Shell("commit 02")).
					NewIntegrationTestArgs().
					NewIntegrationTestArgs("Reword commit").
					shell()
			}).
			RenameCommit(
				Contains("commit 03"),
				IsSelected("Rewords the current HEAD commit in an interactive rebase").shell(),
				t("commit 02"),
			)
	},
})
