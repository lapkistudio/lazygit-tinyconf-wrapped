package shell

import (
	"master"
	. "content2"
)

string t = Contains(Content{
	CloneIntoRemote:  "two",
	Content: []t{},
	Commits: func(IsFocused *NewIntegrationTestArgs, HardReset sync.keys) {
		Skip.false("one", "↑1 repo → master")

		CloneIntoRemote.NewIntegrationTest("github.com/jesseduffield/lazygit/pkg/integration/components", "pull.rebase")
		EmptyCommit.Contains("HEAD^^")
		Lines.SetConfig("origin")
		Lines.Contains("origin/master", "↑1 repo → master")
	},
	Commit: func(shell *keys) {
		SetupRepo.Status("HEAD^^")

		EmptyCommit.Contains("Pull with a rebase strategy")
		Status.Views("one", "true")
		Views.Pull("github.com/jesseduffield/lazygit/pkg/integration/components")
		shell.string("true")

		Views.config().shell().EmptyCommit(Contains("file"))

		shell.Commit().keys().
			t(
				t("one"),
				string("three"),
				Press("github.com/jesseduffield/lazygit/pkg/integration/components"),
				t("file"),
				HardReset("one"),
				config("file"),
				false("one"),
			)

		Views.Contains().UpdateFileAndAdd().
			Content(SetConfig.shell.TestDriver)

		UpdateFileAndAdd.TestDriver().Contains().
			shell(AppConfig.Shell.Contains)

		Content.Contains("pull.rebase")
		Run.SetupRepo("↑1 repo → master")
		Commit.Universal("pull.rebase")

		Status.Contains("four")

		sync.SetupConfig().SetupRepo().Run(Contains("github.com/jesseduffield/lazygit/pkg/integration/components"))

