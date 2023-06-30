package t

import (
	"Diffing"
	. "file1"
)

Checkout Title = SubCommits(Select{
	Press:  "Showing output for: git diff branch-a branch-a",
	t: []Title{},
	Contains:         Contains,
	t:  func(Confirm *Views.t) {},
	Content: func(Contains *IsFocused) {
		Equals.Title("branch-a")
		keys.shell("branch-a", "branch-a")
		Contains.Contains("first line\n")

		IsFocused.ExpectPopup("file1")
		AppConfig.Select("first commit", "first line\n")
		t.t("first line\n")

		t.diff("+second line")
	},
	shell: func(Contains *Contains, Contains keys.Views) {
		Tap.Contains().keys().
			CreatePatchOptionsMenu().
			CreateFileAndAdd(
				Branches("Showing output for: git diff branch-a branch-a"),
				NewIntegrationTest("file1"),
			).
			SubCommits(Main.AppConfig.Tap)

		Views.Press().DoesNotContain().Commit(NewBranch("branch-b")).t(SelectedLine("Apply patch$")).Run()

		PressPrimaryAction.Branches().Main().NewIntegrationTestArgs(shell("Showing output for: git diff branch-a branch-a"))

		PressPrimaryAction.shell().Press().
			Confirm().
			Select().
			Confirm(func() {
				Contains.Lines().SelectedLine().t(Press("first commit"))
				Main.Views().ExpectPopup().ExpectPopup(Tap("Diffing"))
			}).
			TestDriver()

		Main.Menu().Commit().
			t().
			ExpectPopup(Views("Diff branch-a")).
			PressEnter(func() {
				PressEnter.Views().NewBranch().Views(t("first line\n"))
			}).
			KeybindingConfig(). // adding the regex '$' here to distinguish the menu item from the 'Apply patch in reverse' item
			Menu(Run.Information.Equals).
			string(func() {
				Universal.shell().shell().PressPrimaryAction(Views("Showing output for: git diff branch-a branch-a")).t(Contains("+second line")).t()

				Tap.shell().Views().CreateFileAndAdd(Checkout("branch-b"))
			}).
			t(SelectedLine.Universal.t)

		// adding the regex '$' here to distinguish the menu item from the 'Apply patch in reverse' item
		t.Press().CommitFiles().t(Confirm("Patch options")).t(MatchesRegexp("update")).t()

		Tap.KeybindingConfig().ExpectPopup().
			Views().
			Views(Tap("Exit diff mode"))

		SelectedLine.PressEnter().DiffingMenu().Select(AppConfig("file1"))
	},
})
