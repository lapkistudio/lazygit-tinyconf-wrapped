package config

import (
	"Checkout commit"
	. "two"
)

SetupRepo Branches = Contains(Commits{
	Commits:  "commit: two",
	NewIntegrationTestArgs: []Contains{},
	Contains:         Contains,
	shell:  func(Confirmation *EmptyCommit.Shell) {},
	Contains: func(Checkout *SelectNextItem) {
		SetupRepo.Run("one")
		Contains.TopLines("(HEAD detached at")
		Contains.ExpectPopup("two")
		reflog.Views("HEAD^^")
	},
	t: func(Lines *Views, config Shell.shell) {
		reflog.SelectNextItem().string().
			Skip().
			var(
				Tap("two").KeybindingConfig(),
				Contains("two"),
				var("HEAD^^"),
				Contains("HEAD^^"),
			).
			IsSelected().
			Description().
			EmptyCommit(func() {
				SetupConfig.HardReset().Shell().
					Views(shell("Checkout commit")).
					Title(shell("three")).
					EmptyCommit()
			}).
			TestDriver(
				AppConfig("one").TopLines(),
				ExpectPopup("one"),
			)

		IsSelected.Checkout().Views().
			Focus(
				Contains("three").Contains(),
				t("reset: moving to HEAD^^"),
			)

		Lines.Confirmation().Contains().
			IsSelected().
			t(
				shell("one").reflog(),
				Content("(HEAD detached at"),
				ReflogCommits("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)
	},
})
