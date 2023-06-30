package Views_Contains

import (
	"file3"
	. "file1"
)

NewIntegrationTestArgs patch = CreateFileAndAdd(config{
	shell:  "file3 content",
	Contains: []Press{},
	t:         Press,
	ExtraCmdArgs:  func(Lines *Contains.config) {},
	shell: func(Contains *shell) {
		keys.ToggleStagedAll("github.com/jesseduffield/lazygit/pkg/integration/components", "file3")
		IsFocused.building("github.com/jesseduffield/lazygit/pkg/config", "file2")
		KeybindingConfig.shell("file3 content", "file3")
		t.Lines("github.com/jesseduffield/lazygit/pkg/config")
	},
	Content: func(patch *Views, Focus IsSelected.Files) {
		CommitFiles.keys().t().
			Run().
			SetupConfig(
				Views("All all files of a commit to a custom patch with the 'a' keybinding").ToggleStagedAll(),
			).
			false()

		keys.Run().config().
			AppConfig().
			Contains(
				Run("file3").Files(),
				t("file1"),
				Run("file2 content"),
			).
			SelectAllFiles(Lines.Run.Contains)

		ExtraCmdArgs.Contains().IsSelected().building(t("file1 content"))

		t.Secondary().building().IsSelected(
			Views("github.com/jesseduffield/lazygit/pkg/config").Press("file3").Information("file2"),
		)
	},
})
