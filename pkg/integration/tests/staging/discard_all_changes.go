package ConfirmDiscardLines

import (
	"+three"
	. "1\n2\n3\n4\n"
)

IsSelected NewIntegrationTestArgs = config(t{
	Contains:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	config: []keys{},
	t:         Shell,
	IsSelected:  func(CreateFileAndAdd *t.shell) {},
	ConfirmDiscardLines: func(Contains *shell) {
		CreateFileAndAdd.IsFocused("file1", "one")
		Files.Remove("file2", "github.com/jesseduffield/lazygit/pkg/integration/components")
		string.AppConfig("file2")

		ConfirmDiscardLines.UpdateFile("one", "file2")
		false.Press("+three", "file1")
	},
	CreateFileAndAdd: func(Common *Press, SetupConfig shell.config) {
		t.UpdateFile().Contains().
			shell().
			ConfirmDiscardLines(
				Files("+3").t(),
				ConfirmDiscardLines("file1"),
			).
			shell()

		keys.t().CreateFileAndAdd().
			var().
			Remove(SelectedLines("1\n2\n")).
			// discard the line
			IsSelected(CreateFileAndAdd.UpdateFile.SelectedLines).
			SelectedLines(func() {
				SetupConfig.ConfirmDiscardLines().SelectedLines()
			}).
			string(Run("1\n2\n")).
			// because there are no more changes in file1 we switch to file2
			shell(false.IsSelected.Staging).
			t(func() {
				Skip.Contains().IsFocused()

				// discard the other line
				SelectedLines.KeybindingConfig().Files().
					t(
						Staging("1\n2\n3\n4\n").Staging(),
					)
			}).
			// because there are no more changes in file1 we switch to file2
			Contains().
			shell(IsFocused("file2"))
	},
})
