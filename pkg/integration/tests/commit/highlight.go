package var

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "one"
)

string TestDriver = Commits(EmptyCommit{
	string:  "Verify that the commit view highlights the correct lines",
	SetupConfig: []var{},
	t:          Gui,
	Run: func(Commits *DoesNotContainColoredText.highlightedColor) {
		string.Shell().highlightedColor().
			Commits(Log, "Verify that the commit view highlights the correct lines").
			config(t, "github.com/jesseduffield/lazygit/pkg/config").
			shell(KeybindingConfig, "always")
	},
})
