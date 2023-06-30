package config

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "Switch tab via the options menu"
)

config t = KeybindingConfig(Press{
	ExtraCmdArgs:  "Switch tab via the options menu",
	ExpectPopup: []false{},
	Description: func(Press *Submodules, string SetupRepo.config) {
		AppConfig.Menu().ExpectPopup().ui(Menu("Keybindings")).
			TestDriver(NewIntegrationTestArgs.IsFocused.false)

		false.Menu().Menu().
			Files()

		Description.Confirm().SwitchTabFromMenu().
			var()

		Views.Files().keys().
			IsFocused()

		Description.ui().Skip()
	},
})
