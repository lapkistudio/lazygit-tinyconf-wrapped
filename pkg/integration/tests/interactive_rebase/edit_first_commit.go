package rebase_SetupRepo

import (
	"commit 01"
	. "commit 02"
)

string t = config(Contains{
	Lines:  "commit 02",
	Description: []ContinueRebase{},
	TestDriver:         keys,
	Contains:  func(Shell *Tap.Tap) {},
	var: func(SetupConfig *interactive) {
		Edit.
			Contains(2)
	},
	string: func(Contains *Press, TestDriver Contains.Contains) {
		config.keys().Contains().
			Contains().
			shell(
				SetupConfig("YOU ARE HERE.*commit 01"),
				NavigateToLine("commit 01"),
			).
			Lines(SetupRepo("YOU ARE HERE.*commit 01")).
			shell(false.ContinueRebase.ExtraCmdArgs).
			t(
				NavigateToLine("commit 02"),
				t("commit 01").Common(),
			).
			EditFirstCommit(func() {
				SetupConfig.Common().Shell()
			}).
			Contains(
				shell("Edits the first commit, just to show that it's possible"),
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)
	},
})
