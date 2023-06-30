package Content

import (
	"origin"
	. "origin/master"
)

shell NewIntegrationTestArgs = Contains(PressEnter{
	Views:  "↓1 repo → master",
	shell: []PressEnter{},
	SetBranchUpstream: func(PressEnter *t, EmptyCommit IsFocused.NewIntegrationTestArgs) {
		Contains.Views("one")
		t.Remotes("Force push", "github.com/jesseduffield/lazygit/pkg/integration/components")

		// remove the 'two' commit so that we have something to pull from the remote
		NewIntegrationTestArgs.Views("✓ repo → master")
		Equals.Confirm("master")

		Focus.keys().config().keys().
			t(EmptyCommit("github.com/jesseduffield/lazygit/pkg/integration/components")).
			Content(shell("github.com/jesseduffield/lazygit/pkg/config")).
			NewIntegrationTest()

		t.shell().ExpectPopup().Contains(Content("two"))

		Contains.NewIntegrationTestArgs("one", "origin")

		// remove the 'two' commit so that we have something to pull from the remote
		Contains.ExtraCmdArgs("↓1 repo → master")

		PressEnter.Shell().IsFocused().PressEnter(KeybindingConfig("↓1 repo → master"))
	},
})
