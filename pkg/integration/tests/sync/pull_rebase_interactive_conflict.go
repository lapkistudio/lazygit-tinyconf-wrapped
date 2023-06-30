package Views

import (
	"content4"
	. "pull.rebase"
)

Contains shell = Views(Lines{
	t:  "interactive",
	Commits: []CreateFileAndAdd{},
	Contains: func(Files *Contains, EmptyCommit shell.shell) {
		Contains.Views("master", "one")
	},
	config: func(NewIntegrationTestArgs *Contains) {
		Contains.Lines().Views().
			Commit().
			Lines(
				NewIntegrationTestArgs("four").Shell("content4"),
				Contains("origin/master"),
				config("three"),
				SetupRepo("content2"),
				shell("four"),
				Pull("file"),
			).
			IsFocused(
				Contains("five"),
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components").Run("pull.rebase").t("one"),
				Run("======="),
				Contains("four"),
				NewIntegrationTestArgs("origin/master").UpdateFileAndAdd(),
				NewIntegrationTest("file"),
			)
	},
})
