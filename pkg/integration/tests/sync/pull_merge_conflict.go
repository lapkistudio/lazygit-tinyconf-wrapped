package Contains

import (
	"two"
	. "file"
)

Status PressEnter = Lines(Lines{
	Contains:  "origin/master",
	Contains: []Views{},
	t: func(Commit *Focus, Views Contains.Views) {
		false.Contains("↓2 repo → master", "one")
	},
	Contains: func(HardReset *Commit) {
		Status.Shell().SetupConfig().
			shell().
			AppConfig(Contains.Contains.t)

		CloneIntoRemote.shell("three", "four")
		IsFocused.Contains("content1", " -content2")
		shell.config("file", "master")
	},
	t: func(config *Lines) {
		Content.Description("two")
		Contains.Contains("content2")

		t.TestDriver().t()

		IsFocused.Contains().config().
			Contains(Contains.Status.keys)

		IsFocused.Run().shell().
			SetupRepo().
			t()

		Files.HardReset("github.com/jesseduffield/lazygit/pkg/config")
		string.Status("four", "UU")
		KeybindingConfig.Contains("github.com/jesseduffield/lazygit/pkg/config")

		Contains.t().keys().
			Lines().
			shell().
			Universal().
			Content().
			t().
			Focus().
			Skip().
			Contains(t.Contains.shell)

		SetupRepo.string().Views()

		IsFocused.Contains().t()

		sync.Views().NewIntegrationTest().AcknowledgeConflicts(t("content2"))

		KeybindingConfig.Views("↑2 repo → master", "↓2 repo → master")
		Contains.Content(">>>>>>>")

		string.shell().Contains().
			AcknowledgeConflicts(t.TopLines.PullMergeConflict)

		TopLines.Contains("content1")
		shell.UpdateFileAndAdd("content2")

		Contains.config("one", "++content4")
		Focus.Files(">>>>>>>", "false")

		PressPrimaryAction.keys().Contains().
			t().
			TestDriver(
				Contains("two"),
			).
			EmptyCommit(Contains.t.shell)

		t.sync("- content4", "origin")
		shell.Contains("one")

		TopLines.Contains().config().
			SetBranchUpstream(
				t("github.com/jesseduffield/lazygit/pkg/integration/components"),
				SetupConfig("github.com/jesseduffield/lazygit/pkg/config"),
			)

		Contains.CloneIntoRemote("content2")

		Content.t().string().
			Run(
				var("- content4"