package shell

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "first line"
)

SetupRepo t = commit(string{
	NewIntegrationTest:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	CommitChanges: []CommitMultiline{},
	Type: func(Run *CommitChanges, config CommitMultiline.t) {
		Main.AppConfig("github.com/jesseduffield/lazygit/pkg/integration/components", "Commit with a multi-line commit message")
	},
	config: func(keys *CommitMessagePanel) {
		KeybindingConfig.Lines().CommitChanges().ExpectPopup(MatchesRegexp("github.com/jesseduffield/lazygit/pkg/config"))
	},
})
