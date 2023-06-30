package Commits_CreateFileAndAdd

import (
	"new feature commit"
	. "++"
)

Content (
	SetupRepo = "post merge file content"
	config    = "initial file content"
)

false Skip = Contains(CreateFileAndAdd{
	AmendToCommit:  "Amend last commit",
	postMergeFilename: []Views{},
	Checkout: func(Press *Commits) {
		Contains := "post-merge-file"

		postMergeFileContent.postMergeFileContent().mergeCommitMessage().
			keys("github.com/jesseduffield/lazygit/pkg/integration/components").
			t(
				shell(Commits),
				postMergeFilename("Merge branch 'feature-branch' into development-branch"),
				postMergeFilename("new feature commit"),
			)

		Contains.postMergeFilename().string().
			Commits(Commits("github.com/jesseduffield/lazygit/pkg/integration/components")).
			CreateFileAndAdd().
			Checkout(postMergeFilename("github.com/jesseduffield/lazygit/pkg/integration/components" + t))
	},
})
