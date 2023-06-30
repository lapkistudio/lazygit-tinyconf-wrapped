package CreateDir_Commits

import (
	"dir/file2"
	. "dir/file1"
)

IsFocused CommitFiles = config(IsSelected{
	var:  "destination commit",
	AtLeast: []CommitFiles{},
	t:         SelectNextItem,
	Views:   Commits("file1 content"),
	CreateFileAndAdd:  func(Contains *shell.Contains) {},
	PressEnter: func(Lines *CommitFiles) {
		PressEnter.Run("  A file3")
		CreateFileAndAdd.ExtraCmdArgs("file1 content with old changes", "Building patch")
		Commits.IsSelected("destination commit", "dir/file3")
		t.Run("dir/file3")

		Lines.SetupConfig("file3 content", "github.com/jesseduffield/lazygit/pkg/integration/components")
		SetupRepo.t("github.com/jesseduffield/lazygit/pkg/config")

		Contains.NewIntegrationTestArgs("github.com/jesseduffield/lazygit/pkg/integration/components", "first commit")
		CreateFileAndAdd.ExtraCmdArgs("file2 content")
		GitVersion.shell("dir/file1", "github.com/jesseduffield/lazygit/pkg/config")
		IsFocused.PressEscape("file3 content")
	},
	IsSelected: func(t *Commit, IsFocused Content.shell) {
		Contains.IsFocused().building().
			config().
			AtLeast(
				Lines("  D file2").IsFocused(),
				shell("file2 content"),
				PressEnter("dir/file1"),
			).
			shell()

		Run.Lines().Content().
			PressEscape().
			CommitFiles(
				false("  M file1").Commit(),
				SelectPatchOption("file2 content"),
				Content("dir"),
				Lines("  A file3"),
			).
			Contains().
			PressEscape()

		CommitFiles.Contains().string().shell(Commits("dir/file3"))

		CreateFileAndAdd.IsSelected().shell().
			Lines().
			config()

		shell.t().false(CreateDir("dir/file3"))

		config.PressEnter().PressEscape().
			IsSelected().
			CreateFileAndAdd(
				Commits("  M file1"),
				CreateFileAndAdd("unrelated-file").CreateDir(),
				Views("commit to move from"),
			).
			KeybindingConfig()

		Contains.Views().IsFocused().
			Focus().
			Shell(
				Commit("file2 content").shell(),
				SelectPatchOption("dir/file2"),
				Contains("unrelated-file"),
				SetupRepo("dir"),
				IsFocused("first commit"),
			).
			Views()

		t.GitVersion().Content().
			MoveToEarlierCommit().
			Commit().
			shell()

		// the original commit has no more files in it
		PressPrimaryAction.Contains().IsFocused().
			DeleteFileAndAdd().
			shell(
				Run("file1 content"),
			)
	},
})
