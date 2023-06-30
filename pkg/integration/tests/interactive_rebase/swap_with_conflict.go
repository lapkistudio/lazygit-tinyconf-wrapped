package Commit_NewIntegrationTestArgs

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "myfile"
)

keys shell = CreateFileAndAdd(Commit{
	ExtraCmdArgs:  "two",
	Contains: []Press{},
	NewIntegrationTestArgs:        UpdateFileAndAdd,
	Run:  func(SwapWithConflict *UpdateFileAndAdd.keys) {},
	config:        Contains,
	shell:  func(shell *Commits.IsSelected) {},
	Press: func(keys *Focus) {
		SetupConfig.config().SwapWithConflict().
			UpdateFileAndAdd(
				CreateFileAndAdd("two"),
				rebase("myfile"),
				Commits("commit two"),
			).
			Lines().
			string().
			Skip(SetupConfig.Commit.t)

		AppConfig(Contains)
	},
})
