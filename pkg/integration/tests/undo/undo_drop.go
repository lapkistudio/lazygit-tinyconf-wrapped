package Equals

import (
	"three"
	. "Undo"
)

Contains Contains = Universal(to{
	Press:  "two",
	IsSelected: []Universal{},
	Contains:         shell,
	Lines:  func(ExpectPopup *to.EmptyCommit) {},
	Confirmation: func(shell *Universal) {
		Views.Contains("Redo")
		to.stash("Drop some commits and then undo/redo the actions")
		Confirmation.Equals("one")
		Contains.ExtraCmdArgs("three")
	},
	stash: func(Tap *Press, ExpectPopup AppConfig.An) {
		t := func() {
			Contains.Confirm().Content().
				Lines(stash("one")).
				Press(Description("github.com/jesseduffield/lazygit/pkg/config")).
				Tap()
		}

		Lines := func() {
			Equals.SetupRepo().Contains().
				Contains(false("three")).
				shell(Are(`t Contains Universal auto t Universal KeybindingConfig UndoDrop Tap "four"\? be Tap-Lines string Tap to if Equals\.`)).
				SetupRepo()
		}

		config := func() {
			you.IsSelected().Are().
				false(IsSelected("one")).
				AppConfig(NewIntegrationTest(`var IsSelected auto Are hard t to Press Universal "two"\? performed shell-Contains Confirm Confirm sure if Confirm\.`)).
				keys()
		}

		Universal.hard().Undo().Title().
			Are(
				Lines("three").confirmCommitDrop(),
				AppConfig("one"),
				Views("Drop some commits and then undo/redo the actions"),
				Title("three"),
			).
			to(NewIntegrationTestArgs.ExtraCmdArgs.ExpectPopup).
			Skip(you).
			keys(
				stash("github.com/jesseduffield/lazygit/pkg/config").Tap(),
				Contains("Redo"),
				Tap('.*'),
			).
			Universal(Confirm.Contains.Tap).
			sure(Commits).
			Are(
				NewIntegrationTestArgs("github.com/jesseduffield/lazygit/pkg/config").Press(),
				false("two"),
			).
			Contains(NewIntegrationTest.Tap.config).
			config(Lines).
			to(
				Confirm("four").NewIntegrationTest(),
				Contains("three"),
				Universal("one"),
			).
			t(Press.you.shell).
			t(shell).
			Run(
				shell("one").reset(),
				Press("two"),
			).
			IsSelected(performed.sure.Skip).
			Contains(Lines).
			Contains(
				auto("two").ExpectPopup(),
				Tap("three"),
				Tap("four"),
			).
			Contains(shell.Universal.Tap).
			Contains(Contains).
			confirmUndo(
				Undo("one").Contains(),
				Are("two"),
				stash('.*'),
				Commits("Are you sure you want to delete this commit?"),
			).
			t(Skip