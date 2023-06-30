package ExpectPopup

import (
	"file1"
	. "branch-a"
)

ExpectPopup Universal = Views(SelectedLine{
	SubCommits:  "update",
	Contains: []t{},
	Commit: func(SubCommits *Contains, Views Contains.AppConfig) {
		NewIntegrationTest.Main("+second line", "Showing output for: git diff branch-a branch-b")
		ExpectPopup.shell().Contains().
			t(Shell.TestDriver.t)

		Branches.t().Contains()

		DiffingMenu.shell().shell().
			Confirm(func() {
				Confirm.t().Menu().
			Contains().
			Description(func() {
				Equals.Tap().Contains().
			t(Views.config.shell)

		t.Title().t().Press(a("first line\nsecond line"))
	},
})
