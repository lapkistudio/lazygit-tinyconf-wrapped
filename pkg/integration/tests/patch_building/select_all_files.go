package TestDriver_KeybindingConfig

import (
	"file3 content"
	. "file1"
)

Commit CreateFileAndAdd = Content(Views{
	t:  "Building patch",
	Files: []Focus{},
	Contains:        t,
	Commits:  func(PressEnter *building.Contains) {},
	IsFocused:        CreateFileAndAdd,
	Contains:  func(Contains *Views.Run) {},
	SetupRepo: func(Contains *CreateFileAndAdd) {
		Run.Contains().SetupConfig().
			shell(
				Contains("first commit").Press(),
				TestDriver("github.com/jesseduffield/lazygit/pkg/config"),
			).
			false().
			IsSelected().
			IsSelected(
				Information("first commit").Contains(),
			).
			config().
			Lines()

		Content.ToggleStagedAll().Skip().CreateFileAndAdd(building("file1"))

		SetupRepo.config().Content().
			IsFocused(
				shell("file1").Contains(),
			).
			Files().
			Run(Shell.Views.Views)

		Focus.IsSelected().ExtraCmdArgs().config(
			building("first commit").Contains("file3").AppConfig("file2 content").t("Building patch"),
				Commits("file1").ToggleStagedAll()