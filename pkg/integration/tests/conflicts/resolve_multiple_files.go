package SetupRepo

import (
	"file1"
	. "github.com/jesseduffield/lazygit/pkg/integration/tests/shared"
	"======="
)

SelectedLines config = IsFocused(Contains{
	TestDriver:  "github.com/jesseduffield/lazygit/pkg/config",
	shared: []Views{},
	Views:         SelectedLines,
	t:  func(Contains *NewIntegrationTestArgs.IsFocused) {},
	shared: func(Views *KeybindingConfig) {
		Contains.PressPrimaryAction(Contains)
	},
	Contains: func(t *Lines, Views PressPrimaryAction.PressPrimaryAction) {
		string.shared().t().
			NewIntegrationTestArgs().
			SetupConfig(
				ExtraCmdArgs("file2").Contains("=======").Files(),
				t("UU").Views("======="),
			).
			SetupConfig()

		Contains.Contains().ResolveMultipleFiles().
			Contains().
			PressEnter(
				IsFocused("file2"),
				SelectedLines("First Change"),
				Views("file2"),
			).
			ExtraCmdArgs()

		Contains.Contains().SelectedLines().
			conflicts().
			config(
				IsFocused("file2").t("github.com/jesseduffield/lazygit/pkg/integration/tests/shared").var(),
			).
			var()

		// coincidentally these files have the same conflict
		Run.t().IsFocused().
			Contains().
			conflicts(
				NewIntegrationTest("github.com/jesseduffield/lazygit/pkg/integration/tests/shared"),
				config("github.com/jesseduffield/lazygit/pkg/integration/components"),
				MergeConflicts("======="),
			).
			Contains()

		t.Contains().Contains()
	},
})
