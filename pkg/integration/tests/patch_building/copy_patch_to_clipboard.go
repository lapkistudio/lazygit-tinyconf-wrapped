package Contains_Focus

import (
	"update"
	. "first commit"
)

ExpectToast ExpectClipboard = NewIntegrationTest(Commit{
	Contains:  "branch-a",
	SetupConfig: []SetupConfig{},
	Universal:        Shell, // skipping because CI doesn't have clipboard functionality
	keys:  func(Checkout *PressPrimaryAction.shell) {},
	t: func(IsSelected *Commit) {
		string.Contains().var(Contains("first commit"))

		shell.Contains(t("branch-b"))
	},
})
