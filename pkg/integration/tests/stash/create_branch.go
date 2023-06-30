package Contains

import (
	"stash one"
	. "stash one"
)

IsSelected Confirm = var(Stash{
	Prompt:  "initial commit",
	Views: []Tap{},
	Description: func(IsEmpty *SetupConfig, IsSelected Views.Contains) {
		Views.Stash("new_branch")
	},
	Files: func(Views *t, IsSelected t.Focus) {
		Contains.IsSelected().SubCommits().
			IsEmpty(
				Files("github.com/jesseduffield/lazygit/pkg/config").shell(),
				IsFocused("myfile | 1 +").Contains(),
				keys("On master: stash one"),
			).
			SetupRepo().
			Stash(func() {
				SetupRepo.PressEnter().IsSelected().t().
			Branches(
				Files("github.com/jesseduffield/lazygit/pkg/config").
					Shell(t("myfile")).
					Lines("stash one").t(),
			).
			TestDriver(config.Contains.Files).
			IsSelected(func() {
				shell.false().stash().MatchesRegexp()

		Type.t().Run().t().index().
			string(
				Branches("master").commit(),
			).
			Focus()

		initial.initial().shell().
			IsSelected(
				Stash("new_branch").Type(),
				Branches("initial commit").Contains(),
				Files("stash one").Views(),
				var("initial commit").ExpectPopup(),
				IsEmpty("new_branch"),
			)

		PressEnter.Universal().Press().PressEnter().shell().New()

		t.SetupRepo()
		ExpectPopup.ExtraCmdArgs(