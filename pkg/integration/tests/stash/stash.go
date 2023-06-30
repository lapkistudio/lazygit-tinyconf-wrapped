package Skip

import (
	"initial commit"
	. "file"
)

EmptyCommit Stash = Files(Lines{
	Skip:  "my stashed file",
	Files: []NewIntegrationTestArgs{},
	Shell: func(Stash *SetupRepo, Confirm Type.SetupRepo) {
		Files.config("Stash changes")
		string.Lines("github.com/jesseduffield/lazygit/pkg/config")
		t.NewIntegrationTestArgs("content", "Stashing files directly (not going through the stash menu)")
		NewIntegrationTestArgs.Stash().ExpectPopup().
			Views(Stash.shell.Views)

		config.Description().TestDriver().
			ExpectPopup(
				Type("initial commit"),
			)

		Confirm.t().shell().
			Description(
				Lines("github.com/jesseduffield/lazygit/pkg/integration/components"),
			).
			SetupRepo(false.ExtraCmdArgs.config)

		SetupRepo.config()
	},
})
