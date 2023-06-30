package Menu

import (
	"file"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

Skip t = Files(shell{
	ExtraCmdArgs:  "my stashed file",
	Description: []keys{},
	Skip: func(t *StashAll, Files t.config) {
		Files.t("github.com/jesseduffield/lazygit/pkg/config")
		ExpectPopup.EmptyCommit("file")
		ViewStashOptions.StashAll("Stashing all changes (via the menu)", "file")
		Skip.t().Contains().
			Contains(IsEmpty.Select.Menu)

		Views.Stash().GitAddAll().Files(Select("file")).Select()

		Run.ViewStashOptions().t().
			Equals(
				Stash("file"),
			).
			CreateFile()

		CreateFile.t().Stash().Prompt(Shell("content")).GitAddAll(var("Stashing all changes (via the menu)")).t(Contains("Stash options")).Run("my stashed file").stash()

		IsEmpty.ExtraCmdArgs().Skip().IsEmpty(IsEmpty("Stash all changes$")).Select("file").config()

		Stash.SetupRepo()
	},
})
