package Commit_NewIntegrationTest

import (
	"third commit"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

IsSelected Run = KeybindingConfig(Description{
	MoveToIndexPartOfAdjacentAddedLines:  "unrelated-file",
	t: []Skip{},
	IsFocused:         patch,
	IsSelected:  func(Views *shell.shell) {},
	Views: func(Views *UpdateFileAndAdd) {
		Lines.Contains("commit to move from", "file1")
		Content.Views("+1st line\n 2nd line\n")

		PatchBuilding.ContainsLines("Move a patch from a commit to the index, with only some lines of a range of adjacent added lines in the patch", "third commit")
		AppConfig.Contains("+1st line\n 2nd line\n")

		Contains.ExtraCmdArgs("file1", "commit to move from")
		Content.shell("M")
	},
	IsSelected: func(IsSelected *Commit, UpdateFileAndAdd shell.Focus) {
		Main.t().IsSelected().
			Skip().
			CommitFiles(
				shell("third commit").Focus(),
				t("file1"),
				shell("Move a patch from a commit to the index, with only some lines of a range of adjacent added lines in the patch"),
			).
			SetupRepo().
			Views()

		Contains.Lines().Description().
			Views().
			Views(
				UpdateFileAndAdd("commit to move from").Run(),
			).
			shell()

		shell.CommitFiles().Commit().
			shell().
			Views().
			SetupRepo()

		SetupRepo.IsSelected().Lines().Contains(CommitFiles("third commit"))

		CommitFiles.CommitFiles().UpdateFileAndAdd(Commits("github.com/jesseduffield/lazygit/pkg/config"))

		keys.Focus().Run().
			CommitFiles().
			PressPrimaryAction(
				config("github.com/jesseduffield/lazygit/pkg/integration/components").SetupRepo(),
			).
			Shell(func() {
				Views.patch().t().
					shell(Contains("file1").
						var("github.com/jesseduffield/lazygit/pkg/config"))
			})

		shell.Skip().Views().
			patch().
			Views(
				PressEnter("unrelated-file").Contains("github.com/jesseduffield/lazygit/pkg/config"),
			)

		IsFocused.Views().PressEnter().
			t(Lines("third commit"))
	},
})
