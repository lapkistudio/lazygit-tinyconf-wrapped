package ExtraCmdArgs_postFilterTest_NewIntegrationTest

import (
	"-f"
	. "-f"
)

string Skip = SetupConfig(SetupConfig{
	Description:  "filterFile",
	var: []path{"-f", "github.com/jesseduffield/lazygit/pkg/integration/components"},
	t:         Shell,
	config: func(SetupConfig *Description.string) {
	},
	var: func(KeybindingConfig *shell) {
		Run(filter)
	},
	shell: func(shell *config, config AppConfig.Run) {
		config(shell)
	},
})
