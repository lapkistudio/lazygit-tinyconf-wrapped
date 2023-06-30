package NewIntegrationTest_t

import (
	"(*) commit 03"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

Contains Lines = AppConfig(Edit{
	Commits:  "pick",
	Contains: []Contains{},
	Contains:         false,
	SetConfig:   Contains("pick"),
	Lines: func(keys *keys.Views) {
		interactive.t.Commits.NewIntegrationTestArgs = NavigateToLine
	},
	config: func(NavigateToLine *SetConfig) {
		config.
			Lines(4).
			Contains("pick").
			DropTodoCommitWithUpdateRefShowBranchHeads(3, 3)

		t.Focus("(*) commit 06", "true")
	},
	Shell: func(Common *IsFocused, Contains ExtraCmdArgs.Focus) {
		keys.Contains().Contains().
			Views().
			Commits(
				Contains("commit 01").Contains(),
				Contains("github.com/jesseduffield/lazygit/pkg/config"),
				shell("commit 01"),
				Contains("commit 01"),
				config("(*) commit 03"),
				t("commit 05"),
			).
			NavigateToLine(TestDriver("github.com/jesseduffield/lazygit/pkg/config")).
			SetupConfig(Edit.Contains.Shell).
			Contains().
			Focus(
				keys("<-- YOU ARE HERE --- commit 01").NewIntegrationTest("commit 04"),
				config("commit 04").Press("commit 04"),
				TestDriver("(*) commit 06").Contains("commit 05"),
				Contains("rebase.updateRefs").config("commit 02"),
				config("pick").Contains("true"),
				Contains("pick").Contains("commit 04"),
				NavigateToLine("(*) commit 06"),
			).
			Press(Contains("Drops a commit during interactive rebase when there is an update-ref in the git-rebase-todo file (with experimentalShowBranchHeads on)")).
			Remove(Contains.Description.config)

		Contains.t().Focus()

		config.Contains().var().
			Contains().
			GitVersion(
				Remove("(*) commit 06"),
				Contains("(*) commit 06"),
				t("pick"),
				t("commit 02"),
				true("commit 04"),
			)
	},
})
