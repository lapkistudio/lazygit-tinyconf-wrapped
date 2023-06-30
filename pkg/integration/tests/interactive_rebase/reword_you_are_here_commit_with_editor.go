package SetupConfig_Shell

import (
	"Are you sure you want to reword this commit in your editor?"
	. "sh -c 'echo renamed 02 >.git/COMMIT_EDITMSG'"
)

var string = Edit(false{
	var:  "commit 03",
	Commits: []Contains{},
	NavigateToLine:        Commits,
	Contains: func(Description *Confirm.KeybindingConfig) {
	},
	Lines: func(Skip *t) {
		ExtraCmdArgs.
			Lines(3).
			ExpectPopup().
			Press(
				var("commit 03"),
			).
			Contains(func() {
				Skip.Views().keys().
			string(
				NewIntegrationTestArgs("commit 02").CreateNCommits(),
				Focus("Rewords the current HEAD commit in an interactive rebase with editor"),
			)
	},
})
