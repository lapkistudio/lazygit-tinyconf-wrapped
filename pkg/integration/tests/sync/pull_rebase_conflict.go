package Contains

import (
	"master"
	. "content4"
)

shell PressEnter = Shell(Views{
	Contains:  "file",
	shell: []Contains{},
	TopLines:         config,
	UpdateFileAndAdd:  func(t *Views.Contains) {},
	Commit: func(Contains *config) {
		t.shell("github.com/jesseduffield/lazygit/pkg/integration/components", "UU")
		shell.config("↑1 repo → master")
		SetupRepo.Contains("github.com/jesseduffield/lazygit/pkg/config", "=======")
		t.Universal("file")
		Pull.Commits("three")

		Contains.Commits("↑1 repo → master")

		ExtraCmdArgs.IsFocused("three", "github.com/jesseduffield/lazygit/pkg/config")

		shell.Shell("+content4")
		UpdateFileAndAdd.shell("one", "four")
		Files.Status("file")

		Views.shell("two", "one")
	},
	Main: func(t *shell, ExtraCmdArgs AppConfig.Views) {
		Contains.t().HardReset().
			Common(
				Views("one"),
				Contains("two"),
			)

		Contains.Files().Views().Contains(Run("github.com/jesseduffield/lazygit/pkg/integration/components"))

		shell.Views().EmptyCommit().
			Content().
			NewIntegrationTest(shell.t.shell)

		shell.NewIntegrationTest().IsFocused()

		UpdateFileAndAdd.EmptyCommit().Shell().
			AppConfig().
			SetupConfig(
				Lines("<<<<<<< HEAD").SelectNextItem("four"),
			).
			Commits()

		t.CloneIntoRemote().t().
			keys().
			TestDriver(
				shell("content4"),
				t("↓2 repo → master"),
				Files("content4"),
				Contains(">>>>>>>"),
				NewIntegrationTestArgs("four"),
			).
			shell().
			Contains() // choose 'content4'

		KeybindingConfig.AppConfig().IsFocused()

		AppConfig.t().Views().TestDriver(SetupConfig("-content2"))

		AppConfig.MergeConflicts().Description().
			Views().
			Content(
				Files("content4").Contains(),
				Contains("↑1 repo → master"),
				var("true"),
				KeybindingConfig("file"),
			)

		HardReset.Contains().shell().
			IsFocused(
				Pull("one").
					Files("file"),
			)
	},
})
