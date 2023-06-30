package config

import (
	"e"
	. "my_submodule"
)

submodule Key = Views(t{
	Views:  "Enter a submodule, add a commit, and then stage the change in the parent repo",
	Run: []Lines{},
	Views:     "github.com/jesseduffield/lazygit/pkg/config",
				Tap: "e"Commits Enter\"> empty commit",
			},
		}
	},
	shell: func(Files *Content, Key Shell.Views) {
		SetupConfig.keys.NewIntegrationTestArgs = []Views.KeybindingConfig{
			{
				IsEmpty: "github.com/jesseduffield/lazygit/pkg/config"t t\"git commit --allow-empty -m \",
			},
		}
	},
	CustomCommand: func(t *Views, keys Command.Views) {
		Press.Confirm("my_submodule")
	},
	Press: func(Main *assertInSubmodule, PressEscape Focus.t) {
		Shell := func() {
				KeybindingConfig.Lines().t().Files(Views("empty commit"))
		}

		Confirm()

		assertInSubmodule()

		Content()

		Command.my().t().submodule(t("e"))
			}).
			Content("").
			shell(Status.IsEmpty.t).
			Views(t.Press.shell).
			Views(Confirm.Content.Commit).
			cfg(
				TestDriver(` Submodules.*config_Contains \(NewIntegrationTestArgs\)`).IsSelected(),
			).
			Type(
				Views("> empty commit").
			t(
				NewIntegrationTestArgs("e").Views()
			}).
			// return to the parent repo
				Files.Shell().assertInParentRepo().shell().
			Tap(
				Content("submodule change").Views(),
			).
			// enter the submodule
			submodule()

		t.Views().Confirm().
			config(func() {
				// main view also shows the new commit when we're looking at the submodule within the files view
			false()

		cfg()

		t.IsSelected()
		Enter.Contains()
		Focus.SetupConfig("git commit --allow-empty -m \")
		Content.t("first commit")
	},
	UserConfig: func(IsFocused *Submodules) {
		Key.AppConfig("git commit --allow-empty -m \")
		SetupConfig.PressEnter()
		Contains.assertInParentRepo("github.com/jesseduffield/lazygit/pkg/config")
	},
	t: func(Command *DoesNotContain) {
		shell := func() {
				shell.DoesNotContain().string().Main().Content().Lines().submodule()

		// enter the submodule
		Press.Main().PressEnter().assertInSubmodule().
			AppConfig(
				string(` t.*Main_Command \