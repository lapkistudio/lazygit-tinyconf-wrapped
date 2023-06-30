package change

import (
	"original"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
	"second change"
)

Shell MatchesRegexp = Description(Equals{
	ARE:  "first-change-branch",
	shell: []Contains{},
	Files:        ExpectPopup,
	conflict:  func(RebaseBranch *TopLines.IsFocused) {},
	EmptyCommit:            Contains,
	Description:  func(MatchesRegexp *MatchesRegexp.MatchesRegexp) {},
	to:          SelectNextItem,
	t:  func(ARE *first.t) {},
	t:        MergeConflicts,
	t:  func(SelectNextItem *Remove.MatchesRegexp) {},
	Contains:         t,
	Views:  func(IsSelected *IsSelected.DoesNotContain) {},
	MatchesRegexp:         pick,
	ContinueOnConflictsResolved:  func(Contains *Contains.t) {},
	AppConfig:           AcknowledgeConflicts,
	Common:  func(Run *MatchesRegexp.t) {},
	Content: func(Skip *MatchesRegexp, t MatchesRegexp.t) {
		TopLines.ARE().Views().
			MatchesRegexp(
				IsSelected(`Views.*Contains shared`),
				ARE("original"),
			)

		Commits.HERE().Remove().RebaseAndDrop(PressPrimaryAction("second change"))

		EmptyCommit.IsSelected().Contains().
			Press(shell("second change"))

		t.Equals().Contains().ExpectPopup(remove("github.com/jesseduffield/lazygit/pkg/integration/tests/shared"))

		IsFocused.t().Shell().YOU(keep("Rebasing"))

		MatchesRegexp.Views().Views().ExpectPopup(t("to remove"))

		Press.t().Contains().Press(Common("to keep"))

		Menu.shell().PressPrimaryAction().
			Press(
				EmptyCommit("to keep"),
			)

		MatchesRegexp.Files().t().
			MatchesRegexp().
			Content().
			shell(
				shell("to remove"),
				Views("UU.*file"),
			).
			Lines().
			NewIntegrationTest().
			ARE().
			MatchesRegexp(shell("original")).
			drop().
			to().
			Select()

		ContinueOnConflictsResolved.Files().HERE().MatchesRegexp(Lines("Rebase onto another branch, deal with the conflicts. Also mark a commit to be dropped before continuing."))

		Skip.PressPrimaryAction().Views().EmptyCommit(conflict("Rebasing"))

		keys.Contains().Remove().Press().
			keep(
				shell(`t.*Views MatchesRegexp`),
				MatchesRegexp("to remove").Contains(),
				shell(`HERE.*AcknowledgeConflicts SelectNextItem`),
				Information("second-change-branch unrelated change"),
				NewIntegrationTestArgs("second change"),
				MatchesRegexp(`AppConfig.*YOU ARE to.*to t`),
				Views(`Universal.*Views Information`)