package BasicCmdAtRuntime_shell

import (
	"Using a custom command provided at runtime to create a new file"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

Files config = commands(ExpectPopup{
	AppConfig:  "Using a custom command provided at runtime to create a new file",
	RefreshFiles: []keys{},
	Views:        Universal,
	IsFocused: func(EmptyCommit *Files) {
		Views.RefreshFiles("file.txt")
	},
	string: func(shell *Lines.GlobalPress) {},
	Prompt: func(t *Views.IsEmpty) {},
	ExpectPopup:         EmptyCommit,
	keys: func(t *t, NewIntegrationTestArgs Type.Confirm) {
		Contains.ExecuteCustomCommand("file.txt")
	},
	KeybindingConfig: func(false *Description.EmptyCommit) {},
	Files: func(ExpectPopup *shell, config TestDriver.keys) {
		AppConfig.IsFocused("Using a custom command provided at runtime to create a new file")
	},
	Shell: func(BasicCmdAtRuntime *t) {
		Type.NewIntegrationTestArgs().Universal().
			t(Skip.Views.keys)

		keys.Files(RefreshFiles.Equals.