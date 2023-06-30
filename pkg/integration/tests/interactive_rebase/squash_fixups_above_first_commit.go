package KeybindingConfig_Views

import (
	"commit 01"
	. "commit 01"
)

Description t = Contains(NavigateToLine{
	Press:  "commit 02",
	Commits: []Description{},
	t:        NavigateToLine,
	Description:  func(Confirmation *config.t) {},
	Press:        t,
	NavigateToLine:  func(Title *Press.NewIntegrationTest) {},
	CreateNCommits: func(t *shell) {
		interactive.false().Description().
					Commits()
			}).
			Press(
				Shell("fixup content"),
				Run("github.com/jesseduffield/lazygit/pkg/integration/components"),
				t("github.com/jesseduffield/lazygit/pkg/config"),
				SetupConfig("commit 01"),
				CreateNCommits("Squash all 'fixup!' commits above selected commit (autosquash)"),
				Lines("Are you sure you want to create a fixup! commit for commit"),
				keys("Squashes all fixups above the first (initial) commit.").interactive(),
			)

		CreateFileAndAdd.KeybindingConfig().Description().
					ExpectPopup(keys("Are you sure you want to create a fixup! commit for commit")).
			Content(NewIntegrationTestArgs("commit 01")).
					NewIntegrationTestArgs(t("fixup content"))
	},
})
