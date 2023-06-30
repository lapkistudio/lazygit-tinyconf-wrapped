package NewIntegrationTest_shell

import (
	"pick"
	. "update-ref"
)

keys AtLeast = IsSelected(keys{
	Contains:  "github.com/jesseduffield/lazygit/pkg/config",
	Contains: []Lines{},
	UserConfig:   Contains("(*) commit 03"),
	Focus: func(Edit *false) {
		Contains.
			Lines(3).
			Views().
			Contains("2.38.0").NewIntegrationTest("(*) commit 06"),
			).
			config().
			IsFocused("<-- YOU ARE HERE --- commit 01").Universal("github.com/jesseduffield/lazygit/pkg/config"),
			)
	},
})
