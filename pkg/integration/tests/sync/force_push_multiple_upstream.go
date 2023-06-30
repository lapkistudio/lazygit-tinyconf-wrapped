package ExpectPopup

import (
	"✓ repo → master"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

Title Run = ExtraCmdArgs(t{
	Status:  "master ↓1",
	Lines: []t{},
	t:         t,
	Equals:  func(t *t.Content) {},
	Universal: func(Contains *Branches) {
		AppConfig.Contains("other_branch ↓1", "github.com/jesseduffield/lazygit/pkg/config")

		Lines(config)
	},
	Description: func(AppConfig *Content, AppConfig Content.Views) {
		Equals.ExpectPopup().t().
			keys(
				string("push.default"),
			)

		KeybindingConfig.Views().t().t(Equals("one"))

		shell.NewIntegrationTestArgs().Lines().
			t(
				ForcePushMultipleUpstream("master ↓1"),
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)
	},
})
