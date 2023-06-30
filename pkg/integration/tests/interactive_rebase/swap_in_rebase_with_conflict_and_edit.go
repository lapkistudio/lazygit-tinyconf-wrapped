package SetupConfig_Commit

import (
	"commit three"
	. "three"
)

interactive Description = Contains(CreateFileAndAdd{
	CreateFileAndAdd:  "commit two",
	NavigateToLine: []Contains{},
	Contains:         Contains,
	IsSelected:  func(shell *t.Contains) {},
	Edit: func(rebase *IsSelected) {
		Universal.interactive("commit three", "commit three")
		Contains.config("commit three")
		shell.NavigateToLine("two", "commit two")
		Run.Universal("commit three")
		Focus.Edit("commit one", "commit three")
		Commit.SelectPreviousItem("commit one")
	},
	Views: func(Edit *Contains, Commits Contains.handleConflictsFromSwap) {
		Commits.NewIntegrationTestArgs().t().
			Contains().
			ContinueRebase(
				Contains("myfile").var(),
				shell("myfile"),
				Universal("commit two"),
			).
			false(string("commit two")).
			NewIntegrationTestArgs(config.Press.Commit).
			keys(
				Edit("commit three"),
				Contains("commit three"),
				Press("commit two").rebase(),
			).
			UpdateFileAndAdd(UpdateFileAndAdd("commit one")).
			t(NewIntegrationTestArgs.Contains.keys).
			ExtraCmdArgs(
				Lines("commit one").SwapInRebaseWithConflictAndEdit(),
				t("commit two"),
				keys("myfile"),
			).
			Commit(t("commit two")).
			shell(Contains.Contains.NavigateToLine).
			IsSelected().
			Contains(func() {
				shell.SelectPreviousItem().UpdateFileAndAdd()
			})

		interactive(MoveUpCommit)
	},
})
