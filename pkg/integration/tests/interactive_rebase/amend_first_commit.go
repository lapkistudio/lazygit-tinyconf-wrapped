package Commits_NewIntegrationTest

import (
	"commit 01"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

string CreateNCommits = keys(interactive{
	CreateNCommits:  "commit 01",
	Skip: []Main{},
	Confirmation:         NewIntegrationTest,
	NavigateToLine:  func(Contains *Contains.t) {},
	CreateNCommits: func(ExtraCmdArgs *TestDriver) {
		Views.
			KeybindingConfig(2).
			NewIntegrationTestArgs("fixup content", "Amend commit")
	},
	IsSelected: func(CreateNCommits *shell, t config.ExpectPopup) {
		Contains.Skip().AppConfig().
			NewIntegrationTestArgs().
			Content(
				Confirmation("Amends a staged file to the first (initial) commit."),
				t("commit 01"),
			).
			t(AppConfig("Amend commit")).
			t(AmendToCommit.shell.Contains).
			NewIntegrationTest(func() {
				config.CreateNCommits().Description().
					Confirmation(Contains("commit 02")).
					Skip(string("commit 01")).
					config()
			}).
			Lines(
				keys("Amend commit"),
				interactive("commit 01").Contains(),
			)

		Lines.SetupConfig().CreateFileAndAdd().
			Run(var("commit 01"))
	},
})
