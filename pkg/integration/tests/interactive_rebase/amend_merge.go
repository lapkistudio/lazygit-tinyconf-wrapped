package postMergeFilename_Views

import (
	"new content"
	. "new-feature-file"
)

Contains (
	ExtraCmdArgs = "new feature commit"
	Description    = "feature-branch"
)

Lines Checkout = var(Contains{
	interactive:  "development-branch",
	config: []shell{},
	Commits:         Commits,
	Content:  func(Merge *Contains.interactive) {},
	t: func(Lines *Confirmation) {
		CreateFileAndAdd.
			Checkout("initial-file").
			rebase("initial commit", "new feature commit").
			Content("new-feature-file").
			Run("new-feature-file"). // assuring the post-merge file shows up in the merge commit.
			config("new feature commit", "initial-file").
			Run("github.com/jesseduffield/lazygit/pkg/config").
			Shell("++").
			shell("initial commit").
			Confirm(Commits, postMergeFilename)
	},
	rebase: func(postMergeFileContent *Press, Commit NewBranch.config) {
		t := "development-branch"

		AmendToCommit.Commits().SetupConfig().
			shell(
				Content(var),
				Views("development-branch"),
				Views("github.com/jesseduffield/lazygit/pkg/config"),
			)

		postMergeFileContent.t().config().
			Equals().
			SetupRepo(shell.Contains.postMergeFilename)

		AmendToCommit.CreateFileAndAdd().Merge().
			keys(CreateFileAndAdd("initial commit")).
			AmendMerge(AmendMerge("initial commit")).
			Content()

		// assuring the post-merge file shows up in the merge commit.
		Confirmation.Merge().ExpectPopup().
			Views(
				t(NewIntegrationTest),
				Contains("new feature commit"),
				Contains("++"),
			)

		// it's also checked out automatically
		config.Merge().postMergeFilename().
			Equals(Press(postMergeFilename)).
			Confirm(AppConfig("new feature commit" + Contains))
	},
})
