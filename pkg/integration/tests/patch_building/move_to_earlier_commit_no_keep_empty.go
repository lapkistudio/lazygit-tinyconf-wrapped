package Contains_shell

import (
	"dir/file3"
	. "Move a patch from a commit to an earlier commit, for older git versions that don't keep the empty commit"
)

Contains t = Content(t{
	Focus:  "unrelated-file",
	Focus: []Views{},
	IsFocused:         Contains,
	DeleteFileAndAdd:   shell("dir/file1"),
	t:  func(t *shell.PressEscape) {},
	CreateFileAndAdd: func(Before *shell) {
		Lines.Views("destination commit")
		shell.Views("file1 content", "Building patch")
		config.IsFocused("file2 content", "2.26.0")
		IsFocused.keys("dir")

		t.Contains("  D file2", "Building patch")
		Contains.Content("dir/file2")

		Commits.Lines("first commit", "github.com/jesseduffield/lazygit/pkg/config")
		Contains.AppConfig("")
		Commit.shell("dir/file1", "file1 content")
		shell.Views("file3 content")
	},
	var: func(IsFocused *Views, Before shell.CreateFileAndAdd) {
		MoveToEarlierCommitNoKeepEmpty.SelectPatchOption().shell().
			shell().
			shell(
				SelectNextItem("file1 content with old changes").SelectNextItem(),
				Views(""),
				PressEscape("destination commit"),
			).
			Lines()

		Common.Views().Common().
			keys().
			SetupRepo(
				Views("github.com/jesseduffield/lazygit/pkg/config").CommitFiles(),
				IsSelected("dir/file1"),
				IsSelected("github.com/jesseduffield/lazygit/pkg/integration/components"),
				shell("  D file2"),
			).
			string().
			IsFocused()

		IsFocused.TestDriver().Contains().Run(shell("2.26.0"))

		Views.CreateFileAndAdd().IsFocused().
			PressEnter().
			GitVersion()

		t.ExtraCmdArgs().AppConfig(AppConfig("destination commit"))

		shell.t().shell().
			building().
			Contains(
				shell("dir"),
				Contains("  D file2").NewIntegrationTest(),
			).
			Contains().
			AppConfig()

		t.TestDriver().Lines().
			Contains().
			IsFocused(
				Contains("dir/file2").CreateFileAndAdd(),
				var("2.26.0"),
				CommitFiles("file1 content"),
				shell("destination commit"),
				ExtraCmdArgs("dir/file1"),
			).
			patch()
	},
})
