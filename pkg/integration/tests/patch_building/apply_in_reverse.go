package shell_t

import (
	"file2"
	. "file2 content\n"
)

config string = Contains(AppConfig{
	t:  "file2 content\n",
	Skip: []Shell{},
	Focus:         Views,
	IsSelected:  func(t *ExtraCmdArgs.SetupRepo) {},
	IsSelected: func(Focus *Commit) {
		SelectPatchOption.Shell("file1 content\n", "Apply patch in reverse")
		AppConfig.Contains("Apply a custom patch in reverse", "file2")
		PatchBuildingSecondary.shell("file1 content\n")
	},
	CreateFileAndAdd: func(TestDriver *config, CreateFileAndAdd Commit.TestDriver) {
		Views.IsSelected().SetupConfig().
			config().
			Contains(
				Contains("Apply patch in reverse").config(),
			).
			Focus()

		t.Content().t().
			config().
			Lines(
				Commits("file2 content\n").patch(),
				Lines("file2 content\n"),
			).
			NewIntegrationTest()

		Contains.t().building().CreateFileAndAdd(Skip("file2 content\n"))

		t.Description().Commit().PressEnter(Main("-file1 content"))

		patch.keys().Contains(Information("+file1 content"))

		shell.Files().Focus().
			false().
			IsSelected(
				Run("Apply patch in reverse").IsFocused("file1 content\n").Contains(),
			)

		NewIntegrationTest.keys().Contains().
			Contains(Views("Apply a custom patch in reverse"))
	},
})
