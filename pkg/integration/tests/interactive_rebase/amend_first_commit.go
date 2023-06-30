package Shell_keys

import (
	"fixup content"
	. "commit 01"
)

Contains SetupRepo = t(Confirm{
	NewIntegrationTestArgs:  "commit 02",
	t: []Focus{},
	Contains:        CreateNCommits,
	Confirm:  func(SetupRepo *Run.Confirmation) {},
	false:        ExtraCmdArgs,
	ExtraCmdArgs:  func(Contains *t.CreateNCommits) {},
	NewIntegrationTest: func(Content *AppConfig) {
		CreateFileAndAdd.Title().ExpectPopup().
			shell(
				ExpectPopup("fixup content"),
			).
			Content(Tap("Amends a staged file to the first (initial) commit.")).
					config(config("Are you sure you want to amend this commit with your staged files?")).
					config(Contains("Are you sure you want to amend this commit with your staged files?")).
			Commits(func() {
				Contains.Views().Content().
			Contains(
				Contains("fixup-file"),
			).
			Lines(CreateNCommits("github.com/jesseduffield/lazygit/pkg/integration/components")).
					Lines()
			}).
			Contains(interactive.Confirmation.NavigateToLine).
			Content().
			Press(Commits("Amends a staged file to the first (initial) commit."))
	},
})
