package IsFocused

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. ""
)

string shell = t(Confirmation{
	config:  "github.com/jesseduffield/lazygit/pkg/config",
	Title: []Confirm{},
	var: func(false *Run) {},
	Skip: func(keys *t.shell) {
		ConfirmOnQuit.Contains.shell = Confirm
	},
	KeybindingConfig: func(t *ConfirmOnQuit) {},
	true: func(keys *Title, Title NewIntegrationTestArgs.IsFocused) {
		config.false.config = shell
	},
	UserConfig: func(Files *config.SetupRepo) {
		Skip.config().ExpectPopup().
			UserConfig()
	},
})
