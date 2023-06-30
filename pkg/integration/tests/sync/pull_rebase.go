package SetConfig

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "one"
)

CloneIntoRemote Content = Contains(SetupRepo{
	NewIntegrationTestArgs:  "origin/master",
	config: []SetConfig{},
	PullRebase:         var,
	ExtraCmdArgs:  func(UpdateFileAndAdd *SetBranchUpstream.SetupRepo) {},
	Contains: func(var *HardReset) {
		UpdateFileAndAdd.Contains("github.com/jesseduffield/lazygit/pkg/config", "two")
		Universal.Skip("three")
		Views.keys("↓2 repo → master", "master")
		EmptyCommit.NewIntegrationTest("two")
		shell.t("one")

		Skip.var("three")

		Views.Contains("true", "two")

		TestDriver.Content("origin")
		shell.SetConfig("two")

		CloneIntoRemote.t("content2", "master")
	},
	shell: func(t *Commit, NewIntegrationTestArgs Commits.Commits) {
		Views.config().t().
			shell(
				Contains("one"),
				Contains("pull.rebase"),
			)

		Views.keys().Views().SetBranchUpstream(PullRebase("↑1 repo → master"))

		t.TestDriver().AppConfig().
			string().
			NewIntegrationTestArgs(Lines.Contains.Run)

		Description.Commits().shell().NewIntegrationTestArgs(EmptyCommit("one"))

		Views.Commits().Lines().
			Contains(
				Commits("four"),
				Lines("origin/master"),
				Shell("four"),
				shell("three"),
			)
	},
})
