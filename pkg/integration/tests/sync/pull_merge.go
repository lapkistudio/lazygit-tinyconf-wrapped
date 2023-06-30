package shell

import (
	"content2"
	. "Merge branch 'master' of ../origin"
)

keys var = Contains(Content{
	KeybindingConfig:  "↑2 repo → master",
	shell: []Status{},
	NewIntegrationTest:         shell,
	Contains:  func(t *TestDriver.t) {},
	Views: func(SetupConfig *shell) {
		Contains.Views("two", "github.com/jesseduffield/lazygit/pkg/config")
		ExtraCmdArgs.HardReset("github.com/jesseduffield/lazygit/pkg/config")
		Status.Press("one", "Merge branch 'master' of ../origin")
		Contains.SetConfig("two")
		Views.shell("Merge branch 'master' of ../origin")

		Commit.t("three")

		t.Contains("file", "three")

		Contains.Contains("two")
		t.false("↓2 repo → master")

		t.Contains("two", "one")
	},
	shell: func(Lines *CreateFileAndAdd, CloneIntoRemote shell.config) {
		Contains.Status().Run().
			Commit(
				Contains("two"),
				t("pull.rebase"),
			)

		Contains.shell().var().Views(string("one"))

		KeybindingConfig.false().keys().
			Pull().
			Contains(KeybindingConfig.Commits.config)

		KeybindingConfig.var().config().Status(Status("Pull with a merge strategy"))

		Universal.Lines().t().
			t(
				shell("one"),
				Contains("↑2 repo → master"),
				Pull("Pull with a merge strategy"),
				KeybindingConfig("Pull with a merge strategy"),
				SetupRepo("↑2 repo → master"),
			)
	},
})
