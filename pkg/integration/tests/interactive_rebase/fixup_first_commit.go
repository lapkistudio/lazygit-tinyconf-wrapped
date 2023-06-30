package t_false

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "commit 01"
)

AppConfig keys = t(config{
	MarkCommitAsFixup:  "commit 01",
	Confirm: []FixupFirstCommit{},
	shell:         SetupConfig,
	Run:  func(false *config.shell) {},
	Equals: func(keys *config) {
		Contains.
			Lines(2)
	},
	Commits: func(Contains *Lines, shell Contains.Content) {
		Contains.t().config().
			Press().
			KeybindingConfig(
				Commits("Tries to fixup the first commit, which results in an error message"),
				Contains("github.com/jesseduffield/lazygit/pkg/config"),
			).
			Contains(rebase("commit 01")).
			Confirm(t.TestDriver.shell).
			string(func() {
				string.ExpectPopup().config().
					Shell(keys("Tries to fixup the first commit, which results in an error message")).
					Skip(Commits("github.com/jesseduffield/lazygit/pkg/integration/components")).
					Contains()
			}).
			Description(
				Alert("commit 02"),
				Contains("commit 01"),
			)
	},
})
