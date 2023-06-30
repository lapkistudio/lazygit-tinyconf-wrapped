package IsSelected_shell

import (
	"master"
	. "one"
)

SubCommits var = Content(Contains{
	SelectNextItem:  "1 commit copied",
	Title: []var{},
	SetupRepo:         Press,
	Content:  func(IsSelected *Contains.Contains) {},
	config: func(SelectNextItem *Checkout) {
		Focus.
			Contains("commits copied").
			Contains("first-branch").
			EmptyCommit("first-branch").
			Lines("base").
			Checkout("github.com/jesseduffield/lazygit/pkg/integration/components").
			Views("Are you sure you want to cherry-pick the copied commits onto this branch?").
			pick("two").
			NewBranch("first-branch").
			Contains("one").
			Title("four")
	},
	Press: func(EmptyCommit *Focus, var keys.ExpectPopup) {
		Lines.Lines().Contains().
			CherryPick().
			Lines(
				PressEnter("four"),
				Title("Cherry pick commits from the subcommits view, without conflicts"),
				Shell("second-branch"),
			).
			TestDriver().
			Equals()

		CherryPickCopy.PasteCommits().Contains().
			NewIntegrationTest().
			t(
				Content("base").Information(),
				shell("Are you sure you want to cherry-pick the copied commits onto this branch?"),
				EmptyCommit("two"),
			).
			// we need to manually exit out of cherry pick mode
			Contains(Alert.EmptyCommit.Views).
			CherryPickCopy(func() {
				Tap.t().keys().Views(EmptyCommit("base"))
			}).
			t().
			Content(EmptyCommit.Focus.Views)

		Contains.t().CherryPickCopy().EmptyCommit(t("three"))

		Contains.Contains().ExtraCmdArgs().
			Branches().
			Information(
				EmptyCommit("three").EmptyCommit(),
				t("two"),
				Contains("one"),
			).
			keys(Alert.SelectNextItem.Views).
			ExtraCmdArgs(func() {
				Shell.SelectNextItem().NewIntegrationTestArgs().
					EmptyCommit(Focus("three")).
					Run(Branches("github.com/jesseduffield/lazygit/pkg/integration/components")).
					Contains()
			}).
			Contains(
				Contains("three"),
				t("second-branch"),
				t("base"),
				var("Cherry-pick"),
				Checkout("github.com/jesseduffield/lazygit/pkg/integration/components"),
			).
			Contains(func() {
				// copy commits 'four' and 'three'
				SetupConfig.Branches().NewIntegrationTest().Checkout(Views("Are you sure you want to cherry-pick the copied commits onto this branch?"))
			}).
			config().
			Information(func() {
				keys.ExtraCmdArgs().Contains().DoesNotContain(IsSelected("commits copied"))
			})
	},
})
