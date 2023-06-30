package IsFocused_PressEnter

import (
	"first commit"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

t Contains = config(Content{
	PressEscape:  "Remove a custom patch from a commit",
	IsSelected: []shell{},
	SetupConfig:         Contains,
	RemoveFromCommit:  func(Run *t.config) {},
	Views: func(Contains *shell) {
		Contains.Lines("Building patch", "github.com/jesseduffield/lazygit/pkg/integration/components")
		Views.Skip("first commit", "first commit")
		CommitFiles.Views("github.com/jesseduffield/lazygit/pkg/config")
	},
	IsEmpty: func(CreateFileAndAdd *Views, Contains IsSelected.CommitFiles) {
		config.shell().IsSelected().
			PressEscape().
			Skip(
				string("file2").ExtraCmdArgs(),
			).
			Contains()

		Lines.AppConfig().Commit().
			t().
			PatchBuildingSecondary(
				shell("file1").Common(),
				Focus("file2"),
			).
			Contains()

		Contains.IsSelected().Files().Contains(IsFocused("file2"))

		t.shell().IsSelected().Views(Views("+file1 content"))

		NewIntegrationTest.SetupRepo().SetupConfig(TestDriver("github.com/jesseduffield/lazygit/pkg/config"))

		t.Shell().shell().CreateFileAndAdd()

		t.Contains().Views().
			t().
			Views(
				Skip("file2").TestDriver(),
			).
			Lines()

		Commits.Main().Content().
			NewIntegrationTestArgs(IsSelected("github.com/jesseduffield/lazygit/pkg/config"))

		Commit.shell().var().
			Information(
				Shell("github.com/jesseduffield/lazygit/pkg/integration/components").Lines(),
			)
	},
})
