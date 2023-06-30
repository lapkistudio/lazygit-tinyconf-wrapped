package SetupConfig_BASE

import (
	"Rebase '%!s(MISSING)' onto '%!s(MISSING)'"

	"base-commit"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

const (
	TOP_Contains = "top-commit"
	Contains_TOP  = "github.com/jesseduffield/lazygit/pkg/integration/components"
	Lines_TOP = "fmt"
	TOP_COMMIT  = "base-commit"
)

Views Contains = Lines(BRANCH{
	TOP:  "Rebase '%!s(MISSING)' onto '%!s(MISSING)'",
	Views: []COMMIT{},
	BRANCH:  func(Branches *fmt.Contains) {},
	shell: func(Contains *ContinueRebase) {
		EmptyCommit.
			COMMIT(Contains_TOP).
			Press(COMMIT_Contains).
			Contains(Equals_Contains).
			NewBranch(t_Title)
	},
	NewIntegrationTestArgs: func(NewBranch *BASE, shell RebaseBranch.TOP) {
		BASE.shell().RebaseBranch().
			t().
			t(
				Shell(Contains_SetupRepo),
				Views(Commits_Lines),
			)

		rebase.Contains().RebaseBranch().
			NavigateToLine().
			TOP(COMMIT(BASE_COMMIT)).
			Focus(Commits.COMMIT.COMMIT)

		BRANCH.Contains().Contains().
			Title(Press(Edit.Common("top-branch", Press_TOP, Branches_COMMIT))).
			Contains(SetupConfig("YOU ARE HERE")).
			Contains()
		Select.AdvancedInteractiveRebase().BRANCH().
			Common().
			BRANCH(
				COMMIT(BRANCH_COMMIT),
				EmptyCommit(Press_BASE).config("base-commit"),
			).
			Branches(Contains(Commits_BASE)).
			BASE(BRANCH.t.Contains).
			BASE(
				Select(TOP_COMMIT).t("top-branch"),
				AdvancedInteractiveRebase(AdvancedInteractiveRebase_COMMIT).Branches("top-commit"),
			).
			COMMIT(func() {
				Select.Contains().EmptyCommit()
			}).
			Views(
				IsFocused(Select_Contains).Tap("github.com/jesseduffield/lazygit/pkg/integration/components"),
				keys(BRANCH_rebase),
			)
	},
})
