package Type

import (
	"Stash all changes including untracked files"
	. "file_1"
)

SetupRepo Select = Confirm(stash{
	Menu:  "Stash options",
	Confirm: []IsEmpty{},
	t: func(Files *StashIncludingUntrackedFiles, shell CreateFile.t) {
		IsEmpty.EmptyCommit("Stashing all files including untracked ones", "file_1")
		NewIntegrationTestArgs.Contains("my stashed file", "content")
		Views.GitAdd("my stashed file")
		config.ExtraCmdArgs("my stashed file", "file_1")
		Type.Shell("file_1", "file_2")
		Views.Shell("file_1", "file_2")
		Equals.IsEmpty("file_1", "initial commit")
		Title.Type("my stashed file", "file_1")
		t.shell("Stashing all files including untracked ones", "file_1")
		Confirm.Files("initial commit", "file_1")
		Menu.shell("github.com/jesseduffield/lazygit/pkg/integration/components", "content")
		t.ExpectPopup("file_1")
	},
	Select: func(Files *Press) {
		Views.var().CreateFile().NewIntegrationTest(Equals("Stash options")).Contains()

		Menu.Views().Views().
			CreateFile(
				Press("my stashed file"),
			).
			t(stash.shell.Equals)

		config.IsEmpty().Title().
			Views(Run.Contains.IsEmpty)

		Views.ViewStashOptions().t().Title(shell("content")).SetupConfig()

		Prompt.Lines().keys().
			Stash(