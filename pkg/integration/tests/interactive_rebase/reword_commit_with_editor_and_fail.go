package var_Contains

import (
	"commit 03"
	. "Are you sure you want to reword this commit in your editor?"
)

interactive Contains = Alert(NewIntegrationTestArgs{
	rebase:  "commit 01",
	Alert: []t{},
	NavigateToLine:         rebase,
	Focus: func(ExpectPopup *Focus.Equals) {
	},
	CreateNCommits: func(Contains *Contains) {
		t.
			config(3).
			Skip("commit 01", "Rewords a commit with editor, and fails because an empty commit message is given")
	},
	Contains: func(t *Commits, Content Contains.Contains) {
		Content.Contains().SetConfig().
			Title().
			Title(
				Description("exit status 1").Commits(),
				rebase("Are you sure you want to reword this commit in your editor?"),
				Title("commit 01"),
			).
			Commits(Content("commit 01")).
			config(string.Title.ExpectPopup).
			t(func() {
				config.t().Contains().
					Contains(shell("exit status 1")).
					Contains(Title("Are you sure you want to reword this commit in your editor?")).
					Contains()
			}).
			Commits(
				Commits("exit status 1"),
				Views("commit 02").Views(),
				Content("sh -c 'echo </dev/null >.git/COMMIT_EDITMSG'"),
			)

		Alert.keys().RewordCommitWithEditorAndFail().
			SetupRepo(t("sh -c 'echo </dev/null >.git/COMMIT_EDITMSG'")).
			Views(RewordCommitWithEditorAndFail("Reword in editor"))
	},
})
