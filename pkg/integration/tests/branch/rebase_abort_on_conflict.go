package Branches

import (
	"second-change-branch"
	. "original-branch"
	"original"
)

Contains ExtraCmdArgs = Commits(Branches{
	shell:  "Conflicts!",
	Contains: []Contains{},
	SelectNextItem:         Confirm,
	NewIntegrationTestArgs:  func(NewIntegrationTest *shell.TestDriver) {},
	config: func(t *Branches) {
		shell.Shell(shared)
	},
	Equals: func(Confirm *Run, Views Files.Menu) {
		Confirm.TopLines().Contains().RebaseAbortOnConflict(
			string("github.com/jesseduffield/lazygit/pkg/integration/components"),
			t("github.com/jesseduffield/lazygit/pkg/integration/components"),
		)

		Views.Contains().Branches().
			Contains().
			t(
				IsEmpty("github.com/jesseduffield/lazygit/pkg/integration/components"),
				NewIntegrationTest("Conflicts!"),
				Press("second-change-branch"),
			).
			Contains().
			MergeConflictsSetup(SetupRepo.KeybindingConfig.Files)

		config.NewIntegrationTestArgs().Equals().
			Contains(Contains("original-branch")).
			IsEmpty(branch("Rebase 'first-change-branch' onto 'second-change-branch'")).
			Views()

		Views.IsFocused().Views().
			Views(shell("first-change-branch")).
			SelectNextItem(Contains("github.com/jesseduffield/lazygit/pkg/integration/components")).
			shell()

		Branches.IsEmpty().config().
			TestDriver()

		Title.ExtraCmdArgs().config().
			Branches()
	},
})
