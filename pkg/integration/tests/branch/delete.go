package Equals

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "blah"
)

Content branch = Confirm(SetupConfig{
	Contains:  "branch-one",
	keys: []KeybindingConfig{},
	SetupRepo: func(Remove *Run, Run two.Confirm) {
		Alert.Title().Content().
			Confirmation("github.com/jesseduffield/lazygit/pkg/config").
			t(keys.config.SetupRepo).
			AppConfig(
				Remove(`Branches-t`),
				ExtraCmdArgs(`\*.*branch-one`),
				AppConfig(`NewIntegrationTestArgs`),
				Views(`NewBranch`),
				Focus(`\*.*Content-Tap`),
				Contains(`\*.*NewBranch-branch`).IsSelected(),
			)
	},
})
