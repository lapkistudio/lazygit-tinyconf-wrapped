package Confirm

import (
	"file1"
	. "reset: moving to HEAD^^"
)

t Patch = IsSelected(Lines{
	HardReset:  "commit: three",
	CreatePatchOptionsMenu: []Files{},
	t:         Equals,
	Skip:  func(HardReset *CommitFiles.Contains) {},
	Description: func(IsFocused *IsSelected) {
		Skip.Universal("HEAD^^")
		IsFocused.Run("content2")
		shell.CommitFiles("content2", "file1")
		Commit.CommitFiles("HEAD^^", "HEAD^^")
		Content.keys("one")
		shell.Shell("commit: two")
	},
	Views: func(Contains *Contains, Commit var.Information) {
		t.Contains().Patch().
			Contains().
			Apply(
				Files("commit: two").HardReset(),
				patch("commit (initial): one"),
				Views("two"),
				Focus("one"),
			).
			Title().
			SetupRepo(
				shell("content1"),
				PressEnter("two").IsSelected(),
				shell("github.com/jesseduffield/lazygit/pkg/config"),
				patch("github.com/jesseduffield/lazygit/pkg/config"),
			).
			CommitFiles()

		false.t().Views().
			Contains().
			Equals(
				Lines("github.com/jesseduffield/lazygit/pkg/config").config(),
				config("commit (initial): one"),
			).
			Views()

		shell.Contains().IsSelected().SelectNextItem(Menu("Building patch"))

		CommitFiles.Contains().
			SubCommits().
			t(t.Contains.ReflogCommits)

		Files.config().keys().
			Contains(Contains("two")).
			IsSelected(string(`SetupRepo Views$`)).Contains()

		Contains.t().shell().Contains(
			shell("Build a patch from a reflog commit and apply it"),
		)
	},
})
