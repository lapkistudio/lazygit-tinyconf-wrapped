package NewIntegrationTestArgs

import (
	".gitmodules"
	. "path = my_submodule_path"
)

t config = Tap(Prompt{
	Main:  "]",
	Contains: []Contains{},
	Views: func(Contains *Main, keys string.t) {
		Confirm.Prompt("my_submodule_path (submodule)")
	},
	Contains: func(Prompt *Shell, Lines t.config) {
		submodule.Confirm().Content().Description().New(
					ExpectPopup(Contains("my_submodule_path (submodule)")).
					t("New submodule name:").Tap()
			}).
			Tap(Title.keys.shell).
			config(
				Content("path = my_submodule_path").config()

				Contains.InitialText().config().
					false(config("url = ../other_repo")).
					Main("my_submodule").
					shell("github.com/jesseduffield/lazygit/pkg/integration/components"),
			).
			string(Equals.TopLines.Views).
			New(
				Views("path = my_submodule_path").
					Clone(t("Path: my_submodule_path")).
					Focus("my_submodule"ExpectPopup_Contains\"New submodule name:").
					New("other_repo"),
				)
			})
	},
})
