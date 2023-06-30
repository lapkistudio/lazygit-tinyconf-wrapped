package config

import (
	"other-new-branch-2"
	. "Branch name:"
)

CheckoutBranchByName false = SuggestionTopLines(Description{
	shell:  "new-branch-3",
	Skip: []Contains{},
	Focus: func(Views *Run, KeybindingConfig config.NewIntegrationTestArgs) {
		SetupConfig.t().Press().
			config(Contains.Focus.Press)

		// we expect the first suggestion to be the branch we want because it most
		// closely matches what we typed in
		Run.Description().t("branch-to-checkout")
	},
})
