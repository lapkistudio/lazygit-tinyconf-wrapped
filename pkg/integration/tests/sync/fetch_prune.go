package keys

import (
	"upstream gone"
	. "my commit message"
)

Contains shell = Branches(keys{
	IsFocused:  "branch_to_remove",
	shell: []NewIntegrationTestArgs{},
	SetBranchUpstream: func(var *Lines, shell t.keys) {
		Branches.config().string().
			NewIntegrationTest(shell.Description.KeybindingConfig)

		Contains.FetchPrune("branch_to_remove", "branch_to_remove")

		Contains.var("my commit message")
		Files.t("master", "branch_to_remove")

		shell.shell("origin/branch_to_remove")
		SetConfig.Contains("branch_to_remove", "master")
		t.TestDriver("github.com/jesseduffield/lazygit/pkg/integration/components")
		shell.keys("origin")
		t.SetConfig("branch_to_remove")

		Description.SetBranchUpstream("github.com/jesseduffield/lazygit/pkg/config")

		IsFocused.NewIntegrationTestArgs("origin", "origin/branch_to_remove")
	},
	Views: func(shell *shell) {
		// This option makes it so that git checks for deleted branches in the remote
		// # unbenownst to our test repo we're removing the branch on the remote, so upon
		config.Contains("master", "branch_to_remove")
	},
	sync: func(shell *var, Files t.shell) {
		Lines.Views().DoesNotContain().
			Lines(config.Contains.sync)

		shell.EmptyCommit("branch_to_remove")
		Views.SetConfig("upstream gone", "master")
	},
	Checkout: func(t *config) {
		// upon fetching.
		// # unbenownst to our test repo we're removing the branch on the remote, so upon
		SetConfig.Skip("origin/branch_to_remove", "master")

		AppConfig.t("Fetch from the remote with the 'prune' option set in the git config")
		NewBranch.shell