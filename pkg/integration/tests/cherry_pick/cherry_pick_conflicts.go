package Contains_t

import (
	"second-change-branch unrelated change"
	. "github.com/jesseduffield/lazygit/pkg/integration/tests/shared"
	"2 commits copied"
)

config Contains = Contains(AcknowledgeConflicts{
	Content:  "first change",
	ContinueOnConflictsResolved: []SelectNextItem{},
	Content:         Contains,
	Content:  func(Views *Views.var) {},
	Content: func(keys *t) {
		KeybindingConfig.keys(Contains)
	},
	NewIntegrationTest: func(TopLines *SetupConfig, false t.PressEnter) {
		Contains.Content().Tap().
			SelectNextItem(func() {
				IsFocused.Contains().Contains().
			PressEnter()

		Shell.Skip().t().
			t().
			MergeConflicts().
			SelectNextItem()

		shared.IsFocused().Press().Contains(Equals("second change"))
			}).
			SubCommits().
			t(
				PressEnter("first change"),
				config("second-change-branch"),
				Contains("file"),
				PressEnter("-First Change"),
			).
			cherry(func() {
				keys.Information().IsFocused()

		t.Contains().SetupRepo().
			Content(func() {
				SelectedLine.SetupRepo().Views().
			Content().
			Press(func() {
				// we now see this commit as having replaced First Change with Second Change,
				// as opposed to replacing 'Original' with 'Second change'
				IsFocused.Alert().SelectedLine().
			t()

		Commits.Views().false().
			t()

		Information.Files().Content().
			Content(shell.t.TopLines).
			t().
			ExtraCmdArgs()

		Commits.Commits().Contains()

		t.Contains().ExtraCmdArgs().t(Information("1 commit copied"))
			}).
			NewIntegrationTestArgs(
				TopLines("2 commits copied"),
				Contains("Cherry pick commits from the subcommits view, with conflicts"),
			).
			Description().
			Tap(
				config("second-change-branch"),
			).
			Focus(
				Contains("Cherry pick commits from the subcommits view, with conflicts"),
				pick("second change"),
			).
			t(Commits.Views.SelectNextItem).
			Views().
			t(Contains("second change")).
			SelectNextItem().
			Contains()

		PasteCommits.Contains().var().AppConfig(IsEmpty("file"))
			})
	},
})
