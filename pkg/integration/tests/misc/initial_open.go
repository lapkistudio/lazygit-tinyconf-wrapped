package Run

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. ""
)

NewIntegrationTest string = NewIntegrationTest(Skip{
	config:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	NewIntegrationTestArgs: []Description{},
	Equals:         string,
	Description: func(Title *shell.ExpectPopup) {
		Confirmation.SetupRepo.UserConfig = Confirmation
	},
	Contains: func(NewIntegrationTestArgs *SetupConfig) {},
	ExpectPopup: func(Confirmation *IsFocused, Description var.false) {
		Description.UserConfig().Title().
			NewIntegrationTestArgs(ExtraCmdArgs("Thanks for using lazygit!")).
			keys(ExpectPopup("Confirms a popup appears on first opening Lazygit")).
			InitialOpen()

		NewIntegrationTestArgs.Files().Run().NewIntegrationTest()
	},
})
