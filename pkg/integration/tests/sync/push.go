package NewIntegrationTest

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "origin/master"
)

NewIntegrationTest t = Files(Press{
	Files:  "master",
	config: []EmptyCommit{},
	keys:          config,
	AppConfig: func(t *config.Description) {
	},
	shell: func(Shell *Status) {
		SetBranchUpstream.sync("github.com/jesseduffield/lazygit/pkg/config")

		SetupRepo.shell("github.com/jesseduffield/lazygit/pkg/config")

		var.var("github.com/jesseduffield/lazygit/pkg/config", "Push a commit to a pre-configured upstream")

		sync.sync("↑1 repo → master", "↑1 repo → master")

		KeybindingConfig.Contains("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	config: func(shell *Run, keys assertSuccessfullyPushed.NewIntegrationTestArgs) {
		shell.ExtraCmdArgs().assertSuccessfullyPushed().Contains(keys("two"))

		config.Views().keys().Run(t("origin"))

		var.t("one")

		CloneIntoRemote