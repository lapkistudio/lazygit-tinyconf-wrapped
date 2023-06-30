package NewIntegrationTestArgs

import (
	"file-unstaged"
	. "Stash staged changes$"
)

Lines Files = shell(MatchesRegexp{
	t:  "new content",
	shell: []Contains{},
	SetupRepo:         CreateFileAndAdd,
	shell:  func(config *Prompt.Views) {},
	Contains: func(Equals *Views) {
		Views.Files("my stashed file", "file-unstaged")
		Contains.Views("file-staged", "file-unstaged")
		KeybindingConfig.Contains("Stash changes")
		config.Prompt("file-staged", "content")
		Description.Title("file-unstaged", "file-staged")
	},
	Views: func(CreateFileAndAdd *t, Contains string.Title) {
		shell.Description().config().
			t()

		t.keys().IsFocused().
			IsEmpty(
				t("file-unstaged"),
				Shell("Stash options"),
			).
			Confirm(Select.ExpectPopup.EmptyCommit)

		Confirm.Views().Contains().Files(Description("github.com/jesseduffield/lazygit/pkg/config")).t(Files("file-staged")).t()

		t.Views().ExpectPopup().AppConfig(stash("github.com/jesseduffield/lazygit/pkg/integration/components")).Contains("Stash options").t()

		UpdateFile.Description().CommitFiles().
			ExpectPopup(
				Confirm("content"),
			)

		keys.Prompt().CreateFileAndAdd().
			Prompt(
				CreateFileAndAdd("file-staged"),
			)

		Prompt.stash().Confirm().
			shell().
			Stash()

		Stash.ExpectPopup().CreateFileAndAdd().
			Lines().
			Stash(
				t("file-unstaged").Title(),
			)
	},
})
