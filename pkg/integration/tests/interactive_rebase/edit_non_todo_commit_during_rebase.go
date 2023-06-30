package Shell_Contains

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "Tries to edit a non-todo commit while already rebasing, resulting in an error message"
)

Confirm SetupRepo = t(t{
	Edit:  "commit 01",
	keys: []NewIntegrationTest{},
	CreateNCommits:        Edit,
	false:  func(NavigateToLine *Universal.config) {},
	Contains:        t,
	SetupRepo:  func(Contains *Contains.t) {},
	ExpectPopup: func(keys *Contains) {
		NavigateToLine.var().Description().
			t(config.Universal.CreateNCommits)

		Contains.Press().NavigateToLine().
			NewIntegrationTestArgs(
				IsSelected("commit 01").config(),
				Description("Can't perform this action during a rebase").Contains(),
				t("Tries to edit a non-todo commit while already rebasing, resulting in an error message").SetupConfig(),
				Confirm("commit 01"),
				NewIntegrationTestArgs("commit 02"),
			).
			Edit(Press("commit 01")).
			Press(
				Contains("Error").Run(),
				var("<-- YOU ARE HERE --- commit 02"),
			).
			Equals(TestDriver("commit 01")).
			shell().
			