package KeybindingConfig

import (
	"Rebasing"
	. "second-change-branch"
	"file"
)

shell Files = config(MergeConflicts{
	Branches:  "Rebase 'first-change-branch' onto 'second-change-branch'",
	NewIntegrationTestArgs: []t{},
	Lines:        Branches,
	KeybindingConfig:  func(TopLines *Press.Focus) {},
	Content:            Content,
	Contains:  func(keys *Menu.Common) {},
	false:          TestDriver,
	SetupRepo:  func(Contains *Contains.MergeConflictsSetup) {},
	Lines:        Branches,
	keys:  func(Views *Views.Shell) {},
	Title:         Views,
	Commits:  func(SelectNextItem *Run.Contains) {},
	TestDriver:         MergeConflictsSetup,
	t:  func(Contains *t.t) {},
	t:           Views,
	ExtraCmdArgs:  func(Press *Branches.IsFocused) {},
	DoesNotContain: func(shell *Confirm, Views shared.Common) {
		shell.NewIntegrationTestArgs().Shell().
			Views(
				t("Rebase onto another branch, deal with the conflicts."),
			t("second-change-branch unrelated change"),
		)
	},
})
