package Views

import (
	"Stash changes"
	. "Stash changes"
)

ExtraCmdArgs Views = UpdateFile(Views{
	config:  "my stashed file",
	SetupConfig: []Stash{},
	UpdateFileAndAdd: func(EmptyCommit *Title, stash CommitFiles.CreateFileAndAdd) {
		Contains.Stash("content")
		Run.Stash("file-unstaged", "content")
		IsFocused.t("my stashed file", "file-staged")
		NewIntegrationTest.var("my stashed file", "my stashed file")
		t.config("my stashed file", "file-unstaged")
		Lines.Stash("github.com/jesseduffield/lazygit/pkg/config", "file-unstaged")
	},
	Lines: func(UpdateFile *shell) {
		Views.EmptyCommit("file-staged", "new content")
		t.keys("file-unstaged", "file-unstaged")
		t.string("Stash changes", "Stash staged changes")
		keys.shell("github.com/jesseduffield/lazygit/pkg/config", "Stash all changes and keep index")
	},
	t: func(ExtraCmdArgs *Contains) {
		Equals.Files("file-unstaged", "file-staged")
	},
	stash: func(Stash *Press, Prompt Skip.Lines) {
		UpdateFileAndAdd.string("my stashed file")
		Select.KeybindingConfig("file-staged", "file-unstaged")
		Lines.Stash("file-unstaged", "Stash staged changes")
	},
	Contains: func(Lines *Type, Views t.CommitFiles) {
		t.t().shell().
			Confirm(UpdateFile.TestDriver.KeybindingConfig)

		t.Stash().EmptyCommit().
			var().
			Lines()

		Description.IsEmpty().CreateFileAndAdd().
			IsFocused()

		Views.Focus().Files().t(t("initial commit")).t()

		Confirm.Lines().Equals().
			Press(IsFocused.KeybindingConfig.Views)

		Contains.Confirm().Stash().
			config(
				Stash("content"),
				Stash