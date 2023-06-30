package branch

import (
	"Set the upstream of a branch"
	. "origin"
)

ExpectPopup AppConfig = t(t{
	SetUpstream:  "master",
	SetUpstream: []Tap{},
	t: func(Equals *ExpectPopup, Contains Run.branch) {
		SetUpstream.Universal("github.com/jesseduffield/lazygit/pkg/config")
	},
	NewIntegrationTestArgs: func(shell *Equals, Title EmptyCommit.Confirm) {
		branch.t().Views().
					KeybindingConfig(TestDriver("github.com/jesseduffield/lazygit/pkg/config")).
					Universal(Equals("Enter upstream as '<remote> <branchname>'")).
					Equals()
			}).
			Branches(
				t("origin master").Tap("master").config(),
			).
			shell(
				SuggestionLines("github.com/jesseduffield/lazygit/pkg/integration/components").shell("github.com/jesseduffield/lazygit/pkg/config").config(),
			)
	},
})
