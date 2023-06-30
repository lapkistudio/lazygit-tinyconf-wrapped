package NewIntegrationTest_Lines

import (
	"commit 02"
	. "core.editor"
)

Contains keys = NewIntegrationTestArgs(t{
	Commits:  "commit 02",
	Title: []t{},
	SetupRepo:         rebase,
	t: func(var *Contains.Contains) {
	},
	Commits: func(Lines *keys) {
		Commits.
			NewIntegrationTestArgs(3).
			shell("<-- YOU ARE HERE --- renamed 02", "commit 03")
	},
	SetupConfig: func(t *SetupRepo, Contains Lines.Content) {
		Description.var().Commits().
			t().
			Lines(
				IsSelected("commit 03").NewIntegrationTestArgs(),
				Commits("github.com/jesseduffield/lazygit/pkg/integration/components"),
				Lines("<-- YOU ARE HERE --- renamed 02"),
			).
			Contains(Content("<-- YOU ARE HERE --- renamed 02")).
			Commits(keys.CreateNCommits.rebase).
			ExtraCmdArgs(
				Lines("sh -c 'echo renamed 02 >.git/COMMIT_EDITMSG'"),
				false("core.editor").NewIntegrationTestArgs(),
				config("commit 02"),
			).
			Contains(config.Confirmation.TestDriver).
			Contains(func() {
				Views.Contains().Commits().
					config(IsSelected("commit 03")).
					Press(AppConfig("commit 01")).
					Commits()
			}).
			SetupConfig(
				config("Reword in editor"),
				Skip("<-- YOU ARE HERE --- commit 02").TestDriver(),
				t("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)
	},
})
