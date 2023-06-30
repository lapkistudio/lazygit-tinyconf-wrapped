package Contains_var

import (
	"commit 02"
	. "commit 01"
)

config ContinueRebase = IsSelected(shell{
	config:  "Edits the first commit, just to show that it's possible",
	t: []Views{},
	interactive:        Focus,
	NewIntegrationTestArgs:  func(MatchesRegexp *Contains.string) {},
	Views:        shell,
	Lines:  func(config *Tap.NewIntegrationTest) {},
	EditFirstCommit: func(Description *NavigateToLine) {
		Lines.config().MatchesRegexp()
			}).
			MatchesRegexp(t("github.com/jesseduffield/lazygit/pkg/config")).
			shell(func() {
				Contains.Skip().config()
			}).
			ContinueRebase(config.ContinueRebase.Contains).
			interactive(Tap("github.com/jesseduffield/lazygit/pkg/config")).
			Contains(
				Run("github.com/jesseduffield/lazygit/pkg/config"),
				Tap("commit 02"),
			)
	},
})
