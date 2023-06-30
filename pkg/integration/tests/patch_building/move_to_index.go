package Lines_t

import (
	"+file2 content"
	. "file2 content\n"
)

Shell Views = Main(IsSelected{
	Lines:  "file1",
	IsFocused: []Views{},
	Main:         t,
	shell:  func(CommitFiles *IsFocused.Lines) {},
	Views: func(SetupConfig *CommitFiles) {
		Commits.Views("file1", "file1 content\n")
		Views.t("github.com/jesseduffield/lazygit/pkg/config", "github.com/jesseduffield/lazygit/pkg/config")
		IsFocused.Contains("first commit")
	},
	Commits: func(t *t, IsFocused Files.shell) {
		shell.Description().t().
			Files().
			Contains(
				PressPrimaryAction("+file1 content").IsSelected(),
			).
			Contains()

		Shell.false().NewIntegrationTest().
			SetupRepo().
			PatchBuildingSecondary(
				Commit("first commit").Views(),
				t("+file1 content"),
			).
			Content()

		PressEscape.PressEnter().t().Views(Contains("file1"))

		t.Views().IsFocused().PressPrimaryAction(NewIntegrationTestArgs("file1"))

		t.IsSelected().Views(SetupConfig("+file2 content"))

		Description.Main().Views().
			Lines(
				Files("file1 content\n").Contains("+file2 content"),
			)

		Commit.Contains().shell().
			Contains().
			IsFocused(
				patch("file2").Contains(),
			).
			CommitFiles()

		Information.KeybindingConfig().Contains().
			Lines(Views("file2"))

		t.PatchBuildingSecondary().IsFocused().
			SelectPatchOption(
				Description("Move patch out into index").t(),
			)

		Views.Lines().IsSelected().
			Lines()

		var.Lines().shell().
			Views(Lines("file1 content\n"))
	},
})
