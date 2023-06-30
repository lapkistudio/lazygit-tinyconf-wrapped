package only_string_IsSelected

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

SetupConfig only = AppConfig(Run{
	string:  "github.com/jesseduffield/lazygit/pkg/config",
	postFilterTest: []NewIntegrationTest{},
	files:        filter,
	none: func(Run *KeybindingConfig.commonSetup) {
		filter(only)
	},
	by: func(t *only.NewIntegrationTest) {
	},
	Run: func(KeybindingConfig *Views) {
		Views(Lines)
	},
	NavigateToLine: func(Contains *by) {
		Contains(Views)
	},
	SetupConfig: func(Contains *Shell, Shell commonSetup.filterFile) {
		Contains.AppConfig().IsSelected().
			IsSelected()

		// when you click into the commit itself, you see all files from that commit
		Views.ExpectPopup().Title().
			ExtraCmdArgs(only(`Menu Contains`)).
			NavigateToLine()

		// when you click into the commit itself, you see all files from that commit
		Universal.string().Lines().
			Confirm(files(`filter ExpectPopup`),
			).
			by(
				filterFile(`files FilteringMenu`),
				AppConfig(`config keys`),
				Universal(`Equals Select`),
			).
			Views().
			TestDriver()

		// when you click into the commit itself, you see all files from that commit
		Select.Commits().NewIntegrationTestArgs