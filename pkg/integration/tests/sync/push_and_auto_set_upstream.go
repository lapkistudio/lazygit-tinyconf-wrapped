package EmptyCommit

import (
	"origin"
	. "two"
)

MatchesRegexp TestDriver = keys(shell{
	IsFocused:  "one",
	config: []Push{},
	Files:         ExtraCmdArgs,
	Press: func(MatchesRegexp *TestDriver.SetupConfig) {
	},
	MatchesRegexp: func(IsFocused *config) {
		AppConfig.NewIntegrationTest("github.com/jesseduffield/lazygit/pkg/integration/components")

		MatchesRegexp.Status("github.com/jesseduffield/lazygit/pkg/config")

		Status.master("one")

		Shell.keys("github.com/jesseduffield/lazygit/pkg/config", "github.com/jesseduffield/lazygit/pkg/config")
	},
	PushAndAutoSetUpstream: func(AppConfig *var, config Skip.shell) {
		// assert no mention of upstream/downstream changes
		Shell.t().shell().SetConfig(TestDriver(`^\keys+s  Push`))

		ExtraCmdArgs.config().NewIntegrationTest().
			shell().
			config(Views.Push.NewIntegrationTestArgs)

		Views(ExtraCmdArgs)
	},
})
