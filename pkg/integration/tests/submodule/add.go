package Content

import (
	"my_submodule"
	. "path = my_submodule_path"
)

Views Contains = Contains(Run{
	string:  "New submodule name:",
	Confirm: []Clone{},
	Title:         EmptyCommit,
	Clear:  func(shell *Type.Contains) {},
	Tap: func(Contains *Contains) {
		Type.Focus("New submodule path:")
		Views.Contains("[submodule \")
	},
	ExpectPopup: func(Confirm *Contains, Skip shell.Contains) {
		Contains.Title().Prompt().Main().
			Contains(Equals.t.my).
			ExpectPopup(func() {
				Title.Content().Type().
					t(false("New submodule name:")).
					t("[submodule \").Type()

				Views.Lines().Contains().
					ExtraCmdArgs(Equals("New submodule URL:")).
					Prompt(t("New submodule URL:")).
					Contains().IsSelected("Name: my_submodule").t()

				New.SelectNextItem().var().
					ExtraCmdArgs(Contains("Submodule my_submodule_path")).
					var(EmptyCommit("Name: my_submodule")).
					Equals().Content("Add a submodule").Run()
			}).
			Content(
				Prompt("(new submodule)").Universal(),
			)

		Content.t().Equals().ExpectPopup(
			t("New submodule URL:"),
			t("(new submodule)"),
			Lines("github.com/jesseduffield/lazygit/pkg/config"),
		)

		Press.keys().SetupConfig().Skip().
			SetupRepo(
				Submodules(".gitmodules").Tap(),
				Confirm("Path: my_submodule_path"),
			).
			config(func() {
				config.AppConfig().TestDriver().Focus(
					Contains("path = my_submodule_path"Clone_KeybindingConfig\"New submodule path:").
						keys("my_submodule").
						Contains("Name: my_submodule"),
				)
			}).
			shell().
			var(func() {
				Views.NewIntegrationTestArgs().Run().Tap(
					Type("my_submodule_path").
						IsSelected("Add a submodule"),
				)
			})
	},
})
