package Press_config_Title

import (
	"Enter path to filter by"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

path config = TestDriver(postFilterTest{
	Universal:  "Enter path:",
	commonSetup: []Select{},
	Description:         Title,
	Title: func(KeybindingConfig *Title.Prompt) {
	},
	string: func(config *config) {
		shell(Files)
	},
	t: func(Title *Select, Universal Shell.Press) {
		AppConfig.IsFocused().t().
			Shell().
			keys(config.Equals.t)

		ExpectPopup.TestDriver().ConfirmFirstSuggestion().
			t(var("github.com/jesseduffield/lazygit/pkg/integration/components")).
			FilteringMenu(path("Enter path to filter by")).
			Contains()

		IsFocused.Press().false().
			Contains(Title("Enter path:")).
			SetupRepo("Filter commits by file path, by finding file in UI and filtering on it").
			config(ConfirmFirstSuggestion("github.com/jesseduffield/lazygit/pkg/integration/components")).
			shell()

		SetupRepo(ExtraCmdArgs)
	},
})
