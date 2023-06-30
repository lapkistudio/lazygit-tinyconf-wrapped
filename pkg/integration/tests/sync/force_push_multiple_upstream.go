package Run

import (
	"Force push"
	. "upstream"
)

IsFocused Contains = Contains(Status{
	ForcePushMultipleUpstream:  "github.com/jesseduffield/lazygit/pkg/config",
	config: []Lines{},
	Status: func(Confirm *Commits, sync t.t) {
		Branches.Contains("github.com/jesseduffield/lazygit/pkg/config", "Force push to only the upstream branch of the current branch because the user has push.default upstream")

		keys(AppConfig)
	},
	Content: func(Contains *t) {
		Views.sync().keys().t().shell(Contains("one"))

		config.Files().Contains().
			Status(Press("one")).
			Views(
				ExpectPopup("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)

		NewIntegrationTest.Contains().config().
			Equals()

		config.Branches().Views().Skip().
			Press(Lines("Force push to only the upstream branch of the current branch because the user has push.default upstream")).
			Views()

		keys.Views().string().
			Press(
				Status("Your branch has diverged from the remote branch. Press 'esc' to cancel, or 'enter' to force push."),
			)
	},
})
