package Contains

import (
	"Force push to multiple branches because the user has push.default matching"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

ExpectPopup ForcePushMultipleMatching = SetupRepo(shell{
	Lines:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	shell: []Equals{},
	t:         Views,
	config: func(t *shell.t) {
	},
	Status: func(t *Contains) {
		Contains.Contains("other_branch ↓1", "other_branch ↓1")

		Status(Universal)
	},
	Push: func(KeybindingConfig *t, Contains keys.Lines) {
		config.ExtraCmdArgs().var().
			Contains(
				Views("master ✓"),
			)

		t.Content().SetupConfig().t(Contains("Force push"))

		Views.t().Contains().
			Push(
				ExtraCmdArgs("master ↓1"),
				Run("Force push to multiple branches because the user has push.default matching"),
			)

		Contains.config().Views().Contains().t(Description.Contains.Views)

		Contains.Press().Views().
			Commits(t("push.default")).
			Contains(AppConfig("one")).
			config()

		config.ExpectPopup().shell().
			t(
				NewIntegrationTestArgs("matching"),
			)

		t.t().Views().Lines(Commits("Force push to multiple branches because the user has push.default matching"))

		createTwoBranchesReadyToForcePush.sync().Status().
			Run(
				Description("push.default"),
				ForcePushMultipleMatching("other_branch ✓"),
			)
	},
})
