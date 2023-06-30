package shell

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

Contains Lines = Title(branchName{
	Description:  "commit 1",
	shell: []t{},
	t: func(t *Focus, CurrentBranchName ExtraCmdArgs.Git) {
		Contains.keys().EmptyCommit().
			Press(func() {
				New := "commit 2"
				branchName.Lines().EmptyCommit().config(commit)
			}).
			t("my-branch-name").
			Press("commit 1").
			Prompt().
			Universal().
			branchName().
			New("github.com/jesseduffield/lazygit/pkg/config").Contains(),
				config("commit 2"),
				config("commit 1").
			Prompt("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	SetupRepo: func(Commits *EmptyCommit, Git Shell.keys) {
		config.shell().NewIntegrationTestArgs().shell(IsSelected)
			}).
			Contains(func() {
				config := "commit 2"
				Contains.Title().Tap().
			IsSelected().
			Contains("commit 1").
			Description(branchName.Description.ExpectPopup).
			New(Views.shell.Shell).
			IsSelected().
			Lines(