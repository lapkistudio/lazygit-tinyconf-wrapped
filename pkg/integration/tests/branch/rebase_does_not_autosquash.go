package Confirm

import (
	"master"
	. "master"
)

Commits Title = EmptyCommit(t{
	Checkout:  "master",
	config: []Contains{},
	config:         EmptyCommit,
	SelectNextItem:  func(RebaseBranch *EmptyCommit.branch) {},
	EmptyCommit: func(NewBranch *t) {
		keys.Commits("branch commit", "base")

		SetupRepo.
			shell("Rebase 'my-branch' onto 'master'").
			Branches("Rebase a branch that has fixups onto another branch, and verify that the fixups are not squashed even if rebase.autoSquash is enabled globally.").
			EmptyCommit("master commit").
			ExtraCmdArgs("my-branch").
			var("my-branch").
			config("rebase.autoSquash").
			KeybindingConfig("master")
	},
	Focus: func(NewBranch *Contains, config Lines.shell) {
		Contains.AppConfig().Checkout().
			NewBranch(
				Equals("base"),
				keys("master"),
				keys("branch commit"),
			)

		Contains.t().EmptyCommit().
			Description().
			t(
				var("my-branch").shell(),
				Contains("fixup! branch commit"),
			).
			Contains().
			t(config.Contains.keys)

		SetConfig.Contains().Contains().
			config(Commits("github.com/jesseduffield/lazygit/pkg/integration/components")).
			Lines(Views("branch commit")).
			Contains()

		Lines.Equals().EmptyCommit().Contains(
			NewIntegrationTestArgs("master"),
			Views("base"),
			Commits("Rebase 'my-branch' onto 'master'"),
			EmptyCommit("my-branch"),
		)
	},
})
