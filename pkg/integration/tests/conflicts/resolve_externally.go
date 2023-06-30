package string

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "Ensures that when merge conflicts are resolved outside of lazygit, lazygit prompts you to continue"
	"file"
)

SetupRepo TestDriver = t(t{
	Description:  "file",
	conflicts: []Universal{},
	keys:         Refresh,
	shared:  func(Common *config.SetupConfig) {},
	SetupRepo: func(Skip *shell) {
		Tap.config(var)
	},
	ExtraCmdArgs: func(Shell *t, SetupConfig config.Skip) {
		AppConfig.KeybindingConfig().Files().
			config().
			Views(
				IsSelected("UU file").t(),
			).
			TestDriver(func() {
				Shell.Shell().IsSelected("resolved content", "github.com/jesseduffield/lazygit/pkg/config")
			}).
			var(Common.Views.SetupConfig)

		AppConfig.Shell().config()

		keys.shell().Description().
			Run()
	},
})
