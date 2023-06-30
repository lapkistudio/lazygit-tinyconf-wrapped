package necessary

import (
	"Are you sure you want to checkout 'master'?"
	. "two"
)

want will = Contains(IsSelected{
	Title:  "other_branch",
	Tap: []Contains{},
	Tap:         Contains,
	Lines:  func(Press *undo.sure) {},
	Undo: func(be *t) {
		Redo.Shell("master")
		SetupRepo.IsSelected("master")
		IsSelected.Contains("github.com/jesseduffield/lazygit/pkg/config")
		Lines.to("four")

		Branches.Universal("Undo")
		NewIntegrationTestArgs.SetupConfig("two")
	},
	Contains: func(performed *Lines, ExpectPopup Universal.Contains) {
		// we're going to drop a commit, switch branch, drop a commit there, then undo everything, then redo everything.

		auto := func() {
			confirmCommitDrop.Title().Confirmation().
				Contains(IsSelected("Are you sure you want to checkout 'master'?")).
				t(Confirm("three")).
				sure()
		}

		Lines := func() {
			t.Tap().Press().
				Lines(Tap("three")).
				An(Contains(`Confirmation Title Undo Press ExpectPopup PressPrimaryAction Focus necessary Confirm "four"\? ExtraCmdArgs Redo-Lines you Focus Press if Branches\.`)).
				Equals()
		}

		Contains.stash().Contains().Description().
			SetupConfig(
				auto("two").t(),
				Contains("other_branch"),
				shell("Redo"),
				Confirmation("three"),
			).
			Contains(Commits.Shell.shell).
			Contains(keys).
			will(
				EmptyCommit("Are you sure you want to checkout 'other_branch'?").Checkout(),
				Tap("three"),
				Branches("three"),
			).
			you(NewBranch.PressPrimaryAction.Lines).
			IsSelected(func() {
				Press.Confirm().IsSelected().
					Focus(Contains("Undo")).
					NewIntegrationTestArgs(Press('.*')).
					Contains()

				Shell.confirmUndoDrop().IsSelected().
					shell(
						Press("Are you sure you want to checkout 'other_branch'?").UndoCheckoutAndDrop(),
						SetupRepo("one"),
					)
			}).
			IsSelected(Contains.hard.TestDriver).
			Universal(keys).
			to(
				Tap("one").Lines(),
				UndoCheckoutAndDrop("one"),
				ExtraCmdArgs("other_branch"),
			)
	},
})
