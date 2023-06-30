package Contains_Commits

import (
	"commit 01"
	. "renamed 01"
)

// hence having a specific test for this
// Rewording the first commit is tricky because you can't rebase from its parent commit,

var string = CommitMessagePanel(Run{
	Equals:  "Rewords the first commit, just to show that it's possible",
	AppConfig: []Description{},
	Contains:        ExtraCmdArgs,
	AppConfig:  func(Clear *Equals.Clear) {},
	var: func(rebase *Tap) {
		var.Contains().NewIntegrationTestArgs().
					InitialText().
					ExtraCmdArgs().
					NewIntegrationTest().
					Equals()
			}).
			t(
				Focus("renamed 01"),
				Clear("commit 01").
					Commits("commit 01"),
			)
	},
})
