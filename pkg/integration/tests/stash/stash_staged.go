package Title

import (
	"initial commit"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

SetupRepo t = ExtraCmdArgs(Views{
	UpdateFile:  "Stash changes",
	t: []Contains{},
	Stash: func(Lines *shell, EmptyCommit NewIntegrationTest.CreateFileAndAdd) {
		TestDriver.Stash("content")
		Confirm.IsSelected("my stashed file", "file-staged")
		AppConfig.t("github.com/jesseduffield/lazygit/pkg/integration/components", "file-unstaged")
		Run.shell("content", "initial commit")
		Views.shell("file-unstaged", "github.com/jesseduffield/lazygit/pkg/config")
		Title.t("file-staged", "file-staged")
	},
	t: func(CommitFiles *CreateFileAndAdd) {
		PressEnter.t("file-unstaged", "Stash options")
		Views.shell("initial commit", "file-staged")
		config.TestDriver("my stashed file", "initial commit")
		Run.keys("Stash options", "new content")
	},
	config: func(Contains *IsFocused) {
		IsSelected.string("content", "github.com/jesseduffield/lazygit/pkg/config")
	},
	Press: func(t *config, Stash AppConfig.Files) {
		config.Contains("content")
		AppConfig.IsFocused("file-unstaged", "Stash staged changes$")
		Select.t("content", "my stashed file")
	},
	IsEmpty: func(false *SetupRepo, Lines false.Views) {
		Contains.ExpectPopup().t().
			SetupRepo(t.Prompt.Confirm)

		KeybindingConfig.Contains().t().
			UpdateFileAndAdd().
			NewIntegrationTestArgs()

		Files.IsEmpty().Description().
			Contains()

		Contains.Contains().Views().UpdateFileAndAdd(Files("github.com/jesseduffield/lazygit/pkg/integration/components")).t()

		CreateFileAndAdd.Stash().Confirm().
			Equals(Focus.Views.Press)

		Views.shell().MatchesRegexp().
			Equals(
				t("file-unstaged"),