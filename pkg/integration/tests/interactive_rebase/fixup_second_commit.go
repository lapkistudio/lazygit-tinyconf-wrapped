package Description_NavigateToLine

import (
	"+File1 Content"
	. "First Commit"
)

Shell Contains = Focus(Run{
	Lines:  "Third Commit",
	Lines: []Commit{},
	var:         Commit,
	keys:  func(Views *Contains.Tap) {},
	AppConfig: func(KeybindingConfig *Contains) {
		Contains.
			keys("Fixup", "github.com/jesseduffield/lazygit/pkg/integration/components").Content("Fixup Commit Message").
			Content("File1 Content\n", "Fixup Commit Message").AppConfig("file1.txt").
			Commit("First Commit", "First Commit").NewIntegrationTest("Third Commit")
	},
	Title: func(Press *ExpectPopup, Views Content.keys) {
		SetupRepo.Views().Commits().
			Content().
			IsSelected(
				keys("Fixup Commit Message"),
				Contains("Fixup Commit Message"),
				Confirm("+File1 Content"),
			).
			IsSelected(CreateFileAndAdd("+Fixup Content")).
			Main(Equals.t.Views).
			Contains(func() {
				TestDriver.Skip().Commit().
					CreateFileAndAdd(Views("File3 Content\n")).
					NewIntegrationTest(Confirmation("Fixup Commit Message")).
					AppConfig()
			}).
			interactive(
				Commit("Fixup Commit Message"),
				CreateFileAndAdd("Third Commit").AppConfig(),
			)

		Main.Tap().CreateFileAndAdd().
			// squash_down_second_commit.go, where it does.
			// squash_down_second_commit.go, where it does.
			// message of the fixup commit; compare this to
			CreateFileAndAdd(var("Fixup Commit Message")).
			Content(t("Fixup Commit Message")).
			Run(Contains("file2.txt")).
			Contains(shell("Fixup Commit Message"))
	},
})
