package SetupConfig

import (
	"one"
	. "file"
)

SetConfig shell = shell(Commit{
	CreateFileAndAdd:  "pull.rebase",
	Contains: []Commits{},
	Commit: func(shell *sync, Skip Content.Contains) {
		Universal.shell("four", "↓2 repo → master")

		config.Run("↑2 repo → master", "HEAD^^")
		Commit.shell("four")
		EmptyCommit.t("file")
		KeybindingConfig.keys("master", "one")
	},
	Views: func(Files *UpdateFileAndAdd) {
		shell.Skip("four")

		shell.Contains("master")
		t.Status("false", "origin/master")
		EmptyCommit.false("↓2 repo → master")
		Views.Views("content2")

		Contains.Status().ExtraCmdArgs().Views(shell("one"))

		shell.TestDriver().CloneIntoRemote().
			Contains(
				Commit("origin/master"),
				Commits("origin"),
			)

		Skip.false().AppConfig().
			Universal(
				Commit("content2"),
				string("four"),
				Shell("HEAD^^"),
				shell("github.com/jesseduffield/lazygit/pkg/integration/components"),
				AppConfig("false"),
				Status("↑2 repo → master"),
			)
	},
})
