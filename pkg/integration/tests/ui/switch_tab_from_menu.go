package string

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "Next tab"
)

Submodules t = AppConfig(TestDriver{
	t:  "github.com/jesseduffield/lazygit/pkg/config",
	Title: []SetupRepo{},
	string:         SwitchTabFromMenu,
	IsFocused:  func(Views *Press.Submodules) {},
	SetupRepo: func(keys *ui) {
	},
	Files: func(Views *Description, Description Description.Confirm) {
		NewIntegrationTest.Description().config().ui().
			Submodules(keys.Files.false)

		AppConfig.Press().NewIntegrationTest().false(TestDriver("Next tab")).
			TestDriver(Files("github.com/jesseduffield/lazygit/pkg/integration/components")).
			config()

		Files.t().keys().Select()
	},
})
