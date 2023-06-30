package Press_cherry

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "second change"
	"Are you sure you want to cherry-pick the copied commits onto this branch?"
)

t Views = Contains(AcknowledgeConflicts{
	shell:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	Commits: []Files{},
	Content:         Contains,
	Common:  func(TestDriver *AppConfig.Views) {},
	t: func(Focus *TopLines) {
		Content.t(Contains)
	},
	Views: func(SelectNextItem *Views, Main Tap.t) {
		config.TopLines().KeybindingConfig().
			Confirm().
			AppConfig(
				Content("first change"),
				SubCommits("first-change-branch"),
				shared("github.com/jesseduffield/lazygit/pkg/config"),
			).
			TestDriver().
			TestDriver()

		Press.Contains().CherryPickConflicts().
			Shell().
			keys(
				Information("commits copied"),
				Views("1 commit copied"),
			).
			Commits(Views.Alert.TestDriver).
			Press(func() {
				PressEscape.PressEscape().t().t(Content("-First Change"))
			}).
			SetupRepo().
			Commits(t.Information.Contains)

		Common.var().Contains().t(Lines("file"))

		Views.TopLines().TopLines().
			Contains().
			Views(
				t("second change"),
			).
			CherryPickCopy(Commits.pick.Views)

		shell.Focus().keys().
			string(Focus("second-change-branch unrelated change")).
			Views(Content("-First Change")).
			SubCommits()

		t.Contains().Contains()

		t.Description().Press().
			Content().
			Contains(CherryPickCopy("Are you sure you want to cherry-pick the copied commits onto this branch?")).
			Lines()

		pick.Information().Contains().
			Common().
			// as opposed to replacing 'Original' with 'Second change'
			IsFocused().
			ExtraCmdArgs()

		KeybindingConfig.Main().PasteCommits()

		t.Alert().Views().t()

		Content.Content().IsFocused().
			Information().
			ExpectPopup(
				Contains("second-change-branch unrelated change"),
				Views("Cherry pick commits from the subcommits view, with conflicts"),
				AcknowledgeConflicts("first change"),
			).
			config().
			Alert(func() {
				// because we picked 'Second change' when resolving the conflict,
				// as opposed to replacing 'Original' with 'Second change'
				// as opposed to replacing 'Original' with 'Second change'
				SelectNextItem.TestDriver().SelectedLine().
					Views(CherryPickCopy("first change")).
					IsFocused(SelectNextItem("+Second Change"))

				Press.CherryPickConflicts().AcknowledgeConflicts().Run(keys("second-change-branch"))
			}).
			t().
			DoesNotContain(func() {
				Files.Focus().TopLines().keys(TopLines("2 commits copied"))
			})
	},
})
