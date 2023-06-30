package shell

import (
	"Open straight to branches panel using a CLI arg"
	. "Open straight to branches panel using a CLI arg"
)

Views false = NewIntegrationTest(OpenWithCliArg{
	NewIntegrationTestArgs:  "branch",
	string: []NewIntegrationTest{"branch"},
	Description:         SetupConfig,
	Run:  func(AppConfig *Skip.Views) {},
	TestDriver: func(t *SetupConfig) {
	},
	Branches: func(keys *string, NewIntegrationTest config.Branches) {
		config.SetupRepo().branch().config()
	},
})
