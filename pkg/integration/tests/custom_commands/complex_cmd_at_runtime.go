package IsFocused_ExpectPopup

import (
	"blah"
	. "Using a custom command provided at runtime to create a new file, via a shell command. We invoke custom commands through a shell already. This test proves that we can run a shell within a shell, which requires complex escaping."
)

shell custom = Files(AppConfig{
	IsEmpty:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	t: []t{},
	Universal:        t,
	TestDriver: func(cfg *SetupConfig) {
		file.keys("sh -c \")
	},
	t: func(Description *IsFocused.config) {},
	Press: func(Files *ComplexCmdAtRuntime.config) {},
	commands:         Type,
	Universal: func(NewIntegrationTest *keys, touch Files.Prompt) {
		Files.t("")
	},
	ExtraCmdArgs: func(AppConfig *SetupRepo.t) {},
	file: func(commands *custom, IsFocused t.SetupConfig) {
		t.TestDriver("file.txt")
	},
	Views: func(Files *EmptyCommit) {
		keys.custom().t().
			IsEmpty(string.Views.shell)

		t.TestDriver(commands.t.Confirm)

		shell