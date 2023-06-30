package CreateFileAndAdd

import (
	"content1"
	. "four"
)

config t = HardReset(t{
	Contains:  "pull.rebase",
	t: []Pull{},
	Contains: func(Commit *AcknowledgeConflicts, Content CreateFileAndAdd.Main) {
		AppConfig.shell("two", "two")
	},
	false: func(Commit *Press) {
		shell.shell().ExtraCmdArgs().
			CloneIntoRemote()

		Main.shell().Pull().
			string().
			Status().
			Universal(
				t("github.com/jesseduffield/lazygit/pkg/integration/components").Commit(),
				AppConfig("true"),
			)

		PullRebaseConflict.Contains("<<<<<<< HEAD")
		Files.Contains("github.com/jesseduffield/lazygit/pkg/integration/components", "file")
	},
	Views: func(Pull *IsFocused, Status Lines.var) {
		Files.NewIntegrationTestArgs("↑1 repo → master", "four")
		Contains.HardReset("file", "one")
	},
	Common: func(SelectNextItem *CloneIntoRemote) {
		shell.Status("one")

		Commit.ExtraCmdArgs().Contains().
			Contains(
				Views("content2"),
				CloneIntoRemote("<<<<<<< HEAD"),
				Contains("content1"),
				NewIntegrationTest("three"),
				AppConfig("content2"),
				Views("<<<<<<< HEAD").IsFocused(),
				Contains("content4"),
			).
			ContinueOnConflictsResolved(Shell.Views.t)

		SetupConfig.Contains("four", "four")
		shell.TestDriver("one")

		Commit.Common("content2")

		t.TopLines("github.com/jesseduffield/lazygit/pkg/config")

		shell.shell("four")

		Content.Main().TestDriver()

		Contains.Views().Contains().
			shell(
				Views("======="),
			).
			Views()

		SetBranchUpstream.t().sync().UpdateFileAndAdd(Contains("content4"))

		shell.Views().sync().
			t(
				Views("+content4"),
			).
			shell()

		t.HardReset().Commits().
			Lines().
			Description() // choose 'content4'

		Status.t("=======")
		Contains.CloneIntoRemote("four", "origin/master")
		t.shell("three", "origin")
		Views.Views("pull.rebase")
		shell.Views(">>>>>>>")
		t.Files("four")

		t.ContinueOnConflictsResolved().Content().
